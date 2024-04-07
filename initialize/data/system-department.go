package data

import (
	"errors"
	"gorm.io/gorm"
	"server/common"
	"server/model"
)

// 测试部门数据
var departments = []model.Department{
	{
		BaseModel: model.BaseModel{Id: 1},
		Name:      "集团总部",
		ParentId:  0,
		Users: []model.User{
			users[0],
		},
		ManageUsers: []model.User{
			users[0],
		},
		Children: []model.Department{
			{
				BaseModel: model.BaseModel{Id: 2},
				Name:      "访客",
				ParentId:  1,
				Users: []model.User{
					users[1],
				},
				ManageUsers: []model.User{
					users[0],
					users[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 3},
				Name:      "研发中心",
				ParentId:  1,
				Children: []model.Department{
					{
						BaseModel: model.BaseModel{Id: 4},
						Name:      "开发部",
						ParentId:  3,
					},
					{
						BaseModel: model.BaseModel{Id: 5},
						Name:      "测试部",
						ParentId:  3,
					},
					{
						BaseModel: model.BaseModel{Id: 6},
						Name:      "运维部",
						ParentId:  3,
						Users: []model.User{
							users[0],
						},
					},
				},
			},
		},
	},
}

// 递归插入数据方法
func insertDepartmentsData(departments []model.Department) {
	var department model.Department
	for _, item := range departments {
		// 查看数据是否存在，如果不存在才执行创建
		err := common.DB.Where("id = ? or name = ?", item.Id, item.Name).First(&department).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
		// 递归插入子部门
		if len(item.Children) > 0 {
			insertDepartmentsData(item.Children)
		}
	}
}

// 部门数据初始化
func DepartmentInit() {
	insertDepartmentsData(departments)
}
