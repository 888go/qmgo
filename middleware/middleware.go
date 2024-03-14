package middleware

import (
	"context"
	"github.com/888go/qmgo/field"
	"github.com/888go/qmgo/hook"
	"github.com/888go/qmgo/operator"
	"github.com/888go/qmgo/validator"
)

// callback 定义回调函数类型
type callback func(ctx context.Context, doc interface{}, opType 操作符.OpType, opts ...interface{}) error

// middlewareCallback 注册回调函数切片
// 为保证执行顺序，这里初始化了一些无需通过 Register() 注册的回调函数
var middlewareCallback = []callback{
	hook.Do,
	field.Do,
	validator.Do,
}

// Register 将回调函数注册到中间件
func Register(cb callback) {
	middlewareCallback = append(middlewareCallback, cb)
}

// 对所有已注册的项进行调用
// doc 操作时始终是操作文档
func Do(ctx context.Context, content interface{}, opType 操作符.OpType, opts ...interface{}) error {
	for _, cb := range middlewareCallback {
		if err := cb(ctx, content, opType, opts...); err != nil {
			return err
		}
	}
	return nil
}
