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

	"github.com/qiniu/qmgo/operator"
)

var nilTime time.Time

// filedHandler 定义字段类型和处理器之间的关系 md5:c7cd659bd6a053b2
var fieldHandler = map[operator.OpType]func(doc interface{}) error{
	operator.BeforeInsert:  beforeInsert,
	operator.BeforeUpdate:  beforeUpdate,
	operator.BeforeReplace: beforeUpdate,
	operator.BeforeUpsert:  beforeUpsert,
}

// 函数 init() {
// 注册 middleware，参数为 Do
//}
// md5:4bdefdddb5ec33c1

// Do 调用特定方法根据 fType 处理字段
// 不在这里使用 opts
// md5:01967b5b64a19adb
// [提示]
//// 执行上下文，文档接口
// func ExecuteInContext(操作上下文 context.Context, 文档 interface{}) 
// [结束]
// ff:
// ctx:
// doc:
// opType:
// opts:
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
	//fmt.Println("不支持此类类型") md5:2ba1fad322480d74
	return nil
}

// sliceHandle处理切片文档 md5:92800dd5899836ce
func sliceHandle(docs interface{}, opType operator.OpType) error {
	// []interface{}{UserType实例...} md5:bda81608072dd1ad
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
// 如果文档中的createAt字段的值有效，upsert 不会改变它
// 如果文档中的id字段的值有效，upsert 不会改变它
// 无论如何，改变updateAt字段的值
// md5:f49d81597c8212f6
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

// beforeUpdate处理更新前的字段 md5:a783a1aa99fba490
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

// beforeUpsert 处理字段的before upsert操作
// 如果doc中field createAt的值有效，upsert操作不会改变它
// 如果doc中field id的值有效，upsert操作也不会改变它
// 无论如何都会更新field updateAt的值
// md5:d286cfb6c0a1f1da
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

// 检查opType是否被支持，并调用fieldHandler方法 md5:3bb8cbff6cb4f5e3
func do(doc interface{}, opType operator.OpType) error {
	if f, ok := fieldHandler[opType]; !ok {
		return nil
	} else {
		return f(doc)
	}
}
