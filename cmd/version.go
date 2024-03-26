package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"server/common"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// 版本信息
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看当前版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version is", common.Version)
		os.Exit(0)
	},
}
