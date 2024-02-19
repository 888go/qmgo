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

// InsertOneResult 是 InsertOne 操作返回的结果类型。
type InsertOneResult struct {
	// 插入文档的 _id。由驱动程序生成的值将为 primitive.ObjectID 类型。
	InsertedID interface{}
}

// InsertManyResult 是由 InsertMany 操作返回的一种结果类型。
type InsertManyResult struct {
	// 插入文档的 _id 值。由驱动程序生成的值将为 primitive.ObjectID 类型。
	InsertedIDs []interface{}
}

// UpdateResult 是从 UpdateOne、UpdateMany 和 ReplaceOne 操作返回的结果类型。
type UpdateResult struct {
	MatchedCount  int64       // 匹配过滤器的文档数量。
	ModifiedCount int64       // 此操作修改的文档数量。
	UpsertedCount int64       // 该操作执行的文档更新或插入的数量。
	UpsertedID    interface{} // _id字段是更新文档（upsert document）中的_id字段，如果未执行任何更新操作，则为nil。
}

// DeleteResult 是由 DeleteOne 和 DeleteMany 操作返回的结果类型。
type DeleteResult struct {
	DeletedCount int64 // 删除的文档数量。
}
