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

// Collection 是一个MongoDB集合的句柄 md5:be1b94030609bdd1
type Collection struct {
	collection *mongo.Collection

	registry *bsoncodec.Registry
}

// Find 通过条件过滤并查找，返回QueryI md5:bda4cc0c85d800a1
func (c *Collection) X查询(上下文 context.Context, 查询条件 interface{}, 可选选项 ...opts.FindOptions) QueryI {

	return &Query{
		ctx:        上下文,
		collection: c.collection,
		filter:     查询条件,
		opts:       可选选项,
		registry:   c.registry,
	}
}

// InsertOne 将一个文档插入到集合中
// 如果在 opts 中设置了 InsertHook，钩子将在其上执行，否则钩子会尝试处理 doc 作为钩子
// 参考: https://docs.mongodb.com/manual/reference/command/insert/
// md5:0255181bb812302d
func (c *Collection) X插入(上下文 context.Context, 待插入文档 interface{}, 可选选项 ...opts.InsertOneOptions) (插入结果 *InsertOneResult, 错误 error) {
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
		插入结果 = &InsertOneResult{X插入ID: res.InsertedID}
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 待插入文档, mgo常量.X钩子_插入后, h); 错误 != nil {
		return
	}
	return
}

// InsertMany 执行插入命令，将多个文档插入到集合中。如果opts中的InsertHook被设置，将在其上应用钩子；否则，尝试将doc作为钩子使用。
// 参考：https://docs.mongodb.com/manual/reference/command/insert/
// md5:49f2d7776e74fa79
func (c *Collection) X插入多个(上下文 context.Context, 待插入文档 interface{}, 可选选项 ...opts.InsertManyOptions) (插入结果 *InsertManyResult, 错误 error) {
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
		插入结果 = &InsertManyResult{X插入IDs: res.InsertedIDs}
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 待插入文档, mgo常量.X钩子_插入后, h); 错误 != nil {
		return
	}
	return
}

// interfaceToSliceInterface 将接口类型转换为切片接口类型 md5:49f6ad81d7f669e3
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

