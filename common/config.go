package common

import "embed"

// 配置打包
var FS embed.FS

// 配置引用
var Config Configuration

// 配置结构体
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
}

// 系统配置
type SystemConfiguration struct {
	Listen string `mapstructure:"listen" json:"listen"`
	Port   string `mapstructure:"port" json:"port"`
}
