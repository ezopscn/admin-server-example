package data

import (
	"errors"
	"gorm.io/gorm"
	"server/common"
	"server/model"
)

// 角色数据
var roles = []model.Role{
	{
		BaseModel:   model.BaseModel{Id: 1},
		Name:        "超级管理员",
		Description: "超级管理员，系统预留角色，系统最高管理权限",
	},
	{
		BaseModel:   model.BaseModel{Id: 2},
		Name:        "访客",
		Description: "游客，系统预留角色，系统最低浏览权限",
	},
}

// 角色数据初始化
func RoleInit() {
	for _, item := range roles {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.Role
		err := common.DB.Where("id = ? OR name = ?", item.Id, item.Name).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
	}
}
