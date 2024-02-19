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

package qmgo

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

// 查询
// ctx:上下文
// filter:查询条件
// opts:可选选项
// Find按条件过滤器查找，返回QueryI
func (c *Collection) Find(ctx context.Context, filter interface{}, opts ...opts.FindOptions) QueryI {
	return &Query{
		ctx:        ctx,
		collection: c.collection,
		filter:     filter,
		opts:       opts,
		registry:   c.registry,
	}
}

// 插入
// ctx:上下文
// doc:待插入文档
// opts:可选选项
// result:插入结果
// err:错误
// InsertOne 将一个文档插入集合中
// 如果 opts 中设置了 InsertHook，那么钩子会作用于它，否则钩子尝试将文档当作钩子处理
// 参考：https://docs.mongodb.com/manual/reference/command/insert/
func (c *Collection) InsertOne(ctx context.Context, doc interface{}, opts ...opts.InsertOneOptions) (result *InsertOneResult, err error) {
	h := doc
	insertOneOpts := options.InsertOne()
	if len(opts) > 0 {
		if opts[0].InsertOneOptions != nil {
			insertOneOpts = opts[0].InsertOneOptions
		}
		if opts[0].InsertHook != nil {
			h = opts[0].InsertHook
		}
	}
	if err = middleware.Do(ctx, doc, operator.BeforeInsert, h); err != nil {
		return
	}
	res, err := c.collection.InsertOne(ctx, doc, insertOneOpts)
	if res != nil {
		result = &InsertOneResult{InsertedID: res.InsertedID}
	}
	if err != nil {
		return
	}
	if err = middleware.Do(ctx, doc, operator.AfterInsert, h); err != nil {
		return
	}
	return
}

// 插入多个
// ctx:上下文
// docs:待插入文档
// opts:可选选项
// result:插入结果
// err:错误
// InsertMany 执行一个插入命令，将多个文档插入到集合中。
// 如果 opts 中设置了 InsertHook，则在该 hook 上执行操作；否则尝试将 doc 作为 hook 使用
// 参考文献：https://docs.mongodb.com/manual/reference/command/insert/
func (c *Collection) InsertMany(ctx context.Context, docs interface{}, opts ...opts.InsertManyOptions) (result *InsertManyResult, err error) {
	h := docs
	insertManyOpts := options.InsertMany()
	if len(opts) > 0 {
		if opts[0].InsertManyOptions != nil {
			insertManyOpts = opts[0].InsertManyOptions
		}
		if opts[0].InsertHook != nil {
			h = opts[0].InsertHook
		}
	}
	if err = middleware.Do(ctx, docs, operator.BeforeInsert, h); err != nil {
		return
	}
	sDocs := interfaceToSliceInterface(docs)
	if sDocs == nil {
		return nil, ErrNotValidSliceToInsert
	}

	res, err := c.collection.InsertMany(ctx, sDocs, insertManyOpts)
	if res != nil {
		result = &InsertManyResult{InsertedIDs: res.InsertedIDs}
	}
	if err != nil {
		return
	}
	if err = middleware.Do(ctx, docs, operator.AfterInsert, h); err != nil {
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

// 更新或插入
// ctx:上下文
// filter:更新条件
// replacement:更新内容
// opts:可选选项
// result:更新结果
// err:错误
// Upsert：如果过滤条件匹配，则更新一条文档；如果不匹配，则插入一条文档。当过滤条件无效时，将返回错误。
// replacement 参数必须是一个用于替换所选文档的文档对象，不能为 nil，并且不能包含任何更新操作符。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
// 如果 replacement 中包含 "_id" 字段，并且该文档已存在，请确保初始化时使用现有 id（即使启用了 Qmgo 的默认字段特性）。
// 否则，将会出现 "（不可变字段）'_id' 被修改" 的错误。
func (c *Collection) Upsert(ctx context.Context, filter interface{}, replacement interface{}, opts ...opts.UpsertOptions) (result *UpdateResult, err error) {
	h := replacement
	officialOpts := options.Replace().SetUpsert(true)

	if len(opts) > 0 {
		if opts[0].ReplaceOptions != nil {
			opts[0].ReplaceOptions.SetUpsert(true)
			officialOpts = opts[0].ReplaceOptions
		}
		if opts[0].UpsertHook != nil {
			h = opts[0].UpsertHook
		}
	}
	if err = middleware.Do(ctx, replacement, operator.BeforeUpsert, h); err != nil {
		return
	}

	res, err := c.collection.ReplaceOne(ctx, filter, replacement, officialOpts)

	if res != nil {
		result = translateUpdateResult(res)
	}
	if err != nil {
		return
	}
	if err = middleware.Do(ctx, replacement, operator.AfterUpsert, h); err != nil {
		return
	}
	return
}

// 更新或插入并按ID
// ctx:上下文
// id:更新ID
// replacement:更新内容
// opts:可选选项
// result:更新结果
// err:错误
// UpsertId 如果_id匹配则更新一条文档，如果不匹配则插入一条文档，并将_id注入到该文档中
// replacement参数必须是一个用于替换所选文档的文档，不能为nil
// 且不能包含任何更新操作符
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) UpsertId(ctx context.Context, id interface{}, replacement interface{}, opts ...opts.UpsertOptions) (result *UpdateResult, err error) {
	h := replacement
	officialOpts := options.Replace().SetUpsert(true)

	if len(opts) > 0 {
		if opts[0].ReplaceOptions != nil {
			opts[0].ReplaceOptions.SetUpsert(true)
			officialOpts = opts[0].ReplaceOptions
		}
		if opts[0].UpsertHook != nil {
			h = opts[0].UpsertHook
		}
	}
	if err = middleware.Do(ctx, replacement, operator.BeforeUpsert, h); err != nil {
		return
	}
	res, err := c.collection.ReplaceOne(ctx, bson.M{"_id": id}, replacement, officialOpts)
	if res != nil {
		result = translateUpdateResult(res)
	}
	if err != nil {
		return
	}
	if err = middleware.Do(ctx, replacement, operator.AfterUpsert, h); err != nil {
		return
	}
	return
}

// 更新一条
// ctx:上下文
// filter:更新条件
// update:更新内容
// opts:可选选项
// err:错误
// UpdateOne 执行一个更新命令，用于在集合中最多更新一条文档。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...opts.UpdateOptions) (err error) {
	updateOpts := options.Update()

	if len(opts) > 0 {
		if opts[0].UpdateOptions != nil {
			updateOpts = opts[0].UpdateOptions
		}
		if opts[0].UpdateHook != nil {
			if err = middleware.Do(ctx, opts[0].UpdateHook, operator.BeforeUpdate); err != nil {
				return
			}
		}
	}

	res, err := c.collection.UpdateOne(ctx, filter, update, updateOpts)
	if res != nil && res.MatchedCount == 0 {
		// UpdateOne 支持 Upsert 功能
		if updateOpts.Upsert == nil || !*updateOpts.Upsert {
			err = ErrNoSuchDocuments
		}
	}
	if err != nil {
		return err
	}
	if len(opts) > 0 && opts[0].UpdateHook != nil {
		if err = middleware.Do(ctx, opts[0].UpdateHook, operator.AfterUpdate); err != nil {
			return
		}
	}
	return err
}

