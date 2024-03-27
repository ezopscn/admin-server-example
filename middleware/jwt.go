package middleware

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"server/common"
	"server/dto"
	"server/model"
	"server/pkg/gedis"
	"server/pkg/response"
	"server/pkg/utils"
	"time"
)

// JWT 认证中间件
func JWTAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           common.Config.JWT.Realm,                                // JWT 标识
		Key:             []byte(common.Config.JWT.Key),                          // 签名 Key
		Timeout:         time.Duration(common.Config.JWT.Timeout) * time.Second, // Token 有效期
		Authenticator:   authenticator,                                          // 用户登录校验
		PayloadFunc:     payloadFunc,                                            // Token 封装
		LoginResponse:   loginResponse,                                          // 登录成功响应
		Unauthorized:    unauthorized,                                           // 登录，认证失败响应
		IdentityHandler: identityHandler,                                        // 解析 Token
		Authorizator:    authorizator,                                           // 验证 Token
		LogoutResponse:  logoutResponse,                                         // 注销登录
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",     // Token 查找的字段
		TokenHeadName:   "Bearer",                                               // Token 请求头名称
	})
}

// 隶属 Login 中间件，当调用 LoginHandler 就会触发
// 通过从 ctx 中检索出数据，进行用户登录认证
// 返回包含用户信息的 Map 或者 Struct
func authenticator(ctx *gin.Context) (interface{}, error) {
	// 1.获取用户登录提交的数据
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, errors.New("获取用户登录信息失败")
	}

	// 2.获取客户端 IP，确保代理透传客户端真实 IP，如果获取 IP 失败则使用 None 做标识
	ip := ctx.ClientIP()
	if ip == "" {
		ip = "None"
	}

	// Redis 中保存单个 IP+用户 组成的登录错误次数的 Key
	key := common.GenerateRedisKey(common.RKP.LoginWrongTimes, fmt.Sprintf("%s%s%s", ip, common.RDSKeySeparator, req.Account))

	// 3.获取 redis 中该 IP+用户 错误次数，避免恶意登录
	var conn = gedis.NewOperation()
	times := conn.GetInt(key).UnwrapWithDefaultValue(0)
	if times >= common.Config.Login.WrongTimes {
		return nil, errors.New("认证次数超过上限，账户已锁定")
	}

	// 4.用户未锁定，则验证用户登录账户类型并查询用户，如果没查到则返回账户密码错误
	db := common.DB
	var user model.User
	var err error

	// 判断用户登录采用的方式，支持使用用户名 / 手机号 / Email
	dbt := db.Preload("Role")
	if utils.IsPhone(req.Account) {
		err = dbt.Where("phone = ?", req.Account).First(&user).Error
	} else if utils.IsEmail(req.Account) {
		err = dbt.Where("email = ?", req.Account).First(&user).Error
	} else {
		err = dbt.Where("username = ?", req.Account).First(&user).Error
	}

	// 5.用户查询失败，密码不对，都在原有的 redis 保存的错误次数上 +1，并设置过期时间
	if err != nil || !utils.ComparePassword(user.Password, req.Password) {
		times += 1
		conn.Set(key, times, gedis.WithExpire(time.Duration(common.Config.Login.LockTime)*time.Second))
		return nil, errors.New("用户名或密码错误")
	}

	// 6.密码正确，则进行用户状态校验
	if user.Status == &common.Disable {
		return nil, errors.New("用户已禁用，请联系管理员")
	}

	// 7.登录成功
	// 删除错误 redis 中的次数
	_, _ = conn.Del(key)

	// 更新数据库中登录信息
	common.DB.Model(&model.User{}).
		Where("username = ?", user.Username).
		Updates(map[string]interface{}{
			"last_login_ip":   ip,
			"last_login_time": carbon.Now(),
		})

	// 8.返回登录信息
	// 设置 Context，方便后面使用
	ctx.Set("Username", user.Username)

	// 以指针的方式将数据传递给 PayloadFunc 函数继续处理
	return &user, nil
}

// 隶属 Login 中间件，接收 Authenticator 验证成功后传递过来的数据，进行封装成 Token
// MapClaims 必须包含 IdentityKey
// MapClaims 会被嵌入 Token 中，后续可以通过 ExtractClaims 对 Token 进行解析获取到
func payloadFunc(data interface{}) jwt.MapClaims {
	// 断言判断获取传递过来数据是不是用户数据
	if user, ok := data.(*model.User); ok {
		// 封装一些常用的字段，方便直接使用前端和后端都能直接使用
		return jwt.MapClaims{
			jwt.IdentityKey: user.Username,     // 用户名
			"Username":      user.Username,     // 用户名
			"ENName":        user.ENName,       // 英文名
			"CNName":        user.CNName,       // 中文名
			"Avatar":        user.Avatar,       // 用户头像
			"Phone":         user.Phone,        // 用户手机
			"Email":         user.Email,        // 邮箱
			"RoleId":        user.Role.Id,      // 角色 Id
			"RoleName":      user.Role.Name,    // 角色名称
			"DepartmentId":  user.DepartmentId, // 部门 Id
		}
	}
	return jwt.MapClaims{}
}

// 隶属 Login 中间件，响应用户请求
// 接收 PayloadFunc 传递过来的 Token 信息，返回登录成功
func loginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	// 用户响应数据
	var res dto.LoginResponse
	res.Token = token
	res.Expire = expire.Format(common.SecTimeFormat)

	// 不允许多设备登录配置
	if !common.Config.Login.MultiDevices {
		// 获取前面 Context 设置的值，并验证是否合法
		v, _ := ctx.Get("Username")
		username, ok := v.(string)
		if !ok || username == "" {
			response.FailedWithMessage("用户登录状态异常")
			return
		}

		// 将新的 Token 存到 Redis 中，用户下一次请求的时候就去验证该 Token
		key := common.GenerateRedisKey(common.RKP.LoginToken, username)
		cache := gedis.NewOperation()
		cache.Set(key, token, gedis.WithExpire(time.Duration(common.Config.JWT.Timeout)*time.Second))
	}

	// 响应请求
	response.SuccessWithData(res)
}

// 登录失败，验证失败的响应
func unauthorized(ctx *gin.Context, code int, message string) {
	response.FailedWithCodeAndMessage(response.Unauthorized, message)
}

// 用户登录后的中间件，用于解析 Token
func identityHandler(ctx *gin.Context) interface{} {
	// 从 Context 中获取用户名
	claims := jwt.ExtractClaims(ctx)
	username, _ := claims["identity"].(string)
	return &model.User{
		Username: username,
	}
}

// 用户登录后的中间件，用于验证 Token
func authorizator(data interface{}, ctx *gin.Context) bool {
	user, ok := data.(*model.User)
	if ok {
		// 不允许多设备登录配置
		if !common.Config.Login.MultiDevices {
			// Key
			token := jwt.GetToken(ctx)
			key := common.GenerateRedisKey(common.RKP.LoginToken, user.Username)

			// 验证该用户的 Token 和 Redis 中的是否一致
			cache := gedis.NewOperation()
			if cache.GetString(key).Unwrap() != token {
				return false
			}
		}
		return true
	}
	return false
}

// 注销登录
func logoutResponse(ctx *gin.Context, code int) {
	// 获取用户名
	claims := jwt.ExtractClaims(ctx)
	username, _ := claims["identity"].(string)

	// 清理 Redis 保存的数据
	cache := gedis.NewOperation()
	_, _ = cache.Del(common.GenerateRedisKey(common.RKP.LoginToken, username))
	response.Success()
}
