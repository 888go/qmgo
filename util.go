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

package qmgo

import (
	"math"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 现在返回当前毫秒时间 md5:72f6986042eee15a
// ff:取当前时间
// [提示:] func 现在() time.Time {}
func Now() time.Time {
	return time.Unix(0, time.Now().UnixNano()/1e6*1e6)
}

// NewObjectID 生成一个新的 ObjectID。
// 注意：它生成 ObjectID 的方式与 mgo 不同。
// md5:1225b5a9fc54eeee
// ff:生成对象ID
// [提示:] func 新建对象ID() primitive.ObjectID {}
func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// SplitSortField 处理排序符号：字段前的"+"或"-"
// 如果是"+"，返回排序为1
// 如果是"-"，返回排序为-1
// md5:184424f8495a3828
// ff:
// field:
// key:
// sort:
// [提示:] func 分割排序字段(field 字符串) (键 string, 排序方向 int32) {}
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

// CompareVersions 比较两个版本号字符串（即由点分隔的正整数）。比较是在两个版本中较低精度的基础上进行的。例如，3.2 被认为等于 3.2.11，而 3.2.0 则被认为小于 3.2.11。
//
// 如果 version1 大于 version2，返回一个正整数；如果 version1 小于 version2，返回一个负整数；如果 version1 等于 version2，则返回 0。
// md5:4b6863e7e41f1627
// ff:比较版本号
// v1:版本号1
// v2:版本号2
// [提示:] func 比较版本号(v1 string, v2 string) (int, error) {}
func CompareVersions(v1 string, v2 string) (int, error) {
	n1 := strings.Split(v1, ".")
	n2 := strings.Split(v2, ".")

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
