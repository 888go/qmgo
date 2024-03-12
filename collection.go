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

	"github.com/qiniu/qmgo/middleware"
	"github.com/qiniu/qmgo/operator"
	opts "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection 是对 MongoDB 集合的一个引用句柄
type Collection struct {
	collection *mongo.Collection

	registry *bsoncodec.Registry
}

// Find find by condition filter，return QueryI
func (c *Collection) X查询(上下文 context.Context, 查询条件 interface{}, 可选选项 ...opts.FindOptions) QueryI {

	return &Query{
		ctx:        上下文,
		collection: c.collection,
		filter:     查询条件,
		opts:       可选选项,
		registry:   c.registry,
	}
}

// InsertOne 将一个文档插入集合中
// 如果 opts 中设置了 InsertHook，那么钩子会作用于它，否则钩子尝试将文档当作钩子处理
// 参考：https://docs.mongodb.com/manual/reference/command/insert/
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
	if 错误 = middleware.Do(上下文, 待插入文档, operator.BeforeInsert, h); 错误 != nil {
		return
	}
	res, 错误 := c.collection.InsertOne(上下文, 待插入文档, insertOneOpts)
	if res != nil {
		插入结果 = &InsertOneResult{InsertedID: res.InsertedID}
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 待插入文档, operator.AfterInsert, h); 错误 != nil {
		return
	}
	return
}

// InsertMany 执行一个插入命令，将多个文档插入到集合中。
// 如果 opts 中设置了 InsertHook，则在该 hook 上执行操作；否则尝试将 doc 作为 hook 使用
// 参考文献：https://docs.mongodb.com/manual/reference/command/insert/
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
	if 错误 = middleware.Do(上下文, 待插入文档, operator.BeforeInsert, h); 错误 != nil {
		return
	}
	sDocs := interfaceToSliceInterface(待插入文档)
	if sDocs == nil {
		return nil, ErrNotValidSliceToInsert
	}

	res, 错误 := c.collection.InsertMany(上下文, sDocs, insertManyOpts)
	if res != nil {
		插入结果 = &InsertManyResult{InsertedIDs: res.InsertedIDs}
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 待插入文档, operator.AfterInsert, h); 错误 != nil {
		return
	}
	return
}

// interfaceToSliceInterface 将接口转换为切片接口
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

// Upsert：如果过滤条件匹配，则更新一条文档；如果不匹配，则插入一条文档。当过滤条件无效时，将返回错误。
// replacement 参数必须是一个用于替换所选文档的文档对象，不能为 nil，并且不能包含任何更新操作符。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
// 如果 replacement 中包含 "_id" 字段，并且该文档已存在，请确保初始化时使用现有 id（即使启用了 Qmgo 的默认字段特性）。
// 否则，将会出现 "（不可变字段）'_id' 被修改" 的错误。
func (c *Collection) X更新或插入(上下文 context.Context, 更新条件 interface{}, 更新内容 interface{}, 可选选项 ...opts.UpsertOptions) (更新结果 *UpdateResult, 错误 error) {
	h := 更新内容
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
	if 错误 = middleware.Do(上下文, 更新内容, operator.BeforeUpsert, h); 错误 != nil {
		return
	}

	res, 错误 := c.collection.ReplaceOne(上下文, 更新条件, 更新内容, officialOpts)

	if res != nil {
		更新结果 = translateUpdateResult(res)
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 更新内容, operator.AfterUpsert, h); 错误 != nil {
		return
	}
	return
}

// UpsertId 如果_id匹配则更新一条文档，如果不匹配则插入一条文档，并将_id注入到该文档中
// replacement参数必须是一个用于替换所选文档的文档，不能为nil
// 且不能包含任何更新操作符
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) X更新或插入并按ID(上下文 context.Context, 更新ID interface{}, 更新内容 interface{}, 可选选项 ...opts.UpsertOptions) (更新结果 *UpdateResult, 错误 error) {
	h := 更新内容
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
	if 错误 = middleware.Do(上下文, 更新内容, operator.BeforeUpsert, h); 错误 != nil {
		return
	}
	res, 错误 := c.collection.ReplaceOne(上下文, bson.M{"_id": 更新ID}, 更新内容, officialOpts)
	if res != nil {
		更新结果 = translateUpdateResult(res)
	}
	if 错误 != nil {
		return
	}
	if 错误 = middleware.Do(上下文, 更新内容, operator.AfterUpsert, h); 错误 != nil {
		return
	}
	return
}

