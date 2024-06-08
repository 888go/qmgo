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

// Collection 是一个MongoDB集合的句柄 md5:be1b94030609bdd1
// [提示]
//
//	type 文档集合 struct {
//	    文档集合接口 *mongo.DocumentCollection
//	    编码注册器 *bsoncodec.Registry
//	}
//
// [结束]
type Collection struct { //hm:文档集合  cz:type Collection
	collection *mongo.Collection

	registry *bsoncodec.Registry
}

// Find 通过条件过滤并查找，返回QueryI md5:bda4cc0c85d800a1
// [提示:] func (c *集合) 查找(ctx 上下文, 过滤器 interface{})
// ff:查询
// c:
// ctx:上下文
// filter:查询条件
// opts:可选选项
func (c *Collection) Find(ctx context.Context, filter interface{}, opts ...opts.FindOptions) QueryI {
	return &Query{
		ctx:        ctx,
		collection: c.collection,
		filter:     filter,
		opts:       opts,
		registry:   c.registry,
	}
}

// InsertOne 将一个文档插入到集合中
// 如果在 opts 中设置了 InsertHook，钩子将在其上执行，否则钩子会尝试处理 doc 作为钩子
// 参考: https://docs.mongodb.com/manual/reference/command/insert/
// md5:0255181bb812302d
// [提示:] func (c *集合) 插入一个(ctx 上下文, 文档 interface{})
// ff:插入
// c:
// ctx:上下文
// doc:待插入文档
// opts:可选选项
// result:插入结果
// err:错误
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

// InsertMany 执行插入命令，将多个文档插入到集合中。如果opts中的InsertHook被设置，将在其上应用钩子；否则，尝试将doc作为钩子使用。
// 参考：https://docs.mongodb.com/manual/reference/command/insert/
// md5:49f2d7776e74fa79
// [提示]
// func (c *集合) 插入多条(ctx 上下文, 文档 interface{}) (插入结果 []interface{}, 错误 error) {
//
// }
//
// // 注意：这里仅做简单翻译，具体方法名和参数名在实际编程中应保持英文，以符合Go语言的编程规范和社区习惯。
// [结束]多个
// ff:插入多个
// c:
// ctx:上下文
// docs:待插入文档
// opts:可选选项
// result:插入结果
// err:错误
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
// [提示]
// func (c *集合) 更新或插入(ctx 上下文.Context, 过滤器 interface{}) (写入结果 WriteResult, 错误 error) {
//
// }
// [结束]
// ff:更新插入
// c:
// ctx:上下文
// filter:更新条件
// replacement:更新内容
// opts:可选选项
// result:更新结果
// err:错误
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

// UpsertId 如果id匹配，则更新一个文档，如果id不匹配，则插入一个新的文档，并将id注入到文档中。
// 注意,id是十六进制, 不是文本型, 需要转换后查询.
// 替换参数必须是一个将用于替换所选文档的文档。它不能为空，并且不能包含任何更新操作符。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:2a704aa664092959
// [提示:] func (c *集合) 更新或插入Id(ctx 上下文, id 接口{}并按ID
// ff:更新插入并按ID
// c:
// ctx:上下文
// id:更新ID
// replacement:更新内容
// opts:可选选项
// result:更新结果
// err:错误
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

// UpdateOne 执行一个更新命令，最多更新集合中的一份文档。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:a16e90f28370dc2c
// [提示]
// func (c *集合) 更新一条数据(ctx 上下文, 过滤器 interface{}) (更新结果 UpdateResult, 错误 error) {
//
// }
//
// // UpdateResult 是更新操作的结果类型
//
//	type UpdateResult struct {
//	    MatchedCount int64  // 匹配文档数
//	    ModifiedCount int64  // 修改文档数
//	    UpsertedID    *primitive.ObjectID // 新增文档的_id，如果进行了upsert操作
//	}
//
// func (c *Collection) InsertOne(ctx context.Context, document interface{}) (插入结果 InsertOneResult, 错误 error) {
//
// }
//
// // InsertOneResult 插入操作的结果类型
//
//	type InsertOneResult struct {
//	    InsertedID *primitive.ObjectID // 插入文档的_id
//	}
//
// func (c *Collection) DeleteOne(ctx context.Context, filter interface{}) (删除结果 DeleteResult, 错误 error) {
//
// }
//
// // DeleteResult 删除操作的结果类型
//
//	type DeleteResult struct {
//	    DeletedCount int64 // 删除的文档数
//	}
//
// func (c *Collection) Find(ctx context.Context, filter interface{}) *Query {
//
// }
//
// // Query 是用于构建查询的类型
//
//	type Query struct {
//	    // 包含了多个查询相关的配置和方法
//	}
//
// [结束]
// ff:更新一条
// c:
// ctx:上下文
// filter:更新条件
// update:更新内容
// opts:可选选项
// err:错误
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
		// UpdateOne支持upsert功能 md5:aaec7189323f1660
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