// 更新并按ID
// ctx:上下文
// id:更新ID
// update:更新内容
// opts:可选选项
// err:错误
// UpdateId 执行一个更新命令，用于在集合中最多更新一个文档。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) UpdateId(ctx context.Context, id interface{}, update interface{}, opts ...opts.UpdateOptions) (err error) {
	updateOpts := options.Update()

	if len(opts) > 0 {
		if opts[0].UpdateOptions != nil {
			updateOpts = opts[0].UpdateOptions
		}
		if opts[0].UpdateHook != nil {
			if err = middleware.Do(ctx, opts[0].UpdateHook, operator.BeforeUpdate); err != nil {
				return
			}
		}
	}

	res, err := c.collection.UpdateOne(ctx, bson.M{"_id": id}, update, updateOpts)
	if res != nil && res.MatchedCount == 0 {
		err = ErrNoSuchDocuments
	}
	if err != nil {
		return err
	}
	if len(opts) > 0 && opts[0].UpdateHook != nil {
		if err = middleware.Do(ctx, opts[0].UpdateHook, operator.AfterUpdate); err != nil {
			return
		}
	}
	return err
}

// 更新
// ctx:上下文
// filter:更新条件
// update:更新内容
// opts:可选选项
// result:更新结果
// err:错误
// UpdateAll 执行一个更新命令，用于更新集合中的文档。
// 如果没有文档被更新，UpdateResult 中的 matchedCount 为 0
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
func (c *Collection) UpdateAll(ctx context.Context, filter interface{}, update interface{}, opts ...opts.UpdateOptions) (result *UpdateResult, err error) {
	updateOpts := options.Update()
	if len(opts) > 0 {
		if opts[0].UpdateOptions != nil {
			updateOpts = opts[0].UpdateOptions
		}
		if opts[0].UpdateHook != nil {
			if err = middleware.Do(ctx, opts[0].UpdateHook, operator.BeforeUpdate); err != nil {
				return
			}
		}
	}
	res, err := c.collection.UpdateMany(ctx, filter, update, updateOpts)
	if res != nil {
		result = translateUpdateResult(res)
	}
	if err != nil {
		return
	}
	if len(opts) > 0 && opts[0].UpdateHook != nil {
		if err = middleware.Do(ctx, opts[0].UpdateHook, operator.AfterUpdate); err != nil {
			return
		}
	}
	return
}

