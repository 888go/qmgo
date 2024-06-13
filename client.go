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

// X配置 for initial mongodb instance
type X配置 struct {
	// URI example: [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	// URI Reference: https://docs.mongodb.com/manual/reference/connection-string/
	X连接URI      string `json:"uri"`
	X数据库名 string `json:"database"`
	X集合名     string `json:"coll"`
	// X连接超时毫秒 specifies a timeout that is used for creating connections to the server.
	//	If set to 0, no timeout will be used.
	//	The default is 30 seconds.
	X连接超时毫秒 *int64 `json:"connectTimeoutMS"`
	// X最大连接池大小 specifies that maximum number of connections allowed in the driver's connection pool to each server.
	// If this is 0, it will be set to math.MaxInt64,
	// The default is 100.
	X最大连接池大小 *uint64 `json:"maxPoolSize"`
	// X最小连接池大小 specifies the minimum number of connections allowed in the driver's connection pool to each server. If
	// this is non-zero, each server's pool will be maintained in the background to ensure that the size does not fall below
	// the minimum. This can also be set through the "minPoolSize" URI option (e.g. "minPoolSize=100"). The default is 0.
	X最小连接池大小 *uint64 `json:"minPoolSize"`
	// X套接字超时毫秒 specifies how long the driver will wait for a socket read or write to return before returning a
	// network error. If this is 0 meaning no timeout is used and socket operations can block indefinitely.
	// The default is 300,000 ms.
	X身份凭证 *int64 `json:"socketTimeoutMS"`
	// ReadPreference determines which servers are considered suitable for read operations.
	// default is PrimaryMode
	ReadPreference *X读取偏好 `json:"readPreference"`
	// can be used to provide authentication options when configuring a Client.
	Auth *X身份凭证 `json:"auth"`
}

// X身份凭证 can be used to provide authentication options when configuring a Client.
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
type X身份凭证 struct {
	X认证机制 string `json:"authMechanism"`
	X认证源    string `json:"authSource"`
	X用户名      string `json:"username"`
	X密码      string `json:"password"`
	PasswordSet   bool   `json:"passwordSet"`
}

// X读取偏好 determines which servers are considered suitable for read operations.
type X读取偏好 struct {
	// MaxStaleness is the maximum amount of time to allow a server to be considered eligible for selection.
	// Supported from version 3.4.
	X最大延迟毫秒 int64 `json:"maxStalenessMS"`
	// indicates the user's preference on reads.
	// PrimaryMode as default
	Mode readpref.Mode `json:"mode"`
}

// XMongo客户端 specifies the instance to operate mongoDB
type XMongo客户端 struct {
	*X文档集合
	*X数据库
	*X客户端
}

// X连接 creates client instance according to config
// QmgoClient can operates all qmgo.client 、qmgo.database and qmgo.collection
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

// X客户端 creates client to mongo
type X客户端 struct {
	client *mongo.Client
	conf   X配置

	registry *bsoncodec.Registry
}

// X创建客户端 creates Qmgo MongoDB client
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

// client creates connection to MongoDB
func client(ctx context.Context, opt *officialOpts.ClientOptions) (client *mongo.Client, err error) {
	client, err = mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	// half of default connect timeout
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
func newConnectOpts(conf *X配置, o ...options.ClientOptions) (*officialOpts.ClientOptions, error) {
	option := officialOpts.Client()
	for _, apply := range o {
		option = officialOpts.MergeClientOptions(apply.ClientOptions)
	}
	if conf.X连接超时毫秒 != nil {
		timeoutDur := time.Duration(*conf.X连接超时毫秒) * time.Millisecond
		option.SetConnectTimeout(timeoutDur)

	}
	if conf.X身份凭证 != nil {
		timeoutDur := time.Duration(*conf.X身份凭证) * time.Millisecond
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
	option.ApplyURI(conf.X连接URI)

	return option, nil
}

// newAuth create options.Credential from conf.Auth
func newAuth(auth X身份凭证) (credential officialOpts.Credential, err error) {
	if auth.X认证机制 != "" {
		credential.AuthMechanism = auth.X认证机制
	}
	if auth.X认证源 != "" {
		credential.AuthSource = auth.X认证源
	}
	if auth.X用户名 != "" {
		// Validate and process the username.
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

// newReadPref create readpref.ReadPref from config
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

// X关闭连接 closes sockets to the topology referenced by this Client.
func (c *X客户端) X关闭连接(上下文 context.Context) error {
	err := c.client.Disconnect(上下文)
	return err
}

// X是否存活 confirm connection is alive
func (c *X客户端) X是否存活(超时时长 int64) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(超时时长)*time.Second)
	defer cancel()

	if err = c.client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// X设置数据库 create connection to database
func (c *X客户端) X设置数据库(数据库名称 string, 可选选项 ...*options.DatabaseOptions) *X数据库 {
	opts := make([]*officialOpts.DatabaseOptions, 0, len(可选选项))
	for _, o := range 可选选项 {
		opts = append(opts, o.DatabaseOptions)
	}
	databaseOpts := officialOpts.MergeDatabaseOptions(opts...)
	return &X数据库{database: c.client.Database(数据库名称, databaseOpts), registry: c.registry}
}

// X创建Session事务 create one session on client
// Watch out, close session after operation done
func (c *X客户端) X创建Session事务(可选选项 ...*options.SessionOptions) (*XSession事务, error) {
	sessionOpts := officialOpts.Session()
	if len(可选选项) > 0 && 可选选项[0].SessionOptions != nil {
		sessionOpts = 可选选项[0].SessionOptions
	}
	s, err := c.client.StartSession(sessionOpts)
	return &XSession事务{session: s}, err
}

// X事务 do whole transaction in one function
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

// X取版本号 get the version of mongoDB server, like 4.4.0
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

// transactionAllowed check if transaction is allowed
func (c *X客户端) transactionAllowed() bool {
	vr, err := X比较版本号("4.0", c.X取版本号())
	if err != nil {
		return false
	}
	if vr > 0 {
		fmt.Println("transaction is not supported because mongo server version is below 4.0")
		return false
	}
	// TODO dont know why need to do `cli, err := Open(ctx, &c.conf)` in topology() to get topo,
	// Before figure it out, we only use this function in UT
	//topo, err := c.topology()
	//if topo == description.Single {
	//	fmt.Println("transaction is not supported because mongo server topology is single")
	//	return false
	//}
	return true
}
