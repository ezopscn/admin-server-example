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
		Sort:      91,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 3},
				Label:     "部门管理",
				Icon:      "",
				Key:       "/system/department",
				Sort:      10,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 4},
				Label:     "岗位管理",
				Icon:      "",
				Key:       "/system/job",
				Sort:      20,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 5},
				Label:     "用户管理",
				Icon:      "",
				Key:       "/system/user",
				Sort:      30,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 6},
				Label:     "角色管理",
				Icon:      "",
				Key:       "/system/role",
				Sort:      40,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 7},
				Label:     "菜单管理",
				Icon:      "",
				Key:       "/system/menu",
				Sort:      50,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 8},
				Label:     "接口管理",
				Icon:      "",
				Key:       "/system/api",
				Sort:      60,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 9},
				Label:     "授权管理",
				Icon:      "",
				Key:       "/system/privilege",
				Sort:      70,
				ParentId:  2,
			},
			{
				BaseModel: model.BaseModel{Id: 10},
				Label:     "服务配置",
				Icon:      "",
				Key:       "/system/setting",
				Sort:      80,
				ParentId:  2,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 11},
		Label:     "日志审计",
		Icon:      "InsuranceOutlined",
		Key:       "/log",
		Sort:      92,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 12},
				Label:     "登录日志",
				Icon:      "",
				Key:       "/log/login",
				Sort:      10,
				ParentId:  11,
				Roles: []model.Role{
					roles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 13},
				Label:     "操作日志",
				Icon:      "",
				Key:       "/log/operation",
				Sort:      20,
				ParentId:  11,
				Roles: []model.Role{
					roles[1],
				},
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 14},
		Label:     "个人中心",
		Icon:      "UserOutlined",
		Key:       "/me",
		Sort:      93,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 15},
		Label:     "获取帮助",
		Icon:      "FileProtectOutlined",
		Key:       "/help",
		Sort:      100,
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
