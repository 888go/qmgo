
<原文开始>
// Collection is a handle to a MongoDB collection
<原文结束>

# <翻译开始>
// Collection 是一个MongoDB集合的句柄 md5:be1b94030609bdd1
# <翻译结束>


<原文开始>
// Find find by condition filter，return QueryI
<原文结束>

# <翻译开始>
// Find 通过条件过滤并查找，返回QueryI md5:bda4cc0c85d800a1
# <翻译结束>


<原文开始>
// InsertOne insert one document into the collection
// If InsertHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Reference: https://docs.mongodb.com/manual/reference/command/insert/
<原文结束>

# <翻译开始>
// InsertOne 将一个文档插入到集合中
// 如果在 opts 中设置了 InsertHook，钩子将在其上执行，否则钩子会尝试处理 doc 作为钩子
// 参考: https://docs.mongodb.com/manual/reference/command/insert/
// md5:0255181bb812302d
# <翻译结束>


<原文开始>
// InsertMany executes an insert command to insert multiple documents into the collection.
// If InsertHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Reference: https://docs.mongodb.com/manual/reference/command/insert/
<原文结束>

# <翻译开始>
// InsertMany 执行插入命令，将多个文档插入到集合中。如果opts中的InsertHook被设置，将在其上应用钩子；否则，尝试将doc作为钩子使用。
// 参考：https://docs.mongodb.com/manual/reference/command/insert/
// md5:49f2d7776e74fa79
# <翻译结束>


<原文开始>
// interfaceToSliceInterface convert interface to slice interface
<原文结束>

# <翻译开始>
// interfaceToSliceInterface 将接口类型转换为切片接口类型 md5:49f6ad81d7f669e3
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
// Upsert 如果过滤器匹配，则更新一个文档，如果过滤器不匹配，则插入一个文档。如果过滤器无效，会返回错误。
// 替换参数必须是一个将用于替换所选文档的文档。它不能为nil，且不能包含任何更新运算符。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// 如果替换参数包含"_id"字段并且文档已存在，请使用现有ID初始化（即使使用Qmgo默认字段特性也是如此）。否则，会引发"（不可变）字段 '_id' 被修改"的错误。
// md5:d7464af9e1e36d77
# <翻译结束>


<原文开始>
// UpsertId updates one documents if id match, inserts one document if id is not match and the id will inject into the document
// The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil
// and cannot contain any update operators
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpsertId 如果id匹配，则更新一个文档，如果id不匹配，则插入一个新的文档，并将id注入到文档中。
// 替换参数必须是一个将用于替换所选文档的文档。它不能为空，并且不能包含任何更新操作符。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:2a704aa664092959
# <翻译结束>


<原文开始>
// UpdateOne executes an update command to update at most one document in the collection.
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpdateOne 执行一个更新命令，最多更新集合中的一份文档。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:a16e90f28370dc2c
# <翻译结束>


<原文开始>
// UpdateOne support upsert function
<原文结束>

# <翻译开始>
// UpdateOne支持upsert功能 md5:aaec7189323f1660
# <翻译结束>


<原文开始>
// UpdateId executes an update command to update at most one document in the collection.
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpdateId 执行一个更新命令，最多更新集合中的一个文档。
// 参考：https://docs.mongodb.com/manual/reference/operator/update/
// md5:67764db9e90007e8
# <翻译结束>


<原文开始>
// UpdateAll executes an update command to update documents in the collection.
// The matchedCount is 0 in UpdateResult if no document updated
// Reference: https://docs.mongodb.com/manual/reference/operator/update/
<原文结束>

# <翻译开始>
// UpdateAll 执行更新命令以更新集合中的文档。
// 如果没有文档被更新，UpdateResult 中的 matchedCount 将为 0
// 参考资料: https://docs.mongodb.com/manual/reference/operator/update/
// md5:94c36e4a82312809
# <翻译结束>


<原文开始>
// ReplaceOne executes an update command to update at most one document in the collection.
// If UpdateHook in opts is set, hook works on it, otherwise hook try the doc as hook
// Expect type of the doc is the define of user's document
<原文结束>

# <翻译开始>
// ReplaceOne 执行更新命令，最多更新集合中的一个文档。如果 opts 中的 UpdateHook 被设置，那么 Hook 将在其上执行，否则 Hook 尝试将 doc 作为 Hook。预期 doc 的类型是用户定义的文档的定义。
// md5:1d830477f8b32e37
# <翻译结束>


<原文开始>
// Remove executes a delete command to delete at most one document from the collection.
// if filter is bson.M{}，DeleteOne will delete one document in collection
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
<原文结束>

# <翻译开始>
// Remove 执行删除命令，从集合中最多删除一个文档。
// 如果 filter 是 bson.M{}，DeleteOne 将删除集合中的一个文档。
// 参考：https://docs.mongodb.com/manual/reference/command/delete/
// md5:3b5cf64ce5f460b0
# <翻译结束>


<原文开始>
// RemoveId executes a delete command to delete at most one document from the collection.
<原文结束>

# <翻译开始>
// RemoveId 执行删除命令，从集合中删除最多一个文档。 md5:6516d8a8963d018c
# <翻译结束>


<原文开始>
// RemoveAll executes a delete command to delete documents from the collection.
// If filter is bson.M{}，all ducuments in Collection will be deleted
// Reference: https://docs.mongodb.com/manual/reference/command/delete/
<原文结束>

# <翻译开始>
// RemoveAll 执行一个删除命令，从集合中删除文档。如果 filter 是 bson.M{}（空映射），则会删除集合中的所有文档。
// 参考：https://docs.mongodb.com/manual/reference/command/delete/
// md5:abc51456adc5fc5a
# <翻译结束>


