package model

// 行政区模型
type Area struct {
	Id         uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name       string   `gorm:"comment:区名称" json:"name"`
	CityId     uint     `gorm:"comment:市id" json:"city_id"`
	City       City     `gorm:"foreignKey:CityId;comment:市" json:"-"`
	ProvinceId uint     `gorm:"comment:省份id" json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceId;comment:省" json:"-"`
	Streets    []Street `json:"streets,omitempty"`
}

// 自定义表名
func (Area) TableName() string {
	return "area"
}
