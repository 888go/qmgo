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

// X数据库 is a handle to a MongoDB database
type X数据库 struct {
	database *mongo.Database

	registry *bsoncodec.Registry
}

// X取集合 gets collection from database
func (d *X数据库) X取集合(名称 string, 可选选项 ...*options.CollectionOptions) *X文档集合 {
	var cp *mongo.Collection
	var opt = make([]*officialOpts.CollectionOptions, 0, len(可选选项))
	for _, o := range 可选选项 {
		opt = append(opt, o.CollectionOptions)
	}
	collOpt := officialOpts.MergeCollectionOptions(opt...)
	cp = d.database.Collection(名称, collOpt)

	return &X文档集合{
		collection: cp,
		registry:   d.registry,
	}
}

// X取数据库名称 returns the name of database
func (d *X数据库) X取数据库名称() string {
	return d.database.Name()
}

// X删除数据库 drops database
func (d *X数据库) X删除数据库(上下文 context.Context) error {
	return d.database.Drop(上下文)
}

// X执行命令 executes the given command against the database.
//
// The runCommand parameter must be a document for the command to be executed. It cannot be nil.
// This must be an order-preserving type such as bson.D. Map types such as bson.M are not valid.
// If the command document contains a session ID or any transaction-specific fields, the behavior is undefined.
//
// The opts parameter can be used to specify options for this operation (see the options.RunCmdOptions documentation).
func (d *X数据库) X执行命令(上下文 context.Context, runCommand interface{}, 可选选项 ...options.RunCommandOptions) *mongo.SingleResult {
	option := officialOpts.RunCmd()
	if len(可选选项) > 0 && 可选选项[0].RunCmdOptions != nil {
		option = 可选选项[0].RunCmdOptions
	}
	return d.database.RunCommand(上下文, runCommand, option)
}

// X创建集合 executes a create command to explicitly create a new collection with the specified name on the
// server. If the collection being created already exists, this method will return a mongo.CommandError. This method
// requires driver version 1.4.0 or higher.
//
// The opts parameter can be used to specify options for the operation (see the options.CreateCollectionOptions
// documentation).
func (db *X数据库) X创建集合(上下文 context.Context, 集合名称 string, 可选选项 ...options.CreateCollectionOptions) error {
	var option = make([]*officialOpts.CreateCollectionOptions, 0, len(可选选项))
	for _, opt := range 可选选项 {
		if opt.CreateCollectionOptions != nil {
			option = append(option, opt.CreateCollectionOptions)
		}
	}
	return db.database.CreateCollection(上下文, 集合名称, option...)
}
