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

// Change 用于在 Query.Apply 方法中执行 findAndModify 命令所需的字段。
type Change struct {
	Update    interface{} // 更新/替换文档
	Replace   bool        // 是否替换整个文档而不是更新
	Remove    bool        // 是否移除找到的文档而非更新
	Upsert    bool        // 是否在未找到文档时插入，当Remove为false时生效
	ReturnNew bool        // 是否应返回修改后的文档而非旧文档，仅在Remove为false时生效
}

// CursorI：游标接口
type CursorI interface {
	Next(result interface{}) bool
	Close() error
	Err() error
	All(results interface{}) error
	//ID() int64
}

// QueryI Query interface
type QueryI interface {
	Collation(collation *options.Collation) QueryI
	SetArrayFilters(*options.ArrayFilters) QueryI
	Sort(fields ...string) QueryI
	Select(selector interface{}) QueryI
	Skip(n int64) QueryI
	BatchSize(n int64) QueryI
	NoCursorTimeout(n bool) QueryI
	Limit(n int64) QueryI
	One(result interface{}) error
	All(result interface{}) error
	Count() (n int64, err error)
	EstimatedCount() (n int64, err error)
	Distinct(key string, result interface{}) error
	Cursor() CursorI
	Apply(change Change, result interface{}) error
	Hint(hint interface{}) QueryI
}

// AggregateI 定义了聚合的接口
type AggregateI interface {
	All(results interface{}) error
	One(result interface{}) error
	Iter() CursorI // 已弃用，请改用Cursor
	Cursor() CursorI
}
