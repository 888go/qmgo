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
	"reflect"
	"strings"

	"github.com/888go/qmgo/middleware"
	"github.com/888go/qmgo/operator"
	opts "github.com/888go/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// X文档集合 is a handle to a MongoDB collection
type X文档集合 struct {
	collection *mongo.Collection

	registry *bsoncodec.Registry
}

// X查询 find by condition filter，return QueryI
func (c *X文档集合) X查询(上下文 context.Context, 查询条件 interface{}, 可选选项 ...opts.FindOptions) QueryI {

	return &X查询{
		ctx:        上下文,
		collection: c.collection,
		filter:     查询条件,
		opts:       可选选项,
		registry:   c.registry,
	}
}

// X插入 insert one document into the collection
// If InsertHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Reference: https://docs.mongodb.com/manual/reference/command/insert/
func (c *X文档集合) X插入(上下文 context.Context, 待插入文档 interface{}, 可选选项 ...opts.InsertOneOptions) (插入结果 *X插入结果, 错误 error) {
	h := 待插入文档
	insertOneOpts := options.InsertOne()
	if len(可选选项) > 0 {
		if 可选选项[0].InsertOneOptions != nil {
			insertOneOpts = 可选选项[0].InsertOneOptions
		}
		if 可选选项[0].InsertHook != nil {
			h = 可选选项[0].InsertHook
		}
	}
	if 错误 = middleware.Do(上下文, 待插入文档, mgo常量.X钩子_插入前, h); 错误 != nil {
		return
	}
	res, 错误 := c.collection.InsertOne(上下文, 待插入文档, insertOneOpts)
	if res != nil {
		插入结果 = &X插入结果{X插入ID: res.InsertedID}
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 待插入文档, mgo常量.X钩子_插入后, h); 错误 != nil {
		return
	}
	return
}

// X插入多个 executes an insert command to insert multiple documents into the collection.
// If InsertHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Reference: https://docs.mongodb.com/manual/reference/command/insert/
func (c *X文档集合) X插入多个(上下文 context.Context, 待插入文档 interface{}, 可选选项 ...opts.InsertManyOptions) (插入结果 *X插入多条结果, 错误 error) {
	h := 待插入文档
	insertManyOpts := options.InsertMany()
	if len(可选选项) > 0 {
		if 可选选项[0].InsertManyOptions != nil {
			insertManyOpts = 可选选项[0].InsertManyOptions
		}
		if 可选选项[0].InsertHook != nil {
			h = 可选选项[0].InsertHook
		}
	}
	if 错误 = middleware.Do(上下文, 待插入文档, mgo常量.X钩子_插入前, h); 错误 != nil {
		return
	}
	sDocs := interfaceToSliceInterface(待插入文档)
	if sDocs == nil {
		return nil, X错误_插入_无效切片
	}

	res, 错误 := c.collection.InsertMany(上下文, sDocs, insertManyOpts)
	if res != nil {
		插入结果 = &X插入多条结果{X插入IDs: res.InsertedIDs}
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 待插入文档, mgo常量.X钩子_插入后, h); 错误 != nil {
		return
	}
	return
}

// interfaceToSliceInterface convert interface to slice interface
func interfaceToSliceInterface(docs interface{}) []interface{} {
	if reflect.Slice != reflect.TypeOf(docs).Kind() {
		return nil
	}
	s := reflect.ValueOf(docs)
	if s.Len() == 0 {
		return nil
	}
	var sDocs []interface{}
	for i := 0; i < s.Len(); i++ {
		sDocs = append(sDocs, s.Index(i).Interface())
	}
	return sDocs
}

