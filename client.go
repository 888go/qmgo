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

// 初始MongoDB实例的配置 md5:09dcbab1d00adb46
type X配置 struct {
// URI 示例：[mongodb://][user:pass@]主机1[:端口1][,主机2[:端口2],...][/数据库][?选项]
// URI 参考：https://docs.mongodb.com/manual/reference/connection-string/ 
// 
// 这段注释解释了一个MongoDB连接字符串的格式，包括可选的部分如用户名、密码、多个服务器地址、数据库名以及可选的连接选项。URI以`mongodb://`开头，后面可以包含认证信息、主机列表、数据库路径和查询参数。链接：提供了官方文档的参考。
// md5:038c28929efbdde0
	X连接URI      string `json:"uri"`
	X数据库名 string `json:"database"`
	X集合名     string `json:"coll"`
// X连接超时毫秒 指定了建立到服务器连接时使用的超时时间，以毫秒为单位。
// 如果设置为 0，则不会使用超时。
// 默认值为 30 秒。
// md5:bdc6b23048c25478
	X连接超时毫秒 *int64 `json:"connectTimeoutMS"`
// X最大连接池大小 指定驱动程序连接池到每个服务器的最大连接数。
// 如果设置为 0，则将其设置为 math.MaxInt64，
// 默认值为 100。
// md5:6840c2846a8fad6e
	X最大连接池大小 *uint64 `json:"maxPoolSize"`
// X最小连接池大小 指定了驱动程序到每个服务器的连接池中允许的最小连接数。如果此值不为零，将为每个服务器的连接池在后台维护，以确保其大小不低于最小值。这也可以通过 "minPoolSize" URI 选项（如 "minPoolSize=100"）进行设置。默认值为 0。
// md5:9df8b44a6800236b
	X最小连接池大小 *uint64 `json:"minPoolSize"`
// X套接字超时毫秒 指定了驱动程序在返回网络错误之前，等待套接字读写操作返回的最长时间（以毫秒为单位）。如果此值为0，则表示不使用超时，套接字操作可能无限期阻塞。默认值为300,000毫秒。
// md5:1e1ccf1f35a18417
	X套接字超时毫秒 *int64 `json:"socketTimeoutMS"`
// X读取偏好 确定哪些服务器适合进行读取操作。默认为 PrimaryMode。
// md5:6ca3a191c28443b8
	X读取偏好 *X读取偏好 `json:"readPreference"`
	// 可用于在配置客户端时提供身份验证选项。 md5:99c19d7fabc83d2d
	X身份凭证 *X身份凭证 `json:"auth"`
}

// X身份凭证 用于在配置客户端时提供认证选项。
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
type X身份凭证 struct {
	X认证机制 string `json:"authMechanism"`
	X认证源    string `json:"authSource"`
	X用户名      string `json:"username"`
	X密码      string `json:"password"`
	PasswordSet   bool   `json:"passwordSet"`
}

// X读取偏好确定哪些服务器适合进行读取操作。 md5:d5ae507a40965ac9
type X读取偏好 struct {
// MaxStaleness是允许服务器被认为适合选择的最大时间。从版本3.4开始支持。
// md5:01c3097a5d9a368b
	X最大延迟毫秒 int64 `json:"maxStalenessMS"`
// 表示用户在读取操作上的偏好。
// 默认为PrimaryMode。
// md5:85d94814e6ac8eca
	Mode readpref.Mode `json:"mode"`
}

// XMongo客户端 指定操作MongoDB的实例 md5:ef9044b4ab2af757
type XMongo客户端 struct {
	*X文档集合
	*X数据库
	*X客户端
}

// X连接 根据配置创建客户端实例
// QmgoClient 可以操作所有 qmgo.client、qmgo.database 和 qmgo.collection
// md5:bc872aaa93cf801a
func X连接(上下文 context.Context, 配置 *X配置, 可选选项 ...options.ClientOptions) (Qmgo客户端 *XMongo客户端, 错误 error) {
	client, 错误 := X创建客户端(上下文, 配置, 可选选项...)
	if 错误 != nil {
		fmt.Println("new client fail", 错误)
		return
	}

	db := client.X设置数据库(配置.X数据库名)
	coll := db.X取集合(配置.X集合名)

	Qmgo客户端 = &XMongo客户端{
		X客户端:     client,
		X数据库:   db,
		X文档集合: coll,
	}

	return
}

// X客户端 创建一个到Mongo的客户端 md5:3527d3de272044c3
type X客户端 struct {
	client *mongo.Client
	conf   X配置

	registry *bsoncodec.Registry
}

