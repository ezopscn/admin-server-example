package data

import (
	"errors"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
	"server/common"
	"server/model"
	"server/pkg/utils"
)

// 初始化密码
var password = "ezops.cn"

// 用户初始化数据
var users = []model.User{
	{
		BaseModel:        model.BaseModel{Id: 1},
		Username:         "admin",
		ENName:           "Admin",
		CNName:           "超管",
		Phone:            "18888888888",
		ShowPhone:        &common.False,
		Email:            "admin@ezops.cn",
		Password:         utils.CryptoPassword(password),
		JobId:            1,
		JoinTime:         carbon.Now(),
		DepartmentId:     1,
		OfficeProvinceId: 44,
		OfficeCityId:     4403,
		OfficeAddress:    "下沙社区",
		OfficeStation:    "T1-9F-A-11",
		NativeProvinceId: 51,
		NativeCityId:     5105,
		Gender:           &common.Male,
		Avatar:           "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png",
		Birthday:         carbon.Now(),
		CreatorId:        0,
		Status:           &common.Active,
		RoleId:           1,
	},
	{
		BaseModel:        model.BaseModel{Id: 2},
		Username:         "guest",
		ENName:           "Guest",
		CNName:           "访客",
		Phone:            "19999999999",
		ShowPhone:        &common.True,
		Email:            "guest@ezops.cn",
		Password:         utils.CryptoPassword(password),
		JobId:            2,
		JoinTime:         carbon.Now(),
		DepartmentId:     2,
		OfficeProvinceId: 44,
		OfficeCityId:     4403,
		OfficeAddress:    "下沙社区",
		OfficeStation:    "T1-9F-A-12",
		NativeProvinceId: 51,
		NativeCityId:     5105,
		Gender:           &common.Female,
		Avatar:           "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png",
		Birthday:         carbon.Now(),
		CreatorId:        0,
		Status:           &common.Active,
		RoleId:           2,
	},
}

// 用户数据初始化
func UserInit() {
	for _, item := range users {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.User
		err := common.DB.Where("id = ? OR username = ? OR phone = ? OR email = ?",
			item.Id,
			item.Username,
			item.Phone,
			item.Email).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
	}
}
