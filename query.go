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

// X查询 struct definition
type X查询 struct {
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

func (q *X查询) X设置排序规则(规则 *options.Collation) QueryI {
	newQ := q
	newQ.collation = 规则
	return newQ
}

func (q *X查询) X设置不超时(是否不超时 bool) QueryI {
	newQ := q
	newQ.noCursorTimeout = &是否不超时
	return newQ
}

// X设置批量处理数量 sets the value for the X设置批量处理数量 field.
// Means the maximum number of documents to be included in each batch returned by the server.
func (q *X查询) X设置批量处理数量(数量 int64) QueryI {
	newQ := q
	newQ.batchSize = &数量
	return newQ
}

// X排序 is Used to set the sorting rules for the returned results
// Format: "age" or "+age" means to sort the age field in ascending order, "-age" means in descending order
// When multiple sort fields are passed in at the same time, they are arranged in the order in which the fields are passed in.
// For example, {"age", "-name"}, first sort by age in ascending order, then sort by name in descending order
func (q *X查询) X排序(排序字段 ...string) QueryI {
	if len(排序字段) == 0 {
		// A nil bson.D will not correctly serialize, but this case is no-op
		// so an early return will do.
		return q
	}

	var sorts bson.D
	for _, field := range 排序字段 {
		key, n := X分割排序字段(field)
		if key == "" {
			panic("Sort: empty field name")
		}
		sorts = append(sorts, bson.E{Key: key, Value: n})
	}
	newQ := q
	newQ.sort = sorts
	return newQ
}

//  SetArrayFilter use for apply update array
//  For Example :
//  var res = QueryTestItem{}
//  change := Change{
//	Update:    bson.M{"$set": bson.M{"instock.$[elem].qty": 100}},
//	ReturnNew: false,
//  }
//  cli.Find(context.Background(), bson.M{"name": "Lucas"}).
//      X设置切片过滤(&options.ArrayFilters{Filters: []interface{}{bson.M{"elem.warehouse": bson.M{"$in": []string{"C", "F"}}},}}).
//        Apply(change, &res)
func (q *X查询) X设置切片过滤(过滤条件 *options.ArrayFilters) QueryI {
	newQ := q
	newQ.arrayFilters = 过滤条件
	return newQ
}

// X字段 is used to determine which fields are displayed or not displayed in the returned results
// Format: bson.M{"age": 1} means that only the age field is displayed
// bson.M{"age": 0} means to display other fields except age
// When _id is not displayed and is set to 0, it will be returned to display
func (q *X查询) X字段(字段Map interface{}) QueryI {
	newQ := q
	newQ.project = 字段Map
	return newQ
}

// X跳过 skip n records
func (q *X查询) X跳过(跳过数量 int64) QueryI {
	newQ := q
	newQ.skip = &跳过数量
	return newQ
}

// X指定索引字段 sets the value for the X指定索引字段 field.
// This should either be the index name as a string or the index specification
// as a document. The default value is nil, which means that no hint will be sent.
func (q *X查询) X指定索引字段(索引字段 interface{}) QueryI {
	newQ := q
	newQ.hint = 索引字段
	return newQ
}

// X设置最大返回数 limits the maximum number of documents found to n
// The default value is 0, and 0  means no limit, and all matching results are returned
// When the limit value is less than 0, the negative limit is similar to the positive limit, but the cursor is closed after returning a single batch result.
// Reference https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
func (q *X查询) X设置最大返回数(数量 int64) QueryI {
	newQ := q
	newQ.limit = &数量
	return newQ
}

// X取一条 query a record that meets the filter conditions
// If the search fails, an error will be returned
func (q *X查询) X取一条(结果指针 interface{}) error {
	if len(q.opts) > 0 {
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, 操作符.X钩子_查询前); err != nil {
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
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, 操作符.X钩子_查询后); err != nil {
			return err
		}
	}
	return nil
}

// X取全部 query multiple records that meet the filter conditions
// The static type of result must be a slice pointer
func (q *X查询) X取全部(结果指针 interface{}) error {
	if len(q.opts) > 0 {
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, 操作符.X钩子_查询前); err != nil {
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
		if err := middleware.Do(q.ctx, q.opts[0].QueryHook, 操作符.X钩子_查询后); err != nil {
			return err
		}
	}
	return nil
}

