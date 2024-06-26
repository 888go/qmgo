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

package mgo常量

// 查询修饰符
// 参考：https://docs.mongodb.com/manual/reference/operator/query-modifier/
// md5:d060699c713ecc32
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
