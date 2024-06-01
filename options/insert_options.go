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

package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 插入单个选项 struct {
// 插入钩子 interface{}
// *options.插入单个选项
// } 
// 
// 这里的翻译可能需要一些上下文解释：
// - `InsertOneOptions` 可以翻译为 "插入单个选项"，它代表了一组用于插入单条数据的操作选项。
// - `InsertHook` 可以理解为 "插入钩子"，可能是指在插入操作前或后的回调函数，用于自定义逻辑。
// - `options.InsertOneOptions` 这个是指向相同类型但位于 `options` 包内的引用，保持包名不变，所以仍然是 "插入单个选项"。
// [结束]
type InsertOneOptions struct {
	InsertHook interface{}
	*options.InsertOneOptions
}
// [提示]
//type 插入多条选项 struct {
//     插入钩子 interface{}
//     *options.插入多条选项
// }
// [结束]
type InsertManyOptions struct {
	InsertHook interface{}
	*options.InsertManyOptions
}