// X创建客户端 创建 Qmgo MongoDB 客户端 md5:64c9dc0f30edc1ac
func X创建客户端(上下文 context.Context, 配置 *X配置, 可选选项 ...options.ClientOptions) (客户端 *X客户端, 错误 error) {
	opt, 错误 := newConnectOpts(配置, 可选选项...)
	if 错误 != nil {
		return nil, 错误
	}
	client, 错误 := client(上下文, opt)
	if 错误 != nil {
		fmt.Println("new client fail", 错误)
		return
	}
	客户端 = &X客户端{
		client:   client,
		conf:     *配置,
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
func newConnectOpts(conf *X配置, o ...options.ClientOptions) (*officialOpts.ClientOptions, error) {
	option := officialOpts.Client()
	for _, apply := range o {
		option = officialOpts.MergeClientOptions(apply.ClientOptions)
	}
	if conf.X连接超时毫秒 != nil {
		timeoutDur := time.Duration(*conf.X连接超时毫秒) * time.Millisecond
		option.SetConnectTimeout(timeoutDur)

	}
	if conf.X套接字超时毫秒 != nil {
		timeoutDur := time.Duration(*conf.X套接字超时毫秒) * time.Millisecond
		option.SetSocketTimeout(timeoutDur)
	} else {
		option.SetSocketTimeout(300 * time.Second)
	}
	if conf.X最大连接池大小 != nil {
		option.SetMaxPoolSize(*conf.X最大连接池大小)
	}
	if conf.X最小连接池大小 != nil {
		option.SetMinPoolSize(*conf.X最小连接池大小)
	}
	if conf.X读取偏好 != nil {
		readPreference, err := newReadPref(*conf.X读取偏好)
		if err != nil {
			return nil, err
		}
		option.SetReadPreference(readPreference)
	}
	if conf.X身份凭证 != nil {
		auth, err := newAuth(*conf.X身份凭证)
		if err != nil {
			return nil, err
		}
		option.SetAuth(auth)
	}
	option.ApplyURI(conf.X连接URI)

	return option, nil
}

// newAuth 从conf.Auth创建options.Credential选项 md5:88ce8258f4551f1c
func newAuth(auth X身份凭证) (credential officialOpts.Credential, err error) {
	if auth.X认证机制 != "" {
		credential.AuthMechanism = auth.X认证机制
	}
	if auth.X认证源 != "" {
		credential.AuthSource = auth.X认证源
	}
	if auth.X用户名 != "" {
		// 验证和处理用户名。 md5:3c89ddb7c004c9d6
		if strings.Contains(auth.X用户名, "/") {
			err = X错误_不支持用户名
			return
		}
		credential.Username, err = url.QueryUnescape(auth.X用户名)
		if err != nil {
			err = X错误_不支持用户名
			return
		}
	}
	credential.PasswordSet = auth.PasswordSet
	if auth.X密码 != "" {
		if strings.Contains(auth.X密码, ":") {
			err = X错误_不支持密码
			return
		}
		if strings.Contains(auth.X密码, "/") {
			err = X错误_不支持密码
			return
		}
		credential.Password, err = url.QueryUnescape(auth.X密码)
		if err != nil {
			err = X错误_不支持密码
			return
		}
		credential.Password = auth.X密码
	}
	return
}

// newReadPref 根据配置创建 readpref.ReadPref md5:1c0e9080aed7b202
func newReadPref(pref X读取偏好) (*readpref.ReadPref, error) {
	readPrefOpts := make([]readpref.Option, 0, 1)
	if pref.X最大延迟毫秒 != 0 {
		readPrefOpts = append(readPrefOpts, readpref.WithMaxStaleness(time.Duration(pref.X最大延迟毫秒)*time.Millisecond))
	}
	mode := readpref.PrimaryMode
	if pref.Mode != 0 {
		mode = pref.Mode
	}
	readPreference, err := readpref.New(mode, readPrefOpts...)
	return readPreference, err
}

// X关闭连接 关闭到此客户端引用的拓扑结构相关的套接字。 md5:a2c78aacda5cd470
func (c *X客户端) X关闭连接(上下文 context.Context) error {
	err := c.client.Disconnect(上下文)
	return err
}

// X是否存活确认连接是否还活着 md5:1b88dbe0bbaa6726
func (c *X客户端) X是否存活(超时时长 int64) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(超时时长)*time.Second)
	defer cancel()

	if err = c.client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// X设置数据库 创建到数据库的连接 md5:1aa03639d9adcf41
func (c *X客户端) X设置数据库(数据库名称 string, 可选选项 ...*options.DatabaseOptions) *X数据库 {
	opts := make([]*officialOpts.DatabaseOptions, 0, len(可选选项))
	for _, o := range 可选选项 {
		opts = append(opts, o.DatabaseOptions)
	}
	databaseOpts := officialOpts.MergeDatabaseOptions(opts...)
	return &X数据库{database: c.client.Database(数据库名称, databaseOpts), registry: c.registry}
}

// X创建Session事务：在客户端创建一个会话
// 注意，操作完成后要关闭会话
// md5:a25c6035ffabaf48
func (c *X客户端) X创建Session事务(可选选项 ...*options.SessionOptions) (*XSession事务, error) {
	sessionOpts := officialOpts.Session()
	if len(可选选项) > 0 && 可选选项[0].SessionOptions != nil {
		sessionOpts = 可选选项[0].SessionOptions
	}
	s, err := c.client.StartSession(sessionOpts)
	return &XSession事务{session: s}, err
}

// X事务 在一个函数中执行整个事务
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
func (c *X客户端) X事务(上下文 context.Context, 回调函数 func(事务上下文 context.Context) (interface{}, error), 可选选项 ...*options.TransactionOptions) (interface{}, error) {
	if !c.transactionAllowed() {
		return nil, X错误_事务_不支持
	}
	s, err := c.X创建Session事务()
	if err != nil {
		return nil, err
	}
	defer s.X结束Session(上下文)
	return s.X开始事务(上下文, 回调函数, 可选选项...)
}

// X取版本号 获取MongoDB服务器的版本，如4.4.0 md5:85f19b2205255d3a
func (c *X客户端) X取版本号() string {
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
func (c *X客户端) transactionAllowed() bool {
	vr, err := X比较版本号("4.0", c.X取版本号())
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
