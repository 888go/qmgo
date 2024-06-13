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

// DefaultFieldHook defines the interface to change default fields by hook
type DefaultFieldHook interface {
	DefaultUpdateAt()
	DefaultCreateAt()
	DefaultId()
}

// X默认字段名称 defines the default fields to handle when operation happens
// import the X默认字段名称 in document struct to make it working
type X默认字段名称 struct {
	Id       primitive.ObjectID `bson:"_id"`
	X创建时间 time.Time          `bson:"createAt"`
	X更新时间 time.Time          `bson:"updateAt"`
}

// X默认更新时间 changes the default updateAt field
func (df *X默认字段名称) X默认更新时间() {
	df.X更新时间 = time.Now().Local()
}

// X默认创建时间 changes the default createAt field
func (df *X默认字段名称) X默认创建时间() {
	if df.X创建时间.IsZero() {
		df.X创建时间 = time.Now().Local()
	}
}

// X默认ID changes the default _id field
func (df *X默认字段名称) X默认ID() {
	if df.Id.IsZero() {
		df.Id = primitive.NewObjectID()
	}
}
