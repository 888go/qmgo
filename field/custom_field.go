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
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"time"
)

// CustomFields 定义了支持的自定义字段的结构体 md5:d6f9da51bb1f9550
// [提示]
//
//	type 自定义字段 struct {
//	    创建时间 string
//	    更新时间 string
//	    ID        string
//	}
//
// [结束]
type CustomFields struct {
	createAt string
	updateAt string
	id       string
}

// CustomFieldsHook 定义了接口，CustomFields 返回用户想要自定义的字段 md5:7c964cacd7711950
type CustomFieldsHook interface {
	CustomFields() CustomFieldsBuilder //qm:设置更新时间字段名
}

// CustomFieldsBuilder 定义了用户用来设置自定义字段的接口 md5:68d7427693ef0a18
type CustomFieldsBuilder interface {
	SetUpdateAt(fieldName string) CustomFieldsBuilder //qm:设置更新时间字段名
	SetCreateAt(fieldName string) CustomFieldsBuilder //qm:设置创建时间字段名
	SetId(fieldName string) CustomFieldsBuilder       //qm:设置ID字段名
}

// NewCustom 创建一个新的Builder，用于设置自定义字段 md5:f37c4ab7a682c81a
// ff:
// [提示:] func 新建自定义Fields构建器() 自定义Fields构建器 {}
func NewCustom() CustomFieldsBuilder {
	return &CustomFields{}
}

// SetUpdateAt 设置自定义的UpdateAt字段 md5:fa5a62704b166e17设置更新时间字段名设置更新时间字段名设置更新时间字段名设置更新时间字段名设置更新时间字段名
// ff:设置更新时间字段名
// fieldName:字段名称
// [提示:] func (c *自定义字段) 设置更新时间字段(fieldName 字符串) 自定义字段构建器 {}
func (c *CustomFields) SetUpdateAt(fieldName string) CustomFieldsBuilder {
	c.updateAt = fieldName
	return c
}

// SetCreateAt 设置自定义的创建时间字段 md5:9d021ddb5b3276d1设置创建时间字段名设置创建时间字段名设置创建时间字段名设置创建时间字段名设置创建时间字段名
// ff:设置创建时间字段名
// fieldName:字段名称
// [提示:] func (c *自定义字段) 设置创建时间(fieldName string) 自定义字段构建器 {}
func (c *CustomFields) SetCreateAt(fieldName string) CustomFieldsBuilder {
	c.createAt = fieldName
	return c
}

// SetId 设置自定义Id字段 md5:769568d47e77f5fe设置ID字段名设置ID字段名设置ID字段名设置ID字段名设置ID字段名
// ff:设置ID字段名
// fieldName:字段名称
// [提示:] func (c *自定义字段) 设置唯一标识(fieldName 字符串) 自定义字段构建器 {}
func (c *CustomFields) SetId(fieldName string) CustomFieldsBuilder {
	c.id = fieldName
	return c
}

// CustomCreateTime 更改自定义创建时间 md5:0ddb85c54f1b54c2自定义创建时间自定义创建时间自定义创建时间自定义创建时间自定义创建时间
// ff:自定义创建时间
// doc:待插入文档
// [提示:] func (c 自定义字段) 创建时间定制(doc interface{})
func (c CustomFields) CustomCreateTime(doc interface{}) {
	if c.createAt == "" {
		return
	}
	fieldName := c.createAt
	setTime(doc, fieldName, false)
	return
}

// CustomUpdateTime 修改自定义更新时间 md5:71153ba9520a179c自定义更新时间自定义更新时间自定义更新时间自定义更新时间自定义更新时间
// ff:自定义更新时间
// doc:待插入文档
// [提示:] func (c 自定义字段) 自定义更新时间(doc interface{})
func (c CustomFields) CustomUpdateTime(doc interface{}) {
	if c.updateAt == "" {
		return
	}
	fieldName := c.updateAt
	setTime(doc, fieldName, true)
	return
}

// CustomUpdateTime 修改自定义更新时间 md5:71153ba9520a179c自定义ID自定义ID自定义ID自定义ID自定义ID
// ff:自定义ID
// doc:待插入文档
// [提示:] func (c 自定义字段) 自定义Id(文档 interface{})
func (c CustomFields) CustomId(doc interface{}) {
	if c.id == "" {
		return
	}
	fieldName := c.id
	setId(doc, fieldName)
	return
}

// setTime 修改自定义时间字段
// overWrite 参数定义了当字段已有有效值时是否覆盖原有值
// md5:848f135d11ab4cf8
func setTime(doc interface{}, fieldName string, overWrite bool) {
	if reflect.Ptr != reflect.TypeOf(doc).Kind() {
		fmt.Println("not a point type")
		return
	}
	e := reflect.ValueOf(doc).Elem()
	ca := e.FieldByName(fieldName)
	if ca.CanSet() {
		tt := time.Now()
		switch a := ca.Interface().(type) {
		case time.Time:
			if ca.Interface().(time.Time).IsZero() {
				ca.Set(reflect.ValueOf(tt))
			} else if overWrite {
				ca.Set(reflect.ValueOf(tt))
			}
		case int64:
			if ca.Interface().(int64) == 0 {
				ca.SetInt(tt.Unix())
			} else if overWrite {
				ca.SetInt(tt.Unix())
			}
		default:
			fmt.Println("unsupported type to setTime", a)
		}
	}
}

// setId 修改自定义标识字段 md5:ef4585873d988b50
func setId(doc interface{}, fieldName string) {
	if reflect.Ptr != reflect.TypeOf(doc).Kind() {
		fmt.Println("not a point type")
		return
	}
	e := reflect.ValueOf(doc).Elem()
	ca := e.FieldByName(fieldName)
	if ca.CanSet() {
		switch a := ca.Interface().(type) {
		case primitive.ObjectID:
			if ca.Interface().(primitive.ObjectID).IsZero() {
				ca.Set(reflect.ValueOf(primitive.NewObjectID()))
			}
		case string:
			if ca.String() == "" {
				ca.SetString(primitive.NewObjectID().Hex())
			}
		default:
			fmt.Println("unsupported type to setId", a)
		}
	}
}
