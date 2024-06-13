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

// X自定义字段 defines struct of supported custom fields
type X自定义字段 struct {
	createAt string
	updateAt string
	id       string
}

// CustomFieldsHook defines the interface, CustomFields return custom field user want to change
type CustomFieldsHook interface {
	X设置更新时间字段名() CustomFieldsBuilder
}

// CustomFieldsBuilder defines the interface which user use to set custom fields
type CustomFieldsBuilder interface {
	X设置更新时间字段名(fieldName string) CustomFieldsBuilder
	X设置创建时间字段名(fieldName string) CustomFieldsBuilder
	X设置ID字段名(fieldName string) CustomFieldsBuilder
}

// NewCustom creates new Builder which is used to set the custom fields
func NewCustom() CustomFieldsBuilder {
	return &X自定义字段{}
}

// X设置更新时间字段名 set the custom UpdateAt field
func (c *X自定义字段) X设置更新时间字段名(字段名称 string) CustomFieldsBuilder {
	c.updateAt = 字段名称
	return c
}

// X设置创建时间字段名 set the custom CreateAt field
func (c *X自定义字段) X设置创建时间字段名(字段名称 string) CustomFieldsBuilder {
	c.createAt = 字段名称
	return c
}

// X设置ID字段名 set the custom Id field
func (c *X自定义字段) X设置ID字段名(字段名称 string) CustomFieldsBuilder {
	c.id = 字段名称
	return c
}

// X自定义创建时间 changes the custom create time
func (c X自定义字段) X自定义创建时间(待插入文档 interface{}) {
	if c.createAt == "" {
		return
	}
	fieldName := c.createAt
	setTime(待插入文档, fieldName, false)
	return
}

// X自定义更新时间 changes the custom update time
func (c X自定义字段) X自定义更新时间(待插入文档 interface{}) {
	if c.updateAt == "" {
		return
	}
	fieldName := c.updateAt
	setTime(待插入文档, fieldName, true)
	return
}

// CustomUpdateTime changes the custom update time
func (c X自定义字段) X自定义ID(待插入文档 interface{}) {
	if c.id == "" {
		return
	}
	fieldName := c.id
	setId(待插入文档, fieldName)
	return
}

// setTime changes the custom time fields
// The overWrite defines if change value when the filed has valid value
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

// setId changes the custom Id fields
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
