package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"server/common"
)

func init() {
	rootCmd.AddCommand(startCmd)
	// 指定配置文件参数
	startCmd.Flags().StringVarP(&common.ConfigFile, "config", "f", common.ConfigFile, "specify run config for service")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service with some flags",
	Run: func(cmd *cobra.Command, args []string) {
		// Logo
		fmt.Println(common.Logo)
	},
}
