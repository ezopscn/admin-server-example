package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"server/api"
	v1 "server/api/v1"
)

// 开放路由组
func PublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/health", api.HealthHandler)          // 健康检查接口
	rg.GET("/info", api.InfoHandler)              // 开发者信息接口
	rg.GET("/version", api.VersionHandler)        // 版本信息接口
	rg.GET("/2fa/status", v1.GET2FAStatusHandler) // 获取双因子认证开启状态
	rg.POST("/login", auth.LoginHandler)          // 用户登录接口
	return rg
}

// 认证开放路由组
func PublicAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	arg := rg.Use(auth.MiddlewareFunc())
	arg.GET("/logout", auth.LogoutHandler) // 用户登出接口
	return arg
}
