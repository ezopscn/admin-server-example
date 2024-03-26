package model

// 数据来源：https://github.com/modood/Administrative-divisions-of-China
// 省份模型
type Province struct {
	Id      uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name    string   `gorm:"comment:省份名称" json:"name"`
	Cities  []City   `json:"cities,omitempty"`
	Areas   []Area   `json:"areas,omitempty"`
	Streets []Street `json:"streets,omitempty"`
}

// 自定义表名称
func (Province) TableName() string {
	return "province"
}