// 替换一条
// ctx:上下文
// filter:替换条件
// doc:替换内容
// opts:可选选项
// err:错误
// ReplaceOne 执行一个更新命令，最多替换集合中的一个文档。
// 如果opts中的UpdateHook已设置，则在该hook上执行操作，否则尝试将doc作为hook处理
// 预期doc的类型应为用户定义的文档类型
// 这段Go语言代码注释翻译成中文注释如下：
func (c *Collection) ReplaceOne(ctx context.Context, filter interface{}, doc interface{}, opts ...opts.ReplaceOptions) (err error) {
	h := doc
	replaceOpts := options.Replace()

	if len(opts) > 0 {
		if opts[0].ReplaceOptions != nil {
			replaceOpts = opts[0].ReplaceOptions
			replaceOpts.SetUpsert(false)
		}
		if opts[0].UpdateHook != nil {
			h = opts[0].UpdateHook
		}
	}
	if err = middleware.Do(ctx, doc, operator.BeforeReplace, h); err != nil {
		return
	}
	res, err := c.collection.ReplaceOne(ctx, filter, doc, replaceOpts)
	if res != nil && res.MatchedCount == 0 {
		err = ErrNoSuchDocuments
	}
	if err != nil {
		return err
	}
	if err = middleware.Do(ctx, doc, operator.AfterReplace, h); err != nil {
		return
	}

	return err
}

// 删除一条
// ctx:上下文
// filter:删除条件
// opts:可选选项
// err:错误
// Remove执行delete命令，从集合中最多删除一个文档。
// 如果filter是bson。M{}，DeleteOne将删除集合中的一个文档
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
func (c *Collection) Remove(ctx context.Context, filter interface{}, opts ...opts.RemoveOptions) (err error) {
	deleteOptions := options.Delete()
	if len(opts) > 0 {
		if opts[0].DeleteOptions != nil {
			deleteOptions = opts[0].DeleteOptions
		}
		if opts[0].RemoveHook != nil {
			if err = middleware.Do(ctx, opts[0].RemoveHook, operator.BeforeRemove); err != nil {
				return err
			}
		}
	}
	res, err := c.collection.DeleteOne(ctx, filter, deleteOptions)
	if res != nil && res.DeletedCount == 0 {
		err = ErrNoSuchDocuments
	}
	if err != nil {
		return err
	}
	if len(opts) > 0 && opts[0].RemoveHook != nil {
		if err = middleware.Do(ctx, opts[0].RemoveHook, operator.AfterRemove); err != nil {
			return err
		}
	}
	return err
}

// 删除并按ID
// ctx:上下文
// id:删除ID
// opts:可选选项
// err:错误
// RemoveId 执行一个删除命令，从集合中最多删除一个文档。
func (c *Collection) RemoveId(ctx context.Context, id interface{}, opts ...opts.RemoveOptions) (err error) {
	deleteOptions := options.Delete()
	if len(opts) > 0 {
		if opts[0].DeleteOptions != nil {
			deleteOptions = opts[0].DeleteOptions
		}
		if opts[0].RemoveHook != nil {
			if err = middleware.Do(ctx, opts[0].RemoveHook, operator.BeforeRemove); err != nil {
				return err
			}
		}
	}
	res, err := c.collection.DeleteOne(ctx, bson.M{"_id": id}, deleteOptions)
	if res != nil && res.DeletedCount == 0 {
		err = ErrNoSuchDocuments
	}
	if err != nil {
		return err
	}

	if len(opts) > 0 && opts[0].RemoveHook != nil {
		if err = middleware.Do(ctx, opts[0].RemoveHook, operator.AfterRemove); err != nil {
			return err
		}
	}
	return err
}

// 删除
// ctx:上下文
// filter:删除条件
// opts:可选选项
// result:删除结果
// err:错误
// RemoveAll执行delete命令从集合中删除文档。
// If filter is bson.M{}，all ducuments in Collection will be deleted
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
func (c *Collection) RemoveAll(ctx context.Context, filter interface{}, opts ...opts.RemoveOptions) (result *DeleteResult, err error) {
	deleteOptions := options.Delete()
	if len(opts) > 0 {
		if opts[0].DeleteOptions != nil {
			deleteOptions = opts[0].DeleteOptions
		}
		if opts[0].RemoveHook != nil {
			if err = middleware.Do(ctx, opts[0].RemoveHook, operator.BeforeRemove); err != nil {
				return
			}
		}
	}
	res, err := c.collection.DeleteMany(ctx, filter, deleteOptions)
	if res != nil {
		result = &DeleteResult{DeletedCount: res.DeletedCount}
	}
	if err != nil {
		return
	}
	if len(opts) > 0 && opts[0].RemoveHook != nil {
		if err = middleware.Do(ctx, opts[0].RemoveHook, operator.AfterRemove); err != nil {
			return
		}
	}
	return
}

