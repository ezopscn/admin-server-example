package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

// 用户开放路由组
func UserPublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	prg := rg.Group("/user").Use(auth.MiddlewareFunc())
	prg.GET("/count", v1.GETUserCountHandler)                // 用户数量统计
	prg.GET("/info", v1.GETCurrentUserInfoHandler)           // 获取当前用户的信息
	prg.GET("/info/:username", v1.GETSpecifyUserInfoHandler) // 获取指定用户的信息
	return prg
}

// 用户授权路由组
func UserCasbinRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	crg := rg.Group("/user").Use(auth.MiddlewareFunc())
	return crg
}
