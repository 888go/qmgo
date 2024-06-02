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

package qmgo

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// 初始MongoDB实例的配置 md5:09dcbab1d00adb46
// [提示]
//
//	type 配置 struct {
//	    连接URI        string `json:"uri"`
//	    数据库名       string `json:"database"`
//	    集合名         string `json:"coll"`
//	    连接超时毫秒   *int64 `json:"connectTimeoutMS"`
//	    最大连接池大小 *uint64 `json:"maxPoolSize"`
//	    最小连接池大小 *uint64 `json:"minPoolSize"`
//	    套接字超时毫秒 *int64 `json:"socketTimeoutMS"`
//	    读取首选项     *读取偏好 `json:"readPreference"`
//	    认证信息       *凭证   `json:"auth"`
//	}
//
//	type 读取偏好 struct {
//	    // ... (ReadPref结构体内的字段也需要翻译)
//	}
//
//	type 凭证 struct {
//	    // ... (Credential结构体内的字段也需要翻译)
//	}
//
// [结束]
type Config struct { //hm:配置  cz:type Config
	// URI 示例：[mongodb://][user:pass@]主机1[:端口1][,主机2[:端口2],...][/数据库][?选项]
	// URI 参考：https://docs.mongodb.com/manual/reference/connection-string/
	//
	// 这段注释解释了一个MongoDB连接字符串的格式，包括可选的部分如用户名、密码、多个服务器地址、数据库名以及可选的连接选项。URI以`mongodb://`开头，后面可以包含认证信息、主机列表、数据库路径和查询参数。链接：提供了官方文档的参考。
	// md5:038c28929efbdde0
	Uri      string `json:"uri"`      //qm:连接URI  cz:Uri string `json:"uri"`
	Database string `json:"database"` //qm:数据库名  cz:Database string `json:"database"`
	Coll     string `json:"coll"`     //qm:集合名  cz:Coll string `json:"coll"`
	// ConnectTimeoutMS 指定了建立到服务器连接时使用的超时时间，以毫秒为单位。
	// 如果设置为 0，则不会使用超时。
	// 默认值为 30 秒。
	// md5:bdc6b23048c25478
	ConnectTimeoutMS *int64 `json:"connectTimeoutMS"` //qm:连接超时毫秒  cz:ConnectTimeoutMS *int64 `json:"connectTimeoutMS"`
	// MaxPoolSize 指定驱动程序连接池到每个服务器的最大连接数。
	// 如果设置为 0，则将其设置为 math.MaxInt64，
	// 默认值为 100。
	// md5:6840c2846a8fad6e
	MaxPoolSize *uint64 `json:"maxPoolSize"` //qm:最大连接池大小  cz:MaxPoolSize *uint64 `json:"maxPoolSize"`
	// MinPoolSize 指定了驱动程序到每个服务器的连接池中允许的最小连接数。如果此值不为零，将为每个服务器的连接池在后台维护，以确保其大小不低于最小值。这也可以通过 "minPoolSize" URI 选项（如 "minPoolSize=100"）进行设置。默认值为 0。
	// md5:9df8b44a6800236b
	MinPoolSize *uint64 `json:"minPoolSize"` //qm:最小连接池大小  cz:MinPoolSize *uint64 `json:"minPoolSize"`
	// SocketTimeoutMS 指定了驱动程序在返回网络错误之前，等待套接字读写操作返回的最长时间（以毫秒为单位）。如果此值为0，则表示不使用超时，套接字操作可能无限期阻塞。默认值为300,000毫秒。
	// md5:1e1ccf1f35a18417
	SocketTimeoutMS *int64 `json:"socketTimeoutMS"` //qm:套接字超时毫秒  cz:SocketTimeoutMS *int64 `json:"socketTimeoutMS"`
	// ReadPreference 确定哪些服务器适合进行读取操作。默认为 PrimaryMode。
	// md5:6ca3a191c28443b8
	ReadPreference *ReadPref `json:"readPreference"` //qm:读取偏好  zz:ReadPreference \*[a-zA-Z0-9_\u4e00-\u9fa5 ]+`json:"readPreference"`
	// 可用于在配置客户端时提供身份验证选项。 md5:99c19d7fabc83d2d
	Auth *Credential `json:"auth"` //qm:身份凭证  zz:Auth \*[a-zA-Z0-9_\u4e00-\u9fa5 ]+ `json:"auth"`
}

