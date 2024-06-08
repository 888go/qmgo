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

package qmgo//bm:mgo类

import (
	"context"
	opts "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Pipeline 定义聚合操作的管道 md5:39821c5115607719
type Pipeline []bson.D

// Aggregate是一个聚合的句柄 md5:e06636d2fc45e004
// [提示]
//type 数据聚合 struct {
//     上下文        context.Context
//     管道         interface{}
//     集合         *mongo.Collection
//     聚合选项     []opts.AggregateOptions
// }
// [结束]
type Aggregate struct {
	ctx        context.Context
	pipeline   interface{}
	collection *mongo.Collection
	options    []opts.AggregateOptions
}

// All 遍历聚合的游标，并将每个文档解码为结果。 md5:22b8eb7acebfa36a
// [提示:] func (a *聚合操作) 全部结果(results interface{})
// ff:取全部
// a:
// results:结果指针
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

// One 从聚合结果中遍历游标，并将当前文档解码到结果中。 md5:95d05e20ff85babc
// [提示:] func (a *聚合操作) 单个结果(result interface{})
// ff:取一条
// a:
// result:结果指针
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
// 已弃用，请使用Cursor
// md5:722184e644380849
// [提示:] func (a *聚合操作) 迭代器() 游标接口 {}
// ff:Iter弃用
// a:
func (a *Aggregate) Iter() CursorI {
	return a.Cursor()
}

// Cursor返回聚合后的游标 md5:eac4fdc1facaf217
// [提示:] func (a *聚合操作) 获取游标() 游标接口 {}
// ff:取结果集
// a:
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