// X替换插入 updates one documents if filter match, inserts one document if filter is not match, Error when the filter is invalid
// The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil
// and cannot contain any update operators
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
// If replacement has "_id" field and the document is existed, please initial it with existing id(even with Qmgo default field feature).
// Otherwise, "the (immutable) field '_id' altered" error happens.
func (c *X文档集合) X替换插入(上下文 context.Context, 替换条件 interface{}, 替换内容 interface{}, 可选选项 ...opts.UpsertOptions) (结果 *X更新结果, 错误 error) {
	h := 替换内容
	officialOpts := options.Replace().SetUpsert(true)

	if len(可选选项) > 0 {
		if 可选选项[0].ReplaceOptions != nil {
			可选选项[0].ReplaceOptions.SetUpsert(true)
			officialOpts = 可选选项[0].ReplaceOptions
		}
		if 可选选项[0].UpsertHook != nil {
			h = 可选选项[0].UpsertHook
		}
	}
	if 错误 = middleware.Do(上下文, 替换内容, mgo常量.X钩子_替换插入前, h); 错误 != nil {
		return
	}

	res, 错误 := c.collection.ReplaceOne(上下文, 替换条件, 替换内容, officialOpts)

	if res != nil {
		结果 = translateUpdateResult(res)
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 替换内容, mgo常量.X钩子_替换插入后, h); 错误 != nil {
		return
	}
	return
}

// X替换插入并按ID updates one documents if id match, inserts one document if id is not match and the id will inject into the document
// The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil
// and cannot contain any update operators
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
func (c *X文档集合) X替换插入并按ID(上下文 context.Context, 替换ID interface{}, 替换内容 interface{}, 可选选项 ...opts.UpsertOptions) (结果 *X更新结果, 错误 error) {
	h := 替换内容
	officialOpts := options.Replace().SetUpsert(true)

	if len(可选选项) > 0 {
		if 可选选项[0].ReplaceOptions != nil {
			可选选项[0].ReplaceOptions.SetUpsert(true)
			officialOpts = 可选选项[0].ReplaceOptions
		}
		if 可选选项[0].UpsertHook != nil {
			h = 可选选项[0].UpsertHook
		}
	}
	if 错误 = middleware.Do(上下文, 替换内容, mgo常量.X钩子_替换插入前, h); 错误 != nil {
		return
	}
	res, 错误 := c.collection.ReplaceOne(上下文, bson.M{"_id": 替换ID}, 替换内容, officialOpts)
	if res != nil {
		结果 = translateUpdateResult(res)
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 替换内容, mgo常量.X钩子_替换插入后, h); 错误 != nil {
		return
	}
	return
}

// X更新一条 executes an update command to update at most one document in the collection.
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
func (c *X文档集合) X更新一条(上下文 context.Context, 更新条件 interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (错误 error) {
	updateOpts := options.Update()

	if len(可选选项) > 0 {
		if 可选选项[0].UpdateOptions != nil {
			updateOpts = 可选选项[0].UpdateOptions
		}
		if 可选选项[0].UpdateHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, mgo常量.X钩子_更新前); 错误 != nil {
				return
			}
		}
	}

	res, 错误 := c.collection.UpdateOne(上下文, 更新条件, 更新内容, updateOpts)
	if res != nil && res.MatchedCount == 0 {
		// UpdateOne support upsert function
		if updateOpts.Upsert == nil || !*updateOpts.Upsert {
			错误 = X错误_未找到文档
		}
	}
	if 错误 != nil {
		return 错误
	}
	if len(可选选项) > 0 && 可选选项[0].UpdateHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, mgo常量.X钩子_更新后); 错误 != nil {
			return
		}
	}
	return 错误
}

// X更新并按ID executes an update command to update at most one document in the collection.
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
func (c *X文档集合) X更新并按ID(上下文 context.Context, 更新ID interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (错误 error) {
	updateOpts := options.Update()

	if len(可选选项) > 0 {
		if 可选选项[0].UpdateOptions != nil {
			updateOpts = 可选选项[0].UpdateOptions
		}
		if 可选选项[0].UpdateHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, mgo常量.X钩子_更新前); 错误 != nil {
				return
			}
		}
	}

	res, 错误 := c.collection.UpdateOne(上下文, bson.M{"_id": 更新ID}, 更新内容, updateOpts)
	if res != nil && res.MatchedCount == 0 {
		错误 = X错误_未找到文档
	}
	if 错误 != nil {
		return 错误
	}
	if len(可选选项) > 0 && 可选选项[0].UpdateHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, mgo常量.X钩子_更新后); 错误 != nil {
			return
		}
	}
	return 错误
}

