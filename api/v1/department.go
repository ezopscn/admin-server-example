package v1

import (
	"github.com/gin-gonic/gin"
	"server/common"
	"server/model"
	"server/pkg/response"
)

// 获取部门列表
func GETDepartmentListHandler(ctx *gin.Context) {
	var depts []model.Department
	err := common.DB.Find(&depts).Error
	if err != nil {
		response.FailedWithMessage("查询部门列表失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": depts,
	})
}
