package validator

import (
	"context"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/qiniu/qmgo/operator"
)

// 使用单一实例的Validate，它会缓存结构体信息
var validate = validator.New()

// SetValidate 设置验证器，使其可以使用自定义规则
func SetValidate(v *validator.Validate) {
	validate = v
}

// validatorNeeded 检查对于 opType 是否需要验证器
func validatorNeeded(opType operator.OpType) bool {
	switch opType {
	case operator.BeforeInsert, operator.BeforeUpsert, operator.BeforeReplace:
		return true
	}
	return false
}

// Do 调用验证器进行检查
// 在此处不要使用 opts
func Do(ctx context.Context, doc interface{}, opType operator.OpType, opts ...interface{}) error {
	if !validatorNeeded(opType) {
		return nil
	}
	to := reflect.TypeOf(doc)
	if to == nil {
		return nil
	}
	switch reflect.TypeOf(doc).Kind() {
	case reflect.Slice:
		return sliceHandle(doc, opType)
	case reflect.Ptr:
		v := reflect.ValueOf(doc).Elem()
		switch v.Kind() {
		case reflect.Slice:
			return sliceHandle(v.Interface(), opType)
		default:
			return do(doc)
		}
	default:
		return do(doc)
	}
}

// sliceHandle 处理切片文档
func sliceHandle(docs interface{}, opType operator.OpType) error {
	// []interface{}{UserType{}...} 
// 创建一个接口类型切片，其中包含零个或多个UserType结构体实例。这里的"..."表示可变数量的参数，表示可以传入任意数量的UserType实例到切片中。
	if h, ok := docs.([]interface{}); ok {
		for _, v := range h {
			if err := do(v); err != nil {
				return err
			}
		}
		return nil
	}
	// []UserType{}
	s := reflect.ValueOf(docs)
	for i := 0; i < s.Len(); i++ {
		if err := do(s.Index(i).Interface()); err != nil {

			return err
		}
	}
	return nil
}

// 检查opType是否支持，并调用fieldHandler
func do(doc interface{}) error {
	if !validatorStruct(doc) {
		return nil
	}
	return validate.Struct(doc)
}

// validatorStruct 检查doc的类型是否为validator支持的结构体
// 实现方式与validator相同
func validatorStruct(doc interface{}) bool {
	val := reflect.ValueOf(doc)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct || val.Type() == reflect.TypeOf(time.Time{}) {
		return false
	}
	return true
}
