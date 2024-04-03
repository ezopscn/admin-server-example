package v1

import (
	"github.com/gin-gonic/gin"
	"server/common"
	"server/dto"
	"server/model"
	"server/pkg/response"
	"server/pkg/utils"
	"server/service"
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
	// 获取登录用户用户名
	username, err := utils.GetUsernameFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取登录用户信息异常")
		return
	}

	// 获取用户基本关联信息
	user, err := service.GetBaseUserInfoByUsername(username)
	if err != nil {
		response.FailedWithMessage("查询用户基本信息失败")
		return
	}

	response.SuccessWithData(gin.H{
		"info": user,
	})
}

// 获取指定用户的信息
func GETSpecifyUserInfoHandler(ctx *gin.Context) {
	// 获取登录用户基本关联信息
	cusername, err := utils.GetUsernameFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取登录用户信息异常")
		return
	}
	cuser, err := service.GetBaseUserInfoByUsername(cusername)
	if err != nil {
		response.FailedWithMessage("查询登录用户基本信息失败")
		return
	}

	// 获取指定用户基本关联信息
	username := ctx.Param("username")
	user, err := service.GetBaseUserInfoByUsername(username)
	if err != nil {
		response.FailedWithMessage("查询指定用户基本信息失败")
		return
	}

	// 判断用户是不是隐藏手机号
	if *user.ShowPhone == common.False {
		// 如果查询的用户不是自己，并且自己不是系统预留的超级管理员，则需要隐藏手机号
		if cuser.RoleId != common.SuperAdminRoleId && cuser.Username != user.Username {
			user.Phone = utils.MaskPhone(user.Phone)
		}
	}

	response.SuccessWithData(gin.H{
		"info": user,
	})
}

// 获取基础用户信息列表
func GETAllUserListHandler(ctx *gin.Context) {
	var users []dto.BaseUserInfo
	err := common.DB.Model(model.User{}).Find(&users).Error
	if err != nil {
		response.FailedWithMessage("查询基础用户列表失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": users,
	})
}
