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

	"github.com/888go/qmgo/options"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"
)

// Database 是一个指向 MongoDB 数据库的句柄 md5:9217ae5bd9047e3a
type Database struct {
	database *mongo.Database

	registry *bsoncodec.Registry
}

// Collection 从数据库中获取集合 md5:c5489f5523d5a33d
func (d *Database) X取集合(名称 string, 可选选项 ...*options.CollectionOptions) *Collection {
	var cp *mongo.Collection
	var opt = make([]*officialOpts.CollectionOptions, 0, len(可选选项))
	for _, o := range 可选选项 {
		opt = append(opt, o.CollectionOptions)
	}
	collOpt := officialOpts.MergeCollectionOptions(opt...)
	cp = d.database.Collection(名称, collOpt)

	return &Collection{
		collection: cp,
		registry:   d.registry,
	}
}

// GetDatabaseName 返回数据库的名称 md5:716064a488e6db8b
func (d *Database) X取数据库名称() string {
	return d.database.Name()
}

// DropDatabase 删除数据库 md5:aeac2378daa25d5f
func (d *Database) X删除数据库(上下文 context.Context) error {
	return d.database.Drop(上下文)
}

// RunCommand 在数据库上执行给定的命令。
//
// runCommand 参数必须是将要执行的命令文档。它不能为 nil。这必须是一个保持顺序的类型，如 bson.D。像 bson.M 这样的映射类型是无效的。
// 如果命令文档包含会话 ID 或任何事务特定字段，其行为是未定义的。
//
// 可以使用 opts 参数来指定此操作的选项（参阅 options.RunCmdOptions 的文档）。
// md5:eb93f7217a15650c
func (d *Database) X执行命令(上下文 context.Context, runCommand interface{}, 可选选项 ...options.RunCommandOptions) *mongo.SingleResult {
	option := officialOpts.RunCmd()
	if len(可选选项) > 0 && 可选选项[0].RunCmdOptions != nil {
		option = 可选选项[0].RunCmdOptions
	}
	return d.database.RunCommand(上下文, runCommand, option)
}

// CreateCollection 执行一个创建命令，明确在服务器上使用指定名称创建一个新的集合。如果正在创建的集合已经存在，此方法将返回一个 mongo.CommandError。此方法需要驱动程序版本 1.4.0 或更高版本。
// 
// 参数 opts 可用于指定操作选项（请参阅 options.CreateCollectionOptions 的文档）。
// md5:7bd165db4ed05d28
func (db *Database) X创建集合(上下文 context.Context, 集合名称 string, 可选选项 ...options.CreateCollectionOptions) error {
	var option = make([]*officialOpts.CreateCollectionOptions, 0, len(可选选项))
	for _, opt := range 可选选项 {
		if opt.CreateCollectionOptions != nil {
			option = append(option, opt.CreateCollectionOptions)
		}
	}
	return db.database.CreateCollection(上下文, 集合名称, option...)
}
