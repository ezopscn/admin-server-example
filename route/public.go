package route

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

// 开放路由组
func PublicRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/health", api.HealthHandler)   // 健康检查接口
	rg.GET("/info", api.InfoHandler)       // 开发者信息接口
	rg.GET("/version", api.VersionHandler) // 版本信息接口
	return rg
}
