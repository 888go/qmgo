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

	"github.com/qiniu/qmgo/middleware"
	"github.com/qiniu/qmgo/operator"
	qOpts "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Query 结构体定义
type Query struct {
	filter          interface{}
	sort            interface{}
	project         interface{}
	hint            interface{}
	arrayFilters    *options.ArrayFilters
	limit           *int64
	skip            *int64
	batchSize       *int64
	noCursorTimeout *bool
	collation       *options.Collation

	ctx        context.Context
	collection *mongo.Collection
	opts       []qOpts.FindOptions
	registry   *bsoncodec.Registry
}

// 设置排序规则
// collation:规则
func (q *Query) Collation(collation *options.Collation) QueryI {
	newQ := q
	newQ.collation = collation
	return newQ
}

// 设置不超时
// n:是否不超时
func (q *Query) NoCursorTimeout(n bool) QueryI {
	newQ := q
	newQ.noCursorTimeout = &n
	return newQ
}

// 设置批量处理数量
// n:数量
// BatchSize 设置 BatchSize 字段的值。
// 表示服务器返回的每个批次中包含的最大文档数量。
func (q *Query) BatchSize(n int64) QueryI {
	newQ := q
	newQ.batchSize = &n
	return newQ
}

// 排序
// fields:排序字段
// Sort 用于设置返回结果的排序规则
// 格式： "age" 或 "+age" 表示按年龄字段升序排序，"-age" 表示按年龄字段降序排序
// 当同时传入多个排序字段时，按照字段传入的顺序依次排列
// 例如：{"age", "-name"}，首先按年龄升序排序，然后按姓名降序排序
func (q *Query) Sort(fields ...string) QueryI {
	if len(fields) == 0 {
		// 若 bson.D 为 nil，则无法正确序列化，但由于此处为空操作（no-op），所以提前返回即可。
		return q
	}

	var sorts bson.D
	for _, field := range fields {
		key, n := SplitSortField(field)
		if key == "" {
			panic("Sort: empty field name")
		}
		sorts = append(sorts, bson.E{Key: key, Value: n})
	}
	newQ := q
	newQ.sort = sorts
	return newQ
}

// SetArrayFilter 用于应用更新数组的操作
// 例如：
// 声明一个结果变量
// var res = QueryTestItem{}
// 定义变更内容
//
//	change := Change{
//		Update:    bson.M{"$set": bson.M{"instock.$[elem].qty": 100}}, // 更新数组中符合条件的元素数量为100
//		ReturnNew: false, // 是否返回更新后的文档，默认为false
//	}
//
// 使用cli在上下文中查找指定条件的文档（name为"Lucas"）
// cli.Find(context.Background(), bson.M{"name": "Lucas"}).
// 设置数组过滤器，这里匹配"instock"数组中"warehouse"字段包含"C"或"F"的元素
// .SetArrayFilters(&options.ArrayFilters{Filters: []interface{}{bson.M{"elem.warehouse": bson.M{"$in": []string{"C", "F"}}},}}).
// 应用上述变更到查询结果，并将更新后的内容存入res变量
// .Apply(change, &res)
func (q *Query) SetArrayFilters(filter *options.ArrayFilters) QueryI {
	newQ := q
	newQ.arrayFilters = filter
	return newQ
}

// 字段
// projection:字段Map
// Select 用于确定在返回结果中哪些字段显示或不显示
// 格式：bson.M{"age": 1} 表示只显示 age 字段
// bson.M{"age": 0} 表示除 age 字段外的其他字段均显示
// 当 _id 不显示并设置为 0 时，它将被默认返回显示
func (q *Query) Select(projection interface{}) QueryI {
	newQ := q
	newQ.project = projection
	return newQ
}

// 跳过
// n:跳过数量
// Skip跳过n条记录
func (q *Query) Skip(n int64) QueryI {
	newQ := q
	newQ.skip = &n
	return newQ
}

// 指定索引字段
// hint:索引字段
// Hint 设置Hint字段的值。
// 这个值应该要么是作为字符串的索引名，要么是作为文档的索引规范。
// 默认值为nil，这意味着不会发送任何提示。
func (q *Query) Hint(hint interface{}) QueryI {
	newQ := q
	newQ.hint = hint
	return newQ
}

// 设置最大返回数
// n:数量
// Limit 限制查询结果返回的最大文档数量为 n
// 默认值为 0，当设置为 0 时，表示没有限制，会返回所有匹配的结果
// 当 limit 值小于 0 时，负数的限制与正数类似，但会在返回单批次结果后关闭游标
// 参考文献：https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
func (q *Query) Limit(n int64) QueryI {
	newQ := q
	newQ.limit = &n
	return newQ
}