// UpdateOne 执行一个更新命令，用于在集合中最多更新一条文档。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) X更新一条(上下文 context.Context, 更新条件 interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (错误 error) {
	updateOpts := options.Update()

	if len(可选选项) > 0 {
		if 可选选项[0].UpdateOptions != nil {
			updateOpts = 可选选项[0].UpdateOptions
		}
		if 可选选项[0].UpdateHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, operator.BeforeUpdate); 错误 != nil {
				return
			}
		}
	}

	res, 错误 := c.collection.UpdateOne(上下文, 更新条件, 更新内容, updateOpts)
	if res != nil && res.MatchedCount == 0 {
		// UpdateOne 支持 Upsert 功能
		if updateOpts.Upsert == nil || !*updateOpts.Upsert {
			错误 = ErrNoSuchDocuments
		}
	}
	if 错误 != nil {
		return 错误
	}
	if len(可选选项) > 0 && 可选选项[0].UpdateHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, operator.AfterUpdate); 错误 != nil {
			return
		}
	}
	return 错误
}

// UpdateId 执行一个更新命令，用于在集合中最多更新一个文档。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) X更新并按ID(上下文 context.Context, 更新ID interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (错误 error) {
	updateOpts := options.Update()

	if len(可选选项) > 0 {
		if 可选选项[0].UpdateOptions != nil {
			updateOpts = 可选选项[0].UpdateOptions
		}
		if 可选选项[0].UpdateHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, operator.BeforeUpdate); 错误 != nil {
				return
			}
		}
	}

	res, 错误 := c.collection.UpdateOne(上下文, bson.M{"_id": 更新ID}, 更新内容, updateOpts)
	if res != nil && res.MatchedCount == 0 {
		错误 = ErrNoSuchDocuments
	}
	if 错误 != nil {
		return 错误
	}
	if len(可选选项) > 0 && 可选选项[0].UpdateHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, operator.AfterUpdate); 错误 != nil {
			return
		}
	}
	return 错误
}

// UpdateAll 执行一个更新命令，用于更新集合中的文档。
// 如果没有文档被更新，UpdateResult 中的 matchedCount 为 0
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) X更新(上下文 context.Context, 更新条件 interface{}, 更新内容 interface{}, 可选选项 ...opts.UpdateOptions) (更新结果 *UpdateResult, 错误 error) {
	updateOpts := options.Update()
	if len(可选选项) > 0 {
		if 可选选项[0].UpdateOptions != nil {
			updateOpts = 可选选项[0].UpdateOptions
		}
		if 可选选项[0].UpdateHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, operator.BeforeUpdate); 错误 != nil {
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
		if 错误 = middleware.Do(上下文, 可选选项[0].UpdateHook, operator.AfterUpdate); 错误 != nil {
			return
		}
	}
	return
}

// ReplaceOne 执行一个更新命令，最多替换集合中的一个文档。
// 如果opts中的UpdateHook已设置，则在该hook上执行操作，否则尝试将doc作为hook处理
// 预期doc的类型应为用户定义的文档类型
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
	if 错误 = middleware.Do(上下文, 替换内容, operator.BeforeReplace, h); 错误 != nil {
		return
	}
	res, 错误 := c.collection.ReplaceOne(上下文, 替换条件, 替换内容, replaceOpts)
	if res != nil && res.MatchedCount == 0 {
		错误 = ErrNoSuchDocuments
	}
	if 错误 != nil {
		return 错误
	}
	if 错误 = middleware.Do(上下文, 替换内容, operator.AfterReplace, h); 错误 != nil {
		return
	}

	return 错误
}

// Remove executes a delete command to delete at most one document from the collection.
// if filter is bson.M{}，DeleteOne will delete one document in collection
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
func (c *Collection) X删除一条(上下文 context.Context, 删除条件 interface{}, 可选选项 ...opts.RemoveOptions) (错误 error) {
	deleteOptions := options.Delete()
	if len(可选选项) > 0 {
		if 可选选项[0].DeleteOptions != nil {
			deleteOptions = 可选选项[0].DeleteOptions
		}
		if 可选选项[0].RemoveHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, operator.BeforeRemove); 错误 != nil {
				return 错误
			}
		}
	}
	res, 错误 := c.collection.DeleteOne(上下文, 删除条件, deleteOptions)
	if res != nil && res.DeletedCount == 0 {
		错误 = ErrNoSuchDocuments
	}
	if 错误 != nil {
		return 错误
	}
	if len(可选选项) > 0 && 可选选项[0].RemoveHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, operator.AfterRemove); 错误 != nil {
			return 错误
		}
	}
	return 错误
}

