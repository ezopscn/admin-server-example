package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
	"server/pkg/response"
)

// 健康检测接口
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// 开发者信息接口
func InfoHandler(ctx *gin.Context) {
	response.SuccessWithData(gin.H{
		"developer": "Jayce",
		"email":     "ezops.cn@gmail.com",
	})
}

// 系统版本接口
func VersionHandler(ctx *gin.Context) {
	response.SuccessWithData(gin.H{
		"version": common.Version,
	})
}
