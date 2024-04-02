package initialize

import "server/initialize/data"

// 数据初始化
func MigrateData() {
	data.RegionInit()     // 省市区街道数据
	data.DepartmentInit() // 部门数据
	data.JobInit()        // 工作岗位
	data.RoleInit()       // 角色数据
	data.MenuInit()       // 菜单数据
	data.UserInit()       // 用户数据
	data.APITypeInit()    // 接口类型数据
	data.SettingInit()    // 系统设置数据
}
