/*
 Copyright 2020 The Qmgo Authors.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package field

import (
	"context"
	"reflect"
	"time"
	
	"github.com/888go/qmgo/operator"
)

var nilTime time.Time

// filedHandler 定义了字段类型与处理器之间的关系
var fieldHandler = map[operator.OpType]func(doc interface{}) error{
	operator.BeforeInsert:  beforeInsert,
	operator.BeforeUpdate:  beforeUpdate,
	operator.BeforeReplace: beforeUpdate,
	operator.BeforeUpsert:  beforeUpsert,
}

// 注册函数
//func init() {
//	// 将Do函数注册到中间件
//	middleware.Register(Do)
//}

// 根据fType调用特定方法来处理字段
// 在此处不要使用opts
func Do(ctx context.Context, doc interface{}, opType operator.OpType, opts ...interface{}) error {
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
			return do(doc, opType)
		}
	}
	// 输出："不支持的类型"
	return nil
}

// sliceHandle 处理切片文档
func sliceHandle(docs interface{}, opType operator.OpType) error {
	// []interface{}{UserType{}...} 
// 创建一个接口类型切片，其中包含零个或多个UserType结构体实例。这里的"..."表示可变数量的参数，表示可以传入任意数量的UserType实例到切片中。
	if h, ok := docs.([]interface{}); ok {
		for _, v := range h {
			if err := do(v, opType); err != nil {
				return err
			}
		}
		return nil
	}
	// []UserType{}
	s := reflect.ValueOf(docs)
	for i := 0; i < s.Len(); i++ {
		if err := do(s.Index(i).Interface(), opType); err != nil {
			return err
		}
	}
	return nil
}

// beforeInsert 在插入前处理字段
// 如果doc中createAt字段的值有效，upsert不会改变它
// 如果doc中id字段的值有效，upsert不会改变它
// 无论如何都会更新updateAt字段的值
func beforeInsert(doc interface{}) error {
	if ih, ok := doc.(DefaultFieldHook); ok {
		ih.DefaultId()
		ih.DefaultCreateAt()
		ih.DefaultUpdateAt()
	}
	if ih, ok := doc.(CustomFieldsHook); ok {
		fields := ih.CustomFields()
		fields.(*CustomFields).CustomId(doc)
		fields.(*CustomFields).CustomCreateTime(doc)
		fields.(*CustomFields).CustomUpdateTime(doc)
	}
	return nil
}

// beforeUpdate 在更新字段前进行处理
func beforeUpdate(doc interface{}) error {
	if ih, ok := doc.(DefaultFieldHook); ok {
		ih.DefaultUpdateAt()
	}
	if ih, ok := doc.(CustomFieldsHook); ok {
		fields := ih.CustomFields()
		fields.(*CustomFields).CustomUpdateTime(doc)
	}
	return nil
}

// beforeUpsert 在执行upsert操作前处理字段
// 如果doc中createAt字段的值有效，upsert操作不会改变它
// 如果doc中id字段的值有效，upsert操作也不会改变它
// 无论如何都会更新updateAt字段的值
func beforeUpsert(doc interface{}) error {
	if ih, ok := doc.(DefaultFieldHook); ok {
		ih.DefaultId()
		ih.DefaultCreateAt()
		ih.DefaultUpdateAt()
	}
	if ih, ok := doc.(CustomFieldsHook); ok {
		fields := ih.CustomFields()
		fields.(*CustomFields).CustomId(doc)
		fields.(*CustomFields).CustomCreateTime(doc)
		fields.(*CustomFields).CustomUpdateTime(doc)
	}
	return nil
}

// 检查opType是否支持，并调用fieldHandler
func do(doc interface{}, opType operator.OpType) error {
	if f, ok := fieldHandler[opType]; !ok {
		return nil
	} else {
		return f(doc)
	}
}
