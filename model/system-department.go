package model

// 部门模型
type Department struct {
	BaseModel
	Name     string       `gorm:"comment:部门名称" json:"name"`
	LeaderId uint         `gorm:"comment:负责人id" json:"leader_id"`
	ParentId uint         `gorm:"comment:父id" json:"parent_id"`
	Users    []uint       `gorm:"-" json:"users,omitempty"`    // 用户
	Children []Department `gorm:"-" json:"children,omitempty"` // 子部门关联
}

// 自定义表名
func (Department) TableName() string {
	return "department"
}
