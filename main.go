package main

import (
	"embed"
	"server/cmd"
)

//go:embed config/*
var fs embed.FS // 固定格式，打包的时候会将 config 目录下面的文件都一起打包

func main() {
	cmd.Execute() // 入口
}
