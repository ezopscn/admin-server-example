package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

// 部门开放路由组
func DepartmentPublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	prg := rg.Group("/department").Use(auth.MiddlewareFunc())
	prg.GET("/list", v1.GETDepartmentListHandler) // 获取部门列表
	return prg
}

// 部门授权路由组
func DepartmentCasbinRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	crg := rg.Group("/department").Use(auth.MiddlewareFunc())
	return crg
}
