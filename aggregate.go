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

package mgo类

import (
	"context"
	opts "github.com/888go/qmgo/options"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Pipeline 定义聚合操作的管道 md5:39821c5115607719
type Pipeline []bson.D

// Aggregate是一个聚合的句柄 md5:e06636d2fc45e004
type Aggregate struct {
	ctx        context.Context
	pipeline   interface{}
	collection *mongo.Collection
	options    []opts.AggregateOptions
}

// All 遍历聚合的游标，并将每个文档解码为结果。 md5:22b8eb7acebfa36a
func (a *Aggregate) X取全部(结果指针 interface{}) error {
	opts := options.Aggregate()
	if len(a.options) > 0 {
		opts = a.options[0].AggregateOptions
	}
	c, err := a.collection.Aggregate(a.ctx, a.pipeline, opts)
	if err != nil {
		return err
	}
	return c.All(a.ctx, 结果指针)
}

// One 从聚合结果中遍历游标，并将当前文档解码到结果中。 md5:95d05e20ff85babc
func (a *Aggregate) X取一条(结果指针 interface{}) error {
	opts := options.Aggregate()
	if len(a.options) > 0 {
		opts = a.options[0].AggregateOptions
	}
	c, err := a.collection.Aggregate(a.ctx, a.pipeline, opts)
	if err != nil {
		return err
	}
	cr := Cursor{
		ctx:    a.ctx,
		cursor: c,
		err:    err,
	}
	defer cr.X关闭()
	if !cr.X下一个(结果指针) {
		if err := cr.X取错误(); err != nil {
			return err
		}
		return X错误_未找到文档
	}
	return err
}

// Iter 返回聚合后的游标
// 已弃用，请使用Cursor
// md5:722184e644380849
func (a *Aggregate) Iter弃用() CursorI {
	return a.X取结果集()
}

// Cursor返回聚合后的游标 md5:eac4fdc1facaf217
func (a *Aggregate) X取结果集() CursorI {
	opts := options.Aggregate()
	if len(a.options) > 0 {
		opts = a.options[0].AggregateOptions
	}
	c, err := a.collection.Aggregate(a.ctx, a.pipeline, opts)
	return &Cursor{
		ctx:    a.ctx,
		cursor: c,
		err:    err,
	}
}
