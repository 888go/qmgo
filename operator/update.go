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

// 定义更新操作符
// 参考: https://docs.mongodb.com/manual/reference/operator/update/
// md5:166a573087a3fb27
// [提示]
// const (
//
//	当前日期 = "$currentDate"
//	增加     = "$inc"
//	最小值   = "$min"
//	最大值   = "$max"
//	乘以     = "$mul"
//	重命名   = "$rename"
//	设置     = "$set"
//	插入时设置 = "$setOnInsert"
//	取消设置 = "$unset"
//
//	添加到集合 = "$addToSet"
//	弹出       = "$pop"
//	拉取       = "$pull"
//	推入       = "$push"
//	拉取全部   = "$pullAll"
//
//	每个     = "$each"
//	位置     = "$position"
//	排序     = "$sort"
//
//	位操作 = "$bit"
//
// )
// [结束]
const (
	// Fields
	CurrentDate = "$currentDate" //qm:更新为当前时间  cz:CurrentDate = "$currentDate"
	Inc         = "$inc"         //qm:更新数值递增  cz:Inc = "$inc"
	Min         = "$min"         //qm:更新最小  cz:Min = "$min"
	Max         = "$max"         //qm:更新最大  cz:Max = "$max"
	Mul         = "$mul"         //qm:更新相乘 cz:Mul         = "$mul"
	Rename      = "$rename"      //qm:更新字段名  cz:Rename = "$rename"
	Set         = "$set"         //qm:更新值  cz:Set = "$set"
	SetOnInsert = "$setOnInsert" //qm:更新插入时  cz:SetOnInsert = "$setOnInsert"
	Unset       = "$unset"       //qm:聚合删除字段  cz:Unset = "$unset"

	//	添加到集合 = "$addToSet"
	//	弹出       = "$pop"
	//	拉取       = "$pull"
	//	推入       = "$push"
	//	拉取全部   = "$pullAll"
	//
	//	每个     = "$each"
	//	位置     = "$position"
	//	排序     = "$sort"
	//
	//	位操作 = "$bit"
	// Array Operators
	AddToSet = "$addToSet" //qm:数组不存在追加 cz:AddToSet = "$addToSet"
	Pop      = "$pop"      //qm:数组删首尾 cz:Pop      = "$pop"
	Pull     = "$pull"     //qm:数组删除 cz:Pull     = "$pull"
	Push     = "$push"     //qm:数组追加 cz:Push     = "$push"
	PullAll  = "$pullAll"  //qm:数组删除全部 cz:PullAll  = "$pullAll"

	// Array modifiers
	Each     = "$each"
	Position = "$position"
	Sort     = "$sort" //qm:聚合排序 cz:Sort = "$sort"

	// Array bitwise
	Bit = "$bit"
)