// Upsert 如果过滤器匹配，则更新一个文档，如果过滤器不匹配，则插入一个文档。如果过滤器无效，会返回错误。
// 替换参数必须是一个将用于替换所选文档的文档。它不能为nil，且不能包含任何更新运算符。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// 如果替换参数包含"_id"字段并且文档已存在，请使用现有ID初始化（即使使用Qmgo默认字段特性也是如此）。否则，会引发"（不可变）字段 '_id' 被修改"的错误。
// md5:d7464af9e1e36d77
func (c *Collection) X替换插入(上下文 context.Context, 替换条件 interface{}, 替换内容 interface{}, 可选选项 ...opts.UpsertOptions) (结果 *UpdateResult, 错误 error) {
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

// UpsertId 如果id匹配，则更新一个文档，如果id不匹配，则插入一个新的文档，并将id注入到文档中。
// 注意,id是十六进制, 不是文本型, 需要转换后查询.
// 替换参数必须是一个将用于替换所选文档的文档。它不能为空，并且不能包含任何更新操作符。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:2a704aa664092959
func (c *Collection) X替换插入并按ID(上下文 context.Context, 替换ID interface{}, 替换内容 interface{}, 可选选项 ...opts.UpsertOptions) (结果 *UpdateResult, 错误 error) {
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

// UpdateOne 执行一个更新命令，最多更新集合中的一份文档。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:a16e90f28370dc2c
func (c *Collection) X更新一条(上下文 context.Context, 更新条件 interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (错误 error) {
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
		// UpdateOne支持upsert功能 md5:aaec7189323f1660
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

// UpdateId 执行一个更新命令，最多更新集合中的一个文档。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:67764db9e90007e8
func (c *Collection) X更新并按ID(上下文 context.Context, 更新ID interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (错误 error) {
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

// UpdateAll 执行更新命令以更新集合中的文档。
// 如果没有文档被更新，UpdateResult 中的 matchedCount 将为 0
// 参考资料: https://docs.mongodb.com/manual/reference/operator/update/
// md5:94c36e4a82312809
func (c *Collection) X更新(上下文 context.Context, 更新条件 interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (更新结果 *UpdateResult, 错误 error) {
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

// ReplaceOne 执行更新命令，最多更新集合中的一个文档。如果 opts 中的 UpdateHook 被设置，那么 Hook 将在其上执行，否则 Hook 尝试将 doc 作为 Hook。预期 doc 的类型是用户定义的文档的定义。
// md5:1d830477f8b32e37
func (c *Collection) X替换一条(上下文 context.Context, 替换条件 interface{}, 替换内容 interface{}, 可选选项 ...opts.ReplaceOptions) (错误 error) {
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

// Remove 执行删除命令，从集合中最多删除一个文档。
// 如果 filter 是 bson.M{}，DeleteOne 将删除集合中的一个文档。
// 参考：https://docs.mongodb.com/manual/reference/command/delete/
// md5:3b5cf64ce5f460b0
func (c *Collection) X删除一条(上下文 context.Context, 删除条件 interface{}, 可选选项 ...opts.RemoveOptions) (错误 error) {
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

// RemoveId 执行删除命令，从集合中删除最多一个文档。 md5:6516d8a8963d018c
func (c *Collection) X删除并按ID(上下文 context.Context, 删除ID interface{}, 可选选项 ...opts.RemoveOptions) (错误 error) {
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

// RemoveAll 执行一个删除命令，从集合中删除文档。如果 filter 是 bson.M{}（空映射），则会删除集合中的所有文档。
// 参考：https://docs.mongodb.com/manual/reference/command/delete/
// md5:abc51456adc5fc5a
func (c *Collection) X删除(上下文 context.Context, 删除条件 interface{}, 可选选项 ...opts.RemoveOptions) (删除结果 *DeleteResult, 错误 error) {
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
		删除结果 = &DeleteResult{X删除数量: res.DeletedCount}
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

// Aggregate 在集合上执行聚合命令，并返回一个 AggregateI，用于获取结果文档。 md5:e57ffed517c59fbc
func (c *Collection) X聚合(上下文 context.Context, 聚合管道 interface{}, 可选选项 ...opts.AggregateOptions) AggregateI {
	return &Aggregate{
		ctx:        上下文,
		collection: c.collection,
		pipeline:   聚合管道,
		options:    可选选项,
	}
}

// ensureIndex 在集合上创建多个索引，并返回这些索引的名称。
// 示例：indexes = []string{"idx1", "-idx2", "idx3,idx4"}
// 将创建三个索引，idx1 为升序索引，idx2 为降序索引，idx3 和 idx4 为复合升序索引。
// 参考文档：https://docs.mongodb.com/manual/reference/command/createIndexes/
// md5:50a25575e53025b2
func (c *Collection) ensureIndex(ctx context.Context, indexes []opts.IndexModel) error {
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

// 确保索引（已弃用）
// 建议使用CreateIndexes / CreateOneIndex以获取更多功能）
// EnsureIndexes 在集合中创建唯一和非唯一的索引，与CreateIndexes的组合不同：
// 如果uniques/indexes是`[]string{"name"}`，意味着创建名为"name"的索引
// 如果uniques/indexes是`[]string{"name,-age", "uid"}`，表示创建复合索引：name和-age，然后创建一个索引：uid
// md5:c595ad59f9c60c06
func (c *Collection) EnsureIndexes弃用(ctx context.Context, uniques []string, indexes []string) (err error) {
	var uniqueModel []opts.IndexModel
	var indexesModel []opts.IndexModel
	for _, v := range uniques {
		vv := strings.Split(v, ",")
		indexOpts := options.Index()
		indexOpts.SetUnique(true)
		model := opts.IndexModel{X索引字段: vv, IndexOptions: indexOpts}
		uniqueModel = append(uniqueModel, model)
	}
	if err = c.X创建多条索引(ctx, uniqueModel); err != nil {
		return
	}

	for _, v := range indexes {
		vv := strings.Split(v, ",")
		model := opts.IndexModel{X索引字段: vv}
		indexesModel = append(indexesModel, model)
	}
	if err = c.X创建多条索引(ctx, indexesModel); err != nil {
		return
	}
	return
}

// CreateIndexes 在集合中创建多个索引
// 如果opts.IndexModel中的Key为[]string{"name"}，表示创建索引：name
// 如果opts.IndexModel中的Key为[]string{"name", "-age"}，表示创建复合索引：name和-age
// md5:822a787892c2186f索引s
func (c *Collection) X创建多条索引(上下文 context.Context, 索引s []opts.IndexModel) (错误 error) {
	错误 = c.ensureIndex(上下文, 索引s)
	return
}

// CreateOneIndex 创建一个索引
// 如果opts.IndexModel中的Key为[]string{"name"}，表示创建名为"name"的索引
// 如果opts.IndexModel中的Key为[]string{"name","-age"}，表示创建复合索引：按照"name"升序和"age"降序
// md5:70c27ea42ff3bbbf
func (c *Collection) X创建索引(上下文 context.Context, 索引 opts.IndexModel) error {
	return c.ensureIndex(上下文, []opts.IndexModel{索引})

}

// DropAllIndexes 会删除集合上除了_id字段索引之外的所有索引
// 如果集合上只有_id字段的索引，该函数调用将报告错误
// md5:e7655b40436f93df全部索引
func (c *Collection) X删除全部索引(上下文 context.Context) (错误 error) {
	_, 错误 = c.collection.Indexes().DropAll(上下文)
	return 错误
}

// DropIndex 从集合中删除索引，需要删除的索引应与输入的索引列表匹配
// 索引是 []string{"name"} 表示删除名为 "name" 的单个索引
// 索引是 []string{"name", "-age"} 表示删除复合索引：name 和排除年龄 (-age) 的部分索引
// md5:4ad77e88557061c7索引索引s
func (c *Collection) X删除索引(上下文 context.Context, 索引s []string) error {
	_, err := c.collection.Indexes().DropOne(上下文, generateDroppedIndex(索引s))
	if err != nil {
		return err
	}
	return err
}

// 生成存储在Mongo中的索引，可能包含多个索引（如[]string{"index1","index2"}存储为"index1_1_index2_1"） md5:15332a053c924233
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

// DropIndexDropIndex 会删除索引
// 即使索引不存在，这个操作也是安全的
// md5:e7b65cd93b1de7f7集合
func (c *Collection) X删除集合(上下文 context.Context) error {
	return c.collection.Drop(上下文)
}

// CloneCollection 创建集合的副本 md5:5df787f1c8ebab26
func (c *Collection) X取副本() (*mongo.Collection, error) {
	return c.collection.Clone()
}

// GetCollectionName 返回集合的名字 md5:440484db8f2a466d
func (c *Collection) X取集合名() string {
	return c.collection.Name()
}

// Watch 返回对应集合上所有更改的流。有关更改流的更多信息，请参阅
// md5:7b59cfd256c148f3
func (c *Collection) X取变更流(上下文 context.Context, 管道 interface{}, 可选选项 ...*opts.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	changeStreamOption := options.ChangeStream()
	if len(可选选项) > 0 && 可选选项[0].ChangeStreamOptions != nil {
		changeStreamOption = 可选选项[0].ChangeStreamOptions
	}
	return c.collection.Watch(上下文, 管道, changeStreamOption)
}

// translateUpdateResult 将Mongo的更新结果转换为qmgo定义的UpdateResult md5:cb683a73f25cfe75
func translateUpdateResult(res *mongo.UpdateResult) (result *UpdateResult) {
	result = &UpdateResult{
		X匹配数:  res.MatchedCount,
		X修改数: res.ModifiedCount,
		X替换插入数: res.UpsertedCount,
		X替换插入ID:    res.UpsertedID,
	}
	return
}
