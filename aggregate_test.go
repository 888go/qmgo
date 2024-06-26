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
	opts "github.com/888go/qmgo/options"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAggregate(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), nil, []string{"name"})

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	id3 := primitive.NewObjectID()
	id4 := primitive.NewObjectID()
	id5 := primitive.NewObjectID()
	docs := []interface{}{
		QueryTestItem{Id: id1, Name: "Alice", Age: 10},
		QueryTestItem{Id: id2, Name: "Alice", Age: 12},
		QueryTestItem{Id: id3, Name: "Lucas", Age: 33},
		QueryTestItem{Id: id4, Name: "Lucas", Age: 22},
		QueryTestItem{Id: id5, Name: "Lucas", Age: 44},
	}
	cli.X插入多个(context.Background(), docs)
	matchStage := bson.D{{"$match", []bson.E{{"age", bson.D{{"$gt", 11}}}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$name"}, {"total", bson.D{{"$sum", "$age"}}}}}}
	var showsWithInfo []bson.M

	opt := opts.AggregateOptions{
		AggregateOptions: options.Aggregate().SetAllowDiskUse(true),
	}
	// aggregate ALL()
	err := cli.X聚合(context.Background(), Pipeline{matchStage, groupStage}, opt).X取全部(&showsWithInfo)
	ast.NoError(err)
	ast.Equal(2, len(showsWithInfo))
	for _, v := range showsWithInfo {
		if "Alice" == v["_id"] {
			ast.Equal(int32(12), v["total"])
			continue
		}
		if "Lucas" == v["_id"] {
			ast.Equal(int32(99), v["total"])
			continue
		}
		ast.Error(errors.New("error"), "impossible")
	}
	// Iter()
	iter := cli.X聚合(context.Background(), Pipeline{matchStage, groupStage})
	ast.NotNil(iter)
	err = iter.X取全部(&showsWithInfo)
	ast.NoError(err)
	for _, v := range showsWithInfo {
		if "Alice" == v["_id"] {
			ast.Equal(int32(12), v["total"])
			continue
		}
		if "Lucas" == v["_id"] {
			ast.Equal(int32(99), v["total"])
			continue
		}
		ast.Error(errors.New("error"), "impossible")
	}
	// One()
	var oneInfo bson.M

	opt = opts.AggregateOptions{
		AggregateOptions: options.Aggregate().SetAllowDiskUse(true),
	}
	iter = cli.X聚合(context.Background(), Pipeline{matchStage, groupStage}, opt)
	ast.NotNil(iter)
	iter = cli.X聚合(context.Background(), Pipeline{matchStage, groupStage})
	ast.NotNil(iter)
	err = iter.X取一条(&oneInfo)
	ast.NoError(err)
	ast.Equal(true, oneInfo["_id"] == "Alice" || oneInfo["_id"] == "Lucas")

	// iter
	iter = cli.X聚合(context.Background(), Pipeline{matchStage, groupStage}, opt)
	ast.NotNil(iter)

	i := iter.Iter弃用()

	ct := i.X下一个(&oneInfo)
	ast.Equal(true, oneInfo["_id"] == "Alice" || oneInfo["_id"] == "Lucas")
	ast.Equal(true, ct)
	ct = i.X下一个(&oneInfo)
	ast.Equal(true, oneInfo["_id"] == "Alice" || oneInfo["_id"] == "Lucas")
	ast.Equal(true, ct)
	ct = i.X下一个(&oneInfo)
	ast.Equal(false, ct)

	// err
	ast.Error(cli.X聚合(context.Background(), 1).X取全部(&showsWithInfo))
	ast.Error(cli.X聚合(context.Background(), 1).X取一条(&showsWithInfo))
	ast.Error(cli.X聚合(context.Background(), 1).Iter弃用().X取错误())
	matchStage = bson.D{{"$match", []bson.E{{"age", bson.D{{"$gt", 100}}}}}}
	groupStage = bson.D{{"$group", bson.D{{"_id", "$name"}, {"total", bson.D{{"$sum", "$age"}}}}}}
	ast.Error(cli.X聚合(context.Background(), Pipeline{matchStage, groupStage}).X取一条(&showsWithInfo))

}
