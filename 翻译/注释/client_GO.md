
<原文开始>
// Config for initial mongodb instance
<原文结束>

# <翻译开始>
// Config 是初始 MongoDB 实例的配置
# <翻译结束>


<原文开始>
	// URI example: [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	// URI Reference: https://docs.mongodb.com/manual/reference/connection-string/
<原文结束>

# <翻译开始>
// URI 示例: [mongodb://][用户名:密码@]主机1[:端口1][,主机2[:端口2],...][/数据库][?选项]
// URI 参考文档: https://docs.mongodb.com/manual/reference/connection-string/
// 这段Go语言代码的注释是关于MongoDB数据库连接字符串（URI）的格式说明：
// - `mongodb://`：表示URI的协议部分，表明这是用于连接MongoDB服务器的地址。
// - `[user:pass@]`：可选的认证信息部分，其中`user`代表用户名，`pass`为经过编码的密码。
// - `host1[:port1][,host2[:port2],...]`：必填的服务器地址和端口部分，可以指定一个或多个服务器及对应端口，用逗号分隔。
// - `[/database]`：可选的数据库名称部分，用于指定默认连接的数据库。
// - `[?options]`：可选的连接参数部分，以问号开头，后面跟随一系列键值对（key=value&key=value...），用于设置额外的连接选项。
# <翻译结束>


<原文开始>
	// ConnectTimeoutMS specifies a timeout that is used for creating connections to the server.
	//	If set to 0, no timeout will be used.
	//	The default is 30 seconds.
<原文结束>

# <翻译开始>
// ConnectTimeoutMS 指定一个用于建立到服务器连接的超时时间。
//	如果设置为0，将不使用超时。
//	默认值是30秒。
# <翻译结束>


<原文开始>
	// MaxPoolSize specifies that maximum number of connections allowed in the driver's connection pool to each server.
	// If this is 0, it will be set to math.MaxInt64,
	// The default is 100.
<原文结束>

# <翻译开始>
// MaxPoolSize 指定驱动程序连接池中允许的每个服务器最大连接数。
// 如果该值为0，则会被设置为 math.MaxInt64，
// 默认值是100。
# <翻译结束>


<原文开始>
	// MinPoolSize specifies the minimum number of connections allowed in the driver's connection pool to each server. If
	// this is non-zero, each server's pool will be maintained in the background to ensure that the size does not fall below
	// the minimum. This can also be set through the "minPoolSize" URI option (e.g. "minPoolSize=100"). The default is 0.
<原文结束>

# <翻译开始>
// MinPoolSize 指定驱动程序与每个服务器连接池中允许的最小连接数。如果该值不为零，
// 则会后台维护每个服务器的连接池，确保其大小不低于最小值。也可以通过 "minPoolSize" URI 选项（例如 "minPoolSize=100"）进行设置。
// 默认值为 0。
# <翻译结束>


<原文开始>
	// SocketTimeoutMS specifies how long the driver will wait for a socket read or write to return before returning a
	// network error. If this is 0 meaning no timeout is used and socket operations can block indefinitely.
	// The default is 300,000 ms.
<原文结束>

# <翻译开始>
// SocketTimeoutMS 指定了在返回网络错误之前，驱动程序将等待套接字读写操作返回的时间。如果该值为0，则表示不使用超时，套接字操作可能会无限期阻塞。默认值为300,000毫秒。
# <翻译结束>


<原文开始>
	// ReadPreference determines which servers are considered suitable for read operations.
	// default is PrimaryMode
<原文结束>

# <翻译开始>
// ReadPreference 决定哪些服务器被认为适合读取操作。
// 默认设置为 PrimaryMode
# <翻译结束>


<原文开始>
// can be used to provide authentication options when configuring a Client.
<原文结束>

