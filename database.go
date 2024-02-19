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

	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"
)

// Database 是一个指向 MongoDB 数据库的句柄
type Database struct {
	database *mongo.Database
	registry *bsoncodec.Registry
}

// 取集合
// name:名称
// opts:可选选项
// Collection 从数据库获取集合
func (d *Database) Collection(name string, opts ...*options.CollectionOptions) *Collection {
	var cp *mongo.Collection
	var opt = make([]*officialOpts.CollectionOptions, 0, len(opts))
	for _, o := range opts {
		opt = append(opt, o.CollectionOptions)
	}
	collOpt := officialOpts.MergeCollectionOptions(opt...)
	cp = d.database.Collection(name, collOpt)

	return &Collection{
		collection: cp,
		registry:   d.registry,
	}
}

// 取数据库名称
// GetDatabaseName 返回数据库名称
func (d *Database) GetDatabaseName() string {
	return d.database.Name()
}

// 删除数据库
// ctx:上下文
// DropDatabase 删除数据库
func (d *Database) DropDatabase(ctx context.Context) error {
	return d.database.Drop(ctx)
}

// 执行命令
// ctx:上下文
// RunCommand:命令
// opts:可选选项
// RunCommand 执行给定的命令针对数据库。
//
// runCommand 参数必须是要执行的命令的文档，不能为 nil。它必须是一个保持顺序的类型，如 bson.D。Map 类型如 bson.M 不是有效的。
// 如果命令文档包含会话 ID 或任何事务特定字段，则行为未定义。
//
// opts 参数可用于为此次操作指定选项（请参阅 options.RunCmdOptions 文档）。
func (d *Database) RunCommand(ctx context.Context, runCommand interface{}, opts ...options.RunCommandOptions) *mongo.SingleResult {
	option := officialOpts.RunCmd()
	if len(opts) > 0 && opts[0].RunCmdOptions != nil {
		option = opts[0].RunCmdOptions
	}
	return d.database.RunCommand(ctx, runCommand, option)
}

// 创建集合
// ctx:上下文
// name:集合名称
// opts:可选选项
// CreateCollection 执行创建命令，用于在服务器上明确创建一个指定名称的新集合。如果要创建的集合已存在，则此方法将返回 mongo.CommandError。该方法需要驱动版本 1.4.0 或更高版本。
//
// opts 参数可用于为操作指定选项（请参阅 options.CreateCollectionOptions 文档）。
func (db *Database) CreateCollection(ctx context.Context, name string, opts ...options.CreateCollectionOptions) error {
	var option = make([]*officialOpts.CreateCollectionOptions, 0, len(opts))
	for _, opt := range opts {
		if opt.CreateCollectionOptions != nil {
			option = append(option, opt.CreateCollectionOptions)
		}
	}
	return db.database.CreateCollection(ctx, name, option...)
}
