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

// 定义查询和投影运算符
// 参考：https://docs.mongodb.com/manual/reference/operator/query/
// md5:0562f490d100fe93
// [提示]
// const (
//
//	等于  = "$eq"
//	大于  = "$gt"
//	大于等于 = "$gte"
//	包含  = "$in"
//	小于  = "$lt"
//	小于等于 = "$lte"
//	不等于  = "$ne"
//	不在其中 = "$nin"
//
//	与  = "$and"
//	非  = "$not"
//	或非  = "$nor"
//	或  = "$or"
//
//	存在  = "$exists"
//	类型   = "$type"
//
//	表达式       = "$expr"
//	JSON模式 = "$jsonSchema"
//	取模        = "$mod"
//	正则表达式      = "$regex"
//	文本       = "$text"
//	在哪里      = "$where"
//
//	地理相交 = "$geoIntersects"
//	地理包含 = "$geoWithin"
//	近       = "$near"
//	近似球面 = "$nearSphere"
//
//	全部       = "$all"
//	元素匹配 = "$elemMatch"
//	大小      = "$size"
//
//	位全清零 = "$bitsAllClear"
//	位全置一 = "$bitsAllSet"
//	位任意清零 = "$bitsAnyClear"
//	位任意置一 = "$bitsAnySet"
//
//	注释 = "$comment"
//
//	美元符号 = "$"
//	元数据   = "$meta"
//	切片     = "$slice"
//
// )
// [结束]
const (
	// Comparison
	Eq  = "$eq"  //qm:条件等于  cz:Eq = "$eq"
	Gt  = "$gt"  //qm:条件大于  cz:Gt = "$gt"
	Gte = "$gte" //qm:条件大于等于  cz:Gte = "$gte"
	In  = "$in"  //qm:条件包含  cz:In = "$in"
	Lt  = "$lt"  //qm:条件小于  cz:Lt = "$lt"
	Lte = "$lte" //qm:条件小于等于  cz:Lte = "$lte"
	Ne  = "$ne"  //qm:条件不等于  cz:Ne = "$ne"
	Nin = "$nin" //qm:条件不包含  cz:Nin = "$nin"

	// Logical
	And = "$and" //qm:条件且  cz:And = "$and"
	Not = "$not" //qm:条件非  cz:Not = "$not"
	Nor = "$nor" //qm:条件或非  cz:Nor = "$nor"
	Or  = "$or"  //qm:条件或  cz:Or = "$or"

	// Element
	Exists = "$exists" //qm:条件字段存在  cz:Exists = "$exists"
	Type   = "$type"   //qm:条件类型  cz:Type = "$type"

	// 评价
	Expr       = "$expr"       //qm:条件表达式  cz:Expr = "$expr"
	JsonSchema = "$jsonSchema" //qm:Json效验  cz:JsonSchema = "$jsonSchema"
	Mod        = "$mod"        //qm:取模  cz:Mod = "$mod"
	Regex      = "$regex"      //qm:条件正则  cz:Regex = "$regex"
	Text       = "$text"       //qm:条件全文搜索  cz:Text = "$text"
	Where      = "$where"      //qm:条件Js  cz:Where = "$where"

	// Geo spatial
	GeoIntersects = "$geoIntersects"
	GeoWithin     = "$geoWithin"
	Near          = "$near"
	NearSphere    = "$nearSphere"

	// Array
	All       = "$all"       //qm:数组全部  cz:All = "$all"
	ElemMatch = "$elemMatch" //qm:数组匹配条件  cz:ElemMatch = "$elemMatch"
	Size      = "$size"      //qm:数组数量  cz:Size = "$size"

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
