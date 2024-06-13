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

// X取当前时间 return Millisecond current time
func X取当前时间() time.Time {
	return time.Unix(0, time.Now().UnixNano()/1e6*1e6)
}

// X生成对象ID generates a new ObjectID.
// Watch out: the way it generates objectID is different from mgo
func X生成对象ID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// X分割排序字段 handle sort symbol: "+"/"-" in front of field
// if "+"， return sort as 1
// if "-"， return sort as -1
func X分割排序字段(文本 string) (名称 string, sort int32) {
	sort = 1
	名称 = 文本

	if len(文本) != 0 {
		switch 文本[0] {
		case '+':
			名称 = strings.TrimPrefix(文本, "+")
			sort = 1
		case '-':
			名称 = strings.TrimPrefix(文本, "-")
			sort = -1
		}
	}

	return 名称, sort
}

// X比较版本号 compares two version number strings (i.e. positive integers separated by
// periods). Comparisons are done to the lesser precision of the two versions. For example, 3.2 is
// considered equal to 3.2.11, whereas 3.2.0 is considered less than 3.2.11.
//
// Returns a positive int if version1 is greater than version2, a negative int if version1 is less
// than version2, and 0 if version1 is equal to version2.
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
