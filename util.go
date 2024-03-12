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

package mgo类

import (
	"math"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 现在返回当前时间的毫秒级别
func X取当前时间() time.Time {
	return time.Unix(0, time.Now().UnixNano()/1e6*1e6)
}

// NewObjectID 生成一个新的 ObjectID。
// 注意：它生成 objectID 的方式与 mgo 不同。
func X生成对象ID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// SplitSortField handle sort symbol: "+"/"-" in front of field
// if "+"， return sort as 1
// if "-"， return sort as -1
func SplitSortField(field string) (key string, sort int32) {
	sort = 1
	key = field

	if len(field) != 0 {
		switch field[0] {
		case '+':
			key = strings.TrimPrefix(field, "+")
			sort = 1
		case '-':
			key = strings.TrimPrefix(field, "-")
			sort = -1
		}
	}

	return key, sort
}

// CompareVersions 比较两个版本号字符串（即由点分隔的正整数）。比较操作会以两者中精度较低的那个为准。例如，3.2 被视为等于 3.2.11，而 3.2.0 则被视为小于 3.2.11。
//
// 如果 version1 大于 version2，则返回一个正整数；如果 version1 小于 version2，则返回一个负整数；如果 version1 等于 version2，则返回 0。
func X比较版本号(版本号1 string, 版本号2 string) (int, error) {
	n1 := strings.Split(版本号1, ".")
	n2 := strings.Split(版本号2, ".")

	for i := 0; i < int(math.Min(float64(len(n1)), float64(len(n2)))); i++ {
		i1, err := strconv.Atoi(n1[i])
		if err != nil {
			return 0, err
		}
		i2, err := strconv.Atoi(n2[i])
		if err != nil {
			return 0, err
		}
		difference := i1 - i2
		if difference != 0 {
			return difference, nil
		}
	}

	return 0, nil
}
