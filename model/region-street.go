package model

// 街道模型
type Street struct {
	Id         uint     `gorm:"primaryKey;comment:id" json:"id"`
	Name       string   `gorm:"comment:街道名称" json:"name"`
	AreaId     uint     `gorm:"comment:区id" json:"area_id"`
	Area       Area     `gorm:"foreignKey:CityId;comment:区" json:"-"`
	ProvinceId uint     `gorm:"comment:省份id" json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceId;comment:省" json:"-"`
	CityId     uint     `gorm:"comment:市id" json:"city_id"`
	City       City     `gorm:"foreignKey:CityId;comment:市" json:"-"`
}

// 自定义表名
func (Street) TableName() string {
	return "street"
}
