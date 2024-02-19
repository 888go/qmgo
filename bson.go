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

// 对mongo驱动的bson原生类型进行别名定义
// 这样用户就不需要直接导入go.mongodb.org/mongo-driver/mongo包，所有相关类型都在qmgo包中提供
type (
	// M 是 bson.M 的别名
	M = bson.M
	// A 是 bson.A 的别名
	A = bson.A
	// D 是 bson.D 的别名
	D = bson.D
	// E 是 bson.E 的别名
	E = bson.E
)