// UpdateId 执行一个更新命令，最多更新集合中的一个文档。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:67764db9e90007e8
// [提示:] func (c *集合) 更新Id(ctx 上下文, id 任意类型) (结果 Result, 错误 error)
// ff:更新并按ID
// c:
// ctx:上下文
// id:更新ID
// update:更新内容
// opts:可选选项
// err:错误
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

// UpdateAll 执行更新命令以更新集合中的文档。
// 如果没有文档被更新，UpdateResult 中的 matchedCount 将为 0
// 参考资料: https://docs.mongodb.com/manual/reference/operator/update/
// md5:94c36e4a82312809
// [提示:] func (c *集合) 更新所有(ctx 上下文.Context, 过滤器 interface{})
// ff:更新
// c:
// ctx:上下文
// filter:更新条件
// update:更新内容
// opts:可选选项
// result:更新结果
// err:错误
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

// ReplaceOne 执行更新命令，最多更新集合中的一个文档。如果 opts 中的 UpdateHook 被设置，那么 Hook 将在其上执行，否则 Hook 尝试将 doc 作为 Hook。预期 doc 的类型是用户定义的文档的定义。
// md5:1d830477f8b32e37
// [提示]
// func (c *集合) 替换单个文档(ctx 上下文 контекст, 过滤器 interface{}) (更新结果 UpdateResult, 错误 error) {
//
// }
//
// // UpdateResult 是一个返回结果的结构体，可能包含匹配的文档数和任何操作错误。
//
//	type UpdateResult struct {
//	    MatchedCount int64  // 匹配文档数量
//	    ModifiedCount int64  // 修改文档数量
//	    UpsertedID    interface{} // 新插入文档的ID（如果进行了upsert操作）
//	    Err           error       // 操作错误
//	}
//
// [结束]
// ff:替换一条
// c:
// ctx:上下文
// filter:替换条件
// doc:替换内容
// opts:可选选项
// err:错误
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

// Remove 执行删除命令，从集合中最多删除一个文档。
// 如果 filter 是 bson.M{}，DeleteOne 将删除集合中的一个文档。
// 参考：https://docs.mongodb.com/manual/reference/command/delete/
// md5:3b5cf64ce5f460b0
// [提示:] func (c *集合) 删除(ctx 上下文, 过滤器 interface{})
// ff:删除一条
// c:
// ctx:上下文
// filter:删除条件
// opts:可选选项
// err:错误
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

// RemoveId 执行删除命令，从集合中删除最多一个文档。 md5:6516d8a8963d018c
// [提示:] func (c *集合) 删除ById(ctx 上下文, id interface{}) (删除结果 error)
// ff:删除并按ID
// c:
// ctx:上下文
// id:删除ID
// opts:可选选项
// err:错误
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

// RemoveAll 执行一个删除命令，从集合中删除文档。如果 filter 是 bson.M{}（空映射），则会删除集合中的所有文档。
// 参考：https://docs.mongodb.com/manual/reference/command/delete/
// md5:abc51456adc5fc5a
// [提示:] func (c *集合) 全部移除(ctx 上下文, 过滤器 interface{})
// ff:删除
// c:
// ctx:上下文
// filter:删除条件
// opts:可选选项
// result:删除结果
// err:错误
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

