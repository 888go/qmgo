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

package qmgo

import (
	"context"
	opts "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Pipeline 定义了聚合操作的管道
type Pipeline []bson.D

// Aggregate 是对聚合的处理句柄
type Aggregate struct {
	ctx        context.Context
	pipeline   interface{}
	collection *mongo.Collection
	options    []opts.AggregateOptions
}

// All 遍历从聚合中获取的游标，并将每个文档解码到结果中。
func (a *Aggregate) All(results interface{}) error {
	opts := options.Aggregate()
	if len(a.options) > 0 {
		opts = a.options[0].AggregateOptions
	}
	c, err := a.collection.Aggregate(a.ctx, a.pipeline, opts)
	if err != nil {
		return err
	}
	return c.All(a.ctx, results)
}

// One 通过从聚合中迭代游标，并将当前文档解码到结果中。
func (a *Aggregate) One(result interface{}) error {
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
	defer cr.Close()
	if !cr.Next(result) {
		if err := cr.Err(); err != nil {
			return err
		}
		return ErrNoSuchDocuments
	}
	return err
}

// Iter 返回聚合后的游标
// 已弃用，请使用 Cursor
func (a *Aggregate) Iter() CursorI {
	return a.Cursor()
}

// Cursor 返回聚合后的游标
func (a *Aggregate) Cursor() CursorI {
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
