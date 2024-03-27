package data

import (
	"errors"
	"gorm.io/gorm"
	"server/common"
	"server/model"
)

// 设置数据
var settings = []model.Setting{
	{
		BaseModel:   model.BaseModel{Id: 1},
		Name:        "2FA",
		Value:       "false",
		Description: "是否开启双因子认证设置，开启后用户需要绑定手机，通过实时验证码进行登录",
	},
	{
		BaseModel:   model.BaseModel{Id: 2},
		Name:        "Organization",
		Value:       "深圳运维集团",
		Description: "组织名称（公司名称）",
	},
}

// 系统设置数据初始化
func SettingInit() {
	for _, item := range settings {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.Setting
		err := common.DB.Where("id = ? OR name = ?", item.Id, item.Name).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
	}
}