// Aggregate 在集合上执行聚合命令，并返回一个 AggregateI，用于获取结果文档。 md5:e57ffed517c59fbc
// [提示:] func (c *集合) 批处理聚合操作(ctx 上下文, 管道 interface{})
// ff:聚合
// c:
// ctx:上下文
// pipeline:聚合管道
// opts:可选选项
func (c *Collection) Aggregate(ctx context.Context, pipeline interface{}, opts ...opts.AggregateOptions) AggregateI {
	return &Aggregate{
		ctx:        ctx,
		collection: c.collection,
		pipeline:   pipeline,
		options:    opts,
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

// 确保索引（已弃用）
// 建议使用CreateIndexes / CreateOneIndex以获取更多功能）
// EnsureIndexes 在集合中创建唯一和非唯一的索引，与CreateIndexes的组合不同：
// 如果uniques/indexes是`[]string{"name"}`，意味着创建名为"name"的索引
// 如果uniques/indexes是`[]string{"name,-age", "uid"}`，表示创建复合索引：name和-age，然后创建一个索引：uid
// md5:c595ad59f9c60c06
// [提示:] func (c *集合) 确保索引(ctx 上下文, 唯一索引 []string, 普通索引 []string) (错误 error) {}
// ff:EnsureIndexes弃用
// c:
// ctx:
// uniques:
// indexes:
// err:
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

// CreateIndexes 在集合中创建多个索引
// 如果opts.IndexModel中的Key为[]string{"name"}，表示创建索引：name
// 如果opts.IndexModel中的Key为[]string{"name", "-age"}，表示创建复合索引：name和-age
// md5:822a787892c2186f索引s
// [提示:] func (c *Collection) 创建索引(ctx context.Context, 索引模型 []opts.IndexModel) (错误 error) {}索引s
// ff:创建多条索引
// c:
// ctx:上下文
// indexes:索引s
// err:错误
func (c *Collection) CreateIndexes(ctx context.Context, indexes []opts.IndexModel) (err error) {
	err = c.ensureIndex(ctx, indexes)
	return
}

// CreateOneIndex 创建一个索引
// 如果opts.IndexModel中的Key为[]string{"name"}，表示创建名为"name"的索引
// 如果opts.IndexModel中的Key为[]string{"name","-age"}，表示创建复合索引：按照"name"升序和"age"降序
// md5:70c27ea42ff3bbbf
// [提示:] func (c *集合) 创建单个索引(ctx 上下文, index 索引模型) 错误 {}
// ff:创建索引
// c:
// ctx:上下文
// index:索引
func (c *Collection) CreateOneIndex(ctx context.Context, index opts.IndexModel) error {
	return c.ensureIndex(ctx, []opts.IndexModel{index})

}

// DropAllIndexes 会删除集合上除了_id字段索引之外的所有索引
// 如果集合上只有_id字段的索引，该函数调用将报告错误
// md5:e7655b40436f93df全部索引
// [提示:] func (c *集合) 删除所有索引(ctx 上下文环境) (错误 error) {}全部索引
// ff:删除全部索引
// c:
// ctx:上下文
// err:错误
func (c *Collection) DropAllIndexes(ctx context.Context) (err error) {
	_, err = c.collection.Indexes().DropAll(ctx)
	return err
}

// DropIndex 从集合中删除索引，需要删除的索引应与输入的索引列表匹配
// 索引是 []string{"name"} 表示删除名为 "name" 的单个索引
// 索引是 []string{"name", "-age"} 表示删除复合索引：name 和排除年龄 (-age) 的部分索引
// md5:4ad77e88557061c7索引索引s
// [提示:] func (c *集合) 删除索引(ctx 上下文, 索引列表 []string) error {}索引索引s
// ff:删除索引
// c:
// ctx:上下文
// indexes:索引s
func (c *Collection) DropIndex(ctx context.Context, indexes []string) error {
	_, err := c.collection.Indexes().DropOne(ctx, generateDroppedIndex(indexes))
	if err != nil {
		return err
	}
	return err
}

// 生成存储在Mongo中的索引，可能包含多个索引（如[]string{"index1","index2"}存储为"index1_1_index2_1"） md5:15332a053c924233
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

// DropIndexDropIndex 会删除索引
// 即使索引不存在，这个操作也是安全的
// md5:e7b65cd93b1de7f7集合
// [提示:] func (c *集合) 删除集合(ctx 上下文.Context) 错误 {}集合
// ff:删除集合
// c:
// ctx:上下文
func (c *Collection) DropCollection(ctx context.Context) error {
	return c.collection.Drop(ctx)
}

// CloneCollection 创建集合的副本 md5:5df787f1c8ebab26
// [提示:] func (c *集合) 克隆集合() (*mongo.集合, 错误) {}
// ff:取副本
// c:
func (c *Collection) CloneCollection() (*mongo.Collection, error) {
	return c.collection.Clone()
}

// GetCollectionName 返回集合的名字 md5:440484db8f2a466d
// [提示:] func (c *集合) 获取集合名称() 字符串 {}
// ff:取集合名
// c:
func (c *Collection) GetCollectionName() string {
	return c.collection.Name()
}

// Watch 返回对应集合上所有更改的流。有关更改流的更多信息，请参阅
// md5:7b59cfd256c148f3
// [提示:] func (c *集合) 监听(ctx 上下文.Context, 管道 interface{})
// ff:取变更流
// c:
// ctx:上下文
// pipeline:管道
// opts:可选选项
func (c *Collection) Watch(ctx context.Context, pipeline interface{}, opts ...*opts.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	changeStreamOption := options.ChangeStream()
	if len(opts) > 0 && opts[0].ChangeStreamOptions != nil {
		changeStreamOption = opts[0].ChangeStreamOptions
	}
	return c.collection.Watch(ctx, pipeline, changeStreamOption)
}

// translateUpdateResult 将Mongo的更新结果转换为qmgo定义的UpdateResult md5:cb683a73f25cfe75
func translateUpdateResult(res *mongo.UpdateResult) (result *UpdateResult) {
	result = &UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}
	return
}
