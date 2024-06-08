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

import "go.mongodb.org/mongo-driver/mongo/options"

// CollectionI
// 集合操作接口
// Change 包含了通过 Query.Apply 方法运行 findAndModify 命令时所需字段。 md5:39a15027acb265c1
// [提示]
//
//	type 变更 struct {
//	    更新    interface{}
//	    替换   bool
//	    删除   bool
//	    若不存在插入 bool
//	    返回新文档 bool
//	}
//
// [结束]
type Change struct {
	Update    interface{} //qm:更新替换  cz:Update interface{}   // 更新/替换文档 md5:f186fdee95ec3578
	Replace   bool        //qm:是否替换  cz:Replace bool          // 是否替换文档而不是更新 md5:876d0fb0ea394e91
	Remove    bool        //qm:是否删除  cz:Remove bool          // 是否在找到文档后删除它，而不是更新 md5:af3a9b450dfa43f8
	Upsert    bool        //qm:是否未找到时插入  cz:Upsert bool          // Whether to insert in case the document isn't found, take effect when Remove is false
	ReturnNew bool        //qm:是否返回新文档  cz:ReturnNew bool          // 当Remove为false时，是否返回修改后的文档而不是旧的文档 md5:52269f57ce5c8033
}

// CursorI：Cursor 接口 md5:8a6fa5bfcb19cd93
// [提示]
//
//	type CursorI interface {
//	    // 获取下一个文档
//	    Next() (interface{}, error)
//
//	    // 错误信息
//	    Err() error
//
//	    // 关闭游标
//	    Close() error
//	}
//
// [结束]
type CursorI interface {
	// [提示:] Next(结果 interface{}) bool
	Next(result interface{}) bool //qm:下一个  cz:Next(result interface{}) bool
	// [提示:] 关闭() 错误
	Close() error //qm:关闭  cz:Close() error
	// [提示:] 错误() 错误
	Err() error //qm:取错误  cz:Err() error
	// [提示]
	//所有(result 接口类型) 错误
	All(results interface{}) error //qm:取全部  cz:All(results interface{}) error
	//ID() int64
}

// QueryI Query interface
// [提示:] 类型 QueryI 接口{}
type QueryI interface {
	//zj:type QueryI interface {
	X分页(页码 int, 页大小 int) QueryI
	X取分页数(perPage int) int
	//zj:

	// [提示:] Collation(分片规则 *options.Collation) 查询接口I
	Collation(collation *options.Collation) QueryI //qm:设置排序规则  cz:Collation(collation *options.Collation) QueryI
	// [提示:] Set数组过滤器(*options.ArrayFilters) 查询接口
	SetArrayFilters(*options.ArrayFilters) QueryI //qm:设置切片过滤  cz:SetArrayFilters(*options.ArrayFilters) QueryI
	// [提示]
	//升序排序(字段... string) 查询接口I

	Sort(fields ...string) QueryI //qm:排序  cz:Sort(fields ...string) QueryI
	// [提示:] 选择查询(选择器 interface{}) 查询接口I
	Select(selector interface{}) QueryI //qm:字段  cz:Select(selector interface{}) QueryI
	// [提示:] 跳过(n int64) 查询接口I
	Skip(n int64) QueryI //qm:跳过  cz:Skip(n int64) QueryI
	// [提示]
	//func BatchSize(n int64) QueryI {
	//     // ...
	// }
	// [结束]
	BatchSize(n int64) QueryI //qm:设置批量处理数量  cz:BatchSize(n int64) QueryI
	// [提示]
	//无超时游标(NoCursorTimeout) (n bool) 查询接口I
	//
	// Close() error
	//
	// All(interface{}) error
	//
	// One(interface{}) error
	//
	// Iter() IterI
	//
	// Count() (int64, error)
	//
	// Skip(int) QueryI
	//
	// Limit(int) QueryI
	//
	// Sort(string ...) QueryI
	//
	// Select(bson.D) QueryI
	//
	// Hint(string) QueryI
	//
	// SetMaxScan(int) QueryI
	//
	// SetMaxTimeMS(int) QueryI
	//
	// SetComment(string) QueryI
	//
	// SetSnapshot(bool) QueryI
	//
	// SetAllowDiskUse(bool) QueryI
	//
	// SetCollation(string) QueryI
	//
	// SetBatchSize(int) QueryI
	//
	// SetAwaitData(bool) QueryI
	//
	// Set Exhaust(bool) QueryI
	//
	// SetPartial(bool) QueryI
	// [结束]
	NoCursorTimeout(n bool) QueryI //qm:设置不超时  cz:NoCursorTimeout(n bool) QueryI
	// [提示:] 限制数量(n int64) 查询接口I
	Limit(n int64) QueryI //qm:设置最大返回数  cz:Limit(n int64) QueryI
	// [提示:] 单个查询结果(结果接口类型) (错误)
	// [提示:] 单个查询结果(结果接口类型) (错误)
	One(result interface{}) error //qm:取一条  cz:One(result interface{}) error
	// [提示:] 所有(result 结果) error 错误
	All(result interface{}) error //qm:取全部  cz:All(result interface{}) error
	// [提示:] Count() (数量 int64, 错误 error)
	Count() (n int64, err error) //qm:取数量  cz:Count() (n int64, err error)
	// [提示:] 估算计数() (数量 int64, 错误 error)
	EstimatedCount() (n int64, err error) //qm:取预估数量  cz:EstimatedCount() (n int64, err error)
	// [提示:] 唯一值(Distinct)：根据给定的键(key)获取集合中该字段的所有唯一值，并将结果存储到interface{}类型的变量中，可能返回错误信息。
	Distinct(key string, result interface{}) error //qm:去重  cz:Distinct(key string, result interface{}) error
	// [提示]
	//Cursor() -> 获取游标接口
	//
	// CursorI -> 游标接口
	// [结束]
	// [提示]
	//Cursor() -> 获取游标接口
	//
	// CursorI -> 游标接口
	// [结束]
	Cursor() CursorI //qm:取结果集  cz:Cursor() CursorI
	// [提示:] 应用更改(Change)到文档, 结果存储到result(接口类型), 返回错误信息
	Apply(change Change, result interface{}) error //qm:执行命令  cz:Apply(change Change, result interface{}) error
	// [提示]
	//设置提示(Hint interface{}) 查询接口I
	//
	// Filter(filter interface{}) QueryI
	// [结束]
	Hint(hint interface{}) QueryI //qm:指定索引字段  cz:Hint(hint interface{}) QueryI
}

