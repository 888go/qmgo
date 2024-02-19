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

// DefaultFieldHook 定义了一个接口，用于通过钩子方式修改默认字段
type DefaultFieldHook interface {
	DefaultUpdateAt()
	DefaultCreateAt()
	DefaultId()
}

// DefaultField 定义了在操作发生时默认处理的字段
// 在文档结构体中导入 DefaultField 以使其生效
type DefaultField struct {
	Id       primitive.ObjectID `bson:"_id"`
	CreateAt time.Time          `bson:"createAt"`
	UpdateAt time.Time          `bson:"updateAt"`
}

// DefaultUpdateAt 修改默认的 updateAt 字段
func (df *DefaultField) DefaultUpdateAt() {
	df.UpdateAt = time.Now().Local()
}

// DefaultCreateAt 修改默认的 createAt 字段
func (df *DefaultField) DefaultCreateAt() {
	if df.CreateAt.IsZero() {
		df.CreateAt = time.Now().Local()
	}
}

// DefaultId 更改默认的 _id 字段
func (df *DefaultField) DefaultId() {
	if df.Id.IsZero() {
		df.Id = primitive.NewObjectID()
	}
}
