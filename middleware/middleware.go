package middleware

import (
	"context"
	"github.com/888go/qmgo/field"
	"github.com/888go/qmgo/hook"
	"github.com/888go/qmgo/operator"
	"github.com/888go/qmgo/validator"
)

// callback define the callback function type
type callback func(ctx context.Context, doc interface{}, opType mgo常量.OpType, opts ...interface{}) error

// middlewareCallback the register callback slice
// some callbacks initial here without Register() for order
var middlewareCallback = []callback{
	hook.Do,
	field.Do,
	validator.Do,
}

// Register register callback into middleware
func Register(回调函数 callback) {
	middlewareCallback = append(middlewareCallback, 回调函数)
}

// Do call every registers
// The doc is always the document to operate
func Do(ctx context.Context, content interface{}, opType mgo常量.OpType, opts ...interface{}) error {
	for _, cb := range middlewareCallback {
		if err := cb(ctx, content, opType, opts...); err != nil {
			return err
		}
	}
	return nil
}
