package v1

import (
	"github.com/gin-gonic/gin"
	"server/pkg/response"
	"server/pkg/utils"
	"server/service"
)

// 获取菜单列表
func GETMenuListHandler(ctx *gin.Context) {
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

	// 查询角色的菜单，如果角色是系统预设超级管理员，则获取所有菜单
	menus, err := service.GetMenuListByRoleId(cuser.RoleId)
	if err != nil {
		response.FailedWithMessage("查询菜单列表失败")
		return
	}

	response.SuccessWithData(gin.H{
		"list": menus,
	})
}
