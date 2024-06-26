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
	"fmt"
	"testing"
	"time"

	"github.com/888go/qmgo/operator"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryTestItem struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`

	Instock []struct {
		Warehouse string `bson:"warehouse"`
		Qty       int    `bson:"qty"`
	} `bson:"instock"`
}

type QueryTestItem2 struct {
	Class string `bson:"class"`
}

func TestQuery_One(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), nil, []string{"name"})

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	id3 := primitive.NewObjectID()
	docs := []interface{}{
		bson.D{{Key: "_id", Value: id1}, {Key: "name", Value: "Alice"}, {Key: "age", Value: 18}},
		bson.D{{Key: "_id", Value: id2}, {Key: "name", Value: "Alice"}, {Key: "age", Value: 19}},
		bson.D{{Key: "_id", Value: id3}, {Key: "name", Value: "Lucas"}, {Key: "age", Value: 20}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var res QueryTestItem

	filter1 := bson.M{
		"name": "Alice",
	}
	projection1 := bson.M{
		"age": 0,
	}

	err = cli.X查询(context.Background(), filter1).X字段(projection1).X排序("age").X设置最大返回数(1).X跳过(1).X取一条(&res)
	ast.Nil(err)
	ast.Equal(id2, res.Id)
	ast.Equal("Alice", res.Name)

	res = QueryTestItem{}
	filter2 := bson.M{
		"name": "Lily",
	}

	err = cli.X查询(context.Background(), filter2).X取一条(&res)
	ast.Error(err)
	ast.Empty(res)

	// filter 是 bson.M 类型的空映射，表示匹配所有文档并返回一个结果。 md5:5a0dc74674539e4e
	res = QueryTestItem{}
	filter3 := bson.M{}

	err = cli.X查询(context.Background(), filter3).X取一条(&res)
	ast.NoError(err)
	ast.NotEmpty(res)

	// filter is nil，error
	res = QueryTestItem{}
	err = cli.X查询(context.Background(), nil).X取一条(&res)
	ast.Error(err)
	ast.Empty(res)

	// res 为 nil 或者无法解析 md5:970a874db5a3d5c0
	err = cli.X查询(context.Background(), filter1).X取一条(nil)
	ast.Error(err)

	var tv int
	err = cli.X查询(context.Background(), filter1).X取一条(&tv)
	ast.Error(err)

	// res是一个解析的对象，但是bson标签与mongodb记录不一致，没有报告错误，res的数据结构处于初始化状态。 md5:60d100e8fd5c135d
	var tt QueryTestItem2
	err = cli.X查询(context.Background(), filter1).X取一条(&tt)
	ast.NoError(err)
	ast.Empty(tt)
}

func TestQuery_All(t *testing.T) {
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
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var res []QueryTestItem

	filter1 := bson.M{
		"name": "Alice",
	}
	projection1 := bson.M{
		"name": 0,
	}

	err = cli.X查询(context.Background(), filter1).X字段(projection1).X排序("age").X设置最大返回数(2).X跳过(1).X取全部(&res)
	ast.NoError(err)
	ast.Equal(1, len(res))

	res = make([]QueryTestItem, 0)
	filter2 := bson.M{
		"name": "Lily",
	}

	err = cli.X查询(context.Background(), filter2).X取全部(&res)
	ast.NoError(err)
	ast.Empty(res)

	// filter 是 bson.M{}，这意味着匹配所有，会返回集合中的所有记录 md5:c0c66af96a433502
	res = make([]QueryTestItem, 0)
	filter3 := bson.M{}

	err = cli.X查询(context.Background(), filter3).X取全部(&res)
	ast.NoError(err)
	ast.Equal(4, len(res))

	res = make([]QueryTestItem, 0)
	err = cli.X查询(context.Background(), nil).X取全部(&res)
	ast.Error(err)
	ast.Empty(res)

	err = cli.X查询(context.Background(), filter1).X取全部(nil)
	ast.Error(err)

	var tv int
	err = cli.X查询(context.Background(), filter1).X取全部(&tv)
	ast.Error(err)
// res 是一个可解析的对象，但其 bson 标签与 mongodb 记录不一致，且不会报告错误
// 将根据 res 数据结构的 bson 标签映射相应的值，没有值的标签将使用对应类型的默认值
// res 的长度表示过滤条件筛选出的记录数
// md5:fa2c9312a213eab9
	var tt []QueryTestItem2
	err = cli.X查询(context.Background(), filter1).X取全部(&tt)
	ast.NoError(err)
	ast.Equal(2, len(tt))
}

func TestQuery_Count(t *testing.T) {
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
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var cnt int64

	filter1 := bson.M{
		"name": "Alice",
	}

	cnt, err = cli.X查询(context.Background(), filter1).X设置最大返回数(2).X跳过(1).X取数量()
	ast.NoError(err)
	ast.Equal(int64(1), cnt)

	filter2 := bson.M{
		"name": "Lily",
	}

	cnt, err = cli.X查询(context.Background(), filter2).X取数量()
	ast.NoError(err)
	ast.Zero(cnt)

	filter3 := bson.M{}

	cnt, err = cli.X查询(context.Background(), filter3).X取数量()
	ast.NoError(err)
	ast.Equal(int64(4), cnt)

	cnt, err = cli.X查询(context.Background(), nil).X取数量()
	ast.Error(err)
	ast.Zero(cnt)
}

func TestQuery_Skip(t *testing.T) {
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
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var res []QueryTestItem

	// filter 可以匹配记录，跳过一条记录，并返回剩余的记录。 md5:b966e759fac20d97
	filter1 := bson.M{
		"name": "Alice",
	}

	err = cli.X查询(context.Background(), filter1).X跳过(1).X取全部(&res)
	ast.NoError(err)
	ast.Equal(1, len(res))

	// filter 可以匹配记录，跳过的数量大于现有记录的总数时，res 返回空 md5:d4411346be877b9e
	res = make([]QueryTestItem, 0)

	err = cli.X查询(context.Background(), filter1).X跳过(3).X取全部(&res)
	ast.NoError(err)
	ast.Empty(res)

	res = make([]QueryTestItem, 0)

	err = cli.X查询(context.Background(), filter1).X跳过(-3).X取全部(&res)
	ast.Error(err)
	ast.Empty(res)
}

func TestQuery_Limit(t *testing.T) {
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
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var res []QueryTestItem

	filter1 := bson.M{
		"name": "Alice",
	}

	err = cli.X查询(context.Background(), filter1).X设置最大返回数(1).X取全部(&res)
	ast.NoError(err)
	ast.Equal(1, len(res))

	res = make([]QueryTestItem, 0)

	err = cli.X查询(context.Background(), filter1).X设置最大返回数(3).X取全部(&res)
	ast.NoError(err)
	ast.Equal(2, len(res))

	res = make([]QueryTestItem, 0)
	var cursor CursorI

	cursor = cli.X查询(context.Background(), filter1).X设置最大返回数(-2).X取结果集()
	ast.NoError(cursor.X取错误())
	ast.NotNil(cursor)
}

func TestQuery_Sort(t *testing.T) {
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
		bson.M{"_id": id3, "name": "Lucas", "age": 18},
		bson.M{"_id": id4, "name": "Lucas", "age": 19},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var res []QueryTestItem

	// 按升序对单个字段进行排序 md5:cb85098a3b639ea3
	filter1 := bson.M{
		"name": "Alice",
	}

	err = cli.X查询(context.Background(), filter1).X排序("age").X取全部(&res)
	ast.NoError(err)
	ast.Equal(2, len(res))
	ast.Equal(id1, res[0].Id)
	ast.Equal(id2, res[1].Id)

	// 以降序对单个字段进行排序 md5:e53fe948db01b8ef
	err = cli.X查询(context.Background(), filter1).X排序("-age").X取全部(&res)
	ast.NoError(err)
	ast.Equal(2, len(res))
	ast.Equal(id2, res[0].Id)
	ast.Equal(id1, res[1].Id)

	// 以降序对单个字段进行排序 md5:e53fe948db01b8ef
	err = cli.X查询(context.Background(), bson.M{}).X排序("-age", "+name").X取全部(&res)
	ast.NoError(err)
	ast.Equal(4, len(res))
	ast.Equal(id2, res[0].Id)
	ast.Equal(id4, res[1].Id)
	ast.Equal(id1, res[2].Id)
	ast.Equal(id3, res[3].Id)

	// fields is ""，panic
	res = make([]QueryTestItem, 0)
	ast.Panics(func() {
		cli.X查询(context.Background(), filter1).X排序("").X取全部(&res)
	})

	// fields为空，不会引发恐慌或错误（#128） md5:65471cfbb3cddea4
	err = cli.X查询(context.Background(), bson.M{}).X排序().X取全部(&res)
	ast.NoError(err)
	ast.Equal(4, len(res))

}

func TestQuery_Distinct(t *testing.T) {
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
	id6 := primitive.NewObjectID()
	docs := []interface{}{
		bson.M{"_id": id1, "name": "Alice", "age": 18},
		bson.M{"_id": id2, "name": "Alice", "age": 19},
		bson.M{"_id": id3, "name": "Lucas", "age": 20},
		bson.M{"_id": id4, "name": "Lucas", "age": 21},
		bson.M{"_id": id5, "name": "Kitty", "age": 23, "detail": bson.M{"errInfo": "timeout", "extra": "i/o"}},
		bson.M{"_id": id6, "name": "Kitty", "age": "23", "detail": bson.M{"errInfo": "timeout", "extra": "i/o"}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error

	filter1 := bson.M{
		"name": "Lily",
	}
	var res1 []int32

	err = cli.X查询(context.Background(), filter1).X去重("age", &res1)
	ast.NoError(err)
	ast.Equal(0, len(res1))

	filter2 := bson.M{
		"name": "Alice",
	}
	var res2 []int32

	err = cli.X查询(context.Background(), filter2).X去重("age", &res2)
	ast.NoError(err)
	ast.Equal(2, len(res2))

	var res3 []int32

	err = cli.X查询(context.Background(), filter2).X去重("age", res3)
	ast.EqualError(err, X错误_结果参数_必须切片指针.Error())

	var res4 int

	err = cli.X查询(context.Background(), filter2).X去重("age", &res4)
	ast.EqualError(err, X错误_结果参数_必须切片地址.Error())

	var res5 []string

	err = cli.X查询(context.Background(), filter2).X去重("age", &res5)
	ast.EqualError(err, X错误_查询结果_类型不一致.Error())

// 对于不同版本的mongod（如v4.4.0和v4.0.19），行为有所不同：v4.4.0会返回错误，而v4.0.19则可能返回nil
// 不使用res6
// _, err = cli.Find(context.Background(), filter2).Distinct("", &res6)
// 如果err非nil，则打印错误信息：(Location40352) FieldPath不能使用空字符串构建
// 验证res6的长度为0
// md5:db8b4089027d21a0

	var res7 []int32
	filter3 := 1

	err = cli.X查询(context.Background(), filter3).X去重("age", &res7)
	ast.Error(err)
	ast.Equal(0, len(res7))

	var res8 interface{}

	res8 = []string{}
	err = cli.X查询(context.Background(), filter2).X去重("age", &res8)
	ast.NoError(err)
	ast.NotNil(res8)

	res9, ok := res8.(primitive.A)
	ast.Equal(true, ok)
	ast.Len(res9, 2)

	filter4 := bson.M{}
	var res10 []int32
	err = cli.X查询(context.Background(), filter4).X去重("detail", &res10)
	ast.EqualError(err, X错误_查询结果_类型不一致.Error())

	type tmpStruct struct {
		ErrInfo string `bson:"errInfo"`
		Extra   string `bson:"extra"`
	}
	var res11 []tmpStruct
	err = cli.X查询(context.Background(), filter4).X去重("detail", &res11)
	ast.NoError(err)

	type tmpErrStruct struct {
		ErrInfo string    `bson:"errInfo"`
		Extra   time.Time `bson:"extra"`
	}
	var res12 []tmpErrStruct
	err = cli.X查询(context.Background(), filter4).X去重("detail", &res12)
	ast.EqualError(err, X错误_查询结果_类型不一致.Error())

	var res13 []int32
	err = cli.X查询(context.Background(), filter4).X去重("age", &res13)
	ast.EqualError(err, X错误_查询结果_类型不一致.Error())

	var res14 []interface{}
	err = cli.X查询(context.Background(), filter4).X去重("age", &res14)
	ast.NoError(err)
	ast.Len(res14, 6)
	for _, v := range res14 {
		switch v.(type) {
		case int32:
			fmt.Printf("int32 :%d\n", v)
		case string:
			fmt.Printf("string :%s\n", v)
		default:
			fmt.Printf("defalut err: %v %T\n", v, v)
		}
	}
}

func TestQuery_Select(t *testing.T) {
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
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var res QueryTestItem

	filter1 := bson.M{
		"_id": id1,
	}
	projection1 := bson.M{
		"age": 1,
	}

	err = cli.X查询(context.Background(), filter1).X字段(projection1).X取一条(&res)
	ast.NoError(err)
	ast.NotNil(res)
	ast.Equal("", res.Name)
	ast.Equal(18, res.Age)
	ast.Equal(id1, res.Id)

	res = QueryTestItem{}
	projection2 := bson.M{
		"age": 0,
	}

	err = cli.X查询(context.Background(), filter1).X字段(projection2).X取一条(&res)
	ast.NoError(err)
	ast.NotNil(res)
	ast.Equal("Alice", res.Name)
	ast.Equal(0, res.Age)
	ast.Equal(id1, res.Id)

	res = QueryTestItem{}
	projection3 := bson.M{
		"_id": 0,
	}

	err = cli.X查询(context.Background(), filter1).X字段(projection3).X取一条(&res)
	ast.NoError(err)
	ast.NotNil(res)
	ast.Equal("Alice", res.Name)
	ast.Equal(18, res.Age)
	ast.Equal(primitive.NilObjectID, res.Id)
}

func TestQuery_Cursor(t *testing.T) {
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
		bson.D{{"_id", id1}, {"name", "Alice"}, {"age", 18}},
		bson.D{{"_id", id2}, {"name", "Alice"}, {"age", 19}},
		bson.D{{"_id", id3}, {"name", "Lucas"}, {"age", 20}},
		bson.D{{"_id", id4}, {"name", "Lucas"}, {"age", 21}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var res QueryTestItem

	filter1 := bson.M{
		"name": "Alice",
	}
	projection1 := bson.M{
		"name": 0,
	}

	cursor := cli.X查询(context.Background(), filter1).X字段(projection1).X排序("age").X设置最大返回数(2).X跳过(1).X取结果集()
	ast.NoError(cursor.X取错误())
	ast.NotNil(cursor)

	val := cursor.X下一个(&res)
	ast.Equal(true, val)
	ast.Equal(id2, res.Id)

	val = cursor.X下一个(&res)
	ast.Equal(false, val)

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

	filter3 := 1

	cursor = cli.X查询(context.Background(), filter3).X取结果集()
	ast.Error(cursor.X取错误())
}

func TestQuery_Hint(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), nil, []string{"name", "age"})

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
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	var res []QueryTestItem

	filter1 := bson.M{
		"name": "Alice",
		"age":  18,
	}

	// index name as hint
	err = cli.X查询(context.Background(), filter1).X指定索引字段("age_1").X取全部(&res)
	ast.NoError(err)
	ast.Equal(1, len(res))

	// index name as hint
	var resOne QueryTestItem
	err = cli.X查询(context.Background(), filter1).X指定索引字段("name_1").X取一条(&resOne)
	ast.NoError(err)

	// not index name as hint
	err = cli.X查询(context.Background(), filter1).X指定索引字段("age").X取全部(&res)
	ast.Error(err)

	// nil hint
	err = cli.X查询(context.Background(), filter1).X指定索引字段(nil).X取全部(&res)
	ast.NoError(err)

}

func TestQuery_Apply(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), nil, []string{"name"})

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	id3 := primitive.NewObjectID()
	docs := []interface{}{
		bson.M{"_id": id1, "name": "Alice", "age": 18},
		bson.M{"_id": id2, "name": "Alice", "age": 19},
		bson.M{"_id": id3, "name": "Lucas", "age": 20, "instock": []bson.M{
			{"warehouse": "B", "qty": 15},
			{"warehouse": "C", "qty": 35},
			{"warehouse": "E", "qty": 15},
			{"warehouse": "F", "qty": 45},
		}}}

	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	res1 := QueryTestItem{}
	filter1 := bson.M{
		"name": "Tom",
	}
	change1 := Change{}

	err = cli.X查询(context.Background(), filter1).X执行命令(change1, &res1)
	ast.EqualError(err, mongo.ErrNilDocument.Error())

	change1.X更新替换 = bson.M{
		mgo常量.X更新值: bson.M{
			"name": "Tom",
			"age":  18,
		},
	}
	err = cli.X查询(context.Background(), filter1).X执行命令(change1, &res1)
	ast.EqualError(err, mongo.ErrNoDocuments.Error())

	change1.X是否返回新文档 = true
	err = cli.X查询(context.Background(), filter1).X执行命令(change1, &res1)
	ast.EqualError(err, mongo.ErrNoDocuments.Error())

	change1.X是否返回新文档 = false
	change1.X是否未找到时插入 = true
	err = cli.X查询(context.Background(), filter1).X执行命令(change1, &res1)
	ast.NoError(err)
	ast.Equal("", res1.Name)
	ast.Equal(0, res1.Age)

	change1.X更新替换 = bson.M{
		mgo常量.X更新值: bson.M{
			"name": "Tom",
			"age":  19,
		},
	}
	change1.X是否返回新文档 = true
	change1.X是否未找到时插入 = true
	err = cli.X查询(context.Background(), filter1).X执行命令(change1, &res1)
	ast.NoError(err)
	ast.Equal("Tom", res1.Name)
	ast.Equal(19, res1.Age)

	res2 := QueryTestItem{}
	filter2 := bson.M{
		"name": "Alice",
	}
	change2 := Change{
		X是否返回新文档: true,
		X更新替换: bson.M{
			mgo常量.X更新值: bson.M{
				"name": "Alice",
				"age":  22,
			},
		},
	}
	projection2 := bson.M{
		"age": 1,
	}
	err = cli.X查询(context.Background(), filter2).X排序("age").X字段(projection2).X执行命令(change2, &res2)
	ast.NoError(err)
	ast.Equal("", res2.Name)
	ast.Equal(22, res2.Age)

	res3 := QueryTestItem{}
	filter3 := bson.M{
		"name": "Bob",
	}
	change3 := Change{
		X是否删除: true,
	}
	err = cli.X查询(context.Background(), filter3).X执行命令(change3, &res3)
	ast.EqualError(err, mongo.ErrNoDocuments.Error())

	res3 = QueryTestItem{}
	filter3 = bson.M{
		"name": "Alice",
	}
	projection3 := bson.M{
		"age": 1,
	}
	err = cli.X查询(context.Background(), filter3).X排序("age").X字段(projection3).X执行命令(change3, &res3)
	ast.NoError(err)
	ast.Equal("", res3.Name)
	ast.Equal(19, res3.Age)

	res4 := QueryTestItem{}
	filter4 := bson.M{
		"name": "Bob",
	}
	change4 := Change{
		X是否替换: true,
		X更新替换: bson.M{
			mgo常量.X更新值: bson.M{
				"name": "Bob",
				"age":  23,
			},
		},
	}
	err = cli.X查询(context.Background(), filter4).X执行命令(change4, &res4)
	ast.EqualError(err, X错误_替换_文档含更新操作符.Error())

	change4.X更新替换 = bson.M{"name": "Bob", "age": 23}
	err = cli.X查询(context.Background(), filter4).X执行命令(change4, &res4)
	ast.EqualError(err, mongo.ErrNoDocuments.Error())

	change4.X是否返回新文档 = true
	err = cli.X查询(context.Background(), filter4).X执行命令(change4, &res4)
	ast.EqualError(err, mongo.ErrNoDocuments.Error())

	change4.X是否未找到时插入 = true
	change4.X是否返回新文档 = true
	err = cli.X查询(context.Background(), filter4).X执行命令(change4, &res4)
	ast.NoError(err)
	ast.Equal("Bob", res4.Name)
	ast.Equal(23, res4.Age)

	change4 = Change{
		X是否替换:   true,
		X更新替换:    bson.M{"name": "Bob", "age": 25},
		X是否未找到时插入:    true,
		X是否返回新文档: false,
	}
	projection4 := bson.M{
		"age":  1,
		"name": 1,
	}
	err = cli.X查询(context.Background(), filter4).X排序("age").X字段(projection4).X执行命令(change4, &res4)
	ast.NoError(err)
	ast.Equal("Bob", res4.Name)
	ast.Equal(23, res4.Age)

	res4 = QueryTestItem{}
	filter4 = bson.M{
		"name": "James",
	}
	change4 = Change{
		X是否替换:   true,
		X更新替换:    bson.M{"name": "James", "age": 26},
		X是否未找到时插入:    true,
		X是否返回新文档: false,
	}
	err = cli.X查询(context.Background(), filter4).X执行命令(change4, &res4)
	ast.NoError(err)
	ast.Equal("", res4.Name)
	ast.Equal(0, res4.Age)

	var res5 = QueryTestItem{}
	filter5 := bson.M{"name": "Lucas"}
	change5 := Change{
		X更新替换:    bson.M{"$set": bson.M{"instock.$[elem].qty": 100}},
		X是否返回新文档: true,
	}
	err = cli.X查询(context.Background(), filter5).X设置切片过滤(&options.ArrayFilters{Filters: []interface{}{
		bson.M{"elem.warehouse": bson.M{"$in": []string{"C", "F"}}},
	}}).X执行命令(change5, &res5)
	ast.NoError(err)

	for _, item := range res5.Instock {
		switch item.Warehouse {
		case "C", "F":
			ast.Equal(100, item.Qty)
		case "B", "E":
			ast.Equal(15, item.Qty)
		}
	}
}

func TestQuery_BatchSize(t *testing.T) {
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
	_, _ = cli.X插入多个(context.Background(), docs)
	var res []QueryTestItem

	err := cli.X查询(context.Background(), bson.M{"name": "Alice"}).X设置批量处理数量(1).X取全部(&res)
	ast.NoError(err)
	ast.Len(res, 2)

}
