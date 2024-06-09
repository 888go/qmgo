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
//type 中文IndexModel struct {
//     Key        []string 
//     IndexOption *options.中文IndexOptions
// }
// 
// // options包内的结构体类型翻译
// type 中文IndexOptions struct {
//     Name       string // 名称
//     Background bool   // 背景创建
//     DropDups   bool   // 删除重复值
//     ExpireAfter int32  // 过期时间（秒）
//     Sparse     bool   // 稀疏索引
//     Unique     bool   // 唯一索引
//     TextWeight float64 // 文本搜索权重
// }
// [结束]
type IndexModel struct {//hm:索引选项  cz:type IndexModel  
	Key []string//qm:索引字段  cz:Key []string   // 指定索引键字段；以减号（-）前缀名称表示降序排列 md5:69763ade41fb7152
	*options.IndexOptions
}
