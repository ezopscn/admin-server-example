package model

// 系统设置模型
type Setting struct {
	BaseModel
	Name        string `gorm:"uniqueIndex:uidx_name;comment:设置名称" json:"name"`
	Value       string `gorm:"comment:设置值" json:"value"`
	Description string `gorm:"comment:设置说明" json:"description"`
}

// 自定义表名
func (Setting) TableName() string {
	return "setting"
}
