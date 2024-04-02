package service

import (
	"github.com/bytedance/sonic"
	"gorm.io/gorm"
	"server/common"
	"server/model"
	"server/pkg/gedis"
	"strconv"
)

// 根据角色 Id 查询菜单列表
func GetMenuListByRoleId(roleId uint) (menus []model.Menu, err error) {
	// 查询缓存
	var cache = gedis.NewOperation()
	menuListKey := common.GenerateRedisKey(common.RKP.MenuList, strconv.Itoa(int(roleId)))
	dataStr := cache.GetString(menuListKey).Unwrap()
	if dataStr != "" {
		err = sonic.Unmarshal([]byte(dataStr), &menus)
		if err == nil {
			return
		}
	}

	// 查询数据库
	if roleId == common.SuperAdminRoleId {
		err = common.DB.Order("sort").Find(&menus).Error
	} else {
		var role model.Role
		err = common.DB.Preload("Menus", func(db *gorm.DB) *gorm.DB {
			return db.Order("menu.sort")
		}).Where("id = ?", roleId).First(&role).Error
		if err == nil {
			menus = role.Menus
		}
	}
	if err != nil {
		return
	}

	// 保存缓存
	jsonData, _ := sonic.Marshal(menus)
	cache.Set(menuListKey, jsonData, gedis.WithExpire(common.MenuListExpireTime))
	return
}