// 查询一条
// result:结果指针
// 根据过滤条件查询一条记录
// 若查询未找到匹配项，则返回错误
func (q *Query) One(result interface{}) error {
	if len(q.opts) > 0 {
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, operator.BeforeQuery); err != nil {
			return err
		}
	}
	opt := options.FindOne()

	if q.collation != nil {
		opt.SetCollation(q.collation)
	}
	if q.sort != nil {
		opt.SetSort(q.sort)
	}
	if q.project != nil {
		opt.SetProjection(q.project)
	}
	if q.skip != nil {
		opt.SetSkip(*q.skip)
	}
	if q.hint != nil {
		opt.SetHint(q.hint)
	}

	err := q.collection.FindOne(q.ctx, q.filter, opt).Decode(result)

	if err != nil {
		return err
	}
	if len(q.opts) > 0 {
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, operator.AfterQuery); err != nil {
			return err
		}
	}
	return nil
}

// 查询全部
// result:结果指针
// 根据过滤条件查询满足条件的多条记录
// 结果的静态类型必须是指向切片的指针
func (q *Query) All(result interface{}) error {
	if len(q.opts) > 0 {
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, operator.BeforeQuery); err != nil {
			return err
		}
	}
	opt := options.Find()
	if q.collation != nil {
		opt.SetCollation(q.collation)
	}
	if q.sort != nil {
		opt.SetSort(q.sort)
	}
	if q.project != nil {
		opt.SetProjection(q.project)
	}
	if q.limit != nil {
		opt.SetLimit(*q.limit)
	}
	if q.skip != nil {
		opt.SetSkip(*q.skip)
	}
	if q.hint != nil {
		opt.SetHint(q.hint)
	}
	if q.batchSize != nil {
		opt.SetBatchSize(int32(*q.batchSize))
	}
	if q.noCursorTimeout != nil {
		opt.SetNoCursorTimeout(*q.noCursorTimeout)
	}

	var err error
	var cursor *mongo.Cursor

	cursor, err = q.collection.Find(q.ctx, q.filter, opt)

	c := Cursor{
		ctx:    q.ctx,
		cursor: cursor,
		err:    err,
	}
	err = c.All(result)
	if err != nil {
		return err
	}
	if len(q.opts) > 0 {
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, operator.AfterQuery); err != nil {
			return err
		}
	}
	return nil
}

// 取数量
// n:数量
// err:错误
// Count 计算符合条件的条目数量
func (q *Query) Count() (n int64, err error) {
	opt := options.Count()

	if q.limit != nil {
		opt.SetLimit(*q.limit)
	}
	if q.skip != nil {
		opt.SetSkip(*q.skip)
	}

	return q.collection.CountDocuments(q.ctx, q.filter, opt)
}

// 取预估数量
// n:数量
// err:错误
// EstimatedCount 通过使用元数据估算集合的数量
func (q *Query) EstimatedCount() (n int64, err error) {
	return q.collection.EstimatedDocumentCount(q.ctx)
}

// 去重
// key:字段名
// result:数组指针
// Distinct 获取集合中指定字段的唯一值，并以切片形式返回
// 结果应通过指针传递给切片
// 该函数将验证结果切片中元素的静态类型与在mongodb中获取的数据类型是否一致
// 参考文献：https://docs.mongodb.com/manual/reference/command/distinct/
func (q *Query) Distinct(key string, result interface{}) error {
	resultVal := reflect.ValueOf(result)

	if resultVal.Kind() != reflect.Ptr {
		return ErrQueryNotSlicePointer
	}

	resultElmVal := resultVal.Elem()
	if resultElmVal.Kind() != reflect.Interface && resultElmVal.Kind() != reflect.Slice {
		return ErrQueryNotSliceType
	}

	opt := options.Distinct()
	res, err := q.collection.Distinct(q.ctx, key, q.filter, opt)
	if err != nil {
		return err
	}
	registry := q.registry
	if registry == nil {
		registry = bson.DefaultRegistry
	}
	valueType, valueBytes, err_ := bson.MarshalValueWithRegistry(registry, res)
	if err_ != nil {
		fmt.Printf("bson.MarshalValue err: %+v\n", err_)
		return err_
	}

	rawValue := bson.RawValue{Type: valueType, Value: valueBytes}
	err = rawValue.Unmarshal(result)
	if err != nil {
		fmt.Printf("rawValue.Unmarshal err: %+v\n", err)
		return ErrQueryResultTypeInconsistent
	}

	return nil
}

// 取结果集
// 获取一个Cursor对象，可用于遍历查询结果集
// 在获取到CursorI对象后，应主动调用Close接口关闭游标
func (q *Query) Cursor() CursorI {
	opt := options.Find()

	if q.sort != nil {
		opt.SetSort(q.sort)
	}
	if q.project != nil {
		opt.SetProjection(q.project)
	}
	if q.limit != nil {
		opt.SetLimit(*q.limit)
	}
	if q.skip != nil {
		opt.SetSkip(*q.skip)
	}

	if q.batchSize != nil {
		opt.SetBatchSize(int32(*q.batchSize))
	}
	if q.noCursorTimeout != nil {
		opt.SetNoCursorTimeout(*q.noCursorTimeout)
	}

	var err error
	var cur *mongo.Cursor
	cur, err = q.collection.Find(q.ctx, q.filter, opt)
	return &Cursor{
		ctx:    q.ctx,
		cursor: cur,
		err:    err,
	}
}

