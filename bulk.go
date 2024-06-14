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

// X批量操作结果 是由Bulk.Run操作返回的结果类型。 md5:3a422d6b1b20649c
type X批量操作结果 struct {
	// 插入的文档数量。 md5:f44082352897f08b
	X插入数 int64

	// 更新和替换操作中，被过滤器匹配的文档数量。 md5:90fab681d83f2e97
	X匹配数 int64

	// 被更新和替换操作修改的文档数量。 md5:1e4886e32c8092e3
	X修改数 int64

	// 删除的文档数量。 md5:8872e8629ebbcf3c
	X删除数 int64

	// 通过update和replace操作插入的文档数量。 md5:3074b4c76263ae0c
	X替换插入数 int64

	// 一个操作索引到每个插入文档的_id的映射。 md5:b4c301dceb41d860
	X替换插入IDs map[int64]interface{}
}

// X批量操作 是用于批量操作的上下文，这些操作将一次性发送到数据库进行批量写入。
//
// X批量操作 不适用于并发使用。
//
// 注意：
//
// 在批量操作中的单个操作目前不会触发中间件或钩子。
//
// 与原版 mgo 不同，qmgo 实现的 X批量操作 并不会在不支持原生批量操作的老版本 MongoDB 服务器上模拟逐个执行批量操作。
//
// 只有官方驱动支持的操作被暴露出来，因此方法中缺少 InsertMany。
// md5:97e7f3c645b8ba7f
type X批量操作 struct {
	coll *X文档集合

	queue   []mongo.WriteModel
	ordered *bool
}

// X创建批量执行返回一个新的上下文，用于准备批量执行操作。 md5:e39897d617450e92
func (c *X文档集合) X创建批量执行() *X批量操作 {
	return &X批量操作{
		coll:    c,
		queue:   nil,
		ordered: nil,
	}
}

// X设置有序执行 将批量设置为有序或无序。
//
// 如果设置为有序，写操作在单个写操作失败后不会继续。默认为有序。
// md5:caf2eac3fe50a750
func (b *X批量操作) X设置有序执行(开启有序 bool) *X批量操作 {
	b.ordered = &开启有序
	return b
}

// X插入 将一个 X插入 操作加入到批量执行队列中。 md5:65abbf989aa97556
func (b *X批量操作) X插入(待插入文档 interface{}) *X批量操作 {
	wm := mongo.NewInsertOneModel().SetDocument(待插入文档)
	b.queue = append(b.queue, wm)
	return b
}

// X删除一条 队列一个删除操作，用于批量执行。 md5:a9c84e1a291eea0f
func (b *X批量操作) X删除一条(删除条件 interface{}) *X批量操作 {
	wm := mongo.NewDeleteOneModel().SetFilter(删除条件)
	b.queue = append(b.queue, wm)
	return b
}

// X删除并按ID 队列一个 X删除并按ID 操作以进行批量执行。 md5:f3fbfef26bde41fc
func (b *X批量操作) X删除并按ID(删除ID interface{}) *X批量操作 {
	b.X删除一条(bson.M{"_id": 删除ID})
	return b
}

// X删除 会将一个 X删除 操作加入到批量执行的队列中。 md5:df548d516b324574
func (b *X批量操作) X删除(删除条件 interface{}) *X批量操作 {
	wm := mongo.NewDeleteManyModel().SetFilter(删除条件)
	b.queue = append(b.queue, wm)
	return b
}

// X替换插入将X替换插入操作排队进行批量执行。替换应该是没有操作符的文档
// md5:1115932f50b88737
func (b *X批量操作) X替换插入(替换条件 interface{}, 替换内容 interface{}) *X批量操作 {
	wm := mongo.NewReplaceOneModel().SetFilter(替换条件).SetReplacement(替换内容).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// X替换插入一条 为批量执行队列一个 X替换插入一条 操作。更新操作应该包含运算符
// md5:7052a86d53229aab一条
func (b *X批量操作) X替换插入一条(替换条件 interface{}, 替换内容 interface{}) *X批量操作 {
	wm := mongo.NewUpdateOneModel().SetFilter(替换条件).SetUpdate(替换内容).SetUpsert(true)
	b.queue = append(b.queue, wm)
	return b
}

// X替换插入并按ID 用于批量执行的X替换插入并按ID操作进行排队。
// 替换的文档应该不包含操作符。
// md5:c5d9cc678823f8e5并按ID
func (b *X批量操作) X替换插入并按ID(替换ID interface{}, 替换内容 interface{}) *X批量操作 {
	b.X替换插入(bson.M{"_id": 替换ID}, 替换内容)
	return b
}

// X更新一条 为批量执行队列一个 X更新一条 操作。更新操作应该包含操作符
// md5:0e587045b560687a
func (b *X批量操作) X更新一条(更新条件 interface{}, 更新内容 interface{}) *X批量操作 {
	wm := mongo.NewUpdateOneModel().SetFilter(更新条件).SetUpdate(更新内容)
	b.queue = append(b.queue, wm)
	return b
}

// X更新并按ID 为批量执行排队一个 X更新并按ID 操作。更新应该包含操作符
// md5:968d7d02f007ae39
func (b *X批量操作) X更新并按ID(更新ID interface{}, 更新内容 interface{}) *X批量操作 {
	b.X更新一条(bson.M{"_id": 更新ID}, 更新内容)
	return b
}

// X更新 队列一个 X更新 操作，用于批量执行。
// 更新应该包含操作符
// md5:b1fdc26a48273948
func (b *X批量操作) X更新(更新条件 interface{}, 更新内容 interface{}) *X批量操作 {
	wm := mongo.NewUpdateManyModel().SetFilter(更新条件).SetUpdate(更新内容)
	b.queue = append(b.queue, wm)
	return b
}

// X执行 执行收集到的单个批量操作。
//
// 调用成功会重置 Bulk。如果返回错误，内部操作队列保持不变，包含成功和失败的操作。
// md5:c3ce14d8defe8da0
func (b *X批量操作) X执行(上下文 context.Context) (*X批量操作结果, error) {
	opts := options.BulkWriteOptions{
		Ordered: b.ordered,
	}
	result, err := b.coll.collection.BulkWrite(上下文, b.queue, &opts)
	if err != nil {
		// 在原始的mgo中，如果发生错误，队列不会被重置。 md5:b7f801e955f364a8
		return nil, err
	}

	// 清空队列以备可能的重用，遵循mgo的行为。 md5:ac1070c096c485e8
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
