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
	
	"github.com/888go/qmgo/middleware"
	"github.com/888go/qmgo/operator"
	qOpts "github.com/888go/qmgo/options"
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

func (q *Query) X设置排序规则(规则 *options.Collation) QueryI {
	newQ := q
	newQ.collation = 规则
	return newQ
}

func (q *Query) X设置不超时(是否不超时 bool) QueryI {
	newQ := q
	newQ.noCursorTimeout = &是否不超时
	return newQ
}

// BatchSize 设置 BatchSize 字段的值。
// 表示服务器返回的每个批次中包含的最大文档数量。
func (q *Query) X设置批量处理数量(数量 int64) QueryI {
	newQ := q
	newQ.batchSize = &数量
	return newQ
}

// Sort 用于设置返回结果的排序规则
// 格式： "age" 或 "+age" 表示按年龄字段升序排序，"-age" 表示按年龄字段降序排序
// 当同时传入多个排序字段时，按照字段传入的顺序依次排列
// 例如：{"age", "-name"}，首先按年龄升序排序，然后按姓名降序排序
func (q *Query) X排序(排序字段 ...string) QueryI {
	if len(排序字段) == 0 {
// 若 bson.D 为 nil，则无法正确序列化，但由于此处为空操作（no-op），所以提前返回即可。
		return q
	}

	var sorts bson.D
	for _, field := range 排序字段 {
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
// change := Change{
//	Update:    bson.M{"$set": bson.M{"instock.$[elem].qty": 100}}, // 更新数组中符合条件的元素数量为100
//	ReturnNew: false, // 是否返回更新后的文档，默认为false
// }
// 使用cli在上下文中查找指定条件的文档（name为"Lucas"）
// cli.Find(context.Background(), bson.M{"name": "Lucas"}).
// 设置数组过滤器，这里匹配"instock"数组中"warehouse"字段包含"C"或"F"的元素
// .SetArrayFilters(&options.ArrayFilters{Filters: []interface{}{bson.M{"elem.warehouse": bson.M{"$in": []string{"C", "F"}}},}}).
// 应用上述变更到查询结果，并将更新后的内容存入res变量
// .Apply(change, &res)
func (q *Query) X设置数组过滤(过滤条件 *options.ArrayFilters) QueryI {
	newQ := q
	newQ.arrayFilters = 过滤条件
	return newQ
}

// Select 用于确定在返回结果中哪些字段显示或不显示
// 格式：bson.M{"age": 1} 表示只显示 age 字段
// bson.M{"age": 0} 表示除 age 字段外的其他字段均显示
// 当 _id 不显示并设置为 0 时，它将被默认返回显示
func (q *Query) X字段(字段Map interface{}) QueryI {
	newQ := q
	newQ.project = 字段Map
	return newQ
}

// Skip skip n records
func (q *Query) X跳过(跳过数量 int64) QueryI {
	newQ := q
	newQ.skip = &跳过数量
	return newQ
}

// Hint 设置Hint字段的值。
// 这个值应该要么是作为字符串的索引名，要么是作为文档的索引规范。
// 默认值为nil，这意味着不会发送任何提示。
func (q *Query) X指定索引字段(索引字段 interface{}) QueryI {
	newQ := q
	newQ.hint = 索引字段
	return newQ
}

// Limit 限制查询结果返回的最大文档数量为 n
// 默认值为 0，当设置为 0 时，表示没有限制，会返回所有匹配的结果
// 当 limit 值小于 0 时，负数的限制与正数类似，但会在返回单批次结果后关闭游标
// 参考文献：https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
func (q *Query) X设置最大返回数(数量 int64) QueryI {
	newQ := q
	newQ.limit = &数量
	return newQ
}

// 根据过滤条件查询一条记录
// 若查询未找到匹配项，则返回错误
func (q *Query) X取一条(结果指针 interface{}) error {
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

	err := q.collection.FindOne(q.ctx, q.filter, opt).Decode(结果指针)

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

// 根据过滤条件查询满足条件的多条记录
// 结果的静态类型必须是指向切片的指针
func (q *Query) X取全部(结果指针 interface{}) error {
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
	err = c.X取全部(结果指针)
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

// Count 计算符合条件的条目数量
func (q *Query) X取数量() (数量 int64, 错误 error) {
	opt := options.Count()

	if q.limit != nil {
		opt.SetLimit(*q.limit)
	}
	if q.skip != nil {
		opt.SetSkip(*q.skip)
	}

	return q.collection.CountDocuments(q.ctx, q.filter, opt)
}

// EstimatedCount 通过使用元数据估算集合的数量
func (q *Query) X取预估数量() (数量 int64, 错误 error) {
	return q.collection.EstimatedDocumentCount(q.ctx)
}

// Distinct 获取集合中指定字段的唯一值，并以切片形式返回
// 结果应通过指针传递给切片
// 该函数将验证结果切片中元素的静态类型与在mongodb中获取的数据类型是否一致
// 参考文献：https://docs.mongodb.com/manual/reference/command/distinct/
func (q *Query) X去重(字段名 string, 数组指针 interface{}) error {
	resultVal := reflect.ValueOf(数组指针)

	if resultVal.Kind() != reflect.Ptr {
		return ErrQueryNotSlicePointer
	}

	resultElmVal := resultVal.Elem()
	if resultElmVal.Kind() != reflect.Interface && resultElmVal.Kind() != reflect.Slice {
		return ErrQueryNotSliceType
	}

	opt := options.Distinct()
	res, err := q.collection.Distinct(q.ctx, 字段名, q.filter, opt)
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
	err = rawValue.Unmarshal(数组指针)
	if err != nil {
		fmt.Printf("rawValue.Unmarshal err: %+v\n", err)
		return ErrQueryResultTypeInconsistent
	}

	return nil
}

// 获取一个Cursor对象，可用于遍历查询结果集
// 在获取到CursorI对象后，应主动调用Close接口关闭游标
func (q *Query) X取结果集() CursorI {
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

// Apply runs the findAndModify command, which allows updating, replacing
// or removing a document matching a query and atomically returning either the old
// version (the default) or the new version of the document (when ReturnNew is true)
//
// The Sort and Select query methods affect the result of Apply. In case
// multiple documents match the query, Sort enables selecting which document to
// act upon by ordering it first. Select enables retrieving only a selection
// of fields of the new or old document.
//
// When Change.Replace is true, it means replace at most one document in the collection
// and the update parameter must be a document and cannot contain any update operators;
// if no objects are found and Change.Upsert is false, it will returns ErrNoDocuments.
// When Change.Remove is true, it means delete at most one document in the collection
// and returns the document as it appeared before deletion; if no objects are found,
// it will returns ErrNoDocuments.
// When both Change.Replace and Change.Remove are false，it means update at most one document
// in the collection and the update parameter must be a document containing update operators;
// if no objects are found and Change.Upsert is false, it will returns ErrNoDocuments.
//
// reference: https://docs.mongodb.com/manual/reference/command/findAndModify/
func (q *Query) X执行命令(change Change, result interface{}) error {
	var err error

	if change.X是否删除 {
		err = q.findOneAndDelete(change, result)
	} else if change.X是否替换 {
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
//findOneAndReplace 函数用于对 MongoDB 集合执行“查找并替换”操作，
// 它会根据给定的查询条件找到第一条匹配的文档，然后使用新的文档数据进行替换。
func (q *Query) findOneAndReplace(change Change, result interface{}) error {
	opts := options.FindOneAndReplace()
	if q.sort != nil {
		opts.SetSort(q.sort)
	}
	if q.project != nil {
		opts.SetProjection(q.project)
	}
	if change.X未找到是否插入 {
		opts.SetUpsert(change.X未找到是否插入)
	}
	if change.X是否返回新文档 {
		opts.SetReturnDocument(options.After)
	}

	err := q.collection.FindOneAndReplace(q.ctx, q.filter, change.X更新替换, opts).Decode(result)
	if change.X未找到是否插入 && !change.X是否返回新文档 && err == mongo.ErrNoDocuments {
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
	if change.X未找到是否插入 {
		opts.SetUpsert(change.X未找到是否插入)
	}
	if change.X是否返回新文档 {
		opts.SetReturnDocument(options.After)
	}

	if q.arrayFilters != nil {
		opts.SetArrayFilters(*q.arrayFilters)
	}

	err := q.collection.FindOneAndUpdate(q.ctx, q.filter, change.X更新替换, opts).Decode(result)
	if change.X未找到是否插入 && !change.X是否返回新文档 && err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}
