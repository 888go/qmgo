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

// 定义查询和投影运算符
// 参考：https://docs.mongodb.com/manual/reference/operator/query/
// md5:0562f490d100fe93
const (
	// Comparison
	X条件等于  = "$eq"
	X条件大于  = "$gt"
	X条件大于等于 = "$gte"
	X条件包含  = "$in"
	X条件小于  = "$lt"
	X条件小于等于 = "$lte"
	X条件不等于  = "$ne"
	X条件不包含 = "$nin"

	// Logical
	X条件且 = "$and"
	X条件非 = "$not"
	X条件或非 = "$nor"
	X条件或  = "$or"

	// Element
	X条件字段存在 = "$exists"
	X条件类型   = "$type"

	// Evaluation
	X条件表达式       = "$expr"
	Json效验 = "$jsonSchema"
	X取模        = "$mod"
	X条件正则      = "$regex"
	X条件全文搜索       = "$text"
	X条件Js      = "$where"

	// Geo spatial
	GeoIntersects = "$geoIntersects"
	GeoWithin     = "$geoWithin"
	Near          = "$near"
	NearSphere    = "$nearSphere"

	// Array
	X数组全部       = "$all"
	X数组匹配条件 = "$elemMatch"
	X数组数量      = "$size"

	// Bitwise
	BitsAllClear = "$bitsAllClear"
	BitsAllSet   = "$bitsAllSet"
	BitsAnyClear = "$bitsAnyClear"
	BitsAnySet   = "$bitsAnySet"

	// Comments
	Comment = "$comment"

	// Projection operators
	Dollar = "$"
	Meta   = "$meta"
	Slice  = "$slice"
)
