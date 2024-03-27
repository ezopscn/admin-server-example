package gedis

import "server/common"

// Slice 类型结果封装
type SliceResult struct {
	Result []interface{}
	Error  error
}

// 构造函数
func NewSliceResult(result []interface{}, error error) *SliceResult {
	return &SliceResult{Result: result, Error: error}
}

// 解析结果
func (r *SliceResult) Unwrap() []interface{} {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，", r.Error.Error())
	}
	return r.Result
}

// 查询失败返回默认值
func (r *SliceResult) UnwrapWithDefaultValue(v []interface{}) []interface{} {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，返回默认值：", r.Error.Error())
		return v
	}
	return r.Result
}

// 查询失败执行函数
func (r *SliceResult) UnwrapWithFunc(f func() []interface{}) []interface{} {
	if r.Error != nil {
		common.SystemLog.Debug("Redis 缓存查询失败，返回执行函数：", r.Error.Error())
		return f()
	}
	return r.Result
}

// 使用迭代器获取值
//
//	func demo() {
//	    var conn = NewOperation()
//	    var res = conn.Mget("name", "age", "gender").Iter()
//	    for res.HasNext() {
//	        fmt.Println(res.Next())
//	    }
//	}
func (r *SliceResult) Iter() *Iterator {
	return NewIterator(r.Result)
}
