package initialize

import "server/initialize/data"

// 数据初始化
func MigrateData() {
	data.RegionInit() // 省市区街道数据
}
