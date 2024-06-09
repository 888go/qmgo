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

// 定义查询结构体 md5:56541bbc29d4ce15
// [提示]
//type 查询 struct {
//     筛选条件 interface{}
//     排序方式 interface{}
//     选择字段 interface{}
//     索引指引 interface{}
//     数组筛选器 *options.数组筛选选项
//     限制数量 *int64
//     跳过数量 *int64
//     批次大小 *int64
//     不超时 *bool
//     文档集合 *mongo.Collection
//     查询选项 []qOpts.查询选项
//     编码注册器 *bsoncodec.Registry
//     上下文 context.Context
// }
// [结束]
type Query struct {//hm:查询  cz:type Query  
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

// [提示:] func (q *Query) 文档排序规则设置(collation *options.Collation) 查询接口I {}
// ff:设置排序规则
// q:
// collation:规则
func (q *Query) Collation(collation *options.Collation) QueryI {
	newQ := q
	newQ.collation = collation
	return newQ
}

// [提示:] func (q *Query) 不使用超时游标(n bool) QueryI {}
// ff:设置不超时
// q:
// n:是否不超时
func (q *Query) NoCursorTimeout(n bool) QueryI {
	newQ := q
	newQ.noCursorTimeout = &n
	return newQ
}

// BatchSize 设置 BatchSize 字段的值。
// 它表示服务器返回的每批文档的最大数量。
// md5:66277d16095ac151
// [提示:] func (q *Query) 设置批次大小(n int64) QueryI {}
// ff:设置批量处理数量
// q:
// n:数量
func (q *Query) BatchSize(n int64) QueryI {
	newQ := q
	newQ.batchSize = &n
	return newQ
}

// Sort 用于设置返回结果的排序规则
// 格式："age" 或 "+age" 表示按年龄字段升序排序，"-age" 表示降序排序
// 同时传入多个排序字段时，按照字段传递的顺序进行排列。
// 例如，{"age", "-name"}，首先按年龄升序排序，然后按姓名降序排序。
// md5:0b1b9f5345904541
// [提示:] func (q *Query) 排序(字段 ...string) QueryI {}
// ff:排序
// q:
// fields:排序字段
func (q *Query) Sort(fields ...string) QueryI {
	if len(fields) == 0 {
// 一个空的bson.D不会正确地序列化，但这种情况下可以提前返回。
// md5:c94b59dcb408353d
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

// SetArrayFilter 用于应用更新数组的过滤器
// 示例：
// var res = QueryTestItem{}
//
//	change := Change{
//	    Update:    bson.M{"$set": bson.M{"instock.$[elem].qty": 100}},
//	    ReturnNew: false,
//	}
//
// cli.Find(context.Background(), bson.M{"name": "Lucas"}).
//
//	SetArrayFilters(&options.ArrayFilters{Filters: []interface{}{bson.M{"elem.warehouse": bson.M{"$in": []string{"C", "F"}}},}}).
//	  Apply(change, &res)
//
// 这段代码的注释说明了`SetArrayFilter`方法是用于设置更新操作中的数组过滤器。它给出了一个例子，展示了如何使用该方法来更新名为"Lucas"的文档中，符合条件（"elem.warehouse"在"C"或"F"中）的`instock`数组元素的`qty`字段为100。`Apply`方法最后将变更应用到查询结果上。
// md5:3fa80906c918e6a3
// [提示:] func (q *Query) 设置数组过滤器(filter *选项.数组过滤器) 查询接口 {}
// ff:设置切片过滤
// q:
// filter:过滤条件
func (q *Query) SetArrayFilters(filter *options.ArrayFilters) QueryI {
	newQ := q
	newQ.arrayFilters = filter
	return newQ
}

// Select用于确定在返回结果中显示或不显示哪些字段
// 格式：bson.M{"age": 1} 表示只显示age字段
// bson.M{"age": 0} 表示除了age以外的其他字段都显示
// 如果不显示_id并且设置为0，它将被返回并显示
// md5:3beb3c9bd51ad3fe
// [提示]
//func (q *Query) 选择(projection interface{}) 
// 
// // 下面的方法未给出完整信息，我将提供一般性的翻译，但可能需要根据实际代码上下文进行调整
// 
// func (c *Collection) InsertOne(doc interface{}) (insertedIDprimitive, error) {
//     // 翻译：在集合中插入一个文档，返回插入的ID和可能的错误
// }
// 
// func (c *Collection) Find(filter interface{}, opts ...*FindOptions) *Query {
//     // 翻译：根据过滤器找到文档，返回一个查询对象，可选的查找选项
// }
// 
// func (q *Query) Limit(n int) *Query {
//     // 翻译：限制查询结果的数量，返回更新后的查询对象
// }
// 
// func (q *Query) Skip(n int) *Query {
//     // 翻译：跳过查询结果的前n个文档，返回更新后的查询对象
// }
// 
// func (q *Query) Sort(fields ...string) *Query {
//     // 翻译：按指定字段排序查询结果，返回更新后的查询对象
// }
// 
// func (q *Query) All(result interface{}) error {
//     // 翻译：将所有查询结果填充到结果接口中，返回可能出现的错误
// }
// 
// func (q *Query) One(result interface{}) error {
//     // 翻译：获取查询结果中的第一条文档并填充到结果接口中，返回可能出现的错误
// }
// 
// func (q *Query) Count() (int64, error) {
//     // 翻译：计算查询结果的文档数量，返回总数和可能的错误
// }
// [结束]
// ff:字段
// q:
// projection:字段Map
func (q *Query) Select(projection interface{}) QueryI {
	newQ := q
	newQ.project = projection
	return newQ
}

// Skip skip n records
// [提示:] func (q *Query) 跳过(n int64) QueryI {}
// ff:跳过
// q:
// n:跳过数量
func (q *Query) Skip(n int64) QueryI {
	newQ := q
	newQ.skip = &n
	return newQ
}

// Hint 设置Hint字段的值。这应该是字符串形式的索引名称，或者是文档形式的索引规范。默认值为nil，表示不发送提示。
// md5:3d3535508606dd43
// [提示:] func (q *Query) 指定索引(hint interface{})
// ff:指定索引字段
// q:
// hint:索引字段
func (q *Query) Hint(hint interface{}) QueryI {
	newQ := q
	newQ.hint = hint
	return newQ
}

// Limit 将找到的最大文档数量限制为 n
// 默认值为 0，0 表示无限制，返回所有匹配的结果
// 当限制值小于 0 时，负限制类似于正限制，但返回单个批次结果后关闭游标。
// 参考 https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
// md5:9081095bd35be08f
// [提示:] func (q *查询) 限制(n int64) 查询接口 {}
// ff:设置最大返回数
// q:
// n:数量
func (q *Query) Limit(n int64) QueryI {
	newQ := q
	newQ.limit = &n
	return newQ
}

// 对符合过滤条件的记录执行一次查询
// 如果搜索失败，将返回一个错误
// md5:68571c814c5cd088
// [提示:] func (q *Query) 一个(result interface{})
// ff:取一条
// q:
// result:结果指针
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

// 用于查询满足过滤条件的所有记录
// 结果的静态类型必须是切片指针
// md5:5f57d8aff8afe252
// [提示:] func (q *Query) 全部获取(result interface{})
// ff:取全部
// q:
// result:结果指针
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

// Count 计算符合条件的条目数量 md5:7bed3eaaee1ce368
// [提示:] func (q *Query) 计数() (总数 int64, 错误 error) {}
// ff:取数量
// q:
// n:数量
// err:错误
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

// EstimatedCount 通过元数据计算集合的数量,
// EstimatedDocumentCount() 方法比 CountDocuments() 方法更快，因为它使用集合的元数据而不是扫描整个集合。
// EstimatedCount 通过元数据计算集合的数量,
// EstimatedDocumentCount() 方法比 CountDocuments() 方法更快，因为它使用集合的元数据而不是扫描整个集合。
// md5:8c9bd7e463139421
// [提示:] func (q *Query) 估算计数() (总数 int64, 错误 error) {}
// ff:取预估数量
// q:
// n:数量
// err:错误
func (q *Query) EstimatedCount() (n int64, err error) {
	return q.collection.EstimatedDocumentCount(q.ctx)
}

// Distinct 从集合中获取指定字段的唯一值，并以切片形式返回。
// result 应该是一个指向切片的指针。
// 函数会检查result切片元素的静态类型是否与MongoDB中获取的数据类型一致。
// 参考：https://docs.mongodb.com/manual/reference/command/distinct/
// md5:b83f3aa5718b2dfd
// [提示:] func (q *Query) 唯一值(key string, result interface{}
// ff:去重
// q:
// key:字段名
// result:切片指针
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

// Cursor 获取一个 Cursor 对象，可用于遍历查询结果集
// 在获取到 CursorI 对象后，应主动调用 Close 接口来关闭游标
// md5:b1e9fc62a5f777fe
// [提示:] func (q *Query) 获取游标() 游标接口 {}
// ff:取结果集
// q:
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

// Apply 执行 findAndModify 命令，该命令允许根据查询更新、替换或删除文档，并原子性地返回旧版本（默认）或新版本的文档（当 ReturnNew 为 true 时）。
// 
// Sort 和 Select 查询方法会影响 Apply 的结果。如果有多个文档匹配查询，Sort 可以通过排序来决定操作哪个文档。Select 则可以仅获取新或旧文档中选定字段的内容。
// 
// 当 Change.Replace 为 true 时，意味着在集合中最多替换一个文档，且 update 参数必须是一个文档，不能包含任何更新运算符；如果没有找到对象并且 Change.Upsert 为 false，则会返回 ErrNoDocuments 错误。当 Change.Remove 为 true 时，意味着最多删除集合中的一个文档，并返回删除前的文档状态；如果没有找到对象，同样返回 ErrNoDocuments。
// 
// 如果 Change.Replace 和 Change.Remove 都为 false，则表示最多更新集合中的一个文档，update 参数必须是一个包含更新运算符的文档；如果没有找到对象并且 Change.Upsert 为 false，则返回 ErrNoDocuments。
// 
// 参考：https://docs.mongodb.com/manual/reference/command/findAndModify/
// md5:e14e8d7ceac827ac
// [提示:] func (q *Query) 应用变更(change 变更, result 结果接口{}
// ff:执行命令
// q:
// change:
// result:
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
// 参考：https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndDelete/ 
// 
// 这段 Go 代码的注释表示这是一个名为 "findOneAndDelete" 的函数或方法，它用于从 MongoDB 集合中查找并删除第一个匹配的文档。参考链接指向 MongoDB 官方文档，提供了关于该方法的详细说明。
// md5:23b36fd4f1711d7b
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
// 参考：https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndReplace/
// md5:cd4b97a5409057c1
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
// 参考：https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndUpdate/ 
// 
// 这段Go代码中的注释表示这是一个名为`findOneAndUpdate`的方法，它引用了MongoDB文档中关于`db.collection.findOneAndUpdate`方法的手册参考。这个方法在MongoDB数据库操作中用于根据指定的条件查找并更新第一个匹配的文档。
// md5:fe84856a45a0b0f1
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

// zj:
func (q *Query) X分页(页码 int, 页大小 int) QueryI {
	return q.Skip(int64((页大小 * (页码 - 1)))).Limit(int64(页大小))
}

func (q *Query) X取分页数(perPage int) int {
	// 获取预估文档总数
	docCount, _ := q.EstimatedCount()
	// 计算总分页数
	totalPages := int(docCount) / perPage
	if int(docCount)%perPage != 0 {
		totalPages++
	}
	return totalPages
}

//zj: