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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	
	opts "github.com/888go/qmgo/options"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDatabase(t *testing.T) {
	ast := require.New(t)

	var sTimeout int64 = 500000
	var cTimeout int64 = 3000
	var maxPoolSize uint64 = 3000
	var minPoolSize uint64 = 0
	collName := "testopen"
	dbName := "qmgotest"

	cfg := Config{
		Uri:              "mongodb://localhost:27017",
		Database:         dbName,
		Coll:             collName,
		ConnectTimeoutMS: &cTimeout,
		SocketTimeoutMS:  &sTimeout,
		MaxPoolSize:      &maxPoolSize,
		MinPoolSize:      &minPoolSize,
	}

	c, err := X创建客户端(context.Background(), &cfg)
	ast.NoError(err)
	cli := c.X设置数据库(cfg.Database)
	ast.Nil(err)
	ast.Equal(dbName, cli.X取数据库名称())
	coll := cli.X取集合(collName)
	ast.Equal(collName, coll.X取集合名())
	cli.X取集合(collName).X删除集合(context.Background())
	cli.X删除数据库(context.Background())

}

func TestRunCommand(t *testing.T) {
	ast := require.New(t)

	cli := initClient("test")

	opts := opts.RunCommandOptions{RunCmdOptions: options.RunCmd().SetReadPreference(readpref.Primary())}
	res := cli.X执行命令(context.Background(), bson.D{
		{"ping", 1}}, opts)
	ast.NoError(res.Err())
}

// TestCreateCollection 是一个测试函数，用于测试创建时间序列集合的功能。
// 参数 t *testing.T 为 Golang 中的标准测试对象，用于断言和报告测试结果。
// func TestCreateCollection(t *testing.T) {
	// ast := require.New(t) 创建一个新的断言工具对象，方便进行错误判断。
	// initClient("test") 初始化一个客户端，连接到名为 "test" 的数据库或服务。
// 	cli := initClient("test")
	// 定义 TimeSeriesOptions 结构体实例，设置时间字段为 "timestamp"
// 	timeSeriesOpt := options.TimeSeriesOptions{
// 		TimeField: "timestamp",
// 	}
	// 设置元数据字段为 "metadata"
// 	timeSeriesOpt.SetMetaField("metadata")
	// 创建一个空的上下文对象 ctx，用于执行后续操作。
// 	ctx := context.Background()
	// 创建 CreateCollectionOptions 实例，并设置其中的时间序列选项为上面定义的 timeSeriesOpt。
// 	createCollectionOpts := opts.CreateCollectionOptions{
// 		CreateCollectionOptions: options.CreateCollection().SetTimeSeriesOptions(&timeSeriesOpt),
// 	}
	// 使用 cli 客户端尝试创建名为 "syslog" 的集合，并传入 createCollectionOpts 配置选项。
	// 如果创建过程中出现错误，则通过断言工具判断 err 是否为 nil，如果不是则测试失败。
// 	if err := cli.CreateCollection(ctx, "syslog", createCollectionOpts); err != nil {
// 		ast.NoError(err)
// 	}
	// 删除名为 "syslog" 的集合。
// 	cli.DropCollection(ctx)
	// 删除当前使用的数据库（可能在 initClient 函数中指定）。
// 	cli.DropDatabase(ctx)
// }