// 聚合
// ctx:上下文
// pipeline:聚合管道
// opts:可选选项
// Aggregate 对集合执行聚合命令，并返回一个 AggregateI 以便获取结果文档。
func (c *Collection) Aggregate(ctx context.Context, pipeline interface{}, opts ...opts.AggregateOptions) AggregateI {
	return &Aggregate{
		ctx:        ctx,
		collection: c.collection,
		pipeline:   pipeline,
		options:    opts,
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

// EnsureIndexes弃用
// EnsureIndexes 已弃用
// 建议使用 CreateIndexes / CreateOneIndex 以获得更多的功能)
// EnsureIndexes 在集合中创建唯一索引和非唯一索引
// 索引的组合方式与 CreateIndexes 不同：
// 如果 uniques/indexes 是 []string{"name"}，表示创建名为 "name" 的索引
// 如果 uniques/indexes 是 []string{"name,-age","uid"}，表示首先创建复合索引：name 和 -age（按 name 升序、age 降序），然后创建一个名为 uid 的单字段索引
func (c *Collection) EnsureIndexes(ctx context.Context, uniques []string, indexes []string) (err error) {
	var uniqueModel []opts.IndexModel
	var indexesModel []opts.IndexModel
	for _, v := range uniques {
		vv := strings.Split(v, ",")
		indexOpts := options.Index()
		indexOpts.SetUnique(true)
		model := opts.IndexModel{Key: vv, IndexOptions: indexOpts}
		uniqueModel = append(uniqueModel, model)
	}
	if err = c.CreateIndexes(ctx, uniqueModel); err != nil {
		return
	}

	for _, v := range indexes {
		vv := strings.Split(v, ",")
		model := opts.IndexModel{Key: vv}
		indexesModel = append(indexesModel, model)
	}
	if err = c.CreateIndexes(ctx, indexesModel); err != nil {
		return
	}
	return
}

// 索引多条
// ctx:上下文
// indexes:索引s
// err:错误
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
func (c *Collection) CreateIndexes(ctx context.Context, indexes []opts.IndexModel) (err error) {
	err = c.ensureIndex(ctx, indexes)
	return
}

// 索引一条
// ctx:上下文
// index:索引
// CreateOneIndex 创建一个索引
// 如果 opts.IndexModel 中的 Key 为 []string{"name"}，表示创建名为 "name" 的索引
// 如果 opts.IndexModel 中的 Key 为 []string{"name", "-age"}，表示创建组合索引：包含 name 和 -age（按 name 正序、age 倒序）
func (c *Collection) CreateOneIndex(ctx context.Context, index opts.IndexModel) error {
	return c.ensureIndex(ctx, []opts.IndexModel{index})

}

// 删除全部索引
// ctx:上下文
// err:错误
// DropAllIndexes 从集合中删除所有索引，但保留_id字段的索引
// 如果集合上只有_id字段的索引，函数调用将报告错误
func (c *Collection) DropAllIndexes(ctx context.Context) (err error) {
	_, err = c.collection.Indexes().DropAll(ctx)
	return err
}

// 删除索引
// ctx:上下文
// indexes:索引s
// DropIndex 删除集合中的索引，需要删除的索引应与输入的索引一致
// 索引indexes为[]string{"name"}表示删除名为"name"的索引
// 索引indexes为[]string{"name","-age"}表示删除由"name"和"-age"组成的复合索引
func (c *Collection) DropIndex(ctx context.Context, indexes []string) error {
	_, err := c.collection.Indexes().DropOne(ctx, generateDroppedIndex(indexes))
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

// 删除集合
// ctx:上下文
// DropCollection 删除集合
// 即使集合不存在，此操作也是安全的
func (c *Collection) DropCollection(ctx context.Context) error {
	return c.collection.Drop(ctx)
}

// 取副本
// CloneCollection 创建一个 Collection 的副本
func (c *Collection) CloneCollection() (*mongo.Collection, error) {
	return c.collection.Clone()
}

// 取集合名
// GetCollectionName 返回集合名称
func (c *Collection) GetCollectionName() string {
	return c.collection.Name()
}

// 取变更流
// ctx:上下文
// pipeline:管道
// opts:可选选项
// Watch 返回一个变更流，用于接收相应集合的所有变更。有关变更流的更多信息，请参阅
// https://docs.mongodb.com/manual/changeStreams/
func (c *Collection) Watch(ctx context.Context, pipeline interface{}, opts ...*opts.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	changeStreamOption := options.ChangeStream()
	if len(opts) > 0 && opts[0].ChangeStreamOptions != nil {
		changeStreamOption = opts[0].ChangeStreamOptions
	}
	return c.collection.Watch(ctx, pipeline, changeStreamOption)
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
