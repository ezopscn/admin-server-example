package data

import (
	"errors"
	"gorm.io/gorm"
	"server/common"
	"server/model"
)

// 菜单数据
var menus = []model.Menu{
	{
		BaseModel: model.BaseModel{Id: 1},
		Label:     "工作空间",
		Icon:      "HomeOutlined",
		Key:       "/dashboard",
		Sort:      0,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 2},
		Label:     "系统配置",
		Icon:      "SettingOutlined",
		Key:       "/system",
		Sort:      9000,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 3},
				Label:     "部门管理",
				Icon:      "",
				Key:       "/system/department",
				Sort:      9010,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 4},
				Label:     "用户管理",
				Icon:      "",
				Key:       "/system/user",
				Sort:      9020,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 5},
				Label:     "角色管理",
				Icon:      "",
				Key:       "/system/role",
				Sort:      9030,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 6},
				Label:     "菜单管理",
				Icon:      "",
				Key:       "/system/menu",
				Sort:      9040,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 7},
				Label:     "接口管理",
				Icon:      "",
				Key:       "/system/api",
				Sort:      9050,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 8},
				Label:     "授权管理",
				Icon:      "",
				Key:       "/system/privilege",
				Sort:      9060,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 9},
				Label:     "服务配置",
				Icon:      "",
				Key:       "/system/setting",
				Sort:      9070,
				ParentId:  2,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 10},
		Label:     "日志审计",
		Icon:      "InsuranceOutlined",
		Key:       "/log",
		Sort:      9100,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 11},
				Label:     "登录日志",
				Icon:      "",
				Key:       "/log/login",
				Sort:      9110,
				ParentId:  10,
				Roles: []model.Role{
					roles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 12},
				Label:     "操作日志",
				Icon:      "",
				Key:       "/log/operation",
				Sort:      9120,
				ParentId:  10,
				Roles: []model.Role{
					roles[1],
				},
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 13},
		Label:     "个人中心",
		Icon:      "UserOutlined",
		Key:       "/me",
		Sort:      9200,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 14},
		Label:     "获取帮助",
		Icon:      "FileProtectOutlined",
		Key:       "/help",
		Sort:      9999,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
	},
}

// 递归插入数据方法
func insertMenusData(menus []model.Menu) {
	var menu model.Menu
	for _, item := range menus {
		// 查看数据是否存在，如果不存在才执行创建，注意 Key 是关键字，需要特别处理，否则会报错
		err := common.DB.Where("id = ? OR label = ? OR `key` = ?", item.Id, item.Label, item.Key).First(&menu).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
		// 递归插入子菜单
		if len(item.Children) > 0 {
			insertMenusData(item.Children)
		}
	}
}

// 菜单数据初始化
func MenuInit() {
	insertMenusData(menus)
}
