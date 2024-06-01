
<原文开始>
// BulkResult is the result type returned by Bulk.Run operation.
<原文结束>

# <翻译开始>
// BulkResult 是由Bulk.Run操作返回的结果类型。 md5:3a422d6b1b20649c
# <翻译结束>


<原文开始>
// The number of documents inserted.
<原文结束>

# <翻译开始>
// 插入的文档数量。 md5:f44082352897f08b
# <翻译结束>


<原文开始>
// The number of documents matched by filters in update and replace operations.
<原文结束>

# <翻译开始>
// 更新和替换操作中，被过滤器匹配的文档数量。 md5:90fab681d83f2e97
# <翻译结束>


<原文开始>
// The number of documents modified by update and replace operations.
<原文结束>

# <翻译开始>
// 被更新和替换操作修改的文档数量。 md5:1e4886e32c8092e3
# <翻译结束>


<原文开始>
// The number of documents deleted.
<原文结束>

# <翻译开始>
// 删除的文档数量。 md5:8872e8629ebbcf3c
# <翻译结束>


<原文开始>
// The number of documents upserted by update and replace operations.
<原文结束>

# <翻译开始>
// 通过update和replace操作插入的文档数量。 md5:3074b4c76263ae0c
# <翻译结束>


<原文开始>
// A map of operation index to the _id of each upserted document.
<原文结束>

# <翻译开始>
// 一个操作索引到每个插入文档的_id的映射。 md5:b4c301dceb41d860
# <翻译结束>


<原文开始>
// Bulk is context for batching operations to be sent to database in a single
// bulk write.
//
// Bulk is not safe for concurrent use.
//
// Notes:
//
// Individual operations inside a bulk do not trigger middlewares or hooks
// at present.
//
// Different from original mgo, the qmgo implementation of Bulk does not emulate
// bulk operations individually on old versions of MongoDB servers that do not
// natively support bulk operations.
//
// Only operations supported by the official driver are exposed, that is why
// InsertMany is missing from the methods.
<原文结束>

# <翻译开始>
// Bulk 是用于批量操作的上下文，这些操作将一次性发送到数据库进行批量写入。
//
// Bulk 不适用于并发使用。
//
// 注意：
//
// 在批量操作中的单个操作目前不会触发中间件或钩子。
//
// 与原版 mgo 不同，qmgo 实现的 Bulk 并不会在不支持原生批量操作的老版本 MongoDB 服务器上模拟逐个执行批量操作。
//
// 只有官方驱动支持的操作被暴露出来，因此方法中缺少 InsertMany。
// md5:97e7f3c645b8ba7f
# <翻译结束>


<原文开始>
// Bulk returns a new context for preparing bulk execution of operations.
<原文结束>

# <翻译开始>
// Bulk返回一个新的上下文，用于准备批量执行操作。 md5:e39897d617450e92
# <翻译结束>


<原文开始>
// SetOrdered marks the bulk as ordered or unordered.
//
// If ordered, writes does not continue after one individual write fails.
// Default is ordered.
<原文结束>

# <翻译开始>
// SetOrdered 将批量设置为有序或无序。
//
// 如果设置为有序，写操作在单个写操作失败后不会继续。默认为有序。
// md5:caf2eac3fe50a750
# <翻译结束>


<原文开始>
// InsertOne queues an InsertOne operation for bulk execution.
<原文结束>

# <翻译开始>
// InsertOne 将一个 InsertOne 操作加入到批量执行队列中。 md5:65abbf989aa97556
# <翻译结束>


<原文开始>
// Remove queues a Remove operation for bulk execution.
<原文结束>

# <翻译开始>
// Remove 队列一个删除操作，用于批量执行。 md5:a9c84e1a291eea0f
# <翻译结束>


<原文开始>
// RemoveId queues a RemoveId operation for bulk execution.
<原文结束>

# <翻译开始>
// RemoveId 队列一个 RemoveId 操作以进行批量执行。 md5:f3fbfef26bde41fc
# <翻译结束>


<原文开始>
// RemoveAll queues a RemoveAll operation for bulk execution.
<原文结束>

# <翻译开始>
// RemoveAll 会将一个 RemoveAll 操作加入到批量执行的队列中。 md5:df548d516b324574
# <翻译结束>


<原文开始>
// Upsert queues an Upsert operation for bulk execution.
// The replacement should be document without operator
<原文结束>

# <翻译开始>
// Upsert将Upsert操作排队进行批量执行。替换应该是没有操作符的文档
// md5:1115932f50b88737
# <翻译结束>


<原文开始>
// UpsertOne queues an UpsertOne operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpsertOne 为批量执行队列一个 UpsertOne 操作。更新操作应该包含运算符
// md5:7052a86d53229aab
# <翻译结束>


<原文开始>
// UpsertId queues an UpsertId operation for bulk execution.
// The replacement should be document without operator
<原文结束>

# <翻译开始>
// UpsertId 用于批量执行的UpsertId操作进行排队。
// 替换的文档应该不包含操作符。
// md5:c5d9cc678823f8e5
# <翻译结束>


<原文开始>
// UpdateOne queues an UpdateOne operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpdateOne 为批量执行队列一个 UpdateOne 操作。更新操作应该包含操作符
// md5:0e587045b560687a
# <翻译结束>


<原文开始>
// UpdateId queues an UpdateId operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpdateId 为批量执行排队一个 UpdateId 操作。更新应该包含操作符
// md5:968d7d02f007ae39
# <翻译结束>


<原文开始>
// UpdateAll queues an UpdateAll operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpdateAll 队列一个 UpdateAll 操作，用于批量执行。
// 更新应该包含操作符
// md5:b1fdc26a48273948
# <翻译结束>


<原文开始>
// Run executes the collected operations in a single bulk operation.
//
// A successful call resets the Bulk. If an error is returned, the internal
// queue of operations is unchanged, containing both successful and failed
// operations.
<原文结束>

# <翻译开始>
// Run 执行收集到的单个批量操作。
//
// 调用成功会重置 Bulk。如果返回错误，内部操作队列保持不变，包含成功和失败的操作。
// md5:c3ce14d8defe8da0
# <翻译结束>


<原文开始>
// In original mgo, queue is not reset in case of error.
<原文结束>

# <翻译开始>
// 在原始的mgo中，如果发生错误，队列不会被重置。 md5:b7f801e955f364a8
# <翻译结束>


<原文开始>
// Empty the queue for possible reuse, as per mgo's behavior.
<原文结束>

# <翻译开始>
// 清空队列以备可能的重用，遵循mgo的行为。 md5:ac1070c096c485e8
# <翻译结束>

