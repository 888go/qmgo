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
	"go.mongodb.org/mongo-driver/mongo"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/888go/qmgo/operator"
	"github.com/888go/qmgo/options"
)

func TestCollection_EnsureIndex(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())

	cli.ensureIndex(context.Background(), nil)
	indexOpts := officialOpts.Index()
	indexOpts.SetUnique(true)
	cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"id1"}, IndexOptions: indexOpts}})
	cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"id2", "id3"}}})
	cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"id4", "-id5"}}})

	// same index，error
	ast.Error(cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"id1"}}}))

	// 检查唯一索引是否正常工作 md5:9b2257b60d7b5998
	var err error
	doc := bson.M{
		"id1": 1,
	}
	_, err = cli.X插入(context.Background(), doc)
	ast.NoError(err)

	coll, err := cli.X取副本()
	ast.NoError(err)
	_, err = coll.InsertOne(context.Background(), doc)
	ast.Equal(true, X是否为重复键错误(err))
}

func TestCollection_EnsureIndexes(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())

	unique := []string{"id1"}
	common := []string{"id2,id3", "id4,-id5"}
	cli.EnsureIndexes弃用(context.Background(), unique, common)

	// same index，error
	ast.Error(cli.EnsureIndexes弃用(context.Background(), nil, unique))

	// 检查唯一索引是否正常工作 md5:9b2257b60d7b5998
	var err error
	doc := bson.M{
		"id1": 1,
	}

	_, err = cli.X插入(context.Background(), doc)
	ast.NoError(err)
	_, err = cli.X插入(context.Background(), doc)
	ast.Equal(true, X是否为重复键错误(err))
}

func TestCollection_CreateIndexes(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())

	var expireS int32 = 100
	unique := []string{"id1"}
	indexOpts := officialOpts.Index()
	indexOpts.SetUnique(true).SetExpireAfterSeconds(expireS)
	ast.NoError(cli.X创建索引(context.Background(), options.X索引选项{X索引字段: unique, IndexOptions: indexOpts}))

	ast.NoError(cli.X创建多条索引(context.Background(), []options.X索引选项{{X索引字段: []string{"id2", "id3"}},
		{X索引字段: []string{"id4", "-id5"}}}))
	// same index，error
	ast.Error(cli.X创建索引(context.Background(), options.X索引选项{X索引字段: unique}))

	// 检查唯一索引是否正常工作 md5:9b2257b60d7b5998
	var err error
	doc := bson.M{
		"id1": 1,
	}

	_, err = cli.X插入(context.Background(), doc)
	ast.NoError(err)
	_, err = cli.X插入(context.Background(), doc)
	ast.Equal(true, X是否为重复键错误(err))
}

func TestCollection_DropAllIndexes(t *testing.T) {
	ast := require.New(t)

	cli := initClient("test")
	defer cli.X删除集合(context.Background())

	var err error
	err = cli.X删除全部索引(context.Background())
	ast.Error(err)

	unique := []string{"id1"}
	common := []string{"id2,id3", "id4,-id5"}
	cli.EnsureIndexes弃用(context.Background(), unique, common)

	err = cli.X删除全部索引(context.Background())
	ast.NoError(err)
}

func TestCollection_DropIndex(t *testing.T) {
	ast := require.New(t)

	cli := initClient("test")
	defer cli.X删除集合(context.Background())

	indexOpts := officialOpts.Index()
	indexOpts.SetUnique(true)
	cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"index1"}, IndexOptions: indexOpts}})

	// same index，error
	ast.Error(cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"index1"}}}))

	err := cli.X删除索引(context.Background(), []string{"index1"})
	ast.NoError(err)
	ast.NoError(cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"index1"}}}))

	indexOpts = officialOpts.Index()
	indexOpts.SetUnique(true)
	cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"-index1"}, IndexOptions: indexOpts}})
	// same index，error
	ast.Error(cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"-index1"}}}))

	err = cli.X删除索引(context.Background(), []string{"-index1"})
	ast.NoError(err)
	ast.NoError(cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"-index1"}}}))

	err = cli.X删除索引(context.Background(), []string{""})
	ast.Error(err)

	err = cli.X删除索引(context.Background(), []string{"index2"})
	ast.Error(err)

	indexOpts = officialOpts.Index()
	indexOpts.SetUnique(true)
	cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"index2", "-index1"}, IndexOptions: indexOpts}})
	ast.Error(cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"index2", "-index1"}}}))
	err = cli.X删除索引(context.Background(), []string{"index2", "-index1"})
	ast.NoError(err)
	ast.NoError(cli.ensureIndex(context.Background(), []options.X索引选项{{X索引字段: []string{"index2", "-index1"}}}))

	err = cli.X删除索引(context.Background(), []string{"-"})
	ast.Error(err)
}

