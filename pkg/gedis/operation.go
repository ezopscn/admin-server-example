package gedis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"server/common"
	"time"
)

// 用户操作
type Operation struct {
	Redis   *redis.Client
	Context context.Context
}

// 构造函数
func NewOperation() *Operation {
	return &Operation{Redis: common.Cache, Context: context.Background()}
}

// 获取单个值为 string 的 Key
func (o *Operation) GetString(key string) *StringResult {
	return NewStringResult(o.Redis.Get(o.Context, key).Result())
}

// 获取单个值为 int 的 Key
func (o *Operation) GetInt(key string) *IntResult {
	return NewIntResult(o.Redis.Get(o.Context, key).Int())
}

// 获取多个 Key 的值
func (o *Operation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(o.Redis.MGet(o.Context, keys...).Result())
}

// 删除 Key
func (o *Operation) Del(key string) (int64, error) {
	return o.Redis.Del(o.Context, key).Result()
}

// 设置 Key / Value
// 用法：gedis.Set("key", "value", gedis.WithExpire(time.Second * 10), gedis.WithNX())
func (o *Operation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
	// 参数列表
	oas := OperationAttrs(attrs)

	// 判断是否设置过期时间，没有则设置永不过期
	expire := oas.Find("expire").UnwrapWithDefaultValue(time.Second * 0).(time.Duration)

	// 判断是否 NX 锁，两种锁只能有一个
	nx := oas.Find("nx").UnwrapWithDefaultValue(nil)
	if nx != nil {
		return NewInterfaceResult(o.Redis.SetNX(o.Context, key, value, expire).Result())
	}

	// 判断是否 XX 锁
	xx := oas.Find("xx").UnwrapWithDefaultValue(nil)
	if xx != nil {
		return NewInterfaceResult(o.Redis.SetXX(o.Context, key, value, expire).Result())
	}

	// 默认
	return NewInterfaceResult(o.Redis.Set(o.Context, key, value, expire).Result())
}
