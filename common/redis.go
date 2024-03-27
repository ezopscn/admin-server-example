package common

import "fmt"

// 分隔符
const RDSKeySeparator = ":"

// Redis Key Prefix
type RDSKeyPrefix struct {
	LoginToken      string // 用户登录 Token 前缀
	LoginWrongTimes string // 用户登录错误次数
}

// 配置 Redis Key Prefix
var RKP = RDSKeyPrefix{
	LoginToken:      "LOGIN-TOKEN",
	LoginWrongTimes: "LOGIN-WRONG-TIMES",
}

// 生成 Key
func GenerateRedisKey(keyPrefix string, keyTag string) string {
	return fmt.Sprintf("%s%s%s", keyPrefix, RDSKeySeparator, keyTag)
}
