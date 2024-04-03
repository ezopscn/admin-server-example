package model

// 角色模型
type Role struct {
	BaseModel
	Name        string `gorm:"uniqueIndex:uidx_name;comment:角色名称" json:"name"`
	Description string `gorm:"comment:角色说明" json:"description"`
	Users       []User `gorm:"-" json:"users,omitempty"`                   // 用户
	Menus       []Menu `gorm:"many2many:role_menu" json:"menus,omitempty"` // 菜单和角色多对多
}

// 自定义表名
func (Role) TableName() string {
	return "role"
}
