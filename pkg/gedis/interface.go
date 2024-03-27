package gedis

import "server/common"

// Interface 类型结果封装
type InterfaceResult struct {
	Result interface{}
	Error  error
}

// 构造函数
func NewInterfaceResult(result interface{}, error error) *InterfaceResult {
	return &InterfaceResult{Result: result, Error: error}
}

// 解析结果
func (r *InterfaceResult) Unwrap() interface{} {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，", r.Error.Error())
	}
	return r.Result
}

// 查询失败返回默认值
func (r *InterfaceResult) UnwrapWithDefaultValue(v interface{}) interface{} {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，返回默认值：", r.Error.Error())
		return v
	}
	return r.Result
}

// 查询失败执行函数
func (r *InterfaceResult) UnwrapWithFunc(f func() interface{}) interface{} {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，返回执行函数：", r.Error.Error())
		return f()
	}
	return r.Result
}
