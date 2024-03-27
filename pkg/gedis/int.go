package gedis

import "server/common"

// Int 类型结果封装
type IntResult struct {
	Result int
	Error  error
}

// 构造函数
func NewIntResult(result int, error error) *IntResult {
	return &IntResult{Result: result, Error: error}
}

// 解析结果
func (r *IntResult) Unwrap() int {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，", r.Error.Error())
	}
	return r.Result
}

// 查询失败返回默认值
func (r *IntResult) UnwrapWithDefaultValue(v int) int {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，返回默认值：", r.Error.Error())
		return v
	}
	return r.Result
}

// 查询失败执行函数
func (r *IntResult) UnwrapWithFunc(f func() int) int {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，返回执行函数：", r.Error.Error())
		return f()
	}
	return r.Result
}
