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

func initTransactionClient(coll string) *QmgoClient {
	cfg := Config{
		Uri:      "mongodb://localhost:27017",
		Database: "transaction",
		Coll:     coll,
	}
	var cTimeout int64 = 0
	var sTimeout int64 = 500000
	var maxPoolSize uint64 = 30000
	var minPoolSize uint64 = 0
	cfg.ConnectTimeoutMS = &cTimeout
	cfg.SocketTimeoutMS = &sTimeout
	cfg.MaxPoolSize = &maxPoolSize
	cfg.MinPoolSize = &minPoolSize
	cfg.ReadPreference = &ReadPref{Mode: readpref.PrimaryMode}
	qClient, err := X创建(context.Background(), &cfg)
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
	s, err := cli.X创建Session(o)
	ast.NoError(err)
	ctx := context.Background()
	defer s.EndSession(ctx)

	callback := func(sCtx context.Context) (interface{}, error) {
		if _, err := cli.X插入(sCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		err = s.AbortTransaction(sCtx)

		return nil, nil
	}

	_, err = s.StartTransaction(ctx, callback)
	ast.NoError(err)

	r := bson.M{}
	err = cli.X查询(ctx, bson.M{"abc": 1}).X取一条(&r)
	ast.Error(err)
// 中止已执行的操作，无法中止后续操作
// 看起来像是mongodb-go-driver驱动的一个bug
	err = cli.X查询(ctx, bson.M{"xyz": 999}).X取一条(&r)
	ast.Error(err)
}

func TestSession_Cancel(t *testing.T) {
	ast := require.New(t)
	cli := initTransactionClient("test")

	defer cli.X删除集合(context.Background())
	s, err := cli.X创建Session()
	ast.NoError(err)
	ctx := context.Background()
	defer s.EndSession(ctx)

	callback := func(sCtx context.Context) (interface{}, error) {
		if _, err := cli.X插入(sCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		return nil, errors.New("cancel operations")
	}
	_, err = s.StartTransaction(ctx, callback)
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
	s, err := cli.X创建Session()
	ast.NoError(err)
	ctx := context.Background()
	defer s.EndSession(ctx)

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
			return nil, ErrTransactionRetry
		}
		return nil, nil
	}
	_, err = s.StartTransaction(ctx, callback)
	ast.NoError(err)
	r := bson.M{}
	cli.X查询(ctx, bson.M{"abc": 1}).X取一条(&r)
	ast.Equal(r["abc"], int32(1))
	cli.X查询(ctx, bson.M{"xyz": 999}).X取一条(&r)
	ast.Equal(r["xyz"], int32(999))
	ast.Equal(count, 1)
}
