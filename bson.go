/*
 Copyright 2021 The Qmgo Authors.
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

import "go.mongodb.org/mongo-driver/bson"

// 别名mongo驱动的bson原语
// 因此用户不需要导入go.mongodb.org/mongo-driver/mongo，所有内容都在qmgo中可用
// md5:2f6e3ba77edc7a63
type (
	// map[string]interface{} , 如:bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}, M是 bson.M 的别名 md5:66b7bee0d7904542
	M = bson.M
	// []interface{},如:bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}} , A是bson.A的别名 md5:7a6f09b99ea36324
	A = bson.A
	// Key/Value结构体数组, 如:bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}} ,D是 bson.D 的别名  md5:a2fd7b05e87775b6
	D = bson.D
	// Key/Value结构体, E 内部的单个元素,是bson.E的别名 md5:d1a800789b88ac58
	E = bson.E
)
