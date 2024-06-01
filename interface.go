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
	Update    interface{}// 更新/替换文档 md5:f186fdee95ec3578
	Replace   bool//qm:是否替换  cz:Replace bool          // 是否替换文档而不是更新 md5:876d0fb0ea394e91
	Remove    bool//qm:是否删除  cz:Remove bool          // 是否在找到文档后删除它，而不是更新 md5:af3a9b450dfa43f8
	Upsert    bool//qm:未找到是否插入  cz:Upsert bool          // Whether to insert in case the document isn't found, take effect when Remove is false
	ReturnNew bool//qm:是否返回新文档  cz:ReturnNew bool          // 当Remove为false时，是否返回修改后的文档而不是旧的文档 md5:52269f57ce5c8033
}

// CursorI：Cursor 接口 md5:8a6fa5bfcb19cd93
type CursorI interface {
	Next(result interface{}) bool//qm:下一个  cz:Next(result interface{}) bool  
	Close() error//qm:关闭  cz:Close() error  
	Err() error//qm:取错误  cz:Err() error  
	All(results interface{}) error//qm:取全部  cz:All(results interface{}) error  
	//ID() int64
}

// QueryI Query interface
type QueryI interface {
	Collation(collation *options.Collation) QueryI//qm:设置排序规则  cz:Collation(collation *options.Collation) QueryI  
	SetArrayFilters(*options.ArrayFilters) QueryI//qm:设置切片过滤  cz:SetArrayFilters(*options.ArrayFilters) QueryI  
	Sort(fields ...string) QueryI//qm:排序  cz:Sort(fields ...string) QueryI  
	Select(selector interface{}) QueryI//qm:字段  cz:Select(selector interface{}) QueryI  
	Skip(n int64) QueryI//qm:跳过  cz:Skip(n int64) QueryI  
	BatchSize(n int64) QueryI//qm:设置批量处理数量  cz:BatchSize(n int64) QueryI  
	NoCursorTimeout(n bool) QueryI//qm:设置不超时  cz:NoCursorTimeout(n bool) QueryI  
	Limit(n int64) QueryI//qm:设置最大返回数  cz:Limit(n int64) QueryI  
	One(result interface{}) error//qm:取一条  cz:One(result interface{}) error  
	All(result interface{}) error//qm:取全部  cz:All(result interface{}) error  
	Count() (n int64, err error)//qm:取数量  cz:Count() (n int64, err error)  
	EstimatedCount() (n int64, err error)//qm:取预估数量  cz:EstimatedCount() (n int64, err error)  
	Distinct(key string, result interface{}) error//qm:去重  cz:Distinct(key string, result interface{}) error  
	Cursor() CursorI//qm:取结果集  cz:Cursor() CursorI  
	Apply(change Change, result interface{}) error//qm:执行命令  cz:Apply(change Change, result interface{}) error  
	Hint(hint interface{}) QueryI//qm:指定索引字段  cz:Hint(hint interface{}) QueryI  
}

// AggregateI 定义聚合接口 md5:e67c5263d98eafa6
type AggregateI interface {
	All(results interface{}) error//qm:取全部  cz:All(results interface{}) error  
	One(result interface{}) error//qm:取一条  cz:One(result interface{}) error  
	Iter() CursorI// 被弃用，请使用Cursor替代 md5:56d9bc403e9aa9a9
	Cursor() CursorI//qm:取结果集  cz:Cursor() CursorI  
}
