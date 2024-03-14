/*
 Copyright 2020 The Qmgo Authors.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package mgo类

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
	
	"github.com/888go/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Config 是初始 MongoDB 实例的配置
type Config struct {
// URI 示例: [mongodb://][用户名:密码@]主机1[:端口1][,主机2[:端口2],...][/数据库][?选项]
// URI 参考文档: https://docs.mongodb.com/manual/reference/connection-string/
// 这段Go语言代码的注释是关于MongoDB数据库连接字符串（URI）的格式说明：
// - `mongodb://`：表示URI的协议部分，表明这是用于连接MongoDB服务器的地址。
// - `[user:pass@]`：可选的认证信息部分，其中`user`代表用户名，`pass`为经过编码的密码。
// - `host1[:port1][,host2[:port2],...]`：必填的服务器地址和端口部分，可以指定一个或多个服务器及对应端口，用逗号分隔。
// - `[/database]`：可选的数据库名称部分，用于指定默认连接的数据库。
// - `[?options]`：可选的连接参数部分，以问号开头，后面跟随一系列键值对（key=value&key=value...），用于设置额外的连接选项。
	Uri      string `json:"uri"`
	Database string `json:"database"`
	Coll     string `json:"coll"`
// ConnectTimeoutMS 指定一个用于建立到服务器连接的超时时间。
//	如果设置为0，将不使用超时。
//	默认值是30秒。
	ConnectTimeoutMS *int64 `json:"connectTimeoutMS"`
// MaxPoolSize 指定驱动程序连接池中允许的每个服务器最大连接数。
// 如果该值为0，则会被设置为 math.MaxInt64，
// 默认值是100。
	MaxPoolSize *uint64 `json:"maxPoolSize"`
// MinPoolSize 指定驱动程序与每个服务器连接池中允许的最小连接数。如果该值不为零，
// 则会后台维护每个服务器的连接池，确保其大小不低于最小值。也可以通过 "minPoolSize" URI 选项（例如 "minPoolSize=100"）进行设置。
// 默认值为 0。
	MinPoolSize *uint64 `json:"minPoolSize"`
// SocketTimeoutMS 指定了在返回网络错误之前，驱动程序将等待套接字读写操作返回的时间。如果该值为0，则表示不使用超时，套接字操作可能会无限期阻塞。默认值为300,000毫秒。
	SocketTimeoutMS *int64 `json:"socketTimeoutMS"`
// ReadPreference 决定哪些服务器被认为适合读取操作。
// 默认设置为 PrimaryMode
	ReadPreference *ReadPref `json:"readPreference"`
	// 可用于在配置 Client 时提供身份验证选项。
	Auth *Credential `json:"auth"`
}

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
type Credential struct {
	AuthMechanism string `json:"authMechanism"`
	AuthSource    string `json:"authSource"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	PasswordSet   bool   `json:"passwordSet"`
}

// ReadPref 决定哪些服务器适合读取操作。
type ReadPref struct {
// MaxStaleness 表示服务器被视为可选的最大过时时间。
// 该特性从版本 3.4 开始支持。
	MaxStalenessMS int64 `json:"maxStalenessMS"`
// 表示用户对读取操作的偏好。
// 默认为 PrimaryMode
	Mode readpref.Mode `json:"mode"`
}

// QmgoClient 指定了操作 MongoDB 的实例
type QmgoClient struct {
	*Collection
	*Database
	*Client
}

// Open creates client instance according to config
// QmgoClient can operates all qmgo.client 、qmgo.database and qmgo.collection
func X创建(上下文 context.Context, 配置 *Config, 可选选项 ...options.ClientOptions) (Qmgo客户端 *QmgoClient, 错误 error) {
	client, 错误 := X创建客户端(上下文, 配置, 可选选项...)
	if 错误 != nil {
		fmt.Println("new client fail", 错误)
		return
	}

	db := client.X设置数据库(配置.Database)
	coll := db.X取集合(配置.Coll)

	Qmgo客户端 = &QmgoClient{
		Client:     client,
		Database:   db,
		Collection: coll,
	}

	return
}

// Client 创建用于连接 MongoDB 的客户端
type Client struct {
	client *mongo.Client
	conf   Config

	registry *bsoncodec.Registry
}

// NewClient 创建 Qmgo MongoDB 客户端
func X创建客户端(上下文 context.Context, 配置 *Config, 可选选项 ...options.ClientOptions) (客户端 *Client, 错误 error) {
	opt, 错误 := newConnectOpts(配置, 可选选项...)
	if 错误 != nil {
		return nil, 错误
	}
	client, 错误 := client(上下文, opt)
	if 错误 != nil {
		fmt.Println("new client fail", 错误)
		return
	}
	客户端 = &Client{
		client:   client,
		conf:     *配置,
		registry: opt.Registry,
	}
	return
}

// client 创建与 MongoDB 的连接
func client(ctx context.Context, opt *officialOpts.ClientOptions) (client *mongo.Client, err error) {
	client, err = mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 连接超时时间的一半
	pCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err = client.Ping(pCtx, readpref.Primary()); err != nil {
		fmt.Println(err)
		return
	}
	return
}

// newConnectOpts creates client options from conf
// Qmgo will follow this way official mongodb driver do：
// - the configuration in uri takes precedence over the configuration in the setter
// - Check the validity of the configuration in the uri, while the configuration in the setter is basically not checked
func newConnectOpts(conf *Config, o ...options.ClientOptions) (*officialOpts.ClientOptions, error) {
	option := officialOpts.Client()
	for _, apply := range o {
		option = officialOpts.MergeClientOptions(apply.ClientOptions)
	}
	if conf.ConnectTimeoutMS != nil {
		timeoutDur := time.Duration(*conf.ConnectTimeoutMS) * time.Millisecond
		option.SetConnectTimeout(timeoutDur)

	}
	if conf.SocketTimeoutMS != nil {
		timeoutDur := time.Duration(*conf.SocketTimeoutMS) * time.Millisecond
		option.SetSocketTimeout(timeoutDur)
	} else {
		option.SetSocketTimeout(300 * time.Second)
	}
	if conf.MaxPoolSize != nil {
		option.SetMaxPoolSize(*conf.MaxPoolSize)
	}
	if conf.MinPoolSize != nil {
		option.SetMinPoolSize(*conf.MinPoolSize)
	}
	if conf.ReadPreference != nil {
		readPreference, err := newReadPref(*conf.ReadPreference)
		if err != nil {
			return nil, err
		}
		option.SetReadPreference(readPreference)
	}
	if conf.Auth != nil {
		auth, err := newAuth(*conf.Auth)
		if err != nil {
			return nil, err
		}
		option.SetAuth(auth)
	}
	option.ApplyURI(conf.Uri)

	return option, nil
}

// newAuth 从 conf.Auth 创建 options.Credential
func newAuth(auth Credential) (credential officialOpts.Credential, err error) {
	if auth.AuthMechanism != "" {
		credential.AuthMechanism = auth.AuthMechanism
	}
	if auth.AuthSource != "" {
		credential.AuthSource = auth.AuthSource
	}
	if auth.Username != "" {
		// 验证并处理用户名。
		if strings.Contains(auth.Username, "/") {
			err = ErrNotSupportedUsername
			return
		}
		credential.Username, err = url.QueryUnescape(auth.Username)
		if err != nil {
			err = ErrNotSupportedUsername
			return
		}
	}
	credential.PasswordSet = auth.PasswordSet
	if auth.Password != "" {
		if strings.Contains(auth.Password, ":") {
			err = ErrNotSupportedPassword
			return
		}
		if strings.Contains(auth.Password, "/") {
			err = ErrNotSupportedPassword
			return
		}
		credential.Password, err = url.QueryUnescape(auth.Password)
		if err != nil {
			err = ErrNotSupportedPassword
			return
		}
		credential.Password = auth.Password
	}
	return
}

// newReadPref 从配置中创建 readpref.ReadPref
func newReadPref(pref ReadPref) (*readpref.ReadPref, error) {
	readPrefOpts := make([]readpref.Option, 0, 1)
	if pref.MaxStalenessMS != 0 {
		readPrefOpts = append(readPrefOpts, readpref.WithMaxStaleness(time.Duration(pref.MaxStalenessMS)*time.Millisecond))
	}
	mode := readpref.PrimaryMode
	if pref.Mode != 0 {
		mode = pref.Mode
	}
	readPreference, err := readpref.New(mode, readPrefOpts...)
	return readPreference, err
}

// Close关闭与此Client关联的拓扑结构的所有套接字连接。
func (c *Client) X关闭(上下文 context.Context) error {
	err := c.client.Disconnect(上下文)
	return err
}

// Ping：确认连接是否存活
func (c *Client) X是否存活(超时时长 int64) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(超时时长)*time.Second)
	defer cancel()

	if err = c.client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// 创建与数据库的连接
func (c *Client) X设置数据库(数据库名称 string, 可选选项 ...*options.DatabaseOptions) *Database {
	opts := make([]*officialOpts.DatabaseOptions, 0, len(可选选项))
	for _, o := range 可选选项 {
		opts = append(opts, o.DatabaseOptions)
	}
	databaseOpts := officialOpts.MergeDatabaseOptions(opts...)
	return &Database{database: c.client.Database(数据库名称, databaseOpts), registry: c.registry}
}

// Session 在客户端创建一个会话
// 注意，在操作完成后关闭会话
func (c *Client) X创建Session(可选选项 ...*options.SessionOptions) (*Session, error) {
	sessionOpts := officialOpts.Session()
	if len(可选选项) > 0 && 可选选项[0].SessionOptions != nil {
		sessionOpts = 可选选项[0].SessionOptions
	}
	s, err := c.client.StartSession(sessionOpts)
	return &Session{session: s}, err
}

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
func (c *Client) X事务(上下文 context.Context, 回调函数 func(sessCtx context.Context) (interface{}, error), 可选选项 ...*options.TransactionOptions) (interface{}, error) {
	if !c.transactionAllowed() {
		return nil, ErrTransactionNotSupported
	}
	s, err := c.X创建Session()
	if err != nil {
		return nil, err
	}
	defer s.EndSession(上下文)
	return s.StartTransaction(上下文, 回调函数, 可选选项...)
}

// ServerVersion 获取MongoDB服务器的版本，如4.4.0
func (c *Client) X取版本号() string {
	var buildInfo bson.Raw
	err := c.client.Database("admin").RunCommand(
		context.Background(),
		bson.D{{"buildInfo", 1}},
	).Decode(&buildInfo)
	if err != nil {
		fmt.Println("run command err", err)
		return ""
	}
	v, err := buildInfo.LookupErr("version")
	if err != nil {
		fmt.Println("look up err", err)
		return ""
	}
	return v.StringValue()
}

// transactionAllowed 检查交易是否被允许
func (c *Client) transactionAllowed() bool {
	vr, err := X比较版本号("4.0", c.X取版本号())
	if err != nil {
		return false
	}
	if vr > 0 {
		fmt.Println("transaction is not supported because mongo server version is below 4.0")
		return false
	}
// TODO 未知为何需要在topology()函数中执行`cli, err := Open(ctx, &c.conf)`来获取topo，
// 在查明原因之前，我们仅在单元测试（UT）中使用此函数
//topo, err := c.topology()
//如果topo是description.Single类型 {
//	 fmt.Println("由于Mongo服务器拓扑为单节点，因此不支持事务")
//	 返回false
//}
	return true
}
