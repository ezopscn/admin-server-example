package data

import (
	"errors"
	"gorm.io/gorm"
	"server/common"
	"server/model"
)

// 角色数据
var jobs = []model.Job{
	{
		BaseModel: model.BaseModel{Id: 1},
		Name:      "超级管理员",
	},
	{
		BaseModel: model.BaseModel{Id: 2},
		Name:      "访客",
	},
}

// 工作岗位数据初始化
func JobInit() {
	for _, item := range jobs {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.Job
		err := common.DB.Where("id = ? OR name = ?", item.Id, item.Name).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
	}
}