// X更新 executes an update command to update documents in the collection.
// The matchedCount is 0 in UpdateResult if no document updated
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
func (c *X文档集合) X更新(上下文 context.Context, 更新条件 interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (更新结果 *X更新结果, 错误 error) {
	updateOpts := options.Update()
	if len(可选选项) > 0 {
		if 可选选项[0].UpdateOptions != nil {
			updateOpts = 可选选项[0].UpdateOptions
		}
		if 可选选项[0].UpdateHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, mgo常量.X钩子_更新前); 错误 != nil {
				return
			}
		}
	}
	res, 错误 := c.collection.UpdateMany(上下文, 更新条件, 更新内容, updateOpts)
	if res != nil {
		更新结果 = translateUpdateResult(res)
	}
	if 错误 != nil {
		return
	}
	if len(可选选项) > 0 && 可选选项[0].UpdateHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, mgo常量.X钩子_更新后); 错误 != nil {
			return
		}
	}
	return
}

// X替换一条 executes an update command to update at most one document in the collection.
// If UpdateHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Expect type of the doc is the define of user's document
func (c *X文档集合) X替换一条(上下文 context.Context, 替换条件 interface{}, 替换内容 interface{}, 可选选项 ...opts.ReplaceOptions) (错误 error) {
	h := 替换内容
	replaceOpts := options.Replace()

	if len(可选选项) > 0 {
		if 可选选项[0].ReplaceOptions != nil {
			replaceOpts = 可选选项[0].ReplaceOptions
			replaceOpts.SetUpsert(false)
		}
		if 可选选项[0].UpdateHook != nil {
			h = 可选选项[0].UpdateHook
		}
	}
	if 错误 = middleware.Do(上下文, 替换内容, mgo常量.X钩子_替换前, h); 错误 != nil {
		return
	}
	res, 错误 := c.collection.ReplaceOne(上下文, 替换条件, 替换内容, replaceOpts)
	if res != nil && res.MatchedCount == 0 {
		错误 = X错误_未找到文档
	}
	if 错误 != nil {
		return 错误
	}
	if 错误 = middleware.Do(上下文, 替换内容, mgo常量.X钩子_替换后, h); 错误 != nil {
		return
	}

	return 错误
}

// X删除一条 executes a delete command to delete at most one document from the collection.
// if filter is bson.M{}，DeleteOne will delete one document in collection
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
func (c *X文档集合) X删除一条(上下文 context.Context, 删除条件 interface{}, 可选选项 ...opts.RemoveOptions) (错误 error) {
	deleteOptions := options.Delete()
	if len(可选选项) > 0 {
		if 可选选项[0].DeleteOptions != nil {
			deleteOptions = 可选选项[0].DeleteOptions
		}
		if 可选选项[0].RemoveHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, mgo常量.X钩子_删除前); 错误 != nil {
				return 错误
			}
		}
	}
	res, 错误 := c.collection.DeleteOne(上下文, 删除条件, deleteOptions)
	if res != nil && res.DeletedCount == 0 {
		错误 = X错误_未找到文档
	}
	if 错误 != nil {
		return 错误
	}
	if len(可选选项) > 0 && 可选选项[0].RemoveHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, mgo常量.X钩子_删除后); 错误 != nil {
			return 错误
		}
	}
	return 错误
}

