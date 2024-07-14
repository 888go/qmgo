package validator

import (
	"context"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/qiniu/qmgo/operator"
)

// 使用单例的Validate，它缓存结构体信息 md5:37316caf6446b052
var validate = validator.New()

// SetValidate 允许使用自定义规则进行验证 md5:c45d0acce1bafd26
// ff:
// v:
func SetValidate(v *validator.Validate) {
	validate = v
}

// validatorNeeded 检查操作类型（opType）是否需要验证器 md5:69c24cea9b0cf3e4
func validatorNeeded(opType operator.OpType) bool {
	switch opType {
	case operator.BeforeInsert, operator.BeforeUpsert, operator.BeforeReplace:
		return true
	}
	return false
}

// Do 调用验证器检查
// 不要在這裡使用 opts
// md5:a3e02eb169c74704
// ff:
// ctx:
// doc:
// opType:
// opts:
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

// sliceHandle处理切片文档 md5:92800dd5899836ce
func sliceHandle(docs interface{}, opType operator.OpType) error {
	// []interface{}{UserType实例...} md5:bda81608072dd1ad
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

// 检查opType是否被支持，并调用fieldHandler方法 md5:3bb8cbff6cb4f5e3
func do(doc interface{}) error {
	if !validatorStruct(doc) {
		return nil
	}
	return validate.Struct(doc)
}

// validatorStruct 检查doc的类型是否为validator支持的结构体
// 实现方式与validator相同
// md5:566d3931e3bc9c80
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
