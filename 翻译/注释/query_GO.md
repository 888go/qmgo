
<原文开始>
// Query struct definition
<原文结束>

# <翻译开始>
// 定义查询结构体 md5:56541bbc29d4ce15
# <翻译结束>


<原文开始>
// BatchSize sets the value for the BatchSize field.
// Means the maximum number of documents to be included in each batch returned by the server.
<原文结束>

# <翻译开始>
// BatchSize 设置 BatchSize 字段的值。
// 它表示服务器返回的每批文档的最大数量。
// md5:66277d16095ac151
# <翻译结束>


<原文开始>
// Sort is Used to set the sorting rules for the returned results
// Format: "age" or "+age" means to sort the age field in ascending order, "-age" means in descending order
// When multiple sort fields are passed in at the same time, they are arranged in the order in which the fields are passed in.
// For example, {"age", "-name"}, first sort by age in ascending order, then sort by name in descending order
<原文结束>

# <翻译开始>
// Sort 用于设置返回结果的排序规则
// 格式："age" 或 "+age" 表示按年龄字段升序排序，"-age" 表示降序排序
// 同时传入多个排序字段时，按照字段传递的顺序进行排列。
// 例如，{"age", "-name"}，首先按年龄升序排序，然后按姓名降序排序。
// md5:0b1b9f5345904541
# <翻译结束>


<原文开始>
		// A nil bson.D will not correctly serialize, but this case is no-op
		// so an early return will do.
<原文结束>

# <翻译开始>
		// 一个空的bson.D不会正确地序列化，但这种情况下可以提前返回。
		// md5:c94b59dcb408353d
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
# <翻译结束>


<原文开始>
// Select is used to determine which fields are displayed or not displayed in the returned results
// Format: bson.M{"age": 1} means that only the age field is displayed
// bson.M{"age": 0} means to display other fields except age
// When _id is not displayed and is set to 0, it will be returned to display
<原文结束>

# <翻译开始>
// Select用于确定在返回结果中显示或不显示哪些字段
// 格式：bson.M{"age": 1} 表示只显示age字段
// bson.M{"age": 0} 表示除了age以外的其他字段都显示
// 如果不显示_id并且设置为0，它将被返回并显示
// md5:3beb3c9bd51ad3fe
# <翻译结束>


<原文开始>
// Hint sets the value for the Hint field.
// This should either be the index name as a string or the index specification
// as a document. The default value is nil, which means that no hint will be sent.
<原文结束>

# <翻译开始>
// Hint 设置Hint字段的值。这应该是字符串形式的索引名称，或者是文档形式的索引规范。默认值为nil，表示不发送提示。
// md5:3d3535508606dd43
# <翻译结束>


<原文开始>
// Limit limits the maximum number of documents found to n
// The default value is 0, and 0  means no limit, and all matching results are returned
// When the limit value is less than 0, the negative limit is similar to the positive limit, but the cursor is closed after returning a single batch result.
// Reference https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
<原文结束>

# <翻译开始>
// Limit 将找到的最大文档数量限制为 n
// 默认值为 0，0 表示无限制，返回所有匹配的结果
// 当限制值小于 0 时，负限制类似于正限制，但返回单个批次结果后关闭游标。
// 参考 https://docs.mongodb.com/manual/reference/method/cursor.limit/index.html
// md5:9081095bd35be08f
# <翻译结束>


<原文开始>
// One query a record that meets the filter conditions
// If the search fails, an error will be returned
<原文结束>

# <翻译开始>
// 对符合过滤条件的记录执行一次查询
// 如果搜索失败，将返回一个错误
// md5:68571c814c5cd088
# <翻译结束>


<原文开始>
// All query multiple records that meet the filter conditions
// The static type of result must be a slice pointer
<原文结束>

# <翻译开始>
// 用于查询满足过滤条件的所有记录
// 结果的静态类型必须是切片指针
// md5:5f57d8aff8afe252
# <翻译结束>


<原文开始>
// Count count the number of eligible entries
<原文结束>

# <翻译开始>
// Count 计算符合条件的条目数量 md5:7bed3eaaee1ce368
# <翻译结束>


<原文开始>
// EstimatedCount count the number of the collection by using the metadata
<原文结束>

# <翻译开始>
// EstimatedCount 通过元数据计算集合的数量,
// EstimatedDocumentCount() 方法比 CountDocuments() 方法更快，因为它使用集合的元数据而不是扫描整个集合。
// md5:8c9bd7e463139421
# <翻译结束>


<原文开始>
// Distinct gets the unique value of the specified field in the collection and return it in the form of slice
// result should be passed a pointer to slice
// The function will verify whether the static type of the elements in the result slice is consistent with the data type obtained in mongodb
// reference https://docs.mongodb.com/manual/reference/command/distinct/
<原文结束>

# <翻译开始>
// Distinct 从集合中获取指定字段的唯一值，并以切片形式返回。
// result 应该是一个指向切片的指针。
// 函数会检查result切片元素的静态类型是否与MongoDB中获取的数据类型一致。
// 参考：https://docs.mongodb.com/manual/reference/command/distinct/
// md5:b83f3aa5718b2dfd
# <翻译结束>


<原文开始>
// Cursor gets a Cursor object, which can be used to traverse the query result set
// After obtaining the CursorI object, you should actively call the Close interface to close the cursor
<原文结束>

# <翻译开始>
// Cursor 获取一个 Cursor 对象，可用于遍历查询结果集
// 在获取到 CursorI 对象后，应主动调用 Close 接口来关闭游标
// md5:b1e9fc62a5f777fe
# <翻译结束>


<原文开始>
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
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// findOneAndDelete
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndDelete/
<原文结束>

# <翻译开始>
// findOneAndDelete
// 参考：https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndDelete/ 
// 
// 这段 Go 代码的注释表示这是一个名为 "findOneAndDelete" 的函数或方法，它用于从 MongoDB 集合中查找并删除第一个匹配的文档。参考链接指向 MongoDB 官方文档，提供了关于该方法的详细说明。
// md5:23b36fd4f1711d7b
# <翻译结束>


<原文开始>
// findOneAndReplace
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndReplace/
<原文结束>

# <翻译开始>
// findOneAndReplace
// 参考：https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndReplace/
// md5:cd4b97a5409057c1
# <翻译结束>


<原文开始>
// findOneAndUpdate
// reference: https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndUpdate/
<原文结束>

# <翻译开始>
// findOneAndUpdate
// 参考：https://docs.mongodb.com/manual/reference/method/db.collection.findOneAndUpdate/ 
// 
// 这段Go代码中的注释表示这是一个名为`findOneAndUpdate`的方法，它引用了MongoDB文档中关于`db.collection.findOneAndUpdate`方法的手册参考。这个方法在MongoDB数据库操作中用于根据指定的条件查找并更新第一个匹配的文档。
// md5:fe84856a45a0b0f1
# <翻译结束>