// X删除并按ID executes a delete command to delete at most one document from the collection.
func (c *X文档集合) X删除并按ID(上下文 context.Context, 删除ID interface{}, 可选选项 ...opts.RemoveOptions) (错误 error) {
	deleteOptions := options.Delete()
	if len(可选选项) > 0 {
		if 可选选项[0].DeleteOptions != nil {
			deleteOptions = 可选选项[0].DeleteOptions
		}
		if 可选选项[0].RemoveHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, mgo常量.X钩子_删除前); 错误 != nil {
				return 错误
			}
		}
	}
	res, 错误 := c.collection.DeleteOne(上下文, bson.M{"_id": 删除ID}, deleteOptions)
	if res != nil && res.DeletedCount == 0 {
		错误 = X错误_未找到文档
	}
	if 错误 != nil {
		return 错误
	}

	if len(可选选项) > 0 && 可选选项[0].RemoveHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, mgo常量.X钩子_删除后); 错误 != nil {
			return 错误
		}
	}
	return 错误
}

// X删除 executes a delete command to delete documents from the collection.
// If filter is bson.M{}，all ducuments in Collection will be deleted
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
func (c *X文档集合) X删除(上下文 context.Context, 删除条件 interface{}, 可选选项 ...opts.RemoveOptions) (删除结果 *X删除结果, 错误 error) {
	deleteOptions := options.Delete()
	if len(可选选项) > 0 {
		if 可选选项[0].DeleteOptions != nil {
			deleteOptions = 可选选项[0].DeleteOptions
		}
		if 可选选项[0].RemoveHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, mgo常量.X钩子_删除前); 错误 != nil {
				return
			}
		}
	}
	res, 错误 := c.collection.DeleteMany(上下文, 删除条件, deleteOptions)
	if res != nil {
		删除结果 = &X删除结果{X删除数量: res.DeletedCount}
	}
	if 错误 != nil {
		return
	}
	if len(可选选项) > 0 && 可选选项[0].RemoveHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, mgo常量.X钩子_删除后); 错误 != nil {
			return
		}
	}
	return
}

// X聚合 executes an aggregate command against the collection and returns a AggregateI to get resulting documents.
func (c *X文档集合) X聚合(上下文 context.Context, 聚合管道 interface{}, 可选选项 ...opts.AggregateOptions) AggregateI {
	return &Aggregate{
		ctx:        上下文,
		collection: c.collection,
		pipeline:   聚合管道,
		options:    可选选项,
	}
}

// ensureIndex create multiple indexes on the collection and returns the names of
// Example：indexes = []string{"idx1", "-idx2", "idx3,idx4"}
// Three indexes will be created, index idx1 with ascending order, index idx2 with descending order, idex3 and idex4 are Compound ascending sort index
// Reference: https://docs.mongodb.com/manual/reference/command/createIndexes/
func (c *X文档集合) ensureIndex(ctx context.Context, indexes []opts.X索引选项) error {
	var indexModels []mongo.IndexModel
	for _, idx := range indexes {
		var model mongo.IndexModel
		var keysDoc bson.D

		for _, field := range idx.X索引字段 {
			key, n := X分割排序字段(field)

			keysDoc = append(keysDoc, bson.E{Key: key, Value: n})
		}
		model = mongo.IndexModel{
			Keys:    keysDoc,
			Options: idx.IndexOptions,
		}

		indexModels = append(indexModels, model)
	}

	if len(indexModels) == 0 {
		return nil
	}

	res, err := c.collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil || len(res) == 0 {
		fmt.Println("<MongoDB.C>: ", c.collection.Name(), " Index: ", indexes, " error: ", err, "res: ", res)
		return err
	}
	return nil
}

// EnsureIndexes弃用 Deprecated
// Recommend to use CreateIndexes / CreateOneIndex for more function)
// EnsureIndexes弃用 creates unique and non-unique indexes in collection
// the combination of indexes is different from CreateIndexes:
// if uniques/indexes is []string{"name"}, means create index "name"
// if uniques/indexes is []string{"name,-age","uid"} means create Compound indexes: name and -age, then create one index: uid
func (c *X文档集合) EnsureIndexes弃用(ctx context.Context, uniques []string, indexes []string) (err error) {
	var uniqueModel []opts.X索引选项
	var indexesModel []opts.X索引选项
	for _, v := range uniques {
		vv := strings.Split(v, ",")
		indexOpts := options.Index()
		indexOpts.SetUnique(true)
		model := opts.X索引选项{X索引字段: vv, IndexOptions: indexOpts}
		uniqueModel = append(uniqueModel, model)
	}
	if err = c.X创建多条索引(ctx, uniqueModel); err != nil {
		return
	}

	for _, v := range indexes {
		vv := strings.Split(v, ",")
		model := opts.X索引选项{X索引字段: vv}
		indexesModel = append(indexesModel, model)
	}
	if err = c.X创建多条索引(ctx, indexesModel); err != nil {
		return
	}
	return
}

