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
	// ErrQueryNotSlicePointer 如果结果参数不是一个切片的指针，返回此错误 md5:99bf2bfe686c166d
	X错误_结果参数_必须切片指针 = errors.New("result argument must be a pointer to a slice")
	// ErrQueryNotSliceType 如果结果参数不是切片的地址时返回错误 md5:a70365d8a017acd7
	X错误_结果参数_必须切片地址 = errors.New("result argument must be a slice address")
	// ErrQueryResultTypeInconsistent 如果查询结果类型与MongoDB值类型不一致时返回 md5:940d09b0f6052a9f
	X错误_查询结果_类型不一致 = errors.New("result type is not equal mongodb value type")
	// ErrQueryResultValCanNotChange 如果结果值不能更改，返回这个错误 md5:95f73ebb72c4463a
	X错误_结果值不能更改 = errors.New("the value of result can not be changed")
	// ErrNoSuchDocuments 如果未找到文档，则返回此错误 md5:381e7dce5c56bc42
	X错误_未找到文档 = mongo.ErrNoDocuments
	// ErrTransactionRetry 如果事务需要重试时返回该错误 md5:82e274f71f7c0175
	X错误_事务_重试 = errors.New("retry transaction")
	// ErrTransactionNotSupported 如果不支持事务，则返回该错误 md5:97ad2e1fbe2e7207
	X错误_事务_不支持 = errors.New("transaction not supported")
	// ErrNotSupportedUsername 如果用户名无效，则返回此错误 md5:c9df5d462362cad6
	X错误_不支持用户名 = errors.New("username not supported")
	// ErrNotSupportedPassword 如果密码无效时返回 md5:e9df4f7f8304cc70
	X错误_不支持密码 = errors.New("password not supported")
	// ErrNotValidSliceToInsert 如果插入的参数不是一个有效的切片，返回该错误 md5:2b940e5853972183
	X错误_插入_无效切片 = errors.New("must be valid slice to insert")
	// ErrReplacementContainUpdateOperators 如果替换文档中包含更新操作符，返回错误 md5:4f127578930f91fa
	X错误_替换_文档含更新操作符 = errors.New("replacement document cannot contain keys beginning with '$'")
)

// IsErrNoDocuments 检查错误是否表示没有找到文档，既包括 mongo-go-driver 的错误也包括 qmgo 自定义的错误
// 已弃用，直接判断 err == ErrNoSuchDocuments 或者 err == mongo.ErrNoDocuments 即可
// md5:a9bcbf0c80c5509c
func X是否为无文档错误(错误 error) bool {
	if 错误 == X错误_未找到文档 {
		return true
	}
	return false
}

// IsDup 检查错误是否为MongoDB的E11000（重复错误）。 md5:4a3435e9efa67970
func X是否为重复键错误(错误 error) bool {
	return 错误 != nil && strings.Contains(错误.Error(), "E11000")
}
