package common

import (
	"fmt"
	"time"
)

// 分隔符
const RDSKeySeparator = ":"

// Redis Key Prefix
type RDSKeyPrefix struct {
	LoginToken      string // 用户登录 Token 前缀
	LoginWrongTimes string // 用户登录错误次数
	UserInfo        string // 用户信息
}

// 配置 Redis Key Prefix
var RKP = RDSKeyPrefix{
	LoginToken:      "LOGIN-TOKEN",
	LoginWrongTimes: "LOGIN-WRONG-TIMES",
	UserInfo:        "USER-INFO",
}

// 生成 Key
func GenerateRedisKey(keyPrefix string, keyTag string) string {
	return fmt.Sprintf("%s%s%s", keyPrefix, RDSKeySeparator, keyTag)
}

// Redis Key 超时时间
var (
	UserInfoExpireTime = time.Second * 60 // 用户信息缓存默认有效期 60 秒
)
