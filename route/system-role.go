package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

// 角色开放路由组
func RolePublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	prg := rg.Group("/role").Use(auth.MiddlewareFunc())
	prg.GET("/count", v1.GETRoleCountHandler) // 角色数量统计
	return prg
}

// 角色授权路由组
func RoleCasbinRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	crg := rg.Group("/role").Use(auth.MiddlewareFunc())
	return crg
}
