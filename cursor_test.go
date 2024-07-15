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

	// 如果查询有一条记录，游标可以运行一次Next，下次返回false md5:ede21f9451dd1feb
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
	// 无法匹配记录，游标运行Next方法并返回false md5:765217e0cad2c295
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

	// 1条记录，当游标关闭后，Next 函数返回 false md5:a3f791b1b606935e
	cursor = cli.X查询(context.Background(), filter1).X字段(projection1).X排序("age").X设置最大返回数(2).X跳过(1).X取结果集()
	ast.NoError(cursor.X取错误())
	ast.NotNil(cursor)

	cursor.X关闭()

	ast.Equal(false, cursor.X下一个(&res))
	ast.NoError(cursor.X取错误())

	// 使用错误生成Cursor md5:aa941bfed7793fe7
	cursor = cli.X查询(context.Background(), 1).X字段(projection1).X排序("age").X设置最大返回数(2).X跳过(1).X取结果集()
	ast.Error(cursor.X取错误())
	// ast.Equal 认为 int64(0) 等于 cursor.ID() 的结果 md5:9676af4589eca183
	ast.Error(cursor.X取全部(&res))
	ast.Error(cursor.X关闭())
	ast.Equal(false, cursor.X下一个(&res))
}