// RemoveId 执行一个删除命令，从集合中最多删除一个文档。
func (c *Collection) X删除并按ID(上下文 context.Context, 删除ID interface{}, 可选选项 ...opts.RemoveOptions) (错误 error) {
	deleteOptions := options.Delete()
	if len(可选选项) > 0 {
		if 可选选项[0].DeleteOptions != nil {
			deleteOptions = 可选选项[0].DeleteOptions
		}
		if 可选选项[0].RemoveHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, operator.BeforeRemove); 错误 != nil {
				return 错误
			}
		}
	}
	res, 错误 := c.collection.DeleteOne(上下文, bson.M{"_id": 删除ID}, deleteOptions)
	if res != nil && res.DeletedCount == 0 {
		错误 = ErrNoSuchDocuments
	}
	if 错误 != nil {
		return 错误
	}

	if len(可选选项) > 0 && 可选选项[0].RemoveHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, operator.AfterRemove); 错误 != nil {
			return 错误
		}
	}
	return 错误
}

// RemoveAll executes a delete command to delete documents from the collection.
// If filter is bson.M{}，all ducuments in Collection will be deleted
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
func (c *Collection) X删除(上下文 context.Context, 删除条件 interface{}, 可选选项 ...opts.RemoveOptions) (删除结果 *DeleteResult, 错误 error) {
	deleteOptions := options.Delete()
	if len(可选选项) > 0 {
		if 可选选项[0].DeleteOptions != nil {
			deleteOptions = 可选选项[0].DeleteOptions
		}
		if 可选选项[0].RemoveHook != nil {
			if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, operator.BeforeRemove); 错误 != nil {
				return
			}
		}
	}
	res, 错误 := c.collection.DeleteMany(上下文, 删除条件, deleteOptions)
	if res != nil {
		删除结果 = &DeleteResult{DeletedCount: res.DeletedCount}
	}
	if 错误 != nil {
		return
	}
	if len(可选选项) > 0 && 可选选项[0].RemoveHook != nil {
		if 错误 = middleware.Do(上下文, 可选选项[0].RemoveHook, operator.AfterRemove); 错误 != nil {
			return
		}
	}
	return
}

