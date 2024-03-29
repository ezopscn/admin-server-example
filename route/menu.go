package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

// 菜单开放路由组
func MenuPublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	prg := rg.Group("/menu").Use(auth.MiddlewareFunc())
	prg.GET("/list", v1.GETMenuListHandler) // 获取菜单列表
	return prg
}

// 菜单授权路由组
func MenuCasbinRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	crg := rg.Group("/menu").Use(auth.MiddlewareFunc())
	return crg
}
