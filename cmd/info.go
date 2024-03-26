package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

// 开发者信息
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "查看服务开发者信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开发人员：Jayce <ezops.cn@gmail.com>.")
		os.Exit(0)
	},
}