// Aggregate 对集合执行聚合命令，并返回一个 AggregateI 以便获取结果文档。
func (c *Collection) X聚合(上下文 context.Context, 聚合管道 interface{}, 可选选项 ...opts.AggregateOptions) AggregateI {
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
func (c *Collection) ensureIndex(ctx context.Context, indexes []opts.IndexModel) error {
	var indexModels []mongo.IndexModel
	for _, idx := range indexes {
		var model mongo.IndexModel
		var keysDoc bson.D

		for _, field := range idx.Key {
			key, n := SplitSortField(field)

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

// EnsureIndexes 已弃用
// 建议使用 CreateIndexes / CreateOneIndex 以获得更多的功能)
// EnsureIndexes 在集合中创建唯一索引和非唯一索引
// 索引的组合方式与 CreateIndexes 不同：
// 如果 uniques/indexes 是 []string{"name"}，表示创建名为 "name" 的索引
// 如果 uniques/indexes 是 []string{"name,-age","uid"}，表示首先创建复合索引：name 和 -age（按 name 升序、age 降序），然后创建一个名为 uid 的单字段索引
func (c *Collection) EnsureIndexes弃用(ctx context.Context, uniques []string, indexes []string) (err error) {
	var uniqueModel []opts.IndexModel
	var indexesModel []opts.IndexModel
	for _, v := range uniques {
		vv := strings.Split(v, ",")
		indexOpts := options.Index()
		indexOpts.SetUnique(true)
		model := opts.IndexModel{Key: vv, IndexOptions: indexOpts}
		uniqueModel = append(uniqueModel, model)
	}
	if err = c.X索引多条(ctx, uniqueModel); err != nil {
		return
	}

	for _, v := range indexes {
		vv := strings.Split(v, ",")
		model := opts.IndexModel{Key: vv}
		indexesModel = append(indexesModel, model)
	}
	if err = c.X索引多条(ctx, indexesModel); err != nil {
		return
	}
	return
}

// CreateIndexes 在集合中创建多个索引
// 如果opts.IndexModel中的Key为[]string{"name"}，表示创建名为"name"的索引
// 如果opts.IndexModel中的Key为[]string{"name","-age"}，表示创建复合索引：按"name"和"-age"（降序）字段
// 进一步详细解释：
// ```go
// CreateIndexes 函数用于在指定的数据库集合中创建多个索引。
// 当 opts.IndexModel 中的 Key 字段是一个包含单个元素 "name" 的字符串切片时，例如 []string{"name"}，
// 这意味着将根据字段 "name" 创建一个升序索引。
// 若 opts.IndexModel 中的 Key 字段是一个包含两个元素 "name" 和 "-age" 的字符串切片，例如 []string{"name", "-age"}，
// 这表示将创建一个复合索引，其中先按 "name" 字段升序排序，然后按 "age" 字段降序排序。
func (c *Collection) X索引多条(上下文 context.Context, 索引s []opts.IndexModel) (错误 error) {
	错误 = c.ensureIndex(上下文, 索引s)
	return
}

// CreateOneIndex 创建一个索引
// 如果 opts.IndexModel 中的 Key 为 []string{"name"}，表示创建名为 "name" 的索引
// 如果 opts.IndexModel 中的 Key 为 []string{"name", "-age"}，表示创建组合索引：包含 name 和 -age（按 name 正序、age 倒序）
func (c *Collection) X索引一条(上下文 context.Context, 索引 opts.IndexModel) error {
	return c.ensureIndex(上下文, []opts.IndexModel{索引})

}

// DropAllIndexes 从集合中删除所有索引，但保留_id字段的索引
// 如果集合上只有_id字段的索引，函数调用将报告错误
func (c *Collection) X删除全部索引(上下文 context.Context) (错误 error) {
	_, 错误 = c.collection.Indexes().DropAll(上下文)
	return 错误
}

// DropIndex 删除集合中的索引，需要删除的索引应与输入的索引一致
// 索引indexes为[]string{"name"}表示删除名为"name"的索引
// 索引indexes为[]string{"name","-age"}表示删除由"name"和"-age"组成的复合索引
func (c *Collection) X删除索引(上下文 context.Context, 索引s []string) error {
	_, err := c.collection.Indexes().DropOne(上下文, generateDroppedIndex(索引s))
	if err != nil {
		return err
	}
	return err
}

// 生成存储在MongoDB中的索引，这些索引可能包含多个索引（例如，[]string{"index1","index2"}将被存储为"index1_1_index2_1"）
func generateDroppedIndex(index []string) string {
	var res string
	for _, e := range index {
		key, sort := SplitSortField(e)
		n := key + "_" + fmt.Sprint(sort)
		if len(res) == 0 {
			res = n
		} else {
			res += "_" + n
		}
	}
	return res
}

// DropCollection 删除集合
// 即使集合不存在，此操作也是安全的
func (c *Collection) X删除集合(上下文 context.Context) error {
	return c.collection.Drop(上下文)
}

// CloneCollection 创建一个 Collection 的副本
func (c *Collection) X取副本() (*mongo.Collection, error) {
	return c.collection.Clone()
}

// GetCollectionName 返回集合名称
func (c *Collection) X取集合名() string {
	return c.collection.Name()
}

// Watch 返回一个变更流，用于接收相应集合的所有变更。有关变更流的更多信息，请参阅
// https://docs.mongodb.com/manual/changeStreams/
func (c *Collection) X取变更流(上下文 context.Context, 管道 interface{}, 可选选项 ...*opts.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	changeStreamOption := options.ChangeStream()
	if len(可选选项) > 0 && 可选选项[0].ChangeStreamOptions != nil {
		changeStreamOption = 可选选项[0].ChangeStreamOptions
	}
	return c.collection.Watch(上下文, 管道, changeStreamOption)
}

// translateUpdateResult 将MongoDB更新结果翻译为qmgo定义的UpdateResult
func translateUpdateResult(res *mongo.UpdateResult) (result *UpdateResult) {
	result = &UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}
	return
}
