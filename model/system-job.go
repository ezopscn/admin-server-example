package model

// 工作岗位模型
type Job struct {
	BaseModel        // 基础字段信息
	Name      string `gorm:"uniqueIndex:uidx_name;comment:岗位名称" json:"name"`
	Users     []User `gorm:"-" json:"users"`
}

// 自定义表名
func (Job) TableName() string {
	return "job"
}
