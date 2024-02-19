
<原文开始>
// Collection is a handle to a MongoDB collection
<原文结束>

# <翻译开始>
// Collection 是对 MongoDB 集合的一个引用句柄
# <翻译结束>


<原文开始>
// InsertOne insert one document into the collection
// If InsertHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Reference: https://docs.mongodb.com/manual/reference/command/insert/
<原文结束>

# <翻译开始>
// InsertOne 将一个文档插入集合中
// 如果 opts 中设置了 InsertHook，那么钩子会作用于它，否则钩子尝试将文档当作钩子处理
// 参考：https://docs.mongodb.com/manual/reference/command/insert/
# <翻译结束>


<原文开始>
// InsertMany executes an insert command to insert multiple documents into the collection.
// If InsertHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Reference: https://docs.mongodb.com/manual/reference/command/insert/
<原文结束>

# <翻译开始>
// InsertMany 执行一个插入命令，将多个文档插入到集合中。
// 如果 opts 中设置了 InsertHook，则在该 hook 上执行操作；否则尝试将 doc 作为 hook 使用
// 参考文献：https://docs.mongodb.com/manual/reference/command/insert/
# <翻译结束>


<原文开始>
// interfaceToSliceInterface convert interface to slice interface
<原文结束>

# <翻译开始>
// interfaceToSliceInterface 将接口转换为切片接口
# <翻译结束>


<原文开始>
// Upsert updates one documents if filter match, inserts one document if filter is not match, Error when the filter is invalid
// The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil
// and cannot contain any update operators
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
// If replacement has "_id" field and the document is existed, please initial it with existing id(even with Qmgo default field feature).
// Otherwise, "the (immutable) field '_id' altered" error happens.
<原文结束>

# <翻译开始>
// Upsert：如果过滤条件匹配，则更新一条文档；如果不匹配，则插入一条文档。当过滤条件无效时，将返回错误。
// replacement 参数必须是一个用于替换所选文档的文档对象，不能为 nil，并且不能包含任何更新操作符。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
// 如果 replacement 中包含 "_id" 字段，并且该文档已存在，请确保初始化时使用现有 id（即使启用了 Qmgo 的默认字段特性）。
// 否则，将会出现 "（不可变字段）'_id' 被修改" 的错误。
# <翻译结束>


<原文开始>
// UpsertId updates one documents if id match, inserts one document if id is not match and the id will inject into the document
// The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil
// and cannot contain any update operators
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpsertId 如果_id匹配则更新一条文档，如果不匹配则插入一条文档，并将_id注入到该文档中
// replacement参数必须是一个用于替换所选文档的文档，不能为nil
// 且不能包含任何更新操作符
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
# <翻译结束>


<原文开始>
// UpdateOne executes an update command to update at most one document in the collection.
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpdateOne 执行一个更新命令，用于在集合中最多更新一条文档。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
# <翻译结束>


<原文开始>
// UpdateOne support upsert function
<原文结束>

# <翻译开始>
// UpdateOne 支持 Upsert 功能
# <翻译结束>


<原文开始>
// UpdateId executes an update command to update at most one document in the collection.
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpdateId 执行一个更新命令，用于在集合中最多更新一个文档。
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
# <翻译结束>


<原文开始>
// UpdateAll executes an update command to update documents in the collection.
// The matchedCount is 0 in UpdateResult if no document updated
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpdateAll 执行一个更新命令，用于更新集合中的文档。
// 如果没有文档被更新，UpdateResult 中的 matchedCount 为 0
// 参考文献：https://docs.mongodb.com/manual/reference/operator/update/
# <翻译结束>


<原文开始>
// ReplaceOne executes an update command to update at most one document in the collection.
// If UpdateHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Expect type of the doc is the define of user's document
<原文结束>

# <翻译开始>
// ReplaceOne 执行一个更新命令，最多替换集合中的一个文档。
// 如果opts中的UpdateHook已设置，则在该hook上执行操作，否则尝试将doc作为hook处理
// 预期doc的类型应为用户定义的文档类型
# <翻译结束>


<原文开始>
// RemoveId executes a delete command to delete at most one document from the collection.
<原文结束>

# <翻译开始>
// RemoveId 执行一个删除命令，从集合中最多删除一个文档。
# <翻译结束>


<原文开始>
// Aggregate executes an aggregate command against the collection and returns a AggregateI to get resulting documents.
<原文结束>

