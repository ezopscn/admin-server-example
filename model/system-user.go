package model

import "github.com/golang-module/carbon/v2"

// 用户模型
type User struct {
	BaseModel
	Username               string        `gorm:"uniqueIndex:uidx_username;comment:用户名（工号）" json:"username"`
	ENName                 string        `gorm:"not null;comment:英文名（没有就用拼音）" json:"en_name"`
	CNName                 string        `gorm:"not null;comment:中文名" json:"cn_name"`
	JobId                  uint          `gorm:"comment:工作岗位id" json:"job_id"` // 关联工作岗位
	Job                    *Job          `gorm:"foreignKey:JobId;comment:工作岗位" json:"job,omitempty"`
	Phone                  string        `gorm:"uniqueIndex:uidx_phone;comment:手机号" json:"phone"`
	ShowPhone              *uint         `gorm:"type:tinyint(1);default:1;comment:是否显示手机号(0=不显示,1=显示)" json:"show_phone"`
	Email                  string        `gorm:"uniqueIndex:uidx_email;comment:邮箱" json:"email"`
	Password               string        `gorm:"not null;comment:密码" json:"-"` // json 中不显示该字段
	Secret                 string        `gorm:"comment:双因子认证密钥" json:"-"`     // json 中不显示该字段
	JoinTime               carbon.Carbon `gorm:"comment:入职日期" json:"join_time"`
	DepartmentId           uint          `gorm:"comment:部门id" json:"department_id"` // 关联部门
	Department             *Department   `gorm:"foreignKey:DepartmentId;comment:部门" json:"department,omitempty"`
	OfficeProvinceId       uint          `gorm:"comment:办公地点省id" json:"office_province_id"` // 关联省
	OfficeProvince         *Province     `gorm:"foreignKey:OfficeProvinceId;comment:省" json:"office_province,omitempty"`
	OfficeCityId           uint          `gorm:"comment:办公地点市id" json:"office_city_id"` // 关联市
	OfficeCity             *City         `gorm:"foreignKey:OfficeCityId;comment:市" json:"office_city,omitempty"`
	OfficeAddress          string        `gorm:"comment:办公地点详细地址" json:"office_address"`
	OfficeStation          string        `gorm:"comment:工位" json:"office_station"`
	NativeProvinceId       uint          `gorm:"comment:籍贯省id" json:"native_province_id"` // 关联省
	NativeProvince         *Province     `gorm:"foreignKey:NativeProvinceId;comment:省" json:"native_province,omitempty"`
	NativeCityId           uint          `gorm:"comment:籍贯市id" json:"native_city_id"` // 关联市
	NativeCity             *City         `gorm:"foreignKey:NativeCityId;comment:市" json:"native_city,omitempty"`
	Gender                 *uint         `gorm:"type:tinyint(1);default:1;comment:性别(1=男,2=女,3=未知)" json:"gender"`
	Avatar                 string        `gorm:"comment:头像" json:"avatar"`
	Birthday               carbon.Carbon `gorm:"comment:生日" json:"birthday"`
	CreatorId              uint          `gorm:"comment:创建人id" json:"creator_id"` // 关联用户自身
	Creator                *User         `gorm:"foreignKey:CreatorId;-" json:"creator,omitempty"`
	LastLoginIP            string        `gorm:"comment:最后一次登录IP" json:"last_login_ip"`
	LastLoginTime          carbon.Carbon `gorm:"comment:最后一次登录时间" json:"last_login_time"`
	LastChangePasswordTime carbon.Carbon `gorm:"comment:最后一次修改密码时间" json:"last_change_password_time"`
	Status                 *uint         `gorm:"type:tinyint(1);default:1;comment:用户状态(0=禁用,1=正常)" json:"status"`
	RoleId                 uint          `gorm:"comment:角色id" json:"role_id"` // 关联角色
	Role                   *Role         `gorm:"foreignKey:RoleId;comment:角色" json:"role,omitempty"`
}

// 自定义表名
func (User) TableName() string {
	return "user"
}
