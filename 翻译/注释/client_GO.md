
<原文开始>
// Config for initial mongodb instance
<原文结束>

# <翻译开始>
// 初始MongoDB实例的配置 md5:09dcbab1d00adb46
# <翻译结束>


<原文开始>
	// URI example: [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	// URI Reference: https://docs.mongodb.com/manual/reference/connection-string/
<原文结束>

# <翻译开始>
	// URI 示例：[mongodb:	//][user:pass@]主机1[:端口1][,主机2[:端口2],...][/数据库][?选项]
	// URI 参考：https:	//docs.mongodb.com/manual/reference/connection-string/ 
	// 
	// 这段注释解释了一个MongoDB连接字符串的格式，包括可选的部分如用户名、密码、多个服务器地址、数据库名以及可选的连接选项。URI以`mongodb:	//`开头，后面可以包含认证信息、主机列表、数据库路径和查询参数。链接：提供了官方文档的参考。
	// md5:038c28929efbdde0
# <翻译结束>


<原文开始>
	// ConnectTimeoutMS specifies a timeout that is used for creating connections to the server.
	//	If set to 0, no timeout will be used.
	//	The default is 30 seconds.
<原文结束>

# <翻译开始>
	// ConnectTimeoutMS 指定了建立到服务器连接时使用的超时时间，以毫秒为单位。
	// 如果设置为 0，则不会使用超时。
	// 默认值为 30 秒。
	// md5:bdc6b23048c25478
# <翻译结束>


<原文开始>
	// MaxPoolSize specifies that maximum number of connections allowed in the driver's connection pool to each server.
	// If this is 0, it will be set to math.MaxInt64,
	// The default is 100.
<原文结束>

# <翻译开始>
	// MaxPoolSize 指定驱动程序连接池到每个服务器的最大连接数。
	// 如果设置为 0，则将其设置为 math.MaxInt64，
	// 默认值为 100。
	// md5:6840c2846a8fad6e
# <翻译结束>


<原文开始>
	// MinPoolSize specifies the minimum number of connections allowed in the driver's connection pool to each server. If
	// this is non-zero, each server's pool will be maintained in the background to ensure that the size does not fall below
	// the minimum. This can also be set through the "minPoolSize" URI option (e.g. "minPoolSize=100"). The default is 0.
<原文结束>

# <翻译开始>
	// MinPoolSize 指定了驱动程序到每个服务器的连接池中允许的最小连接数。如果此值不为零，将为每个服务器的连接池在后台维护，以确保其大小不低于最小值。这也可以通过 "minPoolSize" URI 选项（如 "minPoolSize=100"）进行设置。默认值为 0。
	// md5:9df8b44a6800236b
# <翻译结束>


<原文开始>
	// SocketTimeoutMS specifies how long the driver will wait for a socket read or write to return before returning a
	// network error. If this is 0 meaning no timeout is used and socket operations can block indefinitely.
	// The default is 300,000 ms.
<原文结束>

# <翻译开始>
	// SocketTimeoutMS 指定了驱动程序在返回网络错误之前，等待套接字读写操作返回的最长时间（以毫秒为单位）。如果此值为0，则表示不使用超时，套接字操作可能无限期阻塞。默认值为300,000毫秒。
	// md5:1e1ccf1f35a18417
# <翻译结束>


<原文开始>
	// ReadPreference determines which servers are considered suitable for read operations.
	// default is PrimaryMode
<原文结束>

# <翻译开始>
	// ReadPreference 确定哪些服务器适合进行读取操作。默认为 PrimaryMode。
	// md5:6ca3a191c28443b8
# <翻译结束>


<原文开始>
// can be used to provide authentication options when configuring a Client.
<原文结束>

# <翻译开始>
// 可用于在配置客户端时提供身份验证选项。 md5:99c19d7fabc83d2d
# <翻译结束>


<原文开始>
// Credential can be used to provide authentication options when configuring a Client.
//
// AuthMechanism: the mechanism to use for authentication. Supported values include "SCRAM-SHA-256", "SCRAM-SHA-1",
// "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509", and "MONGODB-AWS". This can also be set through the "authMechanism"
// URI option. (e.g. "authMechanism=PLAIN"). For more information, see
// https://docs.mongodb.com/manual/core/authentication-mechanisms/.
// AuthSource: the name of the database to use for authentication. This defaults to "$external" for MONGODB-X509,
// GSSAPI, and PLAIN and "admin" for all other mechanisms. This can also be set through the "authSource" URI option
// (e.g. "authSource=otherDb").
//
// Username: the username for authentication. This can also be set through the URI as a username:password pair before
// the first @ character. For example, a URI for user "user", password "pwd", and host "localhost:27017" would be
// "mongodb://user:pwd@localhost:27017". This is optional for X509 authentication and will be extracted from the
// client certificate if not specified.
//
// Password: the password for authentication. This must not be specified for X509 and is optional for GSSAPI
// authentication.
//
// PasswordSet: For GSSAPI, this must be true if a password is specified, even if the password is the empty string, and
// false if no password is specified, indicating that the password should be taken from the context of the running
// process. For other mechanisms, this field is ignored.
<原文结束>

