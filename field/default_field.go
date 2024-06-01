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
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DefaultFieldHook 定义了一个接口，用于通过钩子修改默认字段 md5:1e0917183e9bb23c
type DefaultFieldHook interface {
	DefaultUpdateAt()
	DefaultCreateAt()
	DefaultId()
}

// ```go
// 默认字段定义了在操作发生时要处理的默认字段
// 将DefaultField导入文档结构体使其生效
// ```
// md5:542fb0f78cfb4fad
// [提示]
//
//	type 默认字段 struct {
//	    ID       primitive.ObjectID `bson:"_id"`
//	    创建时间 time.Time          `bson:"createAt"`
//	    更新时间 time.Time          `bson:"updateAt"`
//	}
//
// [结束]
type DefaultField struct {
	Id       primitive.ObjectID `bson:"_id"`
	CreateAt time.Time          `bson:"createAt"`
	UpdateAt time.Time          `bson:"updateAt"`
}

// DefaultUpdateAt 更改默认的更新时间字段 md5:2aac31da652c649b
// ff:
// [提示:] func (df *DefaultField) 默认UpdateTime() {}
func (df *DefaultField) DefaultUpdateAt() {
	df.UpdateAt = time.Now().Local()
}

// DefaultCreateAt 更改默认的创建时间字段 md5:1438b66e329ae785
// ff:
// [提示:] func (df *DefaultField) 默认CreatedAt() {}
func (df *DefaultField) DefaultCreateAt() {
	if df.CreateAt.IsZero() {
		df.CreateAt = time.Now().Local()
	}
}

// DefaultId 修改默认的 _id 字段 md5:32bb6b194f03905a
// ff:
// [提示:] func (df *DefaultField) 默认Id() {}
func (df *DefaultField) DefaultId() {
	if df.Id.IsZero() {
		df.Id = primitive.NewObjectID()
	}
}