// Credential can be used to provide authentication options when configuring a Client.
//
// "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509", and "MONGODB-AWS". This can also be set through the "authMechanism"
// URI option. (e.g. "authMechanism=PLAIN"). For more information, see
// GSSAPI, and PLAIN and "admin" for all other mechanisms. This can also be set through the "authSource" URI option
// (e.g. "authSource=otherDb").
//
// the first @ character. For example, a URI for user "user", password "pwd", and host "localhost:27017" would be
// "mongodb://user:pwd@localhost:27017". This is optional for X509 authentication and will be extracted from the
// client certificate if not specified.
//
// authentication.
//
// false if no password is specified, indicating that the password should be taken from the context of the running
// process. For other mechanisms, this field is ignored.
// [提示]
//
//	type 身份凭证 struct {
//	    认证机制     string `json:"authMechanism"`
//	    认证源       string `json:"authSource"`
//	    用户名       string `json:"username"`
//	    密码         string `json:"password"`
//	    密码已设置   bool   `json:"passwordSet"`
//	}
//
// [结束]
type Credential struct { //hm:身份凭证  cz:type Credential
	AuthMechanism string `json:"authMechanism"` //qm:认证机制  cz:AuthMechanism string `json:"authMechanism"`
	AuthSource    string `json:"authSource"`    //qm:认证源  cz:AuthSource string `json:"authSource"`
	Username      string `json:"username"`      //qm:用户名  cz:Username string `json:"username"`
	Password      string `json:"password"`      //qm:密码  cz:Password string `json:"password"`
	PasswordSet   bool   `json:"passwordSet"`
}

// ReadPref确定哪些服务器适合进行读取操作。 md5:d5ae507a40965ac9
// [提示]
//
//	type 读取偏好 struct {
//	    最大延迟毫秒 int64 `json:"maxStalenessMS"`
//	    模式 readpref.模式 `json:"mode"`
//	}
//
// [结束]
type ReadPref struct { //hm:读取偏好  cz:type ReadPref
	// MaxStaleness是允许服务器被认为适合选择的最大时间。从版本3.4开始支持。
	// md5:01c3097a5d9a368b
	MaxStalenessMS int64 `json:"maxStalenessMS"` //qm:最大延迟毫秒  cz:MaxStalenessMS int64 `json:"maxStalenessMS"`
	// 表示用户在读取操作上的偏好。
	// 默认为PrimaryMode。
	// md5:85d94814e6ac8eca
	Mode readpref.Mode `json:"mode"`
}

// QmgoClient 指定操作MongoDB的实例 md5:ef9044b4ab2af757
// [提示]
//
//	type 七牛Mongo客户端 struct {
//	    集合 *集合操作
//	    数据库 *数据库操作
//	    客户端 *客户端连接
//	}
//
// [结束]
type QmgoClient struct { //hm:Mongo客户端  cz:type QmgoClient
	*Collection
	*Database
	*Client
}

// Open 根据配置创建客户端实例
// QmgoClient 可以操作所有 qmgo.client、qmgo.database 和 qmgo.collection
// md5:bc872aaa93cf801a
// [提示:] func 连接(ctx 上下文, 配置 *配置, 选项 ...options.ClientOptions) (客户端 *QmgoClient, 错误 error) {}
// ff:连接
// ctx:上下文
// conf:配置
// o:可选选项
// cli:Qmgo客户端
// err:错误
func Open(ctx context.Context, conf *Config, o ...options.ClientOptions) (cli *QmgoClient, err error) {
	client, err := NewClient(ctx, conf, o...)
	if err != nil {
		fmt.Println("new client fail", err)
		return
	}

	db := client.Database(conf.Database)
	coll := db.Collection(conf.Coll)

	cli = &QmgoClient{
		Client:     client,
		Database:   db,
		Collection: coll,
	}

	return
}

// Client 创建一个到Mongo的客户端 md5:3527d3de272044c3
// [提示]
//
//	type 客户端 struct {
//	    连接 *mongo.Client
//	    配置  Config
//	    注册表 *bsoncodec.Registry
//	}
//
// [结束]
type Client struct { //hm:客户端  cz:type Client
	client *mongo.Client
	conf   Config

	registry *bsoncodec.Registry
}

// NewClient 创建 Qmgo MongoDB 客户端 md5:64c9dc0f30edc1ac
// [提示:] func 新建客户端(ctx 上下文, 配置 *配置, 选项 ...选项客户端) (客户端 *Client, 错误 error) {}
// ff:创建客户端
// ctx:上下文
// conf:配置
// o:可选选项
// cli:客户端
// err:错误
func NewClient(ctx context.Context, conf *Config, o ...options.ClientOptions) (cli *Client, err error) {
	opt, err := newConnectOpts(conf, o...)
	if err != nil {
		return nil, err
	}
	client, err := client(ctx, opt)
	if err != nil {
		fmt.Println("new client fail", err)
		return
	}
	cli = &Client{
		client:   client,
		conf:     *conf,
		registry: opt.Registry,
	}
	return
}