# <翻译开始>
// Credential 用于在配置客户端时提供认证选项。
//
// AuthMechanism: 认证机制。支持的值包括 "SCRAM-SHA-256", "SCRAM-SHA-1", "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509" 和 "MONGODB-AWS"。
// 这也可以通过 "authMechanism" URI 选项设置（例如 "authMechanism=PLAIN"）。更多信息请参阅
// https://docs.mongodb.com/manual/core/authentication-mechanisms/。
// AuthSource: 用于认证的数据库名称。对于 MONGODB-X509、GSSAPI 和 PLAIN，默认为 "$external"，对于所有其他机制默认为 "admin"。
// 这也可以通过 "authSource" URI 选项设置（例如 "authSource=otherDb"）。
//
// Username: 认证用的用户名。这也可以通过 URI 在第一个 @ 字符前设置用户名和密码。例如，用户 "user"，密码 "pwd"，
// 主机 "localhost:27017" 的 URI 为 "mongodb://user:pwd@localhost:27017"。对于 X509 认证这是可选的，如果没有指定，将从客户端证书中提取。
//
// Password: 认证用的密码。对于 X509 不允许指定，对于 GSSAPI 是可选的。
//
// PasswordSet: 对于 GSSAPI，如果指定了密码（即使密码为空字符串），此值必须为 true，如果未指定密码，表示应从运行进程的上下文中获取密码。
// 对于其他机制，此字段会被忽略。
// md5:e1c2a73d163c799a
# <翻译结束>


<原文开始>
// ReadPref determines which servers are considered suitable for read operations.
<原文结束>

# <翻译开始>
// ReadPref确定哪些服务器适合进行读取操作。 md5:d5ae507a40965ac9
# <翻译结束>


<原文开始>
	// MaxStaleness is the maximum amount of time to allow a server to be considered eligible for selection.
	// Supported from version 3.4.
<原文结束>

# <翻译开始>
	// MaxStaleness是允许服务器被认为适合选择的最大时间。从版本3.4开始支持。
	// md5:01c3097a5d9a368b
# <翻译结束>


<原文开始>
	// indicates the user's preference on reads.
	// PrimaryMode as default
<原文结束>

# <翻译开始>
	// 表示用户在读取操作上的偏好。
	// 默认为PrimaryMode。
	// md5:85d94814e6ac8eca
# <翻译结束>


<原文开始>
// QmgoClient specifies the instance to operate mongoDB
<原文结束>

# <翻译开始>
// QmgoClient 指定操作MongoDB的实例 md5:ef9044b4ab2af757
# <翻译结束>


<原文开始>
// Open creates client instance according to config
// QmgoClient can operates all qmgo.client 、qmgo.database and qmgo.collection
<原文结束>

# <翻译开始>
// Open 根据配置创建客户端实例
// QmgoClient 可以操作所有 qmgo.client、qmgo.database 和 qmgo.collection
// md5:bc872aaa93cf801a
# <翻译结束>


<原文开始>
// Client creates client to mongo
<原文结束>

# <翻译开始>
// Client 创建一个到Mongo的客户端 md5:3527d3de272044c3
# <翻译结束>


<原文开始>
// NewClient creates Qmgo MongoDB client
<原文结束>

# <翻译开始>
// NewClient 创建 Qmgo MongoDB 客户端 md5:64c9dc0f30edc1ac
# <翻译结束>


<原文开始>
// client creates connection to MongoDB
<原文结束>

# <翻译开始>
// 客户端创建到MongoDB的连接 md5:5ed46d6e6a970651
# <翻译结束>


<原文开始>
// half of default connect timeout
<原文结束>

# <翻译开始>
// 默认连接超时时间的一半 md5:e544afad71f167e7
# <翻译结束>


<原文开始>
// newConnectOpts creates client options from conf
// Qmgo will follow this way official mongodb driver do：
// - the configuration in uri takes precedence over the configuration in the setter
// - Check the validity of the configuration in the uri, while the configuration in the setter is basically not checked
<原文结束>

