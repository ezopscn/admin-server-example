package initialize

import (
	"github.com/gin-gonic/gin"
	"server/common"
	"server/middleware"
	"server/route"
)

// 路由初始化
func Router() *gin.Engine {
	// 创建一个没中间件的路由引擎
	r := gin.New()

	// 中间件
	r.Use(middleware.AccessLog)       // 请求日志中间件
	r.Use(middleware.Cors)            // 跨域访问中间件
	r.Use(middleware.Exception)       // 异常捕获中间件
	auth, err := middleware.JWTAuth() // JWT认证中间件
	if err != nil {
		panic(err)
	}

	rg := r.Group(common.APIPrefix)

	// 无需登录的路由组
	route.PublicRoutes(rg, auth) // 开放路由组

	// 需要登录的路由组
	route.PublicAuthRoutes(rg, auth) // 开放认证路由组
	route.UserPublicRoutes(rg, auth) // 用户开放路由组
	route.RolePublicRoutes(rg, auth) // 角色开放路由组
	route.MenuPublicRoutes(rg, auth) // 菜单开放路由组

	// 需要授权的路由组
	route.UserCasbinRoutes(rg, auth) // 用户授权路由组
	route.RoleCasbinRoutes(rg, auth) // 角色授权路由组
	route.MenuCasbinRoutes(rg, auth) // 菜单授权路由组

	common.SystemLog.Info("系统路由初始化完成")
	return r
}
