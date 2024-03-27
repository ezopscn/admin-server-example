package data

import (
	"fmt"
	"server/common"
	"strings"
)

// 执行 SQL 文件
func ImportSQLFile(filename string, table string, truncate bool) {
	common.SystemLog.Info("开始同步 SQL 文件：", filename)
	bs, err := common.FS.ReadFile(filename)
	if err != nil {
		common.SystemLog.Error(err.Error())
		panic(err)
	}
	// 判断是否需要提前清空表
	if truncate {
		common.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", table))
	}
	// 导入数据
	sqls := strings.Split(string(bs), ";")
	for _, sql := range sqls {
		if len(sql) > 5 { // 解决空行输出的问题
			common.DB.Table(table).Exec(sql)
		}
	}
}

// 初始化省市区街道数据
func RegionInit() {
	ImportSQLFile("initialize/data/sql/province.sql", "province", true)
	ImportSQLFile("initialize/data/sql/city.sql", "city", true)
	ImportSQLFile("initialize/data/sql/area.sql", "area", true)
	ImportSQLFile("initialize/data/sql/street.sql", "street", true)
}
