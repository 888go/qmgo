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

// 定义查询和投影操作符
// 参考：https://docs.mongodb.com/manual/reference/operator/query/
// （这段代码注释表明接下来将定义在MongoDB中用于查询和投影的相关操作符，并提供了官方文档的参考链接，以便查阅更详细的使用说明。）
const (
	// Comparison
	Eq  = "$eq"
	Gt  = "$gt"
	Gte = "$gte"
	In  = "$in"
	Lt  = "$lt"
	Lte = "$lte"
	Ne  = "$ne"
	Nin = "$nin"

	// Logical
	And = "$and"
	Not = "$not"
	Nor = "$nor"
	Or  = "$or"

	// Element
	Exists = "$exists"
	Type   = "$type"

	// Evaluation
	Expr       = "$expr"
	JsonSchema = "$jsonSchema"
	Mod        = "$mod"
	Regex      = "$regex"
	Text       = "$text"
	Where      = "$where"

	// Geo spatial
	GeoIntersects = "$geoIntersects"
	GeoWithin     = "$geoWithin"
	Near          = "$near"
	NearSphere    = "$nearSphere"

	// Array
	All       = "$all"
	ElemMatch = "$elemMatch"
	Size      = "$size"

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
