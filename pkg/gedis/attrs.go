package gedis

import (
	"fmt"
	"time"
)

// 参数处理
type OperationAttr struct {
	Name  string
	Value interface{}
}

// 构造函数
func NewOperationAttr(name string, value interface{}) *OperationAttr {
	return &OperationAttr{Name: name, Value: value}
}

// 多个参数
type OperationAttrs []*OperationAttr

// 查找用法参数
func (attrs OperationAttrs) Find(name string) *InterfaceResult {
	for _, attr := range attrs {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("不支持该方法,", name))
}

// 设置过期时间
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  "expire",
		Value: t,
	}
}

// 设置 NX 锁，Key 不存在才能设置
func WithNX() *OperationAttr {
	return &OperationAttr{
		Name:  "nx",
		Value: struct{}{},
	}
}

// 设置 XX 锁，Key 存在才能设置
func WithXX() *OperationAttr {
	return &OperationAttr{
		Name:  "xx",
		Value: struct{}{},
	}
}
