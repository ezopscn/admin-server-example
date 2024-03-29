package v1

import (
	"github.com/gin-gonic/gin"
	"server/common"
	"server/model"
	"server/pkg/response"
)

// 获取角色数量统计
func GETRoleCountHandler(ctx *gin.Context) {
	// 查询角色数量
	var count int64
	if err := common.DB.Model(&model.Role{}).Count(&count).Error; err != nil {
		response.FailedWithMessage("获取角色统计失败")
		return
	}
	response.SuccessWithData(gin.H{
		"count": count,
	})
}
