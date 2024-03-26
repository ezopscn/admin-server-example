package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"server/common"
	"server/initialize"
	"server/pkg/utils"
	"time"
)

func init() {
	rootCmd.AddCommand(startCmd)
	// 指定配置文件参数
	startCmd.Flags().StringVarP(&common.ConfigFile, "config", "f", common.ConfigFile, "可选，指定服务启动配置文件")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "参数化启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(common.Logo)  // Logo
		initialize.Config()       // 配置初始化
		initialize.SystemLogger() // 系统日志初始化
		initialize.AccessLogger() // 访问日志初始化
		initialize.MySQL()        // 数据库连接初始化
		r := initialize.Router()  // 路由初始化

		// 判断参数是否合法
		if !utils.IsIPAddress(common.Config.System.Listen) {
			common.SystemLog.Error("系统配置监听地址不合法")
			return
		}

		// 检测端口是否合法
		if !utils.IsPort(common.Config.System.Port) {
			common.SystemLog.Error("系统配置监听端口不合法")
			return
		}

		// 监听地址
		listenAddress := fmt.Sprintf("%s:%s", common.Config.System.Listen, common.Config.System.Port)
		common.SystemLog.Info("服务启动监听的地址：", listenAddress)

		// 配置服务
		server := http.Server{
			Addr:    listenAddress,
			Handler: r,
		}

		// 启动服务
		go func() {
			err := server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				common.SystemLog.Error(err.Error())
				panic(err)
			}
		}()

		// 接收优雅关闭信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		// 等待5秒然后停止服务
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			common.SystemLog.Error(err.Error())
			panic(err)
		}
		common.SystemLog.Info("服务正常的停止完成")
	},
}
