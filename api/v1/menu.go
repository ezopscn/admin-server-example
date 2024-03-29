package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/common"
	"server/model"
	"server/pkg/response"
	"server/pkg/utils"
)

// 获取菜单列表
func GETMenuListHandler(ctx *gin.Context) {
	// 获取当前登录用户的角色 Id
	roleId, err := utils.GetRoleIdFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取登录用户的角色信息失败")
		return
	}
	// 查询角色的菜单，如果角色是系统预设超级管理员，则获取所有菜单
	var menus []model.Menu
	var role model.Role
	if roleId == common.SuperAdminRoleId {
		err = common.DB.Order("sort").Find(&menus).Error
	} else {
		// 预加载字段排序，也可以对预加载字段进行条件预加载
		err = common.DB.Preload("Menus", func(db *gorm.DB) *gorm.DB {
			return db.Order("menu.sort")
		}).Where("id = ?", roleId).First(&role).Error
		if err == nil {
			menus = role.Menus
		}
	}
	if err != nil {
		response.FailedWithMessage("查询菜单列表失败")
		return
	}
	response.SuccessWithData(gin.H{
		"list": menus,
	})
}
