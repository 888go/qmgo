package middleware

import (
	"context"
	"github.com/888go/qmgo/field"
	"github.com/888go/qmgo/hook"
	"github.com/888go/qmgo/operator"
	"github.com/888go/qmgo/validator"
)

// 定义回调函数类型 md5:d0532cbca1800c1a
type callback func(ctx context.Context, doc interface{}, opType mgo常量.OpType, opts ...interface{}) error

// middlewareCallback 是注册回调切片
// 一些回调在这里初始化，无需通过 Register() 方法
// md5:449512335518fc4e
var middlewareCallback = []callback{
	hook.Do,
	field.Do,
	validator.Do,
}

// Register 将回调函数注册到中间件中 md5:23bc8366f03c6dbb
func Register(回调函数 callback) {
	middlewareCallback = append(middlewareCallback, 回调函数)
}

// 始终调用每个注册
// 文档始终是操作文档
// md5:f300f2035d7e8114
func Do(ctx context.Context, content interface{}, opType mgo常量.OpType, opts ...interface{}) error {
	for _, cb := range middlewareCallback {
		if err := cb(ctx, content, opType, opts...); err != nil {
			return err
		}
	}
	return nil
}