// X取数量 count the number of eligible entries
func (q *X查询) X取数量() (数量 int64, 错误 error) {
	opt := options.Count()

	if q.limit != nil {
		opt.SetLimit(*q.limit)
	}
	if q.skip != nil {
		opt.SetSkip(*q.skip)
	}

	return q.collection.CountDocuments(q.ctx, q.filter, opt)
}

// X取预估数量 count the number of the collection by using the metadata
func (q *X查询) X取预估数量() (数量 int64, 错误 error) {
	return q.collection.EstimatedDocumentCount(q.ctx)
}

// X去重 gets the unique value of the specified field in the collection and return it in the form of slice
// result should be passed a pointer to slice
// The function will verify whether the static type of the elements in the result slice is consistent with the data type obtained in mongodb
// reference https://docs.mongodb.com/manual/reference/command/distinct/
func (q *X查询) X去重(字段名 string, 切片指针 interface{}) error {
	resultVal := reflect.ValueOf(切片指针)

	if resultVal.Kind() != reflect.Ptr {
		return X错误_结果参数_必须切片指针
	}

	resultElmVal := resultVal.Elem()
	if resultElmVal.Kind() != reflect.Interface && resultElmVal.Kind() != reflect.Slice {
		return X错误_结果参数_必须切片地址
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
	err = rawValue.Unmarshal(切片指针)
	if err != nil {
		fmt.Printf("rawValue.Unmarshal err: %+v\n", err)
		return X错误_查询结果_类型不一致
	}

	return nil
}

// X取结果集 gets a X取结果集 object, which can be used to traverse the query result set
// After obtaining the CursorI object, you should actively call the Close interface to close the cursor
func (q *X查询) X取结果集() CursorI {
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

// X执行命令 runs the findAndModify command, which allows updating, replacing
// or removing a document matching a query and atomically returning either the old
// version (the default) or the new version of the document (when ReturnNew is true)
//
// The Sort and Select query methods affect the result of X执行命令. In case
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
func (q *X查询) X执行命令(change Change, result interface{}) error {
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
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndDelete/
func (q *X查询) findOneAndDelete(change Change, result interface{}) error {
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
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndReplace/
func (q *X查询) findOneAndReplace(change Change, result interface{}) error {
	opts := options.FindOneAndReplace()
	if q.sort != nil {
		opts.SetSort(q.sort)
	}
	if q.project != nil {
		opts.SetProjection(q.project)
	}
	if change.X是否未找到时插入 {
		opts.SetUpsert(change.X是否未找到时插入)
	}
	if change.X是否返回新文档 {
		opts.SetReturnDocument(options.After)
	}

	err := q.collection.FindOneAndReplace(q.ctx, q.filter, change.Update, opts).Decode(result)
	if change.X是否未找到时插入 && !change.X是否返回新文档 && err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}

// findOneAndUpdate
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndUpdate/
func (q *X查询) findOneAndUpdate(change Change, result interface{}) error {
	opts := options.FindOneAndUpdate()
	if q.sort != nil {
		opts.SetSort(q.sort)
	}
	if q.project != nil {
		opts.SetProjection(q.project)
	}
	if change.X是否未找到时插入 {
		opts.SetUpsert(change.X是否未找到时插入)
	}
	if change.X是否返回新文档 {
		opts.SetReturnDocument(options.After)
	}

	if q.arrayFilters != nil {
		opts.SetArrayFilters(*q.arrayFilters)
	}

	err := q.collection.FindOneAndUpdate(q.ctx, q.filter, change.Update, opts).Decode(result)
	if change.X是否未找到时插入 && !change.X是否返回新文档 && err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}

// zj:
func (q *X查询) X分页(页码 int, 页大小 int) QueryI {
	return q.X跳过(int64((页大小 * (页码 - 1)))).X设置最大返回数(int64(页大小))
}

func (q *X查询) X取分页数(perPage int) int {
	// 获取预估文档总数
	docCount, _ := q.X取预估数量()
	// 计算总分页数
	totalPages := int(docCount) / perPage
	if int(docCount)%perPage != 0 {
		totalPages++
	}
	return totalPages
}

//zj:
