package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"server/common"
	"server/initialize"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(tableCmd)
	migrateCmd.AddCommand(dataCmd)
}

// 迁移命令
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据初始化",
}

// 迁移表
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "数据库结构初始化",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(common.Logo)  // Logo
		initialize.Config()       // 配置初始化
		initialize.SystemLogger() // 系统日志初始化
		initialize.AccessLogger() // 访问日志初始化
		initialize.MySQL()        // 数据库连接初始化
		initialize.MigrateTable()
	},
}

// 迁移数据
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "基础数据初始化",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(common.Logo)  // Logo
		initialize.Config()       // 配置初始化
		initialize.SystemLogger() // 系统日志初始化
		initialize.AccessLogger() // 访问日志初始化
		initialize.MySQL()        // 数据库连接初始化
		initialize.MigrateData()
	},
}
