package main

import (
	"embed"
	"io"
	"os"
	"server/cmd"
	"server/common"
)

//go:embed config/*
//go:embed initialize/data/sql/*
var fs embed.FS // 固定格式，打包的时候会将 config 和 initialize/data/sql 目录下面的文件都一起打包

func main() {
	// 设置全局使用
	common.FS = fs

	// 读取版本号
	f, err := os.Open(common.VersionFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 读取文件全部内容
	version, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// 设置全局版本号
	if string(version) != "" {
		common.Version = string(version)
	}

	// 入口
	cmd.Execute()
}
