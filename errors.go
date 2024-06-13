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
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// X错误_结果参数_必须切片指针 return if result argument is not a pointer to a slice
	X错误_结果参数_必须切片指针 = errors.New("result argument must be a pointer to a slice")
	// X错误_结果参数_必须切片地址 return if result argument is not slice address
	X错误_结果参数_必须切片地址 = errors.New("result argument must be a slice address")
	// X错误_查询结果_类型不一致 return if result type is not equal mongodb value type
	X错误_查询结果_类型不一致 = errors.New("result type is not equal mongodb value type")
	// X错误_结果值不能更改 return if the value of result can not be changed
	X错误_结果值不能更改 = errors.New("the value of result can not be changed")
	// X错误_未找到文档 return if no document found
	X错误_未找到文档 = mongo.ErrNoDocuments
	// X错误_事务_重试 return if transaction need to retry
	X错误_事务_重试 = errors.New("retry transaction")
	// X错误_事务_不支持 return if transaction not supported
	X错误_事务_不支持 = errors.New("transaction not supported")
	// X错误_不支持用户名 return if username is invalid
	X错误_不支持用户名 = errors.New("username not supported")
	// X错误_不支持密码 return if password is invalid
	X错误_不支持密码 = errors.New("password not supported")
	// X错误_插入_无效切片 return if insert argument is not valid slice
	X错误_插入_无效切片 = errors.New("must be valid slice to insert")
	// X错误_替换_文档含更新操作符 return if replacement document contain update operators
	X错误_替换_文档含更新操作符 = errors.New("replacement document cannot contain keys beginning with '$'")
)

// X是否为无文档错误 check if err is no documents, both mongo-go-driver error and qmgo custom error
// Deprecated, simply call if err == ErrNoSuchDocuments or if err == mongo.ErrNoDocuments
func X是否为无文档错误(错误 error) bool {
	if 错误 == X错误_未找到文档 {
		return true
	}
	return false
}

// X是否为重复键错误 check if err is mongo E11000 (duplicate err)。
func X是否为重复键错误(错误 error) bool {
	return 错误 != nil && strings.Contains(错误.Error(), "E11000")
}