<原文开始>
// Aggregate executes an aggregate command against the collection and returns a AggregateI to get resulting documents.
<原文结束>

# <翻译开始>
// Aggregate 在集合上执行聚合命令，并返回一个 AggregateI，用于获取结果文档。 md5:e57ffed517c59fbc
# <翻译结束>


<原文开始>
// ensureIndex create multiple indexes on the collection and returns the names of
// Example：indexes = []string{"idx1", "-idx2", "idx3,idx4"}
// Three indexes will be created, index idx1 with ascending order, index idx2 with descending order, idex3 and idex4 are Compound ascending sort index
// Reference: https://docs.mongodb.com/manual/reference/command/createIndexes/
<原文结束>

# <翻译开始>
// ensureIndex 在集合上创建多个索引，并返回这些索引的名称。
// 示例：indexes = []string{"idx1", "-idx2", "idx3,idx4"}
// 将创建三个索引，idx1 为升序索引，idx2 为降序索引，idx3 和 idx4 为复合升序索引。
// 参考文档：https://docs.mongodb.com/manual/reference/command/createIndexes/
// md5:50a25575e53025b2
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
// 确保索引（已弃用）
// 建议使用CreateIndexes / CreateOneIndex以获取更多功能）
// EnsureIndexes 在集合中创建唯一和非唯一的索引，与CreateIndexes的组合不同：
// 如果uniques/indexes是`[]string{"name"}`，意味着创建名为"name"的索引
// 如果uniques/indexes是`[]string{"name,-age", "uid"}`，表示创建复合索引：name和-age，然后创建一个索引：uid
// md5:c595ad59f9c60c06
# <翻译结束>


<原文开始>
// CreateIndexes creates multiple indexes in collection
// If the Key in opts.IndexModel is []string{"name"}, means create index: name
// If the Key in opts.IndexModel is []string{"name","-age"} means create Compound indexes: name and -age
<原文结束>

# <翻译开始>
// CreateIndexes 在集合中创建多个索引
// 如果opts.IndexModel中的Key为[]string{"name"}，表示创建索引：name
// 如果opts.IndexModel中的Key为[]string{"name", "-age"}，表示创建复合索引：name和-age
// md5:822a787892c2186f
# <翻译结束>


<原文开始>
// CreateOneIndex creates one index
// If the Key in opts.IndexModel is []string{"name"}, means create index name
// If the Key in opts.IndexModel is []string{"name","-age"} means create Compound index: name and -age
<原文结束>

# <翻译开始>
// CreateOneIndex 创建一个索引
// 如果opts.IndexModel中的Key为[]string{"name"}，表示创建名为"name"的索引
// 如果opts.IndexModel中的Key为[]string{"name","-age"}，表示创建复合索引：按照"name"升序和"age"降序
// md5:70c27ea42ff3bbbf
# <翻译结束>


<原文开始>
// DropAllIndexes drop all indexes on the collection except the index on the _id field
// if there is only _id field index on the collection, the function call will report an error
<原文结束>

# <翻译开始>
// DropAllIndexes 会删除集合上除了_id字段索引之外的所有索引
// 如果集合上只有_id字段的索引，该函数调用将报告错误
// md5:e7655b40436f93df
# <翻译结束>


<原文开始>
// DropIndex drop indexes in collection, indexes that be dropped should be in line with inputting indexes
// The indexes is []string{"name"} means drop index: name
// The indexes is []string{"name","-age"} means drop Compound indexes: name and -age
<原文结束>

# <翻译开始>
// DropIndex 从集合中删除索引，需要删除的索引应与输入的索引列表匹配
// 索引是 []string{"name"} 表示删除名为 "name" 的单个索引
// 索引是 []string{"name", "-age"} 表示删除复合索引：name 和排除年龄 (-age) 的部分索引
// md5:4ad77e88557061c7
# <翻译结束>


<原文开始>
// generate indexes that store in mongo which may consist more than one index(like []string{"index1","index2"} is stored as "index1_1_index2_1")
<原文结束>

# <翻译开始>
// 生成存储在Mongo中的索引，可能包含多个索引（如[]string{"index1","index2"}存储为"index1_1_index2_1"） md5:15332a053c924233
# <翻译结束>


<原文开始>
// DropCollection drops collection
// it's safe even collection is not exists
<原文结束>

# <翻译开始>
// DropIndexDropIndex 会删除索引
// 即使索引不存在，这个操作也是安全的
// md5:e7b65cd93b1de7f7
# <翻译结束>


<原文开始>
// CloneCollection creates a copy of the Collection
<原文结束>

# <翻译开始>
// CloneCollection 创建集合的副本 md5:5df787f1c8ebab26
# <翻译结束>


<原文开始>
// GetCollectionName returns the name of collection
<原文结束>

# <翻译开始>
// GetCollectionName 返回集合的名字 md5:440484db8f2a466d
# <翻译结束>


<原文开始>
// Watch returns a change stream for all changes on the corresponding collection. See
// https://docs.mongodb.com/manual/changeStreams/ for more information about change streams.
<原文结束>

# <翻译开始>
// Watch 返回对应集合上所有更改的流。有关更改流的更多信息，请参阅
// https://docs.mongodb.com/manual/changeStreams/。
// md5:7b59cfd256c148f3
# <翻译结束>


<原文开始>
// translateUpdateResult translates mongo update result to qmgo define UpdateResult
<原文结束>

# <翻译开始>
// translateUpdateResult 将Mongo的更新结果转换为qmgo定义的UpdateResult md5:cb683a73f25cfe75
# <翻译结束>

