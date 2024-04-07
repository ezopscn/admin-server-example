package model

// 菜单模型
type Menu struct {
	BaseModel        // 基础字段信息
	Label     string `gorm:"uniqueIndex:uidx_label;comment:菜单名称" json:"label"`
	Icon      string `gorm:"comment:菜单图标" json:"icon"`
	Key       string `gorm:"uniqueIndex:uidx_key;comment:菜单路径" json:"key"`
	Sort      uint   `gorm:"comment:排序" json:"sort"`
	ParentId  uint   `gorm:"comment:父id" json:"parent_id"`
	Children  []Menu `gorm:"-" json:"children"`
	Roles     []Role `gorm:"many2many:role_menu" json:"roles,omitempty"` // 菜单和角色多对多
	RoleIds   []uint `gorm:"-" json:"role_ids"`
}

// 自定义表名
func (Menu) TableName() string {
	return "menu"
}
