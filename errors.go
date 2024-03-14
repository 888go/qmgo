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
	// ErrQueryNotSlicePointer 当查询结果参数不是指向切片的指针时返回该错误
	ErrQueryNotSlicePointer = errors.New("result argument must be a pointer to a slice")
	// ErrQueryNotSliceType 当结果参数不是切片地址时返回该错误
	ErrQueryNotSliceType = errors.New("result argument must be a slice address")
	// ErrQueryResultTypeInconsistent 当查询结果类型与mongodb的值类型不相等时返回该错误
	ErrQueryResultTypeInconsistent = errors.New("result type is not equal mongodb value type")
	// ErrQueryResultValCanNotChange 当查询结果的值不能被更改时返回该错误
	ErrQueryResultValCanNotChange = errors.New("the value of result can not be changed")
	// ErrNoSuchDocuments 当没有找到任何文档时返回这个错误
	ErrNoSuchDocuments = mongo.ErrNoDocuments
	// ErrTransactionRetry：如果事务需要重试则返回该错误
	ErrTransactionRetry = errors.New("retry transaction")
	// ErrTransactionNotSupported 当事务不被支持时返回该错误
	ErrTransactionNotSupported = errors.New("transaction not supported")
	// ErrNotSupportedUsername 当用户名无效时返回该错误
	ErrNotSupportedUsername = errors.New("username not supported")
	// ErrNotSupportedPassword 当密码无效时返回该错误
	ErrNotSupportedPassword = errors.New("password not supported")
	// ErrNotValidSliceToInsert 当插入参数不是一个有效的切片时返回该错误
	ErrNotValidSliceToInsert = errors.New("must be valid slice to insert")
	// ErrReplacementContainUpdateOperators 当替换文档包含更新操作符时返回错误
	ErrReplacementContainUpdateOperators = errors.New("replacement document cannot contain keys beginning with '$'")
)

// IsErrNoDocuments 检查 err 是否表示没有文档，包括mongo-go-driver库的错误和qmgo自定义错误
// 已弃用，直接调用 if err == ErrNoSuchDocuments 或 if err == mongo.ErrNoDocuments 即可
func X是否为无文档错误(错误 error) bool {
	if 错误 == ErrNoSuchDocuments {
		return true
	}
	return false
}

// IsDup check if err is mongo E11000 (duplicate err)。
func X是否为重复键错误(错误 error) bool {
	return 错误 != nil && strings.Contains(错误.Error(), "E11000")
}
