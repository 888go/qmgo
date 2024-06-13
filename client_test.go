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
	"fmt"
	"testing"

	"github.com/888go/qmgo/options"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initClient(col string) *XMongo客户端 {
	cfg := X配置{
		X连接URI:      "mongodb://localhost:27017",
		X数据库名: "qmgotest",
		X集合名:     col,
	}
	var cTimeout int64 = 0
	var sTimeout int64 = 500000
	var maxPoolSize uint64 = 30000
	var minPoolSize uint64 = 0
	cfg.X连接超时毫秒 = &cTimeout
	cfg.X身份凭证 = &sTimeout
	cfg.X最大连接池大小 = &maxPoolSize
	cfg.X最小连接池大小 = &minPoolSize
	cfg.ReadPreference = &X读取偏好{Mode: readpref.PrimaryMode}
	qClient, err := X连接(context.Background(), &cfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return qClient
}

func TestQmgoClient(t *testing.T) {
	ast := require.New(t)
	var timeout int64 = 50

	// uri 错误
	cfg := X配置{
		X连接URI:              "://127.0.0.1",
		X连接超时毫秒: &timeout,
	}

	var err error
	_, err = X连接(context.Background(), &cfg)
	ast.NotNil(err)

	// Open 成功
	var maxPoolSize uint64 = 100
	var minPoolSize uint64 = 0

	cfg = X配置{
		X连接URI:              "mongodb://localhost:27017",
		X数据库名:         "qmgotest",
		X集合名:             "testopen",
		X连接超时毫秒: &timeout,
		X最大连接池大小:      &maxPoolSize,
		X最小连接池大小:      &minPoolSize,
		ReadPreference:   &X读取偏好{Mode: readpref.SecondaryMode, X最大延迟毫秒: 500},
	}

	cli, err := X连接(context.Background(), &cfg)
	ast.NoError(err)
	ast.Equal(cli.X取数据库名称(), "qmgotest")
	ast.Equal(cli.X取集合名(), "testopen")

	err = cli.X是否存活(5)
	ast.NoError(err)

	res, err := cli.X插入(context.Background(), bson.D{{Key: "x", Value: 1}})
	ast.NoError(err)
	ast.NotNil(res)

	cli.X删除集合(context.Background())

	// close Client
	cli.X关闭连接(context.TODO())
	_, err = cli.X插入(context.Background(), bson.D{{Key: "x", Value: 1}})
	ast.EqualError(err, "client is disconnected")

	err = cli.X是否存活(5)
	ast.Error(err)

	// primary mode with max stalenessMS, error
	cfg = X配置{
		X连接URI:              "mongodb://localhost:27017",
		X数据库名:         "qmgotest",
		X集合名:             "testopen",
		X连接超时毫秒: &timeout,
		X最大连接池大小:      &maxPoolSize,
		ReadPreference:   &X读取偏好{Mode: readpref.PrimaryMode, X最大延迟毫秒: 500},
	}

	cli, err = X连接(context.Background(), &cfg)
	ast.Error(err)
}

func TestClient(t *testing.T) {
	ast := require.New(t)

	var maxPoolSize uint64 = 100
	var minPoolSize uint64 = 0
	var timeout int64 = 50

	cfg := &X配置{
		X连接URI:              "mongodb://localhost:27017",
		X连接超时毫秒: &timeout,
		X最大连接池大小:      &maxPoolSize,
		X最小连接池大小:      &minPoolSize,
	}

	c, err := X创建客户端(context.Background(), cfg)
	ast.Equal(nil, err)

	opts := &options.DatabaseOptions{DatabaseOptions: officialOpts.Database().SetReadPreference(readpref.PrimaryPreferred())}
	cOpts := &options.CollectionOptions{CollectionOptions: officialOpts.Collection().SetReadPreference(readpref.PrimaryPreferred())}
	coll := c.X设置数据库("qmgotest", opts).X取集合("testopen", cOpts)

	res, err := coll.X插入(context.Background(), bson.D{{Key: "x", Value: 1}})
	ast.NoError(err)
	ast.NotNil(res)
	coll.X删除集合(context.Background())
}

func TestClient_ServerVersion(t *testing.T) {
	ast := require.New(t)

	cfg := &X配置{
		X连接URI:      "mongodb://localhost:27017",
		X数据库名: "qmgotest",
		X集合名:     "transaction",
	}

	ctx := context.Background()
	cli, err := X连接(ctx, cfg)
	ast.NoError(err)

	version := cli.X取版本号()
	ast.NotEmpty(version)
	fmt.Println(version)
}

func TestClient_newAuth(t *testing.T) {
	ast := require.New(t)

	auth := X身份凭证{
		X认证机制: "PLAIN",
		X认证源:    "PLAIN",
		X用户名:      "qmgo",
		X密码:      "123",
		PasswordSet:   false,
	}
	cred, err := newAuth(auth)
	ast.NoError(err)
	ast.Equal(auth.PasswordSet, cred.PasswordSet)
	ast.Equal(auth.X认证源, cred.AuthSource)
	ast.Equal(auth.X认证机制, cred.AuthMechanism)
	ast.Equal(auth.X用户名, cred.Username)
	ast.Equal(auth.X密码, cred.Password)

	auth = X身份凭证{
		X认证机制: "PLAIN",
		X认证源:    "PLAIN",
		X用户名:      "qmg/o",
		X密码:      "123",
		PasswordSet:   false,
	}
	_, err = newAuth(auth)
	ast.Equal(X错误_不支持用户名, err)

	auth = X身份凭证{
		X认证机制: "PLAIN",
		X认证源:    "PLAIN",
		X用户名:      "qmgo",
		X密码:      "12:3",
		PasswordSet:   false,
	}
	_, err = newAuth(auth)
	ast.Equal(X错误_不支持密码, err)

	auth = X身份凭证{
		X认证机制: "PLAIN",
		X认证源:    "PLAIN",
		X用户名:      "qmgo",
		X密码:      "1/23",
		PasswordSet:   false,
	}
	_, err = newAuth(auth)
	ast.Equal(X错误_不支持密码, err)

	auth = X身份凭证{
		X认证机制: "PLAIN",
		X认证源:    "PLAIN",
		X用户名:      "qmgo",
		X密码:      "1%3",
		PasswordSet:   false,
	}
	_, err = newAuth(auth)
	ast.Equal(X错误_不支持密码, err)

	auth = X身份凭证{
		X认证机制: "PLAIN",
		X认证源:    "PLAIN",
		X用户名:      "q%3mgo",
		X密码:      "13",
		PasswordSet:   false,
	}
	_, err = newAuth(auth)
	ast.Equal(X错误_不支持用户名, err)
}
