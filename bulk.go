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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BulkResult 是 Bulk.Run 操作返回的结果类型。
type BulkResult struct {
	// 插入的文档数量。
	InsertedCount int64

	// 在更新和替换操作中，满足过滤条件的文档数量。
	MatchedCount int64

	// update和replace操作修改的文档数量。
	ModifiedCount int64

	// 删除的文档数量。
	DeletedCount int64

	// update和replace操作中更新或替换的文档数量。
	UpsertedCount int64

	// 一个映射表，键为操作索引，值为每个已更新（upserted）文档的 _id。
	UpsertedIDs map[int64]interface{}
}

// Bulk 用于批量操作的上下文，这些批量操作将被一次性发送到数据库进行批量写入。
//
// Bulk 不支持并发安全使用。
//
// 注意事项：
//
// 当前，在一个批量操作内部的单个操作不会触发中间件或钩子。
//
// 与原始 mgo 不同，qmgo 实现的 Bulk 在不原生支持批量操作的老版本 MongoDB 服务器上，并不会模拟逐个执行批量操作。
//
// 只有官方驱动所支持的操作才会被公开提供，这就是为什么方法中缺少 InsertMany 的原因。
type Bulk struct {
	coll *Collection

	queue   []mongo.WriteModel
	ordered *bool
}

// Bulk 返回一个新的上下文，用于批量执行操作的准备工作。
func (c *Collection) Bulk() *Bulk {
	return &Bulk{
		coll:    c,
		queue:   nil,
		ordered: nil,
	}
}

// SetOrdered 标记批量操作为有序或无序。
//
// 若标记为有序，当其中一次独立写入操作失败后，后续的写入操作将不再继续。
// 默认设置为有序。
func (b *Bulk) SetOrdered(ordered bool) *Bulk {
	b.ordered = &ordered
	return b
}

// InsertOne 将一个 InsertOne 操作排队以进行批量执行。
func (b *Bulk) InsertOne(doc interface{}) *Bulk {
	wm := mongo.NewInsertOneModel().SetDocument(doc)
	b.queue = append(b.queue, wm)
	return b
}

// Remove 函数用于批量执行时，将一个 Remove 操作加入队列。
func (b *Bulk) Remove(filter interface{}) *Bulk {
	wm := mongo.NewDeleteOneModel().SetFilter(filter)
	b.queue = append(b.queue, wm)
	return b
}

// RemoveId 为批量执行队列一个RemoveId操作。
func (b *Bulk) RemoveId(id interface{}) *Bulk {
	b.Remove(bson.M{"_id": id})
	return b
}

// RemoveAll 函数用于批量执行，它将一个 RemoveAll 操作添加到待处理队列中。
func (b *Bulk) RemoveAll(filter interface{}) *Bulk {
	wm := mongo.NewDeleteManyModel().SetFilter(filter)
	b.queue = append(b.queue, wm)
	return b
}

// Upsert 在批量执行中安排一个Upsert操作。
// 替换项应为不包含操作符的文档
func (b *Bulk) Upsert(filter interface{}, replacement interface{}) *Bulk {
	wm := mongo.NewReplaceOneModel().SetFilter(filter).SetReplacement(replacement).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// UpsertOne 函数用于将一个 UpsertOne 操作加入到批量执行的队列中。
// 更新操作应当包含操作符
func (b *Bulk) UpsertOne(filter interface{}, update interface{}) *Bulk {
	wm := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// UpsertId 队列一个 UpsertId 操作以便进行批量执行。
// 替换内容应为不包含操作符的文档。
func (b *Bulk) UpsertId(id interface{}, replacement interface{}) *Bulk {
	b.Upsert(bson.M{"_id": id}, replacement)
	return b
}

// UpdateOne 将一个UpdateOne操作排队以进行批量执行。
// 更新内容应包含操作符
func (b *Bulk) UpdateOne(filter interface{}, update interface{}) *Bulk {
	wm := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update)
	b.queue = append(b.queue, wm)
	return b
}

// UpdateId 将一个UpdateId操作加入队列以进行批量执行。
// 更新操作应包含操作符
func (b *Bulk) UpdateId(id interface{}, update interface{}) *Bulk {
	b.UpdateOne(bson.M{"_id": id}, update)
	return b
}

// UpdateAll 将一个UpdateAll操作添加到队列中以进行批量执行。
// 更新操作应包含操作符
func (b *Bulk) UpdateAll(filter interface{}, update interface{}) *Bulk {
	wm := mongo.NewUpdateManyModel().SetFilter(filter).SetUpdate(update)
	b.queue = append(b.queue, wm)
	return b
}

// Run 执行收集到的所有操作，以单一的批量操作方式。
//
// 若调用成功，将会重置 Bulk。如果返回错误，则内部的操作队列保持不变，
// 该队列中包含已成功执行和未成功执行的操作。
func (b *Bulk) Run(ctx context.Context) (*BulkResult, error) {
	opts := options.BulkWriteOptions{
		Ordered: b.ordered,
	}
	result, err := b.coll.collection.BulkWrite(ctx, b.queue, &opts)
	if err != nil {
		// 在原始mgo中，如果出现错误，队列不会被重置。
		return nil, err
	}

	// 根据mgo的行为，清空队列以备可能的重用。
	b.queue = nil

	return &BulkResult{
		InsertedCount: result.InsertedCount,
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		DeletedCount:  result.DeletedCount,
		UpsertedCount: result.UpsertedCount,
		UpsertedIDs:   result.UpsertedIDs,
	}, nil
}
