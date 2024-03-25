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
	Short: "Show service version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git commit id is", common.Version)
		os.Exit(0)
	},
}