// AggregateI 定义聚合接口 md5:e67c5263d98eafa6
// [提示]
// 类型名称：聚合接口 (AggregateInterface)
//
// 没有找到具体的接口方法，通常在Go语言中，接口包含一个或多个方法。如果接口中包含方法，请提供这些方法以便进行翻译。如果没有，那么这个接口只表示一个聚合操作的占位符，无需单独翻译其方法。
// [结束]
type AggregateI interface {
	// [提示]
	//所有(result 接口类型) 错误
	//
	// Find(filter interface{}, results interface{}) error
	//
	// FindOne(filter interface{}, result interface{}) error
	//
	// CountDocuments(filter interface{}, count *int64) error
	//
	// DeleteOne(filter interface{}) error
	//
	// DeleteMany(filter interface{}, deletedCount *int64) error
	//
	// UpdateOne(filter interface{}, update interface{}, upsert bool) error
	//
	// UpdateMany(filter interface{}, update interface{}, upsert bool, matchedCount *int64, modifiedCount *int64) error
	//
	// ReplaceOne(filter interface{}, replacement interface{}, upsert bool, matchedCount *int64, modifiedCount *int64) error
	//
	// InsertOne(document interface{}, insertedID *string) error
	//
	// InsertMany(documents interface{}, insertedIDs **string) error
	// [结束]
	// [提示]
	//所有(result 接口类型) 错误
	//
	// Find(filter interface{}, results interface{}) error
	//
	// FindOne(filter interface{}, result interface{}) error
	//
	// CountDocuments(filter interface{}, count *int64) error
	//
	// DeleteOne(filter interface{}) error
	//
	// DeleteMany(filter interface{}, deletedCount *int64) error
	//
	// UpdateOne(filter interface{}, update interface{}, upsert bool) error
	//
	// UpdateMany(filter interface{}, update interface{}, upsert bool, matchedCount *int64, modifiedCount *int64) error
	//
	// ReplaceOne(filter interface{}, replacement interface{}, upsert bool, matchedCount *int64, modifiedCount *int64) error
	//
	// InsertOne(document interface{}, insertedID *string) error
	//
	// InsertMany(documents interface{}, insertedIDs **string) error
	// [结束]
	All(results interface{}) error //qm:取全部  cz:All(results interface{}) error
	// [提示:] 单个查询结果(结果接口类型) (错误)
	// [提示:] 单个查询结果(结果接口类型) (错误)
	One(result interface{}) error //qm:取一条  cz:One(result interface{}) error
	// [提示]
	//`Iter()` -> `迭代器()`
	// `CursorI` -> `游标接口`
	// [结束]
	Iter() CursorI //qm:Iter弃用  cz:Iter() CursorI   // 被弃用，请使用Cursor替代 md5:56d9bc403e9aa9a9
	// [提示]
	//Cursor() -> 获取游标接口
	//
	// CursorI -> 游标接口
	// [结束]
	// [提示]
	//Cursor() -> 获取游标接口
	//
	// CursorI -> 游标接口
	// [结束]
	Cursor() CursorI //qm:取结果集  cz:Cursor() CursorI
}
