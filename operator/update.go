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

// define the update operators
// [提示]
//const (
// 	当前日期 = "$currentDate"
// 	增加     = "$inc"
// 	最小值   = "$min"
// 	最大值   = "$max"
// 	乘以     = "$mul"
// 	重命名   = "$rename"
// 	设置     = "$set"
// 	插入时设置 = "$setOnInsert"
// 	取消设置 = "$unset"
// 
// 	添加到集合 = "$addToSet"
// 	弹出       = "$pop"
// 	拉取       = "$pull"
// 	推入       = "$push"
// 	拉取全部   = "$pullAll"
// 
// 	每个     = "$each"
// 	位置     = "$position"
// 	排序     = "$sort"
// 
// 	位操作 = "$bit"
// )
// [结束]
const (
	// Fields
	CurrentDate = "$currentDate"
	Inc         = "$inc"
	Min         = "$min"//qm:最小值  cz:Min = "$min"  
	Max         = "$max"//qm:最大值  cz:Max = "$max"  
	Mul         = "$mul"
	Rename      = "$rename"
	Set         = "$set"//qm:设置值  cz:Set = "$set"  
	SetOnInsert = "$setOnInsert"
	Unset       = "$unset"

	// Array Operators
	AddToSet = "$addToSet"
	Pop      = "$pop"
	Pull     = "$pull"
	Push     = "$push"
	PullAll  = "$pullAll"

	// Array modifiers
	Each     = "$each"
	Position = "$position"
	Sort     = "$sort"

	// Array bitwise
	Bit = "$bit"
)
