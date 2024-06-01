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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"

	opts "github.com/qiniu/qmgo/options"
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

	c, err := NewClient(context.Background(), &cfg)
	ast.NoError(err)
	cli := c.Database(cfg.Database)
	ast.Nil(err)
	ast.Equal(dbName, cli.GetDatabaseName())
	coll := cli.Collection(collName)
	ast.Equal(collName, coll.GetCollectionName())
	cli.Collection(collName).DropCollection(context.Background())
	cli.DropDatabase(context.Background())

}

func TestRunCommand(t *testing.T) {
	ast := require.New(t)

	cli := initClient("test")

	opts := opts.RunCommandOptions{RunCmdOptions: options.RunCmd().SetReadPreference(readpref.Primary())}
	res := cli.RunCommand(context.Background(), bson.D{
		{"ping", 1}}, opts)
	ast.NoError(res.Err())
}

// ```go
// 测试创建集合
// func TestCreateCollection(t *testing.T) {
// 初始化断言工具
// 	ast := require.New(t)
// 
// 初始化客户端，连接名为"test"的数据库
// 	cli := initClient("test")
// 
// 设置时间序列选项
// 	timeSeriesOpt := options.TimeSeriesOptions{
// 		TimeField: "timestamp",
// 	}
// 设置元数据字段
// 	timeSeriesOpt.SetMetaField("metadata")
// 
// 创建上下文
// 	ctx := context.Background()
// 创建集合选项，设置时间序列相关选项
// 	createCollectionOpts := opts.CreateCollectionOptions{CreateCollectionOptions: options.CreateCollection().SetTimeSeriesOptions(&timeSeriesOpt)}
// 创建名为"syslog"的集合，检查是否出错
// 	if err := cli.CreateCollection(ctx, "syslog", createCollectionOpts); err != nil {
// 		ast.NoError(err)
// 	}
// 删除集合
// 	cli.DropCollection(ctx)
// 删除数据库
// 	cli.DropDatabase(ctx)
// }
// ```
// 
// 这段代码是一个测试函数，用于测试在ArangoDB中创建带有时间序列选项的集合。它首先初始化测试所需的工具和客户端，然后定义时间序列的配置，接着在上下文中创建集合。如果在创建过程中没有错误，会删除创建的集合和数据库。
// md5:79faec56c35696a6