func TestCollection_Insert(t *testing.T) {
	ast := require.New(t)

	cli := initClient("test")

	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())

	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	var err error
	doc := bson.M{"_id": primitive.NewObjectID(), "name": "Alice"}

	opts := options.InsertOneOptions{}
	opts.InsertOneOptions = officialOpts.InsertOne().SetBypassDocumentValidation(true)
	res, err := cli.X插入(context.Background(), doc, opts)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(doc["_id"], res.X插入ID)

	res, err = cli.X插入(context.Background(), doc)
	ast.Equal(true, X是否为重复键错误(err))
	ast.Empty(res)
}

func TestCollection_InsertMany(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	var err error
	newDocs := []UserInfo{{Id: X生成对象ID(), Name: "Alice", Age: 10}, {Id: X生成对象ID(), Name: "Lucas", Age: 11}}
	res, err := cli.X插入多个(context.Background(), newDocs)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(2, len(res.X插入IDs))

	newPDocs := []*UserInfo{{Id: X生成对象ID(), Name: "Alice03", Age: 10}, {Id: X生成对象ID(), Name: "Lucas03", Age: 11}}
	res, err = cli.X插入多个(context.Background(), newPDocs)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(2, len(res.X插入IDs))

	docs2 := []UserInfo{
		{Name: "Alice"},
		{Name: "Lucas"},
	}
	opts := options.InsertManyOptions{}
	opts.InsertManyOptions = officialOpts.InsertMany().SetBypassDocumentValidation(true)
	res, err = cli.X插入多个(context.Background(), docs2, opts)
	ast.Equal(true, X是否为重复键错误(err))
	ast.Equal(0, len(res.X插入IDs))

	docs4 := []UserInfo{}
	res, err = cli.X插入多个(context.Background(), []interface{}{docs4})
	ast.Error(err)
	ast.Empty(res)

}

func TestCollection_Upsert(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	docs := []interface{}{
		bson.D{{Key: "_id", Value: id1}, {Key: "name", Value: "Alice"}},
		bson.D{{Key: "_id", Value: id2}, {Key: "name", Value: "Lucas"}},
	}
	_, err := cli.X插入多个(context.Background(), docs)

	ast.NoError(err)
	// replace already exist
	filter1 := bson.M{
		"name": "Alice",
	}
	replacement1 := bson.M{
		"name": "Alice1",
		"age":  18,
	}
	opts := options.UpsertOptions{}
	opts.ReplaceOptions = officialOpts.Replace()
	res, err := cli.X替换插入(context.Background(), filter1, replacement1, opts)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(int64(1), res.X匹配数)
	ast.Equal(int64(1), res.X修改数)
	ast.Equal(int64(0), res.X替换插入数)
	ast.Equal(nil, res.X替换插入ID)

	// not exist
	filter2 := bson.M{
		"name": "Lily",
	}
	replacement2 := bson.M{
		"name": "Lily",
		"age":  20,
	}
	res, err = cli.X替换插入(context.Background(), filter2, replacement2)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(int64(0), res.X匹配数)
	ast.Equal(int64(0), res.X修改数)
	ast.Equal(int64(1), res.X替换插入数)
	ast.NotNil(res.X替换插入ID)

	// filter 是空或者不符合正确的BSON文档格式 md5:a55c6ef20a253667
	replacement3 := bson.M{
		"name": "Geek",
		"age":  21,
	}
	res, err = cli.X替换插入(context.Background(), nil, replacement3)
	ast.Error(err)
	ast.Empty(res)

	res, err = cli.X替换插入(context.Background(), 1, replacement3)
	ast.Error(err)
	ast.Empty(res)

	// replacement 是空的或者不符合正确的BSON文档格式 md5:7b0ecb01590a648b
	filter4 := bson.M{
		"name": "Geek",
	}
	res, err = cli.X替换插入(context.Background(), filter4, nil)
	ast.Error(err)
	ast.Empty(res)

	res, err = cli.X替换插入(context.Background(), filter4, 1)
	ast.Error(err)
	ast.Empty(res)
}

