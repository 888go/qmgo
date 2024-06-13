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

package 操作符

// define the update operators
// refer: https://docs.mongodb.com/manual/reference/operator/update/
const (
	// Fields
	X更新为当前时间 = "$currentDate"
	X更新数值递增         = "$inc"
	X更新最小         = "$min"
	X更新最大         = "$max"
	X更新相乘         = "$mul"
	X更新字段名      = "$rename"
	X更新值         = "$set"
	X更新插入时 = "$setOnInsert"
	X聚合删除字段       = "$unset"

	// Array Operators
	X数组不存在追加 = "$addToSet"
	X数组删首尾      = "$pop"
	X数组删除     = "$pull"
	X数组追加     = "$push"
	X数组删除全部  = "$pullAll"

	// Array modifiers
	Each     = "$each"
	Position = "$position"
	X聚合排序     = "$sort"

	// Array bitwise
	Bit = "$bit"
)
