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

// Database 是一个指向 MongoDB 数据库的句柄 md5:9217ae5bd9047e3a
// [提示]
//type 数据库 struct {
//     数据库实例 *mongo.Database
//     编码注册器 *bsoncodec.Registry
// }
// [结束]
type Database struct {//hm:数据库  cz:type Database  
	database *mongo.Database

	registry *bsoncodec.Registry
}

// Collection 从数据库中获取集合 md5:c5489f5523d5a33d
// [提示:] func (d *数据库) Collection(名称 string, 选项 ...*options.集合选项) *集合 {}
// ff:取集合
// d:
// name:名称
// opts:可选选项
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

// GetDatabaseName 返回数据库的名称 md5:716064a488e6db8b
// [提示:] func (d *数据库) 获取数据库名称() 字符串 {}
// ff:取数据库名称
// d:
func (d *Database) GetDatabaseName() string {
	return d.database.Name()
}

// DropDatabase 删除数据库 md5:aeac2378daa25d5f
// [提示:] func (d *数据库) 删除数据库(ctx 上下文.Context) 错误 {}
// ff:删除数据库
// d:
// ctx:上下文
func (d *Database) DropDatabase(ctx context.Context) error {
	return d.database.Drop(ctx)
}

// RunCommand 在数据库上执行给定的命令。
//
// runCommand 参数必须是将要执行的命令文档。它不能为 nil。这必须是一个保持顺序的类型，如 bson.D。像 bson.M 这样的映射类型是无效的。
// 如果命令文档包含会话 ID 或任何事务特定字段，其行为是未定义的。
//
// 可以使用 opts 参数来指定此操作的选项（参阅 options.RunCmdOptions 的文档）。
// md5:eb93f7217a15650c
// [提示:] func (d *数据库) 执行命令(ctx 上下文 контекст, runCommand 命令结构体)
// ff:执行命令
// d:
// ctx:上下文
// runCommand:
// opts:可选选项
func (d *Database) RunCommand(ctx context.Context, runCommand interface{}, opts ...options.RunCommandOptions) *mongo.SingleResult {
	option := officialOpts.RunCmd()
	if len(opts) > 0 && opts[0].RunCmdOptions != nil {
		option = opts[0].RunCmdOptions
	}
	return d.database.RunCommand(ctx, runCommand, option)
}

// CreateCollection 执行一个创建命令，明确在服务器上使用指定名称创建一个新的集合。如果正在创建的集合已经存在，此方法将返回一个 mongo.CommandError。此方法需要驱动程序版本 1.4.0 或更高版本。
// 
// 参数 opts 可用于指定操作选项（请参阅 options.CreateCollectionOptions 的文档）。
// md5:7bd165db4ed05d28
// [提示:] func (db *数据库) 创建集合(ctx 上下文, 名称 string, 选项 ...options.创建集合选项) error {}
// ff:创建集合
// db:
// ctx:上下文
// name:集合名称
// opts:可选选项
func (db *Database) CreateCollection(ctx context.Context, name string, opts ...options.CreateCollectionOptions) error {
	var option = make([]*officialOpts.CreateCollectionOptions, 0, len(opts))
	for _, opt := range opts {
		if opt.CreateCollectionOptions != nil {
			option = append(option, opt.CreateCollectionOptions)
		}
	}
	return db.database.CreateCollection(ctx, name, option...)
}
