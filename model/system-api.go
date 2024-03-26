package model

// API 类型模型
type APIType struct {
	Id   uint   `gorm:"primaryKey;comment:id" json:"id"`
	Name string `gorm:"uniqueIndex:uidx_name;comment:接口名称" json:"name"`
	APIS []API  `json:"apis,omitempty"`
}

// 自定义表名称
func (APIType) TableName() string {
	return "api_type"
}

// 接口模型
type API struct {
	BaseModel
	API         string  `gorm:"uniqueIndex:uidx_api;comment:接口URI" json:"api"`
	Method      string  `gorm:"comment:请求方法" json:"method"`
	Name        string  `gorm:"uniqueIndex:uidx_name;comment:接口名称" json:"name"`
	Description string  `gorm:"comment:接口说明" json:"description"`
	APITypeId   uint    `gorm:"comment:接口类型id" json:"api_type_id"`
	APIType     APIType `gorm:"foreignKey:APITypeId;comment:接口类型" json:"-"`
}

// 自定义表名称
func (API) TableName() string {
	return "api"
}
