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

// X插入结果 是InsertOne 操作返回的结果类型。 md5:6a10fcb030441781
type X插入结果 struct {
	// 插入文档的 `_id`。驱动程序生成的值将为 primitive.ObjectID 类型。 md5:3e182da948916b02
	X插入ID interface{}
}

// X插入多条结果 是一个由 InsertMany 操作返回的结果类型。 md5:9dcc2360964506ed
type X插入多条结果 struct {
	// 插入文档的 `_id` 值。驱动程序生成的值类型为 `primitive.ObjectID`。 md5:7f2e18379f71d07a
	X插入IDs []interface{}
}

// X更新结果是从UpdateOne，UpdateMany和ReplaceOne操作返回的结果类型。 md5:fb4b146a87ac30bf
type X更新结果 struct {
	X匹配数  int64       // 过滤器匹配到的文档数量。 md5:cac9d662a816ba41
	X修改数 int64       // 该操作修改的文档数量。 md5:d5082cfa94d1e2ea
	X替换插入数 int64       // 由该操作上载的文档数量。 md5:0872523e362c0f10
	X替换插入ID    interface{} // upsert文档的_id字段，如果没有进行upsert操作则为nil。 md5:681ec03a43493d1c
}

// X删除结果 是 DeleteOne 和 DeleteMany 操作返回的结果类型。 md5:2c6c8718a901fb46
type X删除结果 struct {
	X删除数量 int64 // 删除的文档数量。 md5:8872e8629ebbcf3c
}
