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

// InsertOneResult 是InsertOne 操作返回的结果类型。 md5:6a10fcb030441781
// [提示]
//
//	type 插入结果 struct {
//	    插入ID interface{}
//	}
//
// [结束]
type InsertOneResult struct { //hm:插入结果  cz:type InsertOneResult
	// 插入文档的 `_id`。驱动程序生成的值将为 primitive.ObjectID 类型。 md5:3e182da948916b02
	InsertedID interface{} //qm:插入ID  cz:InsertedID interface{}
}

// InsertManyResult 是一个由 InsertMany 操作返回的结果类型。 md5:9dcc2360964506ed
// [提示]
// type 插入多条结果 struct {
// 插入ID []interface{}
// }
// [结束]
type InsertManyResult struct { //hm:插入多条结果  cz:type InsertManyResult
	// 插入文档的 `_id` 值。驱动程序生成的值类型为 `primitive.ObjectID`。 md5:7f2e18379f71d07a
	InsertedIDs []interface{} //qm:插入IDs  cz:InsertedIDs []interface{}
}

// UpdateResult是从UpdateOne，UpdateMany和ReplaceOne操作返回的结果类型。 md5:fb4b146a87ac30bf
// [提示]
//
//	type 更新结果 struct {
//	    匹配计数       int64
//	    修改计数       int64
//	    插入更新计数   int64
//	    插入ID        interface{}
//	}
//
// [结束]
type UpdateResult struct { //hm:更新结果  cz:type UpdateResult
	MatchedCount  int64       //qm:匹配数  cz:MatchedCount int64         // 过滤器匹配到的文档数量。 md5:cac9d662a816ba41
	ModifiedCount int64       //qm:修改数  cz:ModifiedCount int64         // 该操作修改的文档数量。 md5:d5082cfa94d1e2ea
	UpsertedCount int64       //qm:替换插入数  cz:UpsertedCount int64         // 由该操作上载的文档数量。 md5:0872523e362c0f10
	UpsertedID    interface{} //qm:替换插入ID  cz:UpsertedID interface{}   // upsert文档的_id字段，如果没有进行upsert操作则为nil。 md5:681ec03a43493d1c
}

// DeleteResult 是 DeleteOne 和 DeleteMany 操作返回的结果类型。 md5:2c6c8718a901fb46
// [提示]
//
//	type 删除结果 struct {
//	    删除数量 int64
//	}
//
// [结束]
type DeleteResult struct { //hm:删除结果  cz:type DeleteResult
	DeletedCount int64 //qm:删除数量  cz:DeletedCount int64   // 删除的文档数量。 md5:8872e8629ebbcf3c
}
