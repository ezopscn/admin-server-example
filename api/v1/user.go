package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/common"
	"server/model"
	"server/pkg/response"
	"server/pkg/utils"
)

// 获取用户数量统计
func GETUserCountHandler(ctx *gin.Context) {
	// 查询用户数量
	var count int64
	if err := common.DB.Model(&model.User{}).Count(&count).Error; err != nil {
		response.FailedWithMessage("获取用户统计失败")
		return
	}
	response.SuccessWithData(gin.H{
		"count": count,
	})
}

// 获取当前用户的信息
func GETCurrentUserInfoHandler(ctx *gin.Context) {
	// 从 context 中获取当前用户的用户名
	username, err := utils.GetUsernameFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取登录用户信息失败")
		return
	}
	// 查询用户
	var user model.User
	err = common.DB.Where("username = ?", username).Preload(clause.Associations).First(&user).Error
	if err != nil {
		response.FailedWithMessage("获取当前用户信息失败")
		return
	}
	response.SuccessWithData(gin.H{
		"info": user,
	})
}

// 获取指定用户的信息
func GETSpecifyUserInfoHandler(ctx *gin.Context) {
	username := ctx.Param("username") // 获取 URI 参数
	// 查询用户信息
	var user model.User
	// clause.Associations 预加载所有字段
	err := common.DB.Where("username = ?", username).Preload(clause.Associations).First(&user).Error
	if err != nil {
		response.FailedWithMessage("查询指定用户信息失败")
		return
	}
	// 判断用户是不是隐藏手机号
	if *user.ShowPhone == common.False {
		cname, roleId, err := utils.GetUsernameAndRoleIdFromContext(ctx)
		if err != nil {
			response.FailedWithMessage("获取登录用户信息失败")
			return
		}
		// 如果查询的用户不是自己，并且自己不是系统预留的超级管理员，则需要隐藏手机号
		if roleId != common.SuperAdminRoleId && cname != user.Username {
			user.Phone = utils.MaskPhone(user.Phone)
		}
	}
	response.SuccessWithData(gin.H{
		"info": user,
	})
}