// 执行命令并按类型
// Apply运行findmodify命令，该命令允许更新、替换或删除与查询匹配的文档，并原子性地返回文档的旧版本(默认)或新版本(当ReturnNew为true时)
//
// Sort和Select查询方法影响Apply的结果。在多个文档匹配查询的情况下，Sort可以通过先排序来选择要操作的文档。Select启用只检索新文档或旧文档的一部分字段。
//
// 当变化时。Replace为true，这意味着在集合中最多替换一个文档，update参数必须是一个文档，并且不能包含任何更新操作符;
// 如果没有找到对象并进行更改。如果Upsert为false，它将返回ErrNoDocuments。
// 当改变时。Remove为true，这意味着在集合中最多删除一个文档，并返回删除之前出现的文档;如果没有找到对象，
// 返回ErrNoDocuments。
// 当两者都改变时。替换和改变。Remove为false，这意味着更新集合中最多一个文档，更新参数必须是包含更新操作符的文档;
// 如果没有找到对象并进行更改。如果Upsert为false，它将返回ErrNoDocuments。
//
// reference: https://docs.mongodb.com/manual/reference/command/findAndModify/
func (q *Query) Apply(change Change, result interface{}) error {
	var err error

	if change.Remove {
		err = q.findOneAndDelete(change, result)
	} else if change.Replace {
		err = q.findOneAndReplace(change, result)
	} else {
		err = q.findOneAndUpdate(change, result)
	}

	return err
}

// findOneAndDelete
// 参考文献: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndDelete/
// 此函数用于在MongoDB中查找并删除一条文档（记录）
// 它首先会根据提供的查询条件找到集合中第一条匹配的文档
// 找到后，立即从集合中删除该文档，并返回被删除的文档内容
// 注意：此操作为原子操作，在多线程或分布式环境下能保证数据一致性
func (q *Query) findOneAndDelete(change Change, result interface{}) error {
	opts := options.FindOneAndDelete()
	if q.sort != nil {
		opts.SetSort(q.sort)
	}
	if q.project != nil {
		opts.SetProjection(q.project)
	}

	return q.collection.FindOneAndDelete(q.ctx, q.filter, opts).Decode(result)
}

// findOneAndReplace
// 参考文献: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndReplace/
// 此函数实现的功能是，在MongoDB数据库中查找并替换一条文档数据。
// 根据提供的查询条件在指定集合中查找匹配的第一条文档，并用新文档替换它。
// findOneAndReplace 函数用于对 MongoDB 集合执行“查找并替换”操作，
// 它会根据给定的查询条件找到第一条匹配的文档，然后使用新的文档数据进行替换。
func (q *Query) findOneAndReplace(change Change, result interface{}) error {
	opts := options.FindOneAndReplace()
	if q.sort != nil {
		opts.SetSort(q.sort)
	}
	if q.project != nil {
		opts.SetProjection(q.project)
	}
	if change.Upsert {
		opts.SetUpsert(change.Upsert)
	}
	if change.ReturnNew {
		opts.SetReturnDocument(options.After)
	}

	err := q.collection.FindOneAndReplace(q.ctx, q.filter, change.Update, opts).Decode(result)
	if change.Upsert && !change.ReturnNew && err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}

// findOneAndUpdate
// 参考文献: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndUpdate/
// 此函数用于在 MongoDB 集合中查找匹配的第一个文档并更新它。
// 它首先会按照给定的查询条件查找文档，如果找到则根据提供的更新操作符进行更新，
// 然后返回更新前的原始文档（默认行为）或更新后的文档（根据方法选项设置）。
// 这是 MongoDB 的一个核心 CRUD 操作，常用于原子性地更新数据。
func (q *Query) findOneAndUpdate(change Change, result interface{}) error {
	opts := options.FindOneAndUpdate()
	if q.sort != nil {
		opts.SetSort(q.sort)
	}
	if q.project != nil {
		opts.SetProjection(q.project)
	}
	if change.Upsert {
		opts.SetUpsert(change.Upsert)
	}
	if change.ReturnNew {
		opts.SetReturnDocument(options.After)
	}

	if q.arrayFilters != nil {
		opts.SetArrayFilters(*q.arrayFilters)
	}

	err := q.collection.FindOneAndUpdate(q.ctx, q.filter, change.Update, opts).Decode(result)
	if change.Upsert && !change.ReturnNew && err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}
