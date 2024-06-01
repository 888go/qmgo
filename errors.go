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
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

// [提示]
// var (
//
//	查询结果非切片指针错误 = errors.New("结果参数必须是一个指向切片的指针")
//
//	查询结果非切片类型错误 = errors.New("结果参数必须是切片的地址")
//
//	查询结果类型不一致错误 = errors.New("结果类型与mongodb值类型不一致")
//
//	查询结果值不可更改错误 = errors.New("结果值不能被改变")
//
//	无此类文档错误 = mongo.ErrNoDocuments
//
//	事务重试错误 = errors.New("事务重试")
//
//	事务不支持错误 = errors.New("事务不被支持")
//
//	用户名不支持错误 = errors.New("用户名不被支持")
//
//	密码不支持错误 = errors.New("密码不被支持")
//
//	无效插入切片错误 = errors.New("必须是有效的切片才能插入")
//
//	替换文档包含更新操作符错误 = errors.New("替换文档不能包含以'$'开头的键")
//
// )
// [结束]
var (
	// ErrQueryNotSlicePointer 如果结果参数不是一个切片的指针，返回此错误 md5:99bf2bfe686c166d
	ErrQueryNotSlicePointer = errors.New("result argument must be a pointer to a slice")
	// ErrQueryNotSliceType 如果结果参数不是切片的地址时返回错误 md5:a70365d8a017acd7
	ErrQueryNotSliceType = errors.New("result argument must be a slice address")
	// ErrQueryResultTypeInconsistent 如果查询结果类型与MongoDB值类型不一致时返回 md5:940d09b0f6052a9f
	ErrQueryResultTypeInconsistent = errors.New("result type is not equal mongodb value type")
	// ErrQueryResultValCanNotChange 如果结果值不能更改，返回这个错误 md5:95f73ebb72c4463a
	ErrQueryResultValCanNotChange = errors.New("the value of result can not be changed")
	// ErrNoSuchDocuments 如果未找到文档，则返回此错误 md5:381e7dce5c56bc42
	ErrNoSuchDocuments = mongo.ErrNoDocuments
	// ErrTransactionRetry 如果事务需要重试时返回该错误 md5:82e274f71f7c0175
	ErrTransactionRetry = errors.New("retry transaction")
	// ErrTransactionNotSupported 如果不支持事务，则返回该错误 md5:97ad2e1fbe2e7207
	ErrTransactionNotSupported = errors.New("transaction not supported")
	// ErrNotSupportedUsername 如果用户名无效，则返回此错误 md5:c9df5d462362cad6
	ErrNotSupportedUsername = errors.New("username not supported")
	// ErrNotSupportedPassword 如果密码无效时返回 md5:e9df4f7f8304cc70
	ErrNotSupportedPassword = errors.New("password not supported")
	// ErrNotValidSliceToInsert 如果插入的参数不是一个有效的切片，返回该错误 md5:2b940e5853972183
	ErrNotValidSliceToInsert = errors.New("must be valid slice to insert")
	// ErrReplacementContainUpdateOperators 如果替换文档中包含更新操作符，返回错误 md5:4f127578930f91fa
	ErrReplacementContainUpdateOperators = errors.New("replacement document cannot contain keys beginning with '$'")
)

// IsErrNoDocuments 检查错误是否表示没有找到文档，既包括 mongo-go-driver 的错误也包括 qmgo 自定义的错误
// 已弃用，直接判断 err == ErrNoSuchDocuments 或者 err == mongo.ErrNoDocuments 即可
// md5:a9bcbf0c80c5509c
// ff:是否为无文档错误
// err:错误
// [提示:] func 是否无文档错误(err error) bool {}
func IsErrNoDocuments(err error) bool {
	if err == ErrNoSuchDocuments {
		return true
	}
	return false
}

// IsDup 检查错误是否为MongoDB的E11000（重复错误）。 md5:4a3435e9efa67970
// ff:是否为重复键错误
// err:错误
// [提示:] func 是否重复(err 错误) bool {}
func IsDup(err error) bool {
	return err != nil && strings.Contains(err.Error(), "E11000")
}
