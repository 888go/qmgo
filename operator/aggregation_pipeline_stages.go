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

// define the aggregation pipeline stages
// [提示]
//const (
// 	添加字段      = "$addFields"
// 	桶             = "$bucket"
// 	自动桶         = "$bucketAuto"
// 	集合统计       = "$collStats"
// 	计数           = "$count"
// 	方面           = "$facet"
// 	地理近邻       = "$geoNear"
// 	图形查找       = "$graphLookup"
// 	分组           = "$group"
// 	索引统计       = "$indexStats"
// 	限制           = "$limit"
// 	列出会话       = "$listSessions"
// 	连接           = "$lookup"
// 	匹配           = "$match"
// 	合并           = "$merge"
// 	输出           = "$out"
// 	计划缓存统计   = "$planCacheStats"
// 	投影           = "$project"
// 	裁剪           = "$redact"
// 	替换根         = "$replaceRoot"
// 	替换为         = "$replaceWith"
// 	采样           = "$sample"
// 	跳过           = "$skip"
// 	按计数排序     = "$sortByCount"
// 	联合           = "$unionWith"
// 	展开           = "$unwind"
// 
// 	当前操作         = "$currentOp"
// 	列出本地会话   = "$listLocalSessions"
// )
// [结束]
const (
	// 集合聚合阶段 md5:37e4a0637949107b
	AddFields      = "$addFields"
	Bucket         = "$bucket"
	BucketAuto     = "$bucketAuto"
	CollStats      = "$collStats"
	Count          = "$count"
	Facet          = "$facet"
	GeoNear        = "$geoNear"
	GraphLookup    = "$graphLookup"
	Group          = "$group"
	IndexStats     = "$indexStats"
	Limit          = "$limit"
	ListSessions   = "$listSessions"
	Lookup         = "$lookup"
	Match          = "$match"
	Merge          = "$merge"
	Out            = "$out"
	PlanCacheStats = "$planCacheStats"
	Project        = "$project"
	Redact         = "$redact"
	ReplaceRoot    = "$replaceRoot"
	ReplaceWith    = "$replaceWith"
	Sample         = "$sample"
	Skip           = "$skip"
	SortByCount    = "$sortByCount"
	UnionWith      = "$unionWith"
	Unwind         = "$unwind"

	// 数据库聚合阶段 md5:b35520d72e009304
	CurrentOp         = "$currentOp"
	ListLocalSessions = "$listLocalSessions"
)
