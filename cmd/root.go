package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// 命令总入口
var rootCmd = &cobra.Command{
	Use:   "admin-server",
	Short: "ADMIN-SERVER 是一个由 Golang 开发的 RBAC 权限管理系统后端.",
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

// 所有子命令添加到 root 命令，输入 cmd 的入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
