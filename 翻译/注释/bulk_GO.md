
<原文开始>
// BulkResult is the result type returned by Bulk.Run operation.
<原文结束>

# <翻译开始>
// BulkResult 是 Bulk.Run 操作返回的结果类型。
# <翻译结束>


<原文开始>
// The number of documents inserted.
<原文结束>

# <翻译开始>
// 插入的文档数量。
# <翻译结束>


<原文开始>
// The number of documents matched by filters in update and replace operations.
<原文结束>

# <翻译开始>
// 在更新和替换操作中，满足过滤条件的文档数量。
# <翻译结束>


<原文开始>
// The number of documents modified by update and replace operations.
<原文结束>

# <翻译开始>
// update和replace操作修改的文档数量。
# <翻译结束>


<原文开始>
// The number of documents deleted.
<原文结束>

# <翻译开始>
// 删除的文档数量。
# <翻译结束>


<原文开始>
// The number of documents upserted by update and replace operations.
<原文结束>

# <翻译开始>
// update和replace操作中更新或替换的文档数量。
# <翻译结束>


<原文开始>
// A map of operation index to the _id of each upserted document.
<原文结束>

# <翻译开始>
// 一个映射表，键为操作索引，值为每个已更新（upserted）文档的 _id。
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
// Bulk 用于批量操作的上下文，这些批量操作将被一次性发送到数据库进行批量写入。
//
// Bulk 不支持并发安全使用。
//
// 注意事项：
//
// 当前，在一个批量操作内部的单个操作不会触发中间件或钩子。
//
// 与原始 mgo 不同，qmgo 实现的 Bulk 在不原生支持批量操作的老版本 MongoDB 服务器上，并不会模拟逐个执行批量操作。
//
// 只有官方驱动所支持的操作才会被公开提供，这就是为什么方法中缺少 InsertMany 的原因。
# <翻译结束>


<原文开始>
// Bulk returns a new context for preparing bulk execution of operations.
<原文结束>

# <翻译开始>
// Bulk 返回一个新的上下文，用于批量执行操作的准备工作。
# <翻译结束>


<原文开始>
// SetOrdered marks the bulk as ordered or unordered.
//
// If ordered, writes does not continue after one individual write fails.
// Default is ordered.
<原文结束>

# <翻译开始>
// SetOrdered 标记批量操作为有序或无序。
//
// 若标记为有序，当其中一次独立写入操作失败后，后续的写入操作将不再继续。
// 默认设置为有序。
# <翻译结束>


<原文开始>
// InsertOne queues an InsertOne operation for bulk execution.
<原文结束>

# <翻译开始>
// InsertOne 将一个 InsertOne 操作排队以进行批量执行。
# <翻译结束>


<原文开始>
// Remove queues a Remove operation for bulk execution.
<原文结束>

# <翻译开始>
// Remove 函数用于批量执行时，将一个 Remove 操作加入队列。
# <翻译结束>


<原文开始>
// RemoveId queues a RemoveId operation for bulk execution.
<原文结束>

# <翻译开始>
// RemoveId 为批量执行队列一个RemoveId操作。
# <翻译结束>


<原文开始>
// RemoveAll queues a RemoveAll operation for bulk execution.
<原文结束>

# <翻译开始>
// RemoveAll 函数用于批量执行，它将一个 RemoveAll 操作添加到待处理队列中。
# <翻译结束>


<原文开始>
// Upsert queues an Upsert operation for bulk execution.
// The replacement should be document without operator
<原文结束>

# <翻译开始>
// Upsert 在批量执行中安排一个Upsert操作。
// 替换项应为不包含操作符的文档
# <翻译结束>


<原文开始>
// UpsertOne queues an UpsertOne operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpsertOne 函数用于将一个 UpsertOne 操作加入到批量执行的队列中。
// 更新操作应当包含操作符
# <翻译结束>


<原文开始>
// UpsertId queues an UpsertId operation for bulk execution.
// The replacement should be document without operator
<原文结束>

# <翻译开始>
// UpsertId 队列一个 UpsertId 操作以便进行批量执行。
// 替换内容应为不包含操作符的文档。
# <翻译结束>


<原文开始>
// UpdateOne queues an UpdateOne operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpdateOne 将一个UpdateOne操作排队以进行批量执行。
// 更新内容应包含操作符
# <翻译结束>


<原文开始>
// UpdateId queues an UpdateId operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpdateId 将一个UpdateId操作加入队列以进行批量执行。
// 更新操作应包含操作符
# <翻译结束>


<原文开始>
// UpdateAll queues an UpdateAll operation for bulk execution.
// The update should contain operator
<原文结束>

# <翻译开始>
// UpdateAll 将一个UpdateAll操作添加到队列中以进行批量执行。
// 更新操作应包含操作符
# <翻译结束>


<原文开始>
// Run executes the collected operations in a single bulk operation.
//
// A successful call resets the Bulk. If an error is returned, the internal
// queue of operations is unchanged, containing both successful and failed
// operations.
<原文结束>

# <翻译开始>
// Run 执行收集到的所有操作，以单一的批量操作方式。
//
// 若调用成功，将会重置 Bulk。如果返回错误，则内部的操作队列保持不变，
// 该队列中包含已成功执行和未成功执行的操作。
# <翻译结束>


<原文开始>
// In original mgo, queue is not reset in case of error.
<原文结束>

# <翻译开始>
// 在原始mgo中，如果出现错误，队列不会被重置。
# <翻译结束>


<原文开始>
// Empty the queue for possible reuse, as per mgo's behavior.
<原文结束>

# <翻译开始>
// 根据mgo的行为，清空队列以备可能的重用。
# <翻译结束>

