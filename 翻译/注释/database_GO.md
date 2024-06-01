
<原文开始>
// Database is a handle to a MongoDB database
<原文结束>

# <翻译开始>
// Database 是一个指向 MongoDB 数据库的句柄 md5:9217ae5bd9047e3a
# <翻译结束>


<原文开始>
// Collection gets collection from database
<原文结束>

# <翻译开始>
// Collection 从数据库中获取集合 md5:c5489f5523d5a33d
# <翻译结束>


<原文开始>
// GetDatabaseName returns the name of database
<原文结束>

# <翻译开始>
// GetDatabaseName 返回数据库的名称 md5:716064a488e6db8b
# <翻译结束>


<原文开始>
// DropDatabase drops database
<原文结束>

# <翻译开始>
// DropDatabase 删除数据库 md5:aeac2378daa25d5f
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
// RunCommand 在数据库上执行给定的命令。
//
// runCommand 参数必须是将要执行的命令文档。它不能为 nil。这必须是一个保持顺序的类型，如 bson.D。像 bson.M 这样的映射类型是无效的。
// 如果命令文档包含会话 ID 或任何事务特定字段，其行为是未定义的。
//
// 可以使用 opts 参数来指定此操作的选项（参阅 options.RunCmdOptions 的文档）。
// md5:eb93f7217a15650c
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
// CreateCollection 执行一个创建命令，明确在服务器上使用指定名称创建一个新的集合。如果正在创建的集合已经存在，此方法将返回一个 mongo.CommandError。此方法需要驱动程序版本 1.4.0 或更高版本。
// 
// 参数 opts 可用于指定操作选项（请参阅 options.CreateCollectionOptions 的文档）。
// md5:7bd165db4ed05d28
# <翻译结束>

