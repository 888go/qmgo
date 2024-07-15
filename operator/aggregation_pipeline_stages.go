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

// 定义聚合管道阶段
// 参考：https://docs.mongodb.com/manual/reference/operator/aggregation-pipeline/
// md5:40bdaebb510f842e
const (
	// 集合聚合阶段 md5:37e4a0637949107b
	AddFields      = "$addFields"
	Bucket         = "$bucket"
	BucketAuto     = "$bucketAuto"
	CollStats      = "$collStats"
	X聚合计数          = "$count"
	Facet          = "$facet"
	GeoNear        = "$geoNear"
	GraphLookup    = "$graphLookup"
	X聚合分组          = "$group"
	IndexStats     = "$indexStats"
	X聚合限制数量          = "$limit"
	ListSessions   = "$listSessions"
	X聚合关联集合         = "$lookup"
	X聚合条件          = "$match"
	X聚合合并          = "$merge"
	Out            = "$out"
	PlanCacheStats = "$planCacheStats"
	X聚合字段        = "$project"
	Redact         = "$redact"
	ReplaceRoot    = "$replaceRoot"
	ReplaceWith    = "$replaceWith"
	Sample         = "$sample"
	X聚合跳过           = "$skip"
	SortByCount    = "$sortByCount"
	UnionWith      = "$unionWith"
	Unwind         = "$unwind"

	// 数据库聚合阶段 md5:b35520d72e009304
	CurrentOp         = "$currentOp"
	ListLocalSessions = "$listLocalSessions"
)