func TestCollection_UpsertId(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	docs := []interface{}{
		bson.D{{Key: "_id", Value: id1}, {Key: "name", Value: "Alice"}},
		bson.D{{Key: "_id", Value: id2}, {Key: "name", Value: "Lucas"}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error

	// replace already exist
	replacement1 := bson.M{
		"name": "Alice1",
		"age":  18,
	}
	res, err := cli.X替换插入并按ID(context.Background(), id1, replacement1)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(int64(1), res.X匹配数)
	ast.Equal(int64(1), res.X修改数)
	ast.Equal(int64(0), res.X替换插入数)
	ast.Equal(nil, res.X替换插入ID)

	// not exist filter id
	replacement2 := bson.M{
		"name": "Lily",
		"age":  20,
	}
	id3 := primitive.NewObjectID()
	opts := options.UpsertOptions{}
	opts.ReplaceOptions = officialOpts.Replace()
	res, err = cli.X替换插入并按ID(context.Background(), id3, replacement2, opts)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(int64(0), res.X匹配数)
	ast.Equal(int64(0), res.X修改数)
	ast.Equal(int64(1), res.X替换插入数)
	ast.Equal(id3, res.X替换插入ID) // id3 将会插入到已插入的文档中 md5:4cdfbeaa6a4c59ce

	// 使用与文档中id不同的过滤器，错误 md5:1864a41611ea40ba
	id4 := primitive.NewObjectID()
	replacement3 := bson.M{
		"_id":  id4,
		"name": "Joe",
		"age":  20,
	}
	id5 := primitive.NewObjectID()
	res, err = cli.X替换插入并按ID(context.Background(), id5, replacement3)
	ast.Error(err)

	// filter is nil
	replacement4 := bson.M{
		"name": "Geek",
		"age":  21,
	}
	res, err = cli.X替换插入并按ID(context.Background(), nil, replacement4)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(int64(0), res.X匹配数)
	ast.Equal(int64(0), res.X修改数)
	ast.Equal(int64(1), res.X替换插入数)
	ast.Nil(res.X替换插入ID)

	// replacement 是空的或者不符合正确的BSON文档格式 md5:7b0ecb01590a648b
	res, err = cli.X替换插入并按ID(context.Background(), id1, nil)
	ast.Error(err)
	ast.Empty(res)

	res, err = cli.X替换插入并按ID(context.Background(), id1, 1)
	ast.Error(err)
	ast.Empty(res)
}

func TestCollection_Update(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	docs := []interface{}{
		bson.D{{Key: "_id", Value: id1}, {Key: "name", Value: "Alice"}},
		bson.D{{Key: "_id", Value: id2}, {Key: "name", Value: "Lucas"}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	// 更新已存在的记录 md5:cc4fac8615b8fc8a
	filter1 := bson.M{
		"name": "Alice",
	}
	update1 := bson.M{
		mgo常量.X更新值: bson.M{
			"name": "Alice1",
			"age":  18,
		},
	}
	opts := options.UpdateOptions{}
	opts.UpdateOptions = officialOpts.Update().SetBypassDocumentValidation(false)
	err = cli.X更新一条(context.Background(), filter1, update1, opts)
	ast.NoError(err)

	// error when not exist
	filter2 := bson.M{
		"name": "Lily",
	}
	update2 := bson.M{
		mgo常量.X更新值: bson.M{
			"name": "Lily",
			"age":  20,
		},
	}
	err = cli.X更新一条(context.Background(), filter2, update2)
	ast.Equal(err, X错误_未找到文档)

	opt := officialOpts.Update().SetUpsert(true)
	opts = options.UpdateOptions{UpdateOptions: opt}
	err = cli.X更新一条(context.Background(), filter2, update2, opts)
	ast.NoError(err)

	// filter 是空或者不符合正确的BSON文档格式 md5:a55c6ef20a253667
	update3 := bson.M{
		"name": "Geek",
		"age":  21,
	}
	err = cli.X更新一条(context.Background(), nil, update3)
	ast.Error(err)

	err = cli.X更新一条(context.Background(), 1, update3)
	ast.Error(err)

	// update 是 nil 或者格式错误的 BSON 文档 md5:8f6e8bd5cf0af638
	filter4 := bson.M{
		"name": "Geek",
	}
	err = cli.X更新一条(context.Background(), filter4, nil)
	ast.Error(err)

	err = cli.X更新一条(context.Background(), filter4, 1)
	ast.Error(err)
}

func TestCollection_UpdateId(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	docs := []interface{}{
		bson.D{{Key: "_id", Value: id1}, {Key: "name", Value: "Alice"}},
		bson.D{{Key: "_id", Value: id2}, {Key: "name", Value: "Lucas"}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	// 更新已存在的记录 md5:cc4fac8615b8fc8a
	update1 := bson.M{
		mgo常量.X更新值: bson.M{
			"name": "Alice1",
			"age":  18,
		},
	}
	opts := options.UpdateOptions{}
	opts.UpdateOptions = officialOpts.Update().SetBypassDocumentValidation(false)
	err = cli.X更新并按ID(context.Background(), id1, update1, opts)
	ast.NoError(err)

	// id is nil or not exist
	update3 := bson.M{
		"name": "Geek",
		"age":  21,
	}
	err = cli.X更新并按ID(context.Background(), nil, update3)
	ast.Error(err)

	err = cli.X更新并按ID(context.Background(), 1, update3)
	ast.Error(err)

	err = cli.X更新并按ID(context.Background(), "not_exist_id", nil)
	ast.Error(err)

	err = cli.X更新并按ID(context.Background(), "not_exist_id", 1)
	ast.Error(err)
}

func TestCollection_UpdateAll(t *testing.T) {
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
	// 更新已存在的记录 md5:cc4fac8615b8fc8a
	filter1 := bson.M{
		"name": "Alice",
	}
	update1 := bson.M{
		mgo常量.X更新值: bson.M{
			"age": 33,
		},
	}
	opts := options.UpdateOptions{}
	opts.UpdateOptions = officialOpts.Update().SetBypassDocumentValidation(false)
	res, err := cli.X更新(context.Background(), filter1, update1, opts)
	ast.NoError(err)
	ast.NotEmpty(res)
	ast.Equal(int64(2), res.X匹配数)
	ast.Equal(int64(2), res.X修改数)
	ast.Equal(int64(0), res.X替换插入数)
	ast.Equal(nil, res.X替换插入ID)

	// 如果记录不存在，err为nil，res中的MatchedCount为0 md5:ffbbcabc3c0f02fe
	filter2 := bson.M{
		"name": "Lily",
	}
	update2 := bson.M{
		mgo常量.X更新值: bson.M{
			"age": 22,
		},
	}
	res, err = cli.X更新(context.Background(), filter2, update2)
	ast.Nil(err)
	ast.NotNil(res)
	ast.Equal(int64(0), res.X匹配数)

	// filter 是空或者不符合正确的BSON文档格式 md5:a55c6ef20a253667
	update3 := bson.M{
		"name": "Geek",
		"age":  21,
	}
	res, err = cli.X更新(context.Background(), nil, update3)
	ast.Error(err)
	ast.Nil(res)

	res, err = cli.X更新(context.Background(), 1, update3)
	ast.Error(err)
	ast.Nil(res)

	// update 是 nil 或者格式错误的 BSON 文档 md5:8f6e8bd5cf0af638
	filter4 := bson.M{
		"name": "Geek",
	}
	res, err = cli.X更新(context.Background(), filter4, nil)
	ast.Error(err)
	ast.Nil(res)

	res, err = cli.X更新(context.Background(), filter4, 1)
	ast.Error(err)
	ast.Nil(res)
}

func TestCollection_Remove(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), nil, []string{"name"})

	id1 := primitive.NewObjectID().Hex()
	id2 := primitive.NewObjectID().Hex()
	id3 := primitive.NewObjectID().Hex()
	id4 := primitive.NewObjectID().Hex()
	id5 := primitive.NewObjectID()
	docs := []interface{}{
		bson.D{{Key: "_id", Value: id1}, {Key: "name", Value: "Alice"}, {Key: "age", Value: 18}},
		bson.D{{Key: "_id", Value: id2}, {Key: "name", Value: "Alice"}, {Key: "age", Value: 19}},
		bson.D{{Key: "_id", Value: id3}, {Key: "name", Value: "Lucas"}, {Key: "age", Value: 20}},
		bson.D{{Key: "_id", Value: id4}, {Key: "name", Value: "Joe"}, {Key: "age", Value: 20}},
		bson.D{{Key: "_id", Value: id5}, {Key: "name", Value: "Ethan"}, {Key: "age", Value: 1}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	// remove id
	err = cli.X删除并按ID(context.Background(), "")
	ast.Error(err)
	err = cli.X删除并按ID(context.Background(), "not-exists-id")
	ast.True(X是否为无文档错误(err))
	ast.NoError(cli.X删除并按ID(context.Background(), id4))
	ast.NoError(cli.X删除并按ID(context.Background(), id5))

	// 删除记录：名称为 "Alice"，之后预期存在一条名称为 "Alice" 的记录。 md5:274874b30e4288bb
	filter1 := bson.M{
		"name": "Alice",
	}
	opts := options.RemoveOptions{}
	opts.DeleteOptions = officialOpts.Delete()
	err = cli.X删除一条(context.Background(), filter1, opts)
	ast.NoError(err)

	cnt, err := cli.X查询(context.Background(), filter1).X取数量()
	ast.NoError(err)
	ast.Equal(int64(1), cnt)

	// 删除不匹配的记录，如果发生错误则报告错误 md5:46e3eb8e95abdfcc
	filter2 := bson.M{
		"name": "Lily",
	}
	err = cli.X删除一条(context.Background(), filter2)
	ast.Equal(err, X错误_未找到文档)

	// filter 是 bson.M{}，删除一个文档 md5:dc8fa3aa9522cd67
	filter3 := bson.M{}
	preCnt, err := cli.X查询(context.Background(), filter3).X取数量()
	ast.NoError(err)
	ast.Equal(int64(2), preCnt)

	err = cli.X删除一条(context.Background(), filter3)
	ast.NoError(err)

	afterCnt, err := cli.X查询(context.Background(), filter3).X取数量()
	ast.NoError(err)
	ast.Equal(preCnt-1, afterCnt)

	// filter 是空或者不符合正确的BSON文档格式 md5:a55c6ef20a253667
	err = cli.X删除一条(context.Background(), nil)
	ast.Error(err)

	err = cli.X删除一条(context.Background(), 1)
	ast.Error(err)
}

func TestCollection_RemoveAll(t *testing.T) {
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
		bson.D{{Key: "_id", Value: id1}, {Key: "name", Value: "Alice"}, {Key: "age", Value: 18}},
		bson.D{{Key: "_id", Value: id2}, {Key: "name", Value: "Alice"}, {Key: "age", Value: 19}},
		bson.D{{Key: "_id", Value: id3}, {Key: "name", Value: "Lucas"}, {Key: "age", Value: 20}},
		bson.D{{Key: "_id", Value: id4}, {Key: "name", Value: "Rocket"}, {Key: "age", Value: 23}},
	}
	_, _ = cli.X插入多个(context.Background(), docs)

	var err error
	// 删除记录：名称为 "Alice"，之后，预期 - 记录：名称为 "Alice" md5:e6ccda4a8c588184
	filter1 := bson.M{
		"name": "Alice",
	}
	opts := options.RemoveOptions{}
	opts.DeleteOptions = officialOpts.Delete()
	res, err := cli.X删除(context.Background(), filter1, opts)
	ast.NoError(err)
	ast.NotNil(res)
	ast.Equal(int64(2), res.X删除数量)

	cnt, err := cli.X查询(context.Background(), filter1).X取数量()
	ast.NoError(err)
	ast.Equal(int64(0), cnt)

	// 使用不匹配的过滤器删除，结果中的DeletedCount为0 md5:61f4e36a0742d763
	filter2 := bson.M{
		"name": "Lily",
	}
	res, err = cli.X删除(context.Background(), filter2)
	ast.NoError(err)
	ast.NotNil(res)
	ast.Equal(int64(0), res.X删除数量)

	// filter 是 bson.M{}，删除所有文档 md5:bbe0ca02d153a930
	filter3 := bson.M{}
	preCnt, err := cli.X查询(context.Background(), filter3).X取数量()
	ast.NoError(err)
	ast.Equal(int64(2), preCnt)

	res, err = cli.X删除(context.Background(), filter3)
	ast.NoError(err)
	ast.NotNil(res)
	ast.Equal(preCnt, res.X删除数量)

	afterCnt, err := cli.X查询(context.Background(), filter3).X取数量()
	ast.NoError(err)
	ast.Equal(int64(0), afterCnt)

	// filter 是空或者不符合正确的BSON文档格式 md5:a55c6ef20a253667
	res, err = cli.X删除(context.Background(), nil)
	ast.Error(err)
	ast.Nil(res)

	res, err = cli.X删除(context.Background(), 1)
	ast.Error(err)
	ast.Nil(res)
}
func TestSliceInsert(t *testing.T) {
	newDocs := []UserInfo{{Name: "Alice", Age: 10}, {Name: "Lucas", Age: 11}}
	di := interface{}(newDocs)
	dis := interfaceToSliceInterface(di)
	ast := require.New(t)
	ast.Len(dis, 2)

	newDocs_1 := []interface{}{UserInfo{Name: "Alice", Age: 10}, UserInfo{Name: "Lucas", Age: 11}}
	di = interface{}(newDocs_1)
	dis = interfaceToSliceInterface(di)
	ast.Len(dis, 2)

	newDocs_2 := UserInfo{Name: "Alice", Age: 10}
	di = interface{}(newDocs_2)
	dis = interfaceToSliceInterface(di)
	ast.Nil(dis)

	newDocs_3 := []UserInfo{}
	di = interface{}(newDocs_3)
	dis = interfaceToSliceInterface(di)
	ast = require.New(t)
	ast.Nil(dis)
}

func TestCollection_ReplaceOne(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	id := primitive.NewObjectID()
	ui := UserInfo{Id: id, Name: "Lucas", Age: 17}
	_, err := cli.X插入(context.Background(), ui)
	ast.NoError(err)
	ui.Id = id
	ui.Age = 27
	err = cli.X替换一条(context.Background(), bson.M{"_id": id}, &ui)
	ast.NoError(err)

	findUi := UserInfo{}
	err = cli.X查询(context.Background(), bson.M{"name": "Lucas"}).X取一条(&findUi)
	ast.NoError(err)
	ast.Equal(ui.Age, findUi.Age)

	opts := options.ReplaceOptions{}
	opts.ReplaceOptions = officialOpts.Replace()
	err = cli.X替换一条(context.Background(), bson.M{"_id": "notexist"}, &ui, opts)
	ast.Equal(X错误_未找到文档, err)

	err = cli.X替换一条(context.Background(), bson.M{"_id": "notexist"}, nil)
	ast.Error(err)
}

func TestChangeStream(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())

	opts := &options.ChangeStreamOptions{officialOpts.ChangeStream()}
	c, e := cli.X取变更流(context.Background(), mongo.Pipeline{}, opts)
	ast.NoError(e)
	defer c.Close(context.Background())

	doneChane := make(chan struct{})
	go func() {
		ok := c.Next(context.Background())
		ast.True(ok)
		doneChane <- struct{}{}
	}()

	id := primitive.NewObjectID()
	ui := UserInfo{Id: id, Name: "Lucas", Age: 17}
	_, err := cli.X插入(context.Background(), ui)
	ast.NoError(err)
	<-doneChane

}
