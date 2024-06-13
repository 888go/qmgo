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
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCursor(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), nil, []string{"name"})

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	id3 := primitive.NewObjectID()
	id4 := primitive.NewObjectID()
	docs := []interface{}{
		bson.M{"_id": id1, "name": "Alice", "age": 18},
		bson.M{"_id": id2, "name": "Alice", "age": 19},
		bson.M{"_id": id3, "name": "Lucas", "age": 20},
		bson.M{"_id": id4, "name": "Lucas", "age": 21},
	}
	_, err := cli.X插入多个(context.Background(), docs)
	ast.NoError(err)

	var res QueryTestItem

	// if query has 1 record，cursor can run Next one time， Next time return false
	filter1 := bson.M{
		"name": "Alice",
	}
	projection1 := bson.M{
		"name": 0,
	}

	cursor := cli.X查询(context.Background(), filter1).X字段(projection1).X排序("age").X设置最大返回数(2).X跳过(1).X取结果集()
	ast.NoError(cursor.X取错误())

	val := cursor.X下一个(&res)
	ast.Equal(true, val)
	ast.Equal(id2, res.Id)

	val = cursor.X下一个(&res)
	ast.Equal(false, val)

	cursor.X关闭()

	// cursor ALL
	cursor = cli.X查询(context.Background(), filter1).X字段(projection1).X排序("age").X设置最大返回数(2).X取结果集()
	ast.NoError(cursor.X取错误())

	var results []QueryTestItem
	cursor.X取全部(&results)
	ast.Equal(2, len(results))
	// can't match record, cursor run Next and return false
	filter2 := bson.M{
		"name": "Lily",
	}

	cursor = cli.X查询(context.Background(), filter2).X取结果集()
	ast.NoError(cursor.X取错误())
	ast.NotNil(cursor)

	res = QueryTestItem{}
	val = cursor.X下一个(&res)
	ast.Equal(false, val)
	ast.Empty(res)

	cursor.X关闭()

	//  1 record，after cursor close，Next return false
	cursor = cli.X查询(context.Background(), filter1).X字段(projection1).X排序("age").X设置最大返回数(2).X跳过(1).X取结果集()
	ast.NoError(cursor.X取错误())
	ast.NotNil(cursor)

	cursor.X关闭()

	ast.Equal(false, cursor.X下一个(&res))
	ast.NoError(cursor.X取错误())

	// generate Cursor with err
	cursor = cli.X查询(context.Background(), 1).X字段(projection1).X排序("age").X设置最大返回数(2).X跳过(1).X取结果集()
	ast.Error(cursor.X取错误())
	//ast.Equal(int64(0), cursor.ID())
	ast.Error(cursor.X取全部(&res))
	ast.Error(cursor.X关闭())
	ast.Equal(false, cursor.X下一个(&res))
}
