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
	"errors"
	"fmt"
	"testing"

	opts "github.com/888go/qmgo/options"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initTransactionClient(coll string) *XMongo客户端 {
	cfg := X配置{
		X连接URI:      "mongodb://localhost:27017",
		X数据库名: "transaction",
		X集合名:     coll,
	}
	var cTimeout int64 = 0
	var sTimeout int64 = 500000
	var maxPoolSize uint64 = 30000
	var minPoolSize uint64 = 0
	cfg.X连接超时毫秒 = &cTimeout
	cfg.X套接字超时毫秒 = &sTimeout
	cfg.X最大连接池大小 = &maxPoolSize
	cfg.X最小连接池大小 = &minPoolSize
	cfg.X读取偏好 = &X读取偏好{Mode: readpref.PrimaryMode}
	qClient, err := X连接(context.Background(), &cfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	qClient.X插入(context.Background(), bson.M{"name": "before_transaction"})
	return qClient

}
func TestClient_DoTransaction(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()
	cli := initTransactionClient("test")
	defer cli.X删除数据库(ctx)

	fn := func(sCtx context.Context) (interface{}, error) {
		if _, err := cli.X插入(sCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		return nil, nil
	}
	tops := options.Transaction()
	op := &opts.TransactionOptions{TransactionOptions: tops}
	_, err := cli.X事务(ctx, fn, op)
	ast.NoError(err)
	r := bson.M{}
	cli.X查询(ctx, bson.M{"abc": 1}).X取一条(&r)
	ast.Equal(r["abc"], int32(1))

	cli.X查询(ctx, bson.M{"xyz": 999}).X取一条(&r)
	ast.Equal(r["xyz"], int32(999))
}

func TestSession_AbortTransaction(t *testing.T) {
	ast := require.New(t)
	cli := initTransactionClient("test")

	defer cli.X删除集合(context.Background())
	sOpts := options.Session().SetSnapshot(false)
	o := &opts.SessionOptions{sOpts}
	s, err := cli.X创建Session事务(o)
	ast.NoError(err)
	ctx := context.Background()
	defer s.X结束Session(ctx)

	callback := func(sCtx context.Context) (interface{}, error) {
		if _, err := cli.X插入(sCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		err = s.X中止事务(sCtx)

		return nil, nil
	}

	_, err = s.X开始事务(ctx, callback)
	ast.NoError(err)

	r := bson.M{}
	err = cli.X查询(ctx, bson.M{"abc": 1}).X取一条(&r)
	ast.Error(err)
	// abort the already worked operation, can't abort the later operation
	// it seems a mongodb-go-driver bug
	err = cli.X查询(ctx, bson.M{"xyz": 999}).X取一条(&r)
	ast.Error(err)
}

func TestSession_Cancel(t *testing.T) {
	ast := require.New(t)
	cli := initTransactionClient("test")

	defer cli.X删除集合(context.Background())
	s, err := cli.X创建Session事务()
	ast.NoError(err)
	ctx := context.Background()
	defer s.X结束Session(ctx)

	callback := func(sCtx context.Context) (interface{}, error) {
		if _, err := cli.X插入(sCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		return nil, errors.New("cancel operations")
	}
	_, err = s.X开始事务(ctx, callback)
	ast.Error(err)
	r := bson.M{}
	err = cli.X查询(ctx, bson.M{"abc": 1}).X取一条(&r)
	ast.True(X是否为无文档错误(err))
	err = cli.X查询(ctx, bson.M{"xyz": 999}).X取一条(&r)
	ast.True(X是否为无文档错误(err))
}

func TestSession_RetryTransAction(t *testing.T) {
	ast := require.New(t)
	cli := initTransactionClient("test")
	defer cli.X删除集合(context.Background())
	s, err := cli.X创建Session事务()
	ast.NoError(err)
	ctx := context.Background()
	defer s.X结束Session(ctx)

	count := 0
	callback := func(sCtx context.Context) (interface{}, error) {
		if _, err := cli.X插入(sCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		if count == 0 {
			count++
			return nil, X错误_事务_重试
		}
		return nil, nil
	}
	_, err = s.X开始事务(ctx, callback)
	ast.NoError(err)
	r := bson.M{}
	cli.X查询(ctx, bson.M{"abc": 1}).X取一条(&r)
	ast.Equal(r["abc"], int32(1))
	cli.X查询(ctx, bson.M{"xyz": 999}).X取一条(&r)
	ast.Equal(r["xyz"], int32(999))
	ast.Equal(count, 1)
}
