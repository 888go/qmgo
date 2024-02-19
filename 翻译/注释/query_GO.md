
<原文开始>
// Query struct definition
<原文结束>

# <翻译开始>
// Query 结构体定义
# <翻译结束>


<原文开始>
// BatchSize sets the value for the BatchSize field.
// Means the maximum number of documents to be included in each batch returned by the server.
<原文结束>

# <翻译开始>
// BatchSize 设置 BatchSize 字段的值。
// 表示服务器返回的每个批次中包含的最大文档数量。
# <翻译结束>


<原文开始>
// Sort is Used to set the sorting rules for the returned results
// Format: "age" or "+age" means to sort the age field in ascending order, "-age" means in descending order
// When multiple sort fields are passed in at the same time, they are arranged in the order in which the fields are passed in.
// For example, {"age", "-name"}, first sort by age in ascending order, then sort by name in descending order
<原文结束>

# <翻译开始>
// Sort 用于设置返回结果的排序规则
// 格式： "age" 或 "+age" 表示按年龄字段升序排序，"-age" 表示按年龄字段降序排序
// 当同时传入多个排序字段时，按照字段传入的顺序依次排列
// 例如：{"age", "-name"}，首先按年龄升序排序，然后按姓名降序排序
# <翻译结束>


<原文开始>
		// A nil bson.D will not correctly serialize, but this case is no-op
		// so an early return will do.
<原文结束>

# <翻译开始>
// 若 bson.D 为 nil，则无法正确序列化，但由于此处为空操作（no-op），所以提前返回即可。
# <翻译结束>


<原文开始>
//  SetArrayFilter use for apply update array
//  For Example :
//  var res = QueryTestItem{}
//  change := Change{
//	Update:    bson.M{"$set": bson.M{"instock.$[elem].qty": 100}},
//	ReturnNew: false,
//  }
//  cli.Find(context.Background(), bson.M{"name": "Lucas"}).
//      SetArrayFilters(&options.ArrayFilters{Filters: []interface{}{bson.M{"elem.warehouse": bson.M{"$in": []string{"C", "F"}}},}}).
//        Apply(change, &res)
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// Select is used to determine which fields are displayed or not displayed in the returned results
// Format: bson.M{"age": 1} means that only the age field is displayed
// bson.M{"age": 0} means to display other fields except age
// When _id is not displayed and is set to 0, it will be returned to display
<原文结束>

# <翻译开始>
// Select 用于确定在返回结果中哪些字段显示或不显示
// 格式：bson.M{"age": 1} 表示只显示 age 字段
// bson.M{"age": 0} 表示除 age 字段外的其他字段均显示
// 当 _id 不显示并设置为 0 时，它将被默认返回显示
# <翻译结束>


<原文开始>
// Hint sets the value for the Hint field.
// This should either be the index name as a string or the index specification
// as a document. The default value is nil, which means that no hint will be sent.
<原文结束>

# <翻译开始>
// Hint 设置Hint字段的值。
// 这个值应该要么是作为字符串的索引名，要么是作为文档的索引规范。
// 默认值为nil，这意味着不会发送任何提示。
# <翻译结束>


<原文开始>
// Limit limits the maximum number of documents found to n
// The default value is 0, and 0  means no limit, and all matching results are returned
// When the limit value is less than 0, the negative limit is similar to the positive limit, but the cursor is closed after returning a single batch result.
// Reference https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
<原文结束>

# <翻译开始>
// Limit 限制查询结果返回的最大文档数量为 n
// 默认值为 0，当设置为 0 时，表示没有限制，会返回所有匹配的结果
// 当 limit 值小于 0 时，负数的限制与正数类似，但会在返回单批次结果后关闭游标
// 参考文献：https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
# <翻译结束>


<原文开始>
// One query a record that meets the filter conditions
// If the search fails, an error will be returned
<原文结束>

# <翻译开始>
// 根据过滤条件查询一条记录
// 若查询未找到匹配项，则返回错误
# <翻译结束>


<原文开始>
// All query multiple records that meet the filter conditions
// The static type of result must be a slice pointer
<原文结束>

# <翻译开始>
// 根据过滤条件查询满足条件的多条记录
// 结果的静态类型必须是指向切片的指针
# <翻译结束>


<原文开始>
// Count count the number of eligible entries
<原文结束>

# <翻译开始>
// Count 计算符合条件的条目数量
# <翻译结束>


<原文开始>
// EstimatedCount count the number of the collection by using the metadata
<原文结束>

# <翻译开始>
// EstimatedCount 通过使用元数据估算集合的数量
# <翻译结束>


<原文开始>
// Distinct gets the unique value of the specified field in the collection and return it in the form of slice
// result should be passed a pointer to slice
// The function will verify whether the static type of the elements in the result slice is consistent with the data type obtained in mongodb
// reference https://docs.mongodb.com/manual/reference/command/distinct/
<原文结束>

# <翻译开始>
// Distinct 获取集合中指定字段的唯一值，并以切片形式返回
// 结果应通过指针传递给切片
// 该函数将验证结果切片中元素的静态类型与在mongodb中获取的数据类型是否一致
// 参考文献：https://docs.mongodb.com/manual/reference/command/distinct/
# <翻译结束>


<原文开始>
// Cursor gets a Cursor object, which can be used to traverse the query result set
// After obtaining the CursorI object, you should actively call the Close interface to close the cursor
<原文结束>

# <翻译开始>
// 获取一个Cursor对象，可用于遍历查询结果集
// 在获取到CursorI对象后，应主动调用Close接口关闭游标
# <翻译结束>


<原文开始>
// findOneAndDelete
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndDelete/
<原文结束>

# <翻译开始>
// findOneAndDelete
// 参考文献: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndDelete/
// 此函数用于在MongoDB中查找并删除一条文档（记录）
// 它首先会根据提供的查询条件找到集合中第一条匹配的文档
// 找到后，立即从集合中删除该文档，并返回被删除的文档内容
// 注意：此操作为原子操作，在多线程或分布式环境下能保证数据一致性
# <翻译结束>


<原文开始>
// findOneAndReplace
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndReplace/
<原文结束>

# <翻译开始>
// findOneAndReplace
// 参考文献: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndReplace/
// 此函数实现的功能是，在MongoDB数据库中查找并替换一条文档数据。
// 根据提供的查询条件在指定集合中查找匹配的第一条文档，并用新文档替换它。
//findOneAndReplace 函数用于对 MongoDB 集合执行“查找并替换”操作，
// 它会根据给定的查询条件找到第一条匹配的文档，然后使用新的文档数据进行替换。
# <翻译结束>


<原文开始>
// findOneAndUpdate
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndUpdate/
<原文结束>

# <翻译开始>
// findOneAndUpdate
// 参考文献: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndUpdate/
// 此函数用于在 MongoDB 集合中查找匹配的第一个文档并更新它。
// 它首先会按照给定的查询条件查找文档，如果找到则根据提供的更新操作符进行更新，
// 然后返回更新前的原始文档（默认行为）或更新后的文档（根据方法选项设置）。
// 这是 MongoDB 的一个核心 CRUD 操作，常用于原子性地更新数据。
# <翻译结束>