// 客户端创建到MongoDB的连接 md5:5ed46d6e6a970651
func client(ctx context.Context, opt *officialOpts.ClientOptions) (client *mongo.Client, err error) {
	client, err = mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 默认连接超时时间的一半 md5:e544afad71f167e7
	pCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err = client.Ping(pCtx, readpref.Primary()); err != nil {
		fmt.Println(err)
		return
	}
	return
}

// newConnectOpts 从 conf 创建客户端选项
// Qmgo 将遵循官方 MongoDB 驱动程序的做法：
// - URI 中的配置优先于 setter 中的配置
// - 检查 URI 中配置的有效性，而 setter 中的配置基本不进行检查
// md5:e686e2f8bec69b3b
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

// newAuth 从conf.Auth创建options.Credential选项 md5:88ce8258f4551f1c
func newAuth(auth Credential) (credential officialOpts.Credential, err error) {
	if auth.AuthMechanism != "" {
		credential.AuthMechanism = auth.AuthMechanism
	}
	if auth.AuthSource != "" {
		credential.AuthSource = auth.AuthSource
	}
	if auth.Username != "" {
		// 验证和处理用户名。 md5:3c89ddb7c004c9d6
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

// newReadPref 根据配置创建 readpref.ReadPref md5:1c0e9080aed7b202
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

// Close 关闭到此客户端引用的拓扑结构相关的套接字。 md5:a2c78aacda5cd470
// [提示:] func (c *客户端) 关闭(ctx 上下文.Context) 错误 {}
// ff:关闭连接
// ctx:上下文
func (c *Client) Close(ctx context.Context) error {
	err := c.client.Disconnect(ctx)
	return err
}

// Ping确认连接是否还活着 md5:1b88dbe0bbaa6726
// [提示:] func (c *客户端) 心跳检查超时时间(timeout 超时时间秒) error {}
// ff:是否存活
// timeout:超时时长
func (c *Client) Ping(timeout int64) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if err = c.client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// Database 创建到数据库的连接 md5:1aa03639d9adcf41
// [提示:] func (c *客户端) 数据库(name string, options ...*选项.DatabaseOptions) *数据库 {}
// ff:设置数据库
// name:数据库名称
// options:可选选项
func (c *Client) Database(name string, options ...*options.DatabaseOptions) *Database {
	opts := make([]*officialOpts.DatabaseOptions, 0, len(options))
	for _, o := range options {
		opts = append(opts, o.DatabaseOptions)
	}
	databaseOpts := officialOpts.MergeDatabaseOptions(opts...)
	return &Database{database: c.client.Database(name, databaseOpts), registry: c.registry}
}

// Session：在客户端创建一个会话
// 注意，操作完成后要关闭会话
// md5:a25c6035ffabaf48
// [提示:] func (c *客户端) 会话(opt ...*选项.SessionOptions) (*会话, 错误) {}
// ff:创建Session事务
// opt:可选选项
func (c *Client) Session(opt ...*options.SessionOptions) (*Session, error) {
	sessionOpts := officialOpts.Session()
	if len(opt) > 0 && opt[0].SessionOptions != nil {
		sessionOpts = opt[0].SessionOptions
	}
	s, err := c.client.StartSession(sessionOpts)
	return &Session{session: s}, err
}

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
// [提示:] func (c *客户端) 执行事务(ctx 上下文.Context, 回调函数 func(会话上下文 context.Context) (结果 interface{}
// ff:事务
// ctx:上下文
// callback:回调函数
// sessCtx:事务上下文
// opts:可选选项
func (c *Client) DoTransaction(ctx context.Context, callback func(sessCtx context.Context) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	if !c.transactionAllowed() {
		return nil, ErrTransactionNotSupported
	}
	s, err := c.Session()
	if err != nil {
		return nil, err
	}
	defer s.EndSession(ctx)
	return s.StartTransaction(ctx, callback, opts...)
}

// ServerVersion 获取MongoDB服务器的版本，如4.4.0 md5:85f19b2205255d3a
// [提示:] func (c *Client) 服务器版本() string {}
// ff:取版本号
func (c *Client) ServerVersion() string {
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

// transactionAllowed 检查交易是否被允许 md5:d9e86f3ad9610912
func (c *Client) transactionAllowed() bool {
	vr, err := CompareVersions("4.0", c.ServerVersion())
	if err != nil {
		return false
	}
	if vr > 0 {
		fmt.Println("transaction is not supported because mongo server version is below 4.0")
		return false
	}
	// TODO：不知道为什么在`topology()`函数中需要通过`cli, err := Open(ctx, &c.conf)`来获取topo，在弄清楚原因之前，我们只在这个UT（单元测试）中使用这个函数
	//topo, err := c.topology() // 从config对象获取topology信息
	//如果topo是description.Single（单点模式）：
	//    打印 "transaction is not supported because mongo server topology is single"
	//    返回false
	// md5:4d3e4bc17382c028
	return true
}
