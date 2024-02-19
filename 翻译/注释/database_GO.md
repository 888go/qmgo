
<原文开始>
// Database is a handle to a MongoDB database
<原文结束>

# <翻译开始>
// Database 是一个指向 MongoDB 数据库的句柄
# <翻译结束>


<原文开始>
// Collection gets collection from database
<原文结束>

# <翻译开始>
// Collection 从数据库获取集合
# <翻译结束>


<原文开始>
// GetDatabaseName returns the name of database
<原文结束>

# <翻译开始>
// GetDatabaseName 返回数据库名称
# <翻译结束>


<原文开始>
// DropDatabase drops database
<原文结束>

# <翻译开始>
// DropDatabase 删除数据库
# <翻译结束>


<原文开始>
// RunCommand executes the given command against the database.
//
// The runCommand parameter must be a document for the command to be executed. It cannot be nil.
// This must be an order-preserving type such as bson.D. Map types such as bson.M are not valid.
// If the command document contains a session ID or any transaction-specific fields, the behavior is undefined.
//
// The opts parameter can be used to specify options for this operation (see the options.RunCmdOptions documentation).
<原文结束>

# <翻译开始>
// RunCommand 执行给定的命令针对数据库。
//
// runCommand 参数必须是要执行的命令的文档，不能为 nil。它必须是一个保持顺序的类型，如 bson.D。Map 类型如 bson.M 不是有效的。
// 如果命令文档包含会话 ID 或任何事务特定字段，则行为未定义。
//
// opts 参数可用于为此次操作指定选项（请参阅 options.RunCmdOptions 文档）。
# <翻译结束>


<原文开始>
// CreateCollection executes a create command to explicitly create a new collection with the specified name on the
// server. If the collection being created already exists, this method will return a mongo.CommandError. This method
// requires driver version 1.4.0 or higher.
//
// The opts parameter can be used to specify options for the operation (see the options.CreateCollectionOptions
// documentation).
<原文结束>

# <翻译开始>
// CreateCollection 执行创建命令，用于在服务器上明确创建一个指定名称的新集合。如果要创建的集合已存在，则此方法将返回 mongo.CommandError。该方法需要驱动版本 1.4.0 或更高版本。
//
// opts 参数可用于为操作指定选项（请参阅 options.CreateCollectionOptions 文档）。
# <翻译结束>

