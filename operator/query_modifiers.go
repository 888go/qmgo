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

package operator

//Query Modifiers
// [提示]
//const (
// 	解释 = "$explain"
// 	提示 = "$hint"
// 	最大时间毫秒 = "$maxTimeMS"
// 	排序 = "$orderby"
// 	查询 = "$query"
// 	返回键 = "$returnKey"
// 	显示磁盘位置 = "$showDiskLoc"
// 	
// 	自然顺序 = "$natural"
// )
// [结束]
const (
	// Modifiers
	Explain     = "$explain"
	Hint        = "$hint"
	MaxTimeMS   = "$maxTimeMS"
	OrderBy     = "$orderby"
	Query       = "$query"
	ReturnKey   = "$returnKey"
	ShowDiskLoc = "$showDiskLoc"

	// Sort Order
	Natural = "$natural"
)