# <翻译开始>
// Aggregate 对集合执行聚合命令，并返回一个 AggregateI 以便获取结果文档。
# <翻译结束>


<原文开始>
// EnsureIndexes Deprecated
// Recommend to use CreateIndexes / CreateOneIndex for more function)
// EnsureIndexes creates unique and non-unique indexes in collection
// the combination of indexes is different from CreateIndexes:
// if uniques/indexes is []string{"name"}, means create index "name"
// if uniques/indexes is []string{"name,-age","uid"} means create Compound indexes: name and -age, then create one index: uid
<原文结束>

# <翻译开始>
// EnsureIndexes 已弃用
// 建议使用 CreateIndexes / CreateOneIndex 以获得更多的功能)
// EnsureIndexes 在集合中创建唯一索引和非唯一索引
// 索引的组合方式与 CreateIndexes 不同：
// 如果 uniques/indexes 是 []string{"name"}，表示创建名为 "name" 的索引
// 如果 uniques/indexes 是 []string{"name,-age","uid"}，表示首先创建复合索引：name 和 -age（按 name 升序、age 降序），然后创建一个名为 uid 的单字段索引
# <翻译结束>


<原文开始>
// CreateIndexes creates multiple indexes in collection
// If the Key in opts.IndexModel is []string{"name"}, means create index: name
// If the Key in opts.IndexModel is []string{"name","-age"} means create Compound indexes: name and -age
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// CreateOneIndex creates one index
// If the Key in opts.IndexModel is []string{"name"}, means create index name
// If the Key in opts.IndexModel is []string{"name","-age"} means create Compound index: name and -age
<原文结束>

# <翻译开始>
// CreateOneIndex 创建一个索引
// 如果 opts.IndexModel 中的 Key 为 []string{"name"}，表示创建名为 "name" 的索引
// 如果 opts.IndexModel 中的 Key 为 []string{"name", "-age"}，表示创建组合索引：包含 name 和 -age（按 name 正序、age 倒序）
# <翻译结束>


<原文开始>
// DropAllIndexes drop all indexes on the collection except the index on the _id field
// if there is only _id field index on the collection, the function call will report an error
<原文结束>

# <翻译开始>
// DropAllIndexes 从集合中删除所有索引，但保留_id字段的索引
// 如果集合上只有_id字段的索引，函数调用将报告错误
# <翻译结束>


<原文开始>
// DropIndex drop indexes in collection, indexes that be dropped should be in line with inputting indexes
// The indexes is []string{"name"} means drop index: name
// The indexes is []string{"name","-age"} means drop Compound indexes: name and -age
<原文结束>

# <翻译开始>
// DropIndex 删除集合中的索引，需要删除的索引应与输入的索引一致
// 索引indexes为[]string{"name"}表示删除名为"name"的索引
// 索引indexes为[]string{"name","-age"}表示删除由"name"和"-age"组成的复合索引
# <翻译结束>


<原文开始>
// generate indexes that store in mongo which may consist more than one index(like []string{"index1","index2"} is stored as "index1_1_index2_1")
<原文结束>

# <翻译开始>
// 生成存储在MongoDB中的索引，这些索引可能包含多个索引（例如，[]string{"index1","index2"}将被存储为"index1_1_index2_1"）
# <翻译结束>


<原文开始>
// DropCollection drops collection
// it's safe even collection is not exists
<原文结束>

# <翻译开始>
// DropCollection 删除集合
// 即使集合不存在，此操作也是安全的
# <翻译结束>


<原文开始>
// CloneCollection creates a copy of the Collection
<原文结束>

# <翻译开始>
// CloneCollection 创建一个 Collection 的副本
# <翻译结束>


<原文开始>
// GetCollectionName returns the name of collection
<原文结束>

# <翻译开始>
// GetCollectionName 返回集合名称
# <翻译结束>


<原文开始>
// Watch returns a change stream for all changes on the corresponding collection. See
// https://docs.mongodb.com/manual/changeStreams/ for more information about change streams.
<原文结束>

# <翻译开始>
// Watch 返回一个变更流，用于接收相应集合的所有变更。有关变更流的更多信息，请参阅
// https://docs.mongodb.com/manual/changeStreams/
# <翻译结束>


<原文开始>
// translateUpdateResult translates mongo update result to qmgo define UpdateResult
<原文结束>

# <翻译开始>
// translateUpdateResult 将MongoDB更新结果翻译为qmgo定义的UpdateResult
# <翻译结束>

