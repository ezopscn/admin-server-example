package v1

import (
	"github.com/gin-gonic/gin"
	"server/common"
	"server/model"
	"server/pkg/response"
)

// 获取双因子登录是否开启
func GET2FAStatusHandler(ctx *gin.Context) {
	var status bool
	var setting model.Setting
	err := common.DB.Where("name = ?", "2FA").First(&setting).Error
	if err == nil && setting.Value == "true" {
		status = true // 只有当值为 true 才表示开启
	}
	response.SuccessWithData(gin.H{
		"status": status,
	})
}