# <翻译开始>
// 可用于在配置 Client 时提供身份验证选项。
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
// Credential 可用于在配置 Client 时提供认证选项。
//
// AuthMechanism: 指定用于认证的机制。支持的值包括 "SCRAM-SHA-256", "SCRAM-SHA-1",
// "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509" 和 "MONGODB-AWS"。也可以通过 "authMechanism"
// URI 选项设置（例如："authMechanism=PLAIN"）。更多信息请参阅：
// https://docs.mongodb.com/manual/core/authentication-mechanisms/.
// AuthSource: 用于认证的数据库名称。对于 MONGODB-X509、GSSAPI 和 PLAIN，该值默认为 "$external"，
// 对于其他机制，默认为 "admin"。也可以通过 "authSource" URI 选项设置（例如："authSource=otherDb"）。
//
// Username: 认证所需的用户名。也可以通过 URI 以 username:password 的形式在第一个 @ 字符前设置。例如，
// 对于用户名为 "user"，密码为 "pwd"，主机为 "localhost:27017" 的情况，URI 应为 "mongodb://user:pwd@localhost:27017"。
// 对于 X509 认证，这是可选的，并且如果不指定，将从客户端证书中提取。
//
// Password: 认证所需的密码。对于 X509 认证，不应指定密码；对于 GSSAPI 认证，密码是可选的。
//
// PasswordSet: 对于 GSSAPI，如果指定了密码（即使密码为空字符串），则必须为 true；如果没有指定密码，
// 表示应从运行进程的上下文中获取密码，则应为 false。对于其他认证机制，此字段将被忽略。
# <翻译结束>


<原文开始>
// ReadPref determines which servers are considered suitable for read operations.
<原文结束>

# <翻译开始>
// ReadPref 决定哪些服务器适合读取操作。
# <翻译结束>


<原文开始>
	// MaxStaleness is the maximum amount of time to allow a server to be considered eligible for selection.
	// Supported from version 3.4.
<原文结束>

# <翻译开始>
// MaxStaleness 表示服务器被视为可选的最大过时时间。
// 该特性从版本 3.4 开始支持。
# <翻译结束>


<原文开始>
	// indicates the user's preference on reads.
	// PrimaryMode as default
<原文结束>

# <翻译开始>
// 表示用户对读取操作的偏好。
// 默认为 PrimaryMode
# <翻译结束>


<原文开始>
// QmgoClient specifies the instance to operate mongoDB
<原文结束>

# <翻译开始>
// QmgoClient 指定了操作 MongoDB 的实例
# <翻译结束>


<原文开始>
// Client creates client to mongo
<原文结束>

# <翻译开始>
// Client 创建用于连接 MongoDB 的客户端
# <翻译结束>


<原文开始>
// NewClient creates Qmgo MongoDB client
<原文结束>

# <翻译开始>
// NewClient 创建 Qmgo MongoDB 客户端
# <翻译结束>


<原文开始>
// client creates connection to MongoDB
<原文结束>

# <翻译开始>
// client 创建与 MongoDB 的连接
# <翻译结束>


<原文开始>
// half of default connect timeout
<原文结束>

# <翻译开始>
// 连接超时时间的一半
# <翻译结束>


<原文开始>
// newAuth create options.Credential from conf.Auth
<原文结束>

# <翻译开始>
// newAuth 从 conf.Auth 创建 options.Credential
# <翻译结束>


<原文开始>
// Validate and process the username.
<原文结束>

# <翻译开始>
// 验证并处理用户名。
# <翻译结束>


<原文开始>
// newReadPref create readpref.ReadPref from config
<原文结束>

# <翻译开始>
// newReadPref 从配置中创建 readpref.ReadPref
# <翻译结束>


<原文开始>
// Close closes sockets to the topology referenced by this Client.
<原文结束>

# <翻译开始>
// Close关闭与此Client关联的拓扑结构的所有套接字连接。
# <翻译结束>


<原文开始>
// Ping confirm connection is alive
<原文结束>

# <翻译开始>
// Ping：确认连接是否存活
# <翻译结束>


<原文开始>
// Database create connection to database
<原文结束>

# <翻译开始>
// 创建与数据库的连接
# <翻译结束>


<原文开始>
// Session create one session on client
// Watch out, close session after operation done
<原文结束>

# <翻译开始>
// Session 在客户端创建一个会话
// 注意，在操作完成后关闭会话
# <翻译结束>


<原文开始>
// ServerVersion get the version of mongoDB server, like 4.4.0
<原文结束>

# <翻译开始>
// ServerVersion 获取MongoDB服务器的版本，如4.4.0
# <翻译结束>


<原文开始>
// transactionAllowed check if transaction is allowed
<原文结束>

# <翻译开始>
// transactionAllowed 检查交易是否被允许
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
// TODO 未知为何需要在topology()函数中执行`cli, err := Open(ctx, &c.conf)`来获取topo，
// 在查明原因之前，我们仅在单元测试（UT）中使用此函数
//topo, err := c.topology()
//如果topo是description.Single类型 {
//	 fmt.Println("由于Mongo服务器拓扑为单节点，因此不支持事务")
//	 返回false
//}
# <翻译结束>

