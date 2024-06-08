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
	"context"

	opts "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
)

// Session 是一个结构体，表示 MongoDB 的逻辑会话 md5:a17367bc3a251e77
// [提示]
//type 会话 struct {
//     会话接口 mongo.会话
// }
// [结束]
type Session struct {//hm:Session事务  cz:type Session  
	session mongo.Session
}

// StartTransaction 开始一个事务
// 预条件：
// - MongoDB服务器版本大于等于v4.0
// - MongoDB服务器的拓扑结构不是单节点
// 同时需要注意：
//   - 确保回调中的所有操作将sessCtx作为上下文参数
//   - 如果不再使用session，别忘了调用EndSession
//   - 如果回调中的操作耗时超过（包括等于）120秒，这些操作将不会生效
//   - 如果回调中的操作返回qmgo.ErrTransactionRetry错误，
//     整个事务将会重试，因此这个事务必须是幂等的
//   - 如果回调中的操作返回qmgo.ErrTransactionNotSupported错误，
//   - 如果ctx参数中已经附加了一个Session，它将被此session替换。
//
// md5:7a854b4c45212490
// [提示]
//// 开始事务
// func (s *Session) 开始事务(ctx 上下文.Context, cb func(会话上下文 context.Context) (结果 interface{})) 
// [结束]
// ff:开始事务
// s:
// ctx:上下文
// cb:回调函数
// sessCtx:事务上下文
// opts:可选选项
func (s *Session) StartTransaction(ctx context.Context, cb func(sessCtx context.Context) (interface{}, error), opts ...*opts.TransactionOptions) (interface{}, error) {
	transactionOpts := options.Transaction()
	if len(opts) > 0 && opts[0].TransactionOptions != nil {
		transactionOpts = opts[0].TransactionOptions
	}
	result, err := s.session.WithTransaction(ctx, wrapperCustomCb(cb), transactionOpts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EndSession 会终止任何现有的事务并关闭会话。 md5:2ee8849531868b7e
// [提示:] func (s *Session) 结束会话(ctx 上下文Context) {}
// ff:结束Session
// s:
// ctx:上下文
func (s *Session) EndSession(ctx context.Context) {
	s.session.EndSession(ctx)
}

// AbortTransaction 会取消此会话中的活动事务。如果此会话没有活动事务，或者事务已经提交或中止，此方法将返回错误。
// md5:ca9bc056086304f0
// [提示:] func (s *Session) 中止事务(ctx 上下文 контекст) 错误 {}
// ff:中止事务
// s:
// ctx:上下文
func (s *Session) AbortTransaction(ctx context.Context) error {
	return s.session.AbortTransaction(ctx)
}

// wrapperCustomF 将调用者的回调函数包装成mongo驱动所需的函数 md5:8df643188861ec8b
func wrapperCustomCb(cb func(ctx context.Context) (interface{}, error)) func(sessCtx mongo.SessionContext) (interface{}, error) {
	return func(sessCtx mongo.SessionContext) (interface{}, error) {
		result, err := cb(sessCtx)
		if err == ErrTransactionRetry {
			return nil, mongo.CommandError{Labels: []string{driver.TransientTransactionError}}
		}
		return result, err
	}
}
