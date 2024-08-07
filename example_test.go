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
	"testing"

	"github.com/888go/qmgo/operator"
	"github.com/888go/qmgo/options"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	opts "go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URI      = "mongodb://localhost:27017"
	DATABASE = "class"
	COLL     = "user"
)

type UserInfo struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Age    uint16             `bson:"age"`
	Weight uint32             `bson:"weight"`
}

var userInfo = UserInfo{
	Id:     X生成对象ID(),
	Name:   "xm",
	Age:    7,
	Weight: 40,
}

var userInfos = []UserInfo{
	{Id: X生成对象ID(), Name: "a1", Age: 6, Weight: 20},
	{Id: X生成对象ID(), Name: "b2", Age: 6, Weight: 25},
	{Id: X生成对象ID(), Name: "c3", Age: 6, Weight: 30},
	{Id: X生成对象ID(), Name: "d4", Age: 6, Weight: 35},
	{Id: X生成对象ID(), Name: "a1", Age: 7, Weight: 40},
	{Id: X生成对象ID(), Name: "a1", Age: 8, Weight: 45},
}

var poolMonitor = &event.PoolMonitor{
	Event: func(evt *event.PoolEvent) {
		switch evt.Type {
		case event.GetSucceeded:
		case event.ConnectionReturned:
		}
	},
}

func TestQmgo(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()

	// create connect
	opt := opts.Client().SetAppName("example")
	cli, err := X连接(ctx, &Config{X连接URI: URI, X数据库名: DATABASE, X集合名: COLL}, options.ClientOptions{ClientOptions: opt})

	ast.Nil(err)
	defer func() {
		if err = cli.X关闭连接(ctx); err != nil {
			panic(err)
		}
	}()
	defer cli.X删除数据库(ctx)

	cli.EnsureIndexes弃用(ctx, []string{}, []string{"age", "name,weight"})
	// insert one document
	_, err = cli.X插入(ctx, userInfo)
	ast.Nil(err)

	// find one document
	one := UserInfo{}
	err = cli.X查询(ctx, bson.M{"name": userInfo.Name}).X取一条(&one)
	ast.Nil(err)
	ast.Equal(userInfo, one)

	// multiple insert
	_, err = cli.Collection.X插入多个(ctx, userInfos)
	ast.Nil(err)

	// 找到所有、排序并限制 md5:63d2a93384ca2556
	batch := []UserInfo{}
	cli.X查询(ctx, bson.M{"age": 6}).X排序("weight").X设置最大返回数(7).X取全部(&batch)
	ast.Equal(4, len(batch))

	count, err := cli.X查询(ctx, bson.M{"age": 6}).X取数量()
	ast.NoError(err)
	ast.Equal(int64(4), count)

	// aggregate
	matchStage := bson.D{{mgo常量.X聚合条件, []bson.E{{"weight", bson.D{{mgo常量.X条件大于, 30}}}}}}
	groupStage := bson.D{{mgo常量.X聚合分组, bson.D{{"_id", "$name"}, {"total", bson.D{{mgo常量.X求和, "$age"}}}}}}
	var showsWithInfo []bson.M
	err = cli.X聚合(context.Background(), Pipeline{matchStage, groupStage}).X取全部(&showsWithInfo)
	ast.Equal(3, len(showsWithInfo))
	for _, v := range showsWithInfo {
		if "a1" == v["_id"] {
			ast.Equal(int32(15), v["total"])
			continue
		}
		if "d4" == v["_id"] {
			ast.Equal(int32(6), v["total"])
			continue
		}
		ast.Error(errors.New("error"), "impossible")
	}
	// Update one
	err = cli.X更新一条(ctx, bson.M{"name": "d4"}, bson.M{"$set": bson.M{"age": 17}})
	ast.NoError(err)
	cli.X查询(ctx, bson.M{"age": 17}).X取一条(&one)
	ast.Equal("d4", one.Name)
	// UpdateAll
	result, err := cli.X更新(ctx, bson.M{"age": 6}, bson.M{"$set": bson.M{"age": 10}})
	ast.NoError(err)
	count, err = cli.X查询(ctx, bson.M{"age": 10}).X取数量()
	ast.NoError(err)
	ast.Equal(result.X修改数, count)
	// select
	one = UserInfo{}
	err = cli.X查询(ctx, bson.M{"age": 10}).X字段(bson.M{"age": 1}).X取一条(&one)
	ast.NoError(err)
	ast.Equal(10, int(one.Age))
	ast.Equal("", one.Name)
	// remove
	err = cli.X删除一条(ctx, bson.M{"age": 7})
	ast.Nil(err)
}
