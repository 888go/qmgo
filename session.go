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
	"context"
	
	opts "github.com/888go/qmgo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
)

// Session 是一个结构体，代表了 MongoDB 的逻辑会话
type Session struct {
	session mongo.Session
}

// StartTransaction starts transaction
//precondition：
//- version of mongoDB server >= v4.0
//- Topology of mongoDB server is not Single
//At the same time, please pay attention to the following
//- make sure all operations in callback use the sessCtx as context parameter
//- Dont forget to call EndSession if session is not used anymore
//- if operations in callback takes more than(include equal) 120s, the operations will not take effect,
//- if operation in callback return qmgo.ErrTransactionRetry,
//  the whole transaction will retry, so this transaction must be idempotent
//- if operations in callback return qmgo.ErrTransactionNotSupported,
//- If the ctx parameter already has a Session attached to it, it will be replaced by this session.
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

// EndSession 将终止任何现有事务并关闭会话。
func (s *Session) EndSession(ctx context.Context) {
	s.session.EndSession(ctx)
}

// AbortTransaction 中断当前会话的活动事务。如果当前会话没有活动事务，或者事务已被提交或中止，则此方法将返回错误。
func (s *Session) AbortTransaction(ctx context.Context) error {
	return s.session.AbortTransaction(ctx)
}

// wrapperCustomF 封装调用者的回调函数到Mongo驱动的
func wrapperCustomCb(cb func(ctx context.Context) (interface{}, error)) func(sessCtx mongo.SessionContext) (interface{}, error) {
	return func(sessCtx mongo.SessionContext) (interface{}, error) {
		result, err := cb(sessCtx)
		if err == ErrTransactionRetry {
			return nil, mongo.CommandError{Labels: []string{driver.TransientTransactionError}}
		}
		return result, err
	}
}
