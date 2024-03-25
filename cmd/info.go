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
	Short: "Show service information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Service is developed by Jayce.")
		fmt.Println("You can contact me with email <ezops.cn@gmail.com>.")
		os.Exit(0)
	},
}
