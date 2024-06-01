
<原文开始>
// Session is an struct that represents a MongoDB logical session
<原文结束>

# <翻译开始>
// Session 是一个结构体，表示 MongoDB 的逻辑会话 md5:a17367bc3a251e77
# <翻译结束>


<原文开始>
// StartTransaction starts transaction
//precondition：
//- version of mongoDB server >= v4.0
//- Topology of mongoDB server is not Single
//At the same time, please pay attention to the following
//- make sure all operations in callback use the sessCtx as context parameter
//- Dont forget to call EndSession if session is not used anymore
//- if operations in callback takes more than(include equal) 120s, the operations will not take effect,
//- if operation in callback return qmgo.ErrTransactionRetry,
//  the whole transaction will retry, so this transaction must be idempotent
//- if operations in callback return qmgo.ErrTransactionNotSupported,
//- If the ctx parameter already has a Session attached to it, it will be replaced by this session.
<原文结束>

# <翻译开始>
// StartTransaction 开始一个事务
// 预条件：
// - MongoDB服务器版本大于等于v4.0
// - MongoDB服务器的拓扑结构不是单节点
// 同时需要注意：
// - 确保回调中的所有操作将sessCtx作为上下文参数
// - 如果不再使用session，别忘了调用EndSession
// - 如果回调中的操作耗时超过（包括等于）120秒，这些操作将不会生效
// - 如果回调中的操作返回qmgo.ErrTransactionRetry错误，
//   整个事务将会重试，因此这个事务必须是幂等的
// - 如果回调中的操作返回qmgo.ErrTransactionNotSupported错误，
// - 如果ctx参数中已经附加了一个Session，它将被此session替换。
// md5:7a854b4c45212490
# <翻译结束>


<原文开始>
// EndSession will abort any existing transactions and close the session.
<原文结束>

# <翻译开始>
// EndSession 会终止任何现有的事务并关闭会话。 md5:2ee8849531868b7e
# <翻译结束>


<原文开始>
// AbortTransaction aborts the active transaction for this session. This method will return an error if there is no
// active transaction for this session or the transaction has been committed or aborted.
<原文结束>

# <翻译开始>
// AbortTransaction 会取消此会话中的活动事务。如果此会话没有活动事务，或者事务已经提交或中止，此方法将返回错误。
// md5:ca9bc056086304f0
# <翻译结束>


<原文开始>
// wrapperCustomF wrapper caller's callback function to mongo dirver's
<原文结束>

# <翻译开始>
// wrapperCustomF 将调用者的回调函数包装成mongo驱动所需的函数 md5:8df643188861ec8b
# <翻译结束>

