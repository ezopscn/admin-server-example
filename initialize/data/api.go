package data

import (
	"errors"
	"gorm.io/gorm"
	"server/common"
	"server/model"
)

// 接口类型数据
var apiTypes = []model.APIType{
	{
		Id:   1,
		Name: "部门相关",
	},
	{
		Id:   2,
		Name: "用户相关",
	},
	{
		Id:   3,
		Name: "角色相关",
	},
	{
		Id:   4,
		Name: "菜单相关",
	},
	{
		Id:   5,
		Name: "接口相关",
	},
}

// 接口类型数据初始化
func APITypeInit() {
	for _, item := range apiTypes {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.APIType
		err := common.DB.Where("id = ? OR name = ?", item.Id, item.Name).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
	}
}
