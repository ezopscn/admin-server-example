package model

// 城市模型
type City struct {
	Id         uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name       string   `gorm:"comment:市名称" json:"name"`
	ProvinceId uint     `gorm:"comment:省份id" json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceId;comment:省" json:"-"`
	Areas      []Area   `json:"areas,omitempty"`
	Streets    []Street `json:"streets,omitempty"`
}

// 自定义表名称
func (City) TableName() string {
	return "city"
}
