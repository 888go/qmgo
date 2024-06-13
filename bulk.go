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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// X批量操作结果 is the result type returned by Bulk.Run operation.
type X批量操作结果 struct {
	// The number of documents inserted.
	X插入数 int64

	// The number of documents matched by filters in update and replace operations.
	X匹配数 int64

	// The number of documents modified by update and replace operations.
	X修改数 int64

	// The number of documents deleted.
	X删除数 int64

	// The number of documents upserted by update and replace operations.
	X替换插入数 int64

	// A map of operation index to the _id of each upserted document.
	X替换插入IDs map[int64]interface{}
}

// X批量操作 is context for batching operations to be sent to database in a single
// bulk write.
//
// X批量操作 is not safe for concurrent use.
//
// Notes:
//
// Individual operations inside a bulk do not trigger middlewares or hooks
// at present.
//
// Different from original mgo, the qmgo implementation of X批量操作 does not emulate
// bulk operations individually on old versions of MongoDB servers that do not
// natively support bulk operations.
//
// Only operations supported by the official driver are exposed, that is why
// InsertMany is missing from the methods.
type X批量操作 struct {
	coll *X文档集合

	queue   []mongo.WriteModel
	ordered *bool
}

// X创建批量执行 returns a new context for preparing bulk execution of operations.
func (c *X文档集合) X创建批量执行() *X批量操作 {
	return &X批量操作{
		coll:    c,
		queue:   nil,
		ordered: nil,
	}
}

// X设置有序执行 marks the bulk as ordered or unordered.
//
// If ordered, writes does not continue after one individual write fails.
// Default is ordered.
func (b *X批量操作) X设置有序执行(开启有序 bool) *X批量操作 {
	b.ordered = &开启有序
	return b
}

// X插入 queues an X插入 operation for bulk execution.
func (b *X批量操作) X插入(待插入文档 interface{}) *X批量操作 {
	wm := mongo.NewInsertOneModel().SetDocument(待插入文档)
	b.queue = append(b.queue, wm)
	return b
}

// X删除一条 queues a X删除一条 operation for bulk execution.
func (b *X批量操作) X删除一条(删除条件 interface{}) *X批量操作 {
	wm := mongo.NewDeleteOneModel().SetFilter(删除条件)
	b.queue = append(b.queue, wm)
	return b
}

// X删除并按ID queues a X删除并按ID operation for bulk execution.
func (b *X批量操作) X删除并按ID(删除ID interface{}) *X批量操作 {
	b.X删除一条(bson.M{"_id": 删除ID})
	return b
}

// X删除 queues a X删除 operation for bulk execution.
func (b *X批量操作) X删除(删除条件 interface{}) *X批量操作 {
	wm := mongo.NewDeleteManyModel().SetFilter(删除条件)
	b.queue = append(b.queue, wm)
	return b
}

// X替换插入 queues an X替换插入 operation for bulk execution.
// The replacement should be document without operator
func (b *X批量操作) X替换插入(替换条件 interface{}, 替换内容 interface{}) *X批量操作 {
	wm := mongo.NewReplaceOneModel().SetFilter(替换条件).SetReplacement(替换内容).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// X替换插入一条 queues an X替换插入一条 operation for bulk execution.
// The update should contain operator
func (b *X批量操作) X替换插入一条(替换条件 interface{}, 替换内容 interface{}) *X批量操作 {
	wm := mongo.NewUpdateOneModel().SetFilter(替换条件).SetUpdate(替换内容).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// X替换插入并按ID queues an X替换插入并按ID operation for bulk execution.
// The replacement should be document without operator
func (b *X批量操作) X替换插入并按ID(替换ID interface{}, 替换内容 interface{}) *X批量操作 {
	b.X替换插入(bson.M{"_id": 替换ID}, 替换内容)
	return b
}

// X更新一条 queues an X更新一条 operation for bulk execution.
// The update should contain operator
func (b *X批量操作) X更新一条(更新条件 interface{}, 更新内容 interface{}) *X批量操作 {
	wm := mongo.NewUpdateOneModel().SetFilter(更新条件).SetUpdate(更新内容)
	b.queue = append(b.queue, wm)
	return b
}

// X更新并按ID queues an X更新并按ID operation for bulk execution.
// The update should contain operator
func (b *X批量操作) X更新并按ID(更新ID interface{}, 更新内容 interface{}) *X批量操作 {
	b.X更新一条(bson.M{"_id": 更新ID}, 更新内容)
	return b
}

// X更新 queues an X更新 operation for bulk execution.
// The update should contain operator
func (b *X批量操作) X更新(更新条件 interface{}, 更新内容 interface{}) *X批量操作 {
	wm := mongo.NewUpdateManyModel().SetFilter(更新条件).SetUpdate(更新内容)
	b.queue = append(b.queue, wm)
	return b
}

// X执行 executes the collected operations in a single bulk operation.
//
// A successful call resets the Bulk. If an error is returned, the internal
// queue of operations is unchanged, containing both successful and failed
// operations.
func (b *X批量操作) X执行(上下文 context.Context) (*X批量操作结果, error) {
	opts := options.BulkWriteOptions{
		Ordered: b.ordered,
	}
	result, err := b.coll.collection.BulkWrite(上下文, b.queue, &opts)
	if err != nil {
		// In original mgo, queue is not reset in case of error.
		return nil, err
	}

	// Empty the queue for possible reuse, as per mgo's behavior.
	b.queue = nil

	return &X批量操作结果{
		X插入数: result.InsertedCount,
		X匹配数:  result.MatchedCount,
		X修改数: result.ModifiedCount,
		X删除数:  result.DeletedCount,
		X替换插入数: result.UpsertedCount,
		X替换插入IDs:   result.UpsertedIDs,
	}, nil
}
