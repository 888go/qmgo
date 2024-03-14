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

package operator //bm:操作符

// 定义查询和投影操作符
// 参考：https://docs.mongodb.com/manual/reference/operator/query/
// （这段代码注释表明接下来将定义在MongoDB中用于查询和投影的相关操作符，并提供了官方文档的参考链接，以便查阅更详细的使用说明。）
const (
	// Comparison
	Eq  = "$eq"  //hs:等于
	Gt  = "$gt"  //hs:大于
	Gte = "$gte" //hs:大于等于
	In  = "$in"  //hs:包含
	Lt  = "$lt"  //hs:小于
	Lte = "$lte" //hs:小于等于
	Ne  = "$ne"  //hs:不等于
	Nin = "$nin" //hs:不包含

	// Logical
	And = "$and" //hs:且
	Not = "$not" //
	Nor = "$nor" //
	Or  = "$or"  //hs:或

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