# <翻译开始>
// newConnectOpts 从 conf 创建客户端选项
// Qmgo 将遵循官方 MongoDB 驱动程序的做法：
// - URI 中的配置优先于 setter 中的配置
// - 检查 URI 中配置的有效性，而 setter 中的配置基本不进行检查
// md5:e686e2f8bec69b3b
# <翻译结束>


<原文开始>
// newAuth create options.Credential from conf.Auth
<原文结束>

# <翻译开始>
// newAuth 从conf.Auth创建options.Credential选项 md5:88ce8258f4551f1c
# <翻译结束>


<原文开始>
// Validate and process the username.
<原文结束>

# <翻译开始>
// 验证和处理用户名。 md5:3c89ddb7c004c9d6
# <翻译结束>


<原文开始>
// newReadPref create readpref.ReadPref from config
<原文结束>

# <翻译开始>
// newReadPref 根据配置创建 readpref.ReadPref md5:1c0e9080aed7b202
# <翻译结束>


<原文开始>
// Close closes sockets to the topology referenced by this Client.
<原文结束>

# <翻译开始>
// Close 关闭到此客户端引用的拓扑结构相关的套接字。 md5:a2c78aacda5cd470
# <翻译结束>


<原文开始>
// Ping confirm connection is alive
<原文结束>

# <翻译开始>
// Ping确认连接是否还活着 md5:1b88dbe0bbaa6726
# <翻译结束>


<原文开始>
// Database create connection to database
<原文结束>

# <翻译开始>
// Database 创建到数据库的连接 md5:1aa03639d9adcf41
# <翻译结束>


<原文开始>
// Session create one session on client
// Watch out, close session after operation done
<原文结束>

# <翻译开始>
// Session：在客户端创建一个会话
// 注意，操作完成后要关闭会话
// md5:a25c6035ffabaf48
# <翻译结束>


<原文开始>
// DoTransaction do whole transaction in one function
// precondition：
// - version of mongoDB server >= v4.0
// - Topology of mongoDB server is not Single
// At the same time, please pay attention to the following
// - make sure all operations in callback use the sessCtx as context parameter
// - if operations in callback takes more than(include equal) 120s, the operations will not take effect,
// - if operation in callback return qmgo.ErrTransactionRetry,
//   the whole transaction will retry, so this transaction must be idempotent
// - if operations in callback return qmgo.ErrTransactionNotSupported,
// - If the ctx parameter already has a Session attached to it, it will be replaced by this session.
<原文结束>

# <翻译开始>
// DoTransaction 在一个函数中执行整个事务
// 前置条件：
// - MongoDB服务器的版本 >= v4.0
// - MongoDB服务器的拓扑结构不是Single
// 同时需要注意以下几点：
// - 确保回调中的所有操作都使用sessCtx作为context参数
// - 如果回调中的操作（包括等于）耗时超过120s，操作将不会生效
// - 如果回调中的操作返回qmgo.ErrTransactionRetry，整个事务将重试，因此该事务必须幂等
// - 如果回调中的操作返回qmgo.ErrTransactionNotSupported，
// - 如果ctx参数已经绑定了Session，它将被这个Session替换。
// md5:f5555fc9e2733cb9
# <翻译结束>


<原文开始>
// ServerVersion get the version of mongoDB server, like 4.4.0
<原文结束>

# <翻译开始>
// ServerVersion 获取MongoDB服务器的版本，如4.4.0 md5:85f19b2205255d3a
# <翻译结束>


<原文开始>
// transactionAllowed check if transaction is allowed
<原文结束>

# <翻译开始>
// transactionAllowed 检查交易是否被允许 md5:d9e86f3ad9610912
# <翻译结束>


<原文开始>
	// TODO dont know why need to do `cli, err := Open(ctx, &c.conf)` in topology() to get topo,
	// Before figure it out, we only use this function in UT
	//topo, err := c.topology()
	//if topo == description.Single {
	//	fmt.Println("transaction is not supported because mongo server topology is single")
	//	return false
	//}
<原文结束>

# <翻译开始>
	// TODO：不知道为什么在`topology()`函数中需要通过`cli, err := Open(ctx, &c.conf)`来获取topo，在弄清楚原因之前，我们只在这个UT（单元测试）中使用这个函数
	//topo, err := c.topology() 	// 从config对象获取topology信息
	//如果topo是description.Single（单点模式）：
	//    打印 "transaction is not supported because mongo server topology is single"
	//    返回false
	// md5:4d3e4bc17382c028
# <翻译结束>

