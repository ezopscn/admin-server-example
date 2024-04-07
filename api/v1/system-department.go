package v1

import (
	"github.com/gin-gonic/gin"
	"server/common"
	"server/dto"
	"server/model"
	"server/pkg/response"
	"strconv"
	"strings"
)

// 获取部门列表
func GETDepartmentListHandler(ctx *gin.Context) {
	var depts []model.Department
	err := common.DB.Preload("ManageUsers").Find(&depts).Error
	if err != nil {
		common.SystemLog.Error(err.Error())
		response.FailedWithMessage("查询部门列表失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": depts,
	})
}

// 添加部门
func AddDepartmentHandler(ctx *gin.Context) {
	var req dto.AddDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.SystemLog.Error(err.Error())
		response.FailedWithMessage("解析添加部门数据失败")
		return
	}

	// 确保字段合法，添加部门不允许新增顶级部门
	if req.ParentId == 0 {
		response.FailedWithMessage("添加部门的父级部门不允许是系统本身")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		response.FailedWithMessage("部门名称不合法，只能包含中文、数字、字母、横线和下划线")
		return
	}

	// 添加数据
	if err := common.DB.Create(&model.Department{
		ParentId: req.ParentId,
		Name:     req.Name,
	}).Error; err != nil {
		common.SystemLog.Error(err.Error())
		response.FailedWithMessage("部门数据添加失败")
		return
	}
	response.Success()
}

// 更新部门
func UpdateDepartmentHandler(ctx *gin.Context) {
	var req dto.UpdateDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.SystemLog.Error(err.Error())
		response.FailedWithMessage("解析更新部门数据失败")
		return
	}
	// 确保字段合法性，预留的部门不能修改父级部门
	if req.Id == 1 && req.ParentId != 0 {
		response.FailedWithMessage("顶级部门的父级部门只能是系统本身")
		return
	}
	if req.Id != 1 && req.ParentId == 0 {
		response.FailedWithMessage("除了顶级部门的父级部门，其它部门的父级部门不能是系统本身")
		return
	}
	if req.Id == 2 && req.ParentId != 1 {
		response.FailedWithMessage("访客部门的父级部门只能是顶级公司")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		response.FailedWithMessage("部门名称不合法，只能包含中文、数字、字母、横线和下划线")
		return
	}

	// 更新部门
	if err := common.DB.Where("id = ?", req.Id).Updates(model.Department{
		Name:     req.Name,
		ParentId: req.ParentId,
	}).Error; err != nil {
		common.SystemLog.Error(err.Error())
		response.FailedWithMessage("部门修改失败，请检查填写的内容是否正确")
		return
	}

	// 获取部门负责人
	if len(req.ManageUsers) > 0 {
		// 查询数据
		var managers []model.User
		if err := common.DB.Find(&managers, req.ManageUsers).Error; err != nil {
			common.SystemLog.Error(err.Error())
			response.FailedWithMessage("查询部门负责人失败")
			return
		}
		// 更新部门负责人数据
		if err := common.DB.Model(&model.Department{}).Where("id = ?", req.Id).Association("ManageUsers").Append(managers); err != nil {
			common.SystemLog.Error(err.Error())
			response.FailedWithMessage("更新部门信息失败")
			return
		}
	}
	response.Success()
}

// 删除部门
func DeleteDepartmentHandler(ctx *gin.Context) {
	sid := ctx.Param("departmentId")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil || id == 0 {
		response.FailedWithMessage("部门删除失败，需要删除的部门id不合法")
		return
	}
}
