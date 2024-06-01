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

// BulkResult 是由Bulk.Run操作返回的结果类型。 md5:3a422d6b1b20649c
type BulkResult struct {//hm:批量操作结果  cz:type BulkResult  
	// 插入的文档数量。 md5:f44082352897f08b
	InsertedCount int64//qm:插入数  cz:InsertedCount int64  

	// 更新和替换操作中，被过滤器匹配的文档数量。 md5:90fab681d83f2e97
	MatchedCount int64//qm:匹配数  cz:MatchedCount int64  

	// 被更新和替换操作修改的文档数量。 md5:1e4886e32c8092e3
	ModifiedCount int64//qm:修改数  cz:ModifiedCount int64  

	// 删除的文档数量。 md5:8872e8629ebbcf3c
	DeletedCount int64//qm:删除数  cz:DeletedCount int64  

	// 通过update和replace操作插入的文档数量。 md5:3074b4c76263ae0c
	UpsertedCount int64//qm:更新插入数  cz:UpsertedCount int64  

	// 一个操作索引到每个插入文档的_id的映射。 md5:b4c301dceb41d860
	UpsertedIDs map[int64]interface{}//qm:更新插入IDs  cz:UpsertedIDs map[int64]interface{}  
}

// Bulk is context for batching operations to be sent to database in a single
// bulk write.
//
// Bulk is not safe for concurrent use.
//
//
// Individual operations inside a bulk do not trigger middlewares or hooks
// at present.
//
// Different from original mgo, the qmgo implementation of Bulk does not emulate
// bulk operations individually on old versions of MongoDB servers that do not
// natively support bulk operations.
//
// Only operations supported by the official driver are exposed, that is why
// InsertMany is missing from the methods.
type Bulk struct {//hm:批量操作  cz:type Bulk  
	coll *Collection

	queue   []mongo.WriteModel
	ordered *bool
}

// Bulk返回一个新的上下文，用于准备批量执行操作。 md5:e39897d617450e92
// ff:创建批量执行
func (c *Collection) Bulk() *Bulk {
	return &Bulk{
		coll:    c,
		queue:   nil,
		ordered: nil,
	}
}

// SetOrdered 将批量设置为有序或无序。
//
// 如果设置为有序，写操作在单个写操作失败后不会继续。默认为有序。
// md5:caf2eac3fe50a750
// ff:设置有序执行
// ordered:开启有序
func (b *Bulk) SetOrdered(ordered bool) *Bulk {
	b.ordered = &ordered
	return b
}

// InsertOne 将一个 InsertOne 操作加入到批量执行队列中。 md5:65abbf989aa97556
// ff:插入
// doc:待插入文档
func (b *Bulk) InsertOne(doc interface{}) *Bulk {
	wm := mongo.NewInsertOneModel().SetDocument(doc)
	b.queue = append(b.queue, wm)
	return b
}

// Remove 队列一个删除操作，用于批量执行。 md5:a9c84e1a291eea0f
// ff:删除一条
// filter:删除条件
func (b *Bulk) Remove(filter interface{}) *Bulk {
	wm := mongo.NewDeleteOneModel().SetFilter(filter)
	b.queue = append(b.queue, wm)
	return b
}

// RemoveId 队列一个 RemoveId 操作以进行批量执行。 md5:f3fbfef26bde41fc
// ff:删除并按ID
// id:删除ID
func (b *Bulk) RemoveId(id interface{}) *Bulk {
	b.Remove(bson.M{"_id": id})
	return b
}

// RemoveAll 会将一个 RemoveAll 操作加入到批量执行的队列中。 md5:df548d516b324574
// ff:删除
// filter:删除条件
func (b *Bulk) RemoveAll(filter interface{}) *Bulk {
	wm := mongo.NewDeleteManyModel().SetFilter(filter)
	b.queue = append(b.queue, wm)
	return b
}

// Upsert将Upsert操作排队进行批量执行。替换应该是没有操作符的文档
// md5:1115932f50b88737
// ff:更新插入
// filter:更新条件
// replacement:更新内容
func (b *Bulk) Upsert(filter interface{}, replacement interface{}) *Bulk {
	wm := mongo.NewReplaceOneModel().SetFilter(filter).SetReplacement(replacement).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// UpsertOne 为批量执行队列一个 UpsertOne 操作。更新操作应该包含运算符
// md5:7052a86d53229aab
// ff:更新插入一条
// filter:更新条件
// update:更新内容
func (b *Bulk) UpsertOne(filter interface{}, update interface{}) *Bulk {
	wm := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// UpsertId 用于批量执行的UpsertId操作进行排队。
// 替换的文档应该不包含操作符。
// md5:c5d9cc678823f8e5
// ff:更新插入并按ID
// id:更新ID
// replacement:更新内容
func (b *Bulk) UpsertId(id interface{}, replacement interface{}) *Bulk {
	b.Upsert(bson.M{"_id": id}, replacement)
	return b
}

// UpdateOne 为批量执行队列一个 UpdateOne 操作。更新操作应该包含操作符
// md5:0e587045b560687a
// ff:更新一条
// filter:更新条件
// update:更新内容
func (b *Bulk) UpdateOne(filter interface{}, update interface{}) *Bulk {
	wm := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update)
	b.queue = append(b.queue, wm)
	return b
}

// UpdateId 为批量执行排队一个 UpdateId 操作。更新应该包含操作符
// md5:968d7d02f007ae39
// ff:更新并按ID
// id:更新ID
// update:更新内容
func (b *Bulk) UpdateId(id interface{}, update interface{}) *Bulk {
	b.UpdateOne(bson.M{"_id": id}, update)
	return b
}

// UpdateAll 队列一个 UpdateAll 操作，用于批量执行。
// 更新应该包含操作符
// md5:b1fdc26a48273948
// ff:更新
// filter:更新条件
// update:更新内容
func (b *Bulk) UpdateAll(filter interface{}, update interface{}) *Bulk {
	wm := mongo.NewUpdateManyModel().SetFilter(filter).SetUpdate(update)
	b.queue = append(b.queue, wm)
	return b
}

// Run 执行收集到的单个批量操作。
//
// 调用成功会重置 Bulk。如果返回错误，内部操作队列保持不变，包含成功和失败的操作。
// md5:c3ce14d8defe8da0
// ff:执行
// ctx:上下文
func (b *Bulk) Run(ctx context.Context) (*BulkResult, error) {
	opts := options.BulkWriteOptions{
		Ordered: b.ordered,
	}
	result, err := b.coll.collection.BulkWrite(ctx, b.queue, &opts)
	if err != nil {
		// 在原始的mgo中，如果发生错误，队列不会被重置。 md5:b7f801e955f364a8
		return nil, err
	}

	// 清空队列以备可能的重用，遵循mgo的行为。 md5:ac1070c096c485e8
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
