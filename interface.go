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

import "go.mongodb.org/mongo-driver/mongo/options"

// CollectionI
// 集合操作接口
//type CollectionI interface {
//	Find(filter interface{}) QueryI
//	InsertOne(doc interface{}) (*mongo.InsertOneResult, error)
//	InsertMany(docs ...interface{}) (*mongo.InsertManyResult, error)
//	Upsert(filter interface{}, replacement interface{}) (*mongo.UpdateResult, error)
//	UpdateOne(filter interface{}, update interface{}) error
//	UpdateAll(filter interface{}, update interface{}) (*mongo.UpdateResult, error)
//	DeleteOne(filter interface{}) error
//	RemoveAll(selector interface{}) (*mongo.DeleteResult, error)
//	EnsureIndex(indexes []string, isUnique bool)
//	EnsureIndexes(uniques []string, indexes []string)
//}

// Change 包含了通过 Query.Apply 方法运行 findAndModify 命令时所需字段。 md5:39a15027acb265c1
type Change struct {
	X更新替换    interface{} // 更新/替换文档 md5:f186fdee95ec3578
	X是否替换   bool        // 是否替换文档而不是更新 md5:876d0fb0ea394e91
	X是否删除    bool        // 是否在找到文档后删除它，而不是更新 md5:af3a9b450dfa43f8
	X是否未找到时插入    bool        // Whether to insert in case the document isn't found, take effect when Remove is false
	X是否返回新文档 bool        // 当Remove为false时，是否返回修改后的文档而不是旧的文档 md5:52269f57ce5c8033
}

// CursorI：Cursor 接口 md5:8a6fa5bfcb19cd93
type CursorI interface {
	X下一个(result interface{}) bool
	X关闭() error
	X取错误() error
	X取全部(results interface{}) error
	//ID() int64
}

// QueryI Query interface
type QueryI interface {
	//zj:type QueryI interface {
	X分页(页码 int, 页大小 int) QueryI
	X取分页数(perPage int) int
	//zj:
	X设置排序规则(collation *options.Collation) QueryI
	X设置切片过滤(*options.ArrayFilters) QueryI
	X排序(fields ...string) QueryI
	X字段(selector interface{}) QueryI
	X跳过(n int64) QueryI
	X设置批量处理数量(n int64) QueryI
	X设置不超时(n bool) QueryI
	X设置最大返回数(n int64) QueryI
	X取一条(result interface{}) error
	X取全部(result interface{}) error
	X取数量() (n int64, err error)
	X取预估数量() (n int64, err error)
	X去重(key string, result interface{}) error
	X取结果集() CursorI
	X执行命令(change Change, result interface{}) error
	X指定索引字段(hint interface{}) QueryI
}

// AggregateI 定义聚合接口 md5:e67c5263d98eafa6
type AggregateI interface {
	X取全部(results interface{}) error
	X取一条(result interface{}) error
	Iter弃用() CursorI // 被弃用，请使用Cursor替代 md5:56d9bc403e9aa9a9
	X取结果集() CursorI
}
