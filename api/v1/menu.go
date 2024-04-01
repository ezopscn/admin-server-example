package v1

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"server/common"
	"server/dto"
	"server/model"
	"server/pkg/response"
	"server/pkg/utils"
	"server/service"
	"strconv"
)

// 获取所有菜单
func GETMenuAllHandler(ctx *gin.Context) {
	var menus []model.Menu
	err := common.DB.Order("sort").Find(&menus).Error
	if err != nil {
		response.FailedWithMessage("获取所有菜单数据失败")
		return
	}
	response.SuccessWithData(gin.H{
		"menus": menus,
	})
}

// 获取用户菜单列表
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

// 添加菜单
func AddMenuHandler(ctx *gin.Context) {
	var req dto.AddMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage("解析添加菜单数据失败")
		return
	}

	// 顶级菜单必须配置图标
	if (req.ParentId == 0 && req.Icon == "") || (req.ParentId != 0 && req.Icon != "") {
		response.FailedWithMessage("顶级菜单必须配置图标，子菜单则不需要配置图标")
		return
	}

	// 结构体转换
	var menu model.Menu
	if err := copier.Copy(&menu, &req); err != nil {
		response.FailedWithMessage("菜单数据不合法")
		return
	}

	// 添加菜单
	if err := common.DB.Create(&menu).Error; err != nil {
		response.FailedWithMessage("菜单添加失败，请检查填写的内容是否正确")
		return
	}

	response.Success()
}

// 更新菜单
func UpdateMenuHandler(ctx *gin.Context) {
	var req dto.UpdateMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage("解析修改菜单数据失败")
		return
	}

	// 顶级菜单必须配置图标
	if (req.ParentId == 0 && req.Icon == "") || (req.ParentId != 0 && req.Icon != "") {
		response.FailedWithMessage("顶级菜单必须配置图标，子菜单则不需要配置图标")
		return
	}

	// 结构体转换成 map[string]interface{
	m := make(map[string]interface{})
	b, _ := sonic.Marshal(&req)
	if err := sonic.Unmarshal(b, &m); err != nil {
		response.FailedWithMessage("菜单数据不合法")
		return
	}

	// 更新菜单
	if err := common.DB.Model(&model.Menu{}).Where("id = ?", req.Id).Updates(m).Error; err != nil {
		response.FailedWithMessage("菜单修改失败，请检查填写的内容是否正确")
		return
	}

	response.Success()
}

// 删除菜单
func DeleteMenuHandler(ctx *gin.Context) {
	sid := ctx.Param("menuId")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil || id == 0 {
		response.FailedWithMessage("菜单删除失败，需要删除的菜单id不合法")
		return
	}
	// Unscoped 永久删除，不是软删除
	err = common.DB.Where("id = ?", id).Unscoped().Delete(&model.Menu{}).Error
	if err != nil {
		response.FailedWithMessage("菜单删除失败，请刷新后重试")
		return
	}
	// 需要删除关联数据

	response.Success()
}
