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

type X索引选项 struct {
	X索引字段 []string // 指定索引键字段；以减号（-）前缀名称表示降序排列 md5:69763ade41fb7152
	*options.IndexOptions
}
