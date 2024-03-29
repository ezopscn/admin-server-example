package common

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Logo 图形生成网站：http://patorjk.com/software/taag/
const Logo = `
 ▄▄▄▄▄▄▄▄▄▄▄  ▄         ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄        ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄       ▄ 
▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░▌      ▐░▌▐░░░░░░░░░░░▌▐░▌     ▐░▌
▐░█▀▀▀▀▀▀▀█░▌▐░▌       ▐░▌▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░▌░▌     ▐░▌ ▀▀▀▀█░█▀▀▀▀  ▐░▌   ▐░▌ 
▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌▐░▌          ▐░▌▐░▌    ▐░▌     ▐░▌       ▐░▌ ▐░▌  
▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░▌ ▐░▌   ▐░▌     ▐░▌        ▐░▐░▌   
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌  ▐░▌     ▐░▌         ▐░▌    
▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░▌       ▐░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░▌   ▐░▌ ▐░▌     ▐░▌        ▐░▌░▌   
▐░▌          ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌          ▐░▌    ▐░▌▐░▌     ▐░▌       ▐░▌ ▐░▌  
▐░▌          ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░▌     ▐░▐░▌ ▄▄▄▄█░█▄▄▄▄  ▐░▌   ▐░▌ 
▐░▌          ▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌      ▐░░▌▐░░░░░░░░░░░▌▐░▌     ▐░▌
 ▀            ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀        ▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀       ▀ 
                                                                                         `

// 全局工具
var (
	SystemLog *zap.SugaredLogger // 系统日志工具
	AccessLog *zap.SugaredLogger // 访问日志工具
	DB        *gorm.DB           // 数据库连接
	Cache     *redis.Client      // 缓存连接
)

// 系统信息
var (
	APPName               = "ADMIN-EXAMPLE"       // 应用名称
	APIPrefix             = "/api/v1"             // API 前缀
	Version               = "1.0"                 // 版本信息
	ConfigFile            = "config/default.yaml" // 默认配置文件
	VersionFile           = "config/version"      // 版本配置文件
	DefaultPageSize  uint = 1                     // 默认每页显示的数据量
	MaxPageSize      uint = 100                   // 每次请求最大的数据量，用于保证数据安全性
	SuperAdminRoleId uint = 1                     // 系统预留超级用户角色 Id
)

// 时间格式
const (
	MsecTimeFormat = "2006-01-02 15:04:05.000"
	SecTimeFormat  = "2006-01-02 15:04:05"
	DateTimeFormat = "2006-01-02"
)

// 常量定义
var (
	Male          uint = 1 // 性别男
	Female        uint = 2 // 性别女
	UnknownGender uint = 3 // 未知性别
	Disable       uint = 0 // 禁用
	Active        uint = 1 // 启用
	False         uint = 0 // 否
	True          uint = 1 // 是
)