// X创建多条索引 creates multiple indexes in collection
// If the Key in opts.IndexModel is []string{"name"}, means create index: name
// If the Key in opts.IndexModel is []string{"name","-age"} means create Compound indexes: name and -age
func (c *X文档集合) X创建多条索引(上下文 context.Context, 索引s []opts.X索引选项) (错误 error) {
	错误 = c.ensureIndex(上下文, 索引s)
	return
}

// X创建索引 creates one index
// If the Key in opts.IndexModel is []string{"name"}, means create index name
// If the Key in opts.IndexModel is []string{"name","-age"} means create Compound index: name and -age
func (c *X文档集合) X创建索引(上下文 context.Context, 索引 opts.X索引选项) error {
	return c.ensureIndex(上下文, []opts.X索引选项{索引})

}

// X删除全部索引 drop all indexes on the collection except the index on the _id field
// if there is only _id field index on the collection, the function call will report an error
func (c *X文档集合) X删除全部索引(上下文 context.Context) (错误 error) {
	_, 错误 = c.collection.Indexes().DropAll(上下文)
	return 错误
}

// X删除索引 drop indexes in collection, indexes that be dropped should be in line with inputting indexes
// The indexes is []string{"name"} means drop index: name
// The indexes is []string{"name","-age"} means drop Compound indexes: name and -age
func (c *X文档集合) X删除索引(上下文 context.Context, 索引s []string) error {
	_, err := c.collection.Indexes().DropOne(上下文, generateDroppedIndex(索引s))
	if err != nil {
		return err
	}
	return err
}

// generate indexes that store in mongo which may consist more than one index(like []string{"index1","index2"} is stored as "index1_1_index2_1")
func generateDroppedIndex(index []string) string {
	var res string
	for _, e := range index {
		key, sort := X分割排序字段(e)
		n := key + "_" + fmt.Sprint(sort)
		if len(res) == 0 {
			res = n
		} else {
			res += "_" + n
		}
	}
	return res
}

// X删除集合 drops collection
// it's safe even collection is not exists
func (c *X文档集合) X删除集合(上下文 context.Context) error {
	return c.collection.Drop(上下文)
}

// X取副本 creates a copy of the Collection
func (c *X文档集合) X取副本() (*mongo.Collection, error) {
	return c.collection.Clone()
}

// X取集合名 returns the name of collection
func (c *X文档集合) X取集合名() string {
	return c.collection.Name()
}

// X取变更流 returns a change stream for all changes on the corresponding collection. See
// https://docs.mongodb.com/manual/changeStreams/ for more information about change streams.
func (c *X文档集合) X取变更流(上下文 context.Context, 管道 interface{}, 可选选项 ...*opts.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	changeStreamOption := options.ChangeStream()
	if len(可选选项) > 0 && 可选选项[0].ChangeStreamOptions != nil {
		changeStreamOption = 可选选项[0].ChangeStreamOptions
	}
	return c.collection.Watch(上下文, 管道, changeStreamOption)
}

// translateUpdateResult translates mongo update result to qmgo define UpdateResult
func translateUpdateResult(res *mongo.UpdateResult) (result *X更新结果) {
	result = &X更新结果{
		X匹配数:  res.MatchedCount,
		X修改数: res.ModifiedCount,
		X替换插入数: res.UpsertedCount,
		X替换插入ID:    res.UpsertedID,
	}
	return
}
