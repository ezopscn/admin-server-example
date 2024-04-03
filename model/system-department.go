package model

// 部门模型
type Department struct {
	BaseModel
	Name       string       `gorm:"comment:部门名称" json:"name"`
	ParentId   uint         `gorm:"comment:父id" json:"parent_id"`
	ManageUser []User       `gorm:"many2many:department_manage_user" json:"manage_user,omitempty"` // 部门负责人
	Users      []User       `gorm:"many2many:department_user" json:"users,omitempty"`              // 部门用户
	Children   []Department `gorm:"-" json:"children,omitempty"`                                   // 子部门关联
}

// 自定义表名
func (Department) TableName() string {
	return "department"
}
