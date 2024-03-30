package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm/clause"
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
	var cache = gedis.NewOperation()
	times := cache.GetInt(key).UnwrapWithDefaultValue(0)
	if times >= common.Config.Login.WrongTimes {
		return nil, errors.New("认证次数超过上限，账户已锁定")
	}

	// 4.用户未锁定，则验证用户登录账户类型并查询用户，如果没查到则返回账户密码错误
	db := common.DB
	var user model.User
	var err error

	// 判断用户登录采用的方式，支持使用用户名 / 手机号 / Email
	dbt := db.Preload(clause.Associations)
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
		cache.Set(key, times, gedis.WithExpire(time.Duration(common.Config.Login.LockTime)*time.Second))
		return nil, errors.New("用户名或密码错误")
	}

	// 6.密码正确，则进行用户状态校验
	if *user.Status == common.Disable {
		return nil, errors.New("用户已禁用，请联系管理员")
	}

	// 7.登录正确
	// 删除错误 redis 中的次数
	_, _ = cache.Del(key)

	// 更新数据库中登录信息
	common.DB.Model(&model.User{}).
		Where("username = ?", user.Username).
		Updates(map[string]interface{}{
			"last_login_ip":   ip,
			"last_login_time": carbon.Now(),
		})

	// 8.设置 Context，方便后面验证使用
	ctx.Set("Username", user.Username)

	// 9.查看系统是否开启双因子认证
	var setting model.Setting
	if err = db.Where("name = ?", "2FA").First(&setting).Error; err != nil {
		return nil, errors.New("查询双因子认证状态失败")
	}
	// 如果开启了双因子认证，则需要判断用户是否绑定设备
	if setting.Value == "true" {
		// 没有绑定，则需要让用户绑定
		if user.Secret == "" {
			return &user, errors.New("NotBind2FA")
		}
		// 已经绑定了，则需要验证用户的验证码
		if !utils.ValidateTOTPCode(req.VerificationCode, user.Secret) {
			times += 1
			cache.Set(key, times, gedis.WithExpire(time.Duration(common.Config.Login.LockTime)*time.Second))
			return nil, errors.New("手机令牌验证码错误")
		}
	}

	// 10.将用户信息存入 Redis 中，方便最近的请求使用，而不需要再去完整的查询数据库
	// 将用户信息编码为 json
	jsonData, err := json.Marshal(user)
	if err != nil {
		common.SystemLog.Error(err.Error())
		return nil, errors.New("用户信息编码失败错误")
	}
	userInfoKey := common.GenerateRedisKey(common.RKP.UserInfo, user.Username)
	cache.Set(userInfoKey, jsonData, gedis.WithExpire(common.UserInfoExpireTime))

	// 以指针的方式将数据传递给 PayloadFunc 函数继续处理
	return &user, nil
}

// 隶属 Login 中间件，接收 Authenticator 验证成功后传递过来的数据，进行封装成 Token
// MapClaims 必须包含 IdentityKey
// MapClaims 会被嵌入 Token 中，后续可以通过 ExtractClaims 对 Token 进行解析获取到
func payloadFunc(data interface{}) jwt.MapClaims {
	// 断言判断获取传递过来数据是不是用户数据
	if user, ok := data.(*model.User); ok {
		// 封装 id 和 username 这种几乎不会变的字段,方便后面直接使用
		return jwt.MapClaims{
			jwt.IdentityKey: user.Id,       // 用户Id
			"Username":      user.Username, // 用户名
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
	// 未绑定设备，需要单独返回
	if message == "NotBind2FA" {
		// 获取前面 Context 设置的值，并验证是否合法
		v, _ := ctx.Get("Username")
		username, ok := v.(string)
		if !ok || username == "" {
			response.FailedWithMessage("获取用户登录信息异常")
			return
		}

		// 生成二维码链接
		key, err := utils.GenerateTOTPKey(username)
		if err != nil {
			response.FailedWithMessage("获取 TOTP 信息异常")
			return
		}
		url := key.URL()
		secret := key.Secret()

		// 保存 Secret 到数据库
		err = common.DB.Model(&model.User{}).Where("username = ?", username).Update("secret", secret).Error
		if err != nil {
			response.FailedWithMessage("保存用户 TOTP Secret 失败")
			return
		}
		response.FailedWithCodeAndData(response.NotBind2FA, gin.H{"url": url})
		return
	}

	response.FailedWithCodeAndMessage(response.Unauthorized, message)
}

// 用户登录后的中间件，用于解析 Token
func identityHandler(ctx *gin.Context) interface{} {
	// 获取登录用户用户名
	username, err := utils.GetUsernameFromContext(ctx)
	if err != nil {
		return nil
	}
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
	// 获取登录用户用户名
	username, err := utils.GetUsernameFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取登录用户信息异常")
		return
	}

	// 清理 Redis 保存的数据
	cache := gedis.NewOperation()
	_, _ = cache.Del(common.GenerateRedisKey(common.RKP.LoginToken, username))
	_, _ = cache.Del(common.GenerateRedisKey(common.RKP.UserInfo, username))
	response.Success()
}
