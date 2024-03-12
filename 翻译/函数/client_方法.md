提示:
ff= 是方法名称
sx= 是属性或者结构体重命名,默认会跳转到行首进行重命名.
    文档内如果有多个相同的,会一起重命名.
bm= 包名称
th= 替换文本

[func Open(ctx context.Context, conf *Config, o ...options.ClientOptions) (cli *QmgoClient, err error) {]
ff=创建
err=错误
cli=Qmgo客户端
o=可选选项
conf=配置
ctx=上下文

[func NewClient(ctx context.Context, conf *Config, o ...options.ClientOptions) (cli *Client, err error) {]
ff=创建客户端
err=错误
cli=客户端
o=可选选项
conf=配置
ctx=上下文

[func (c *Client) Close(ctx context.Context) error {]
ff=关闭
ctx=上下文

[func (c *Client) Ping(timeout int64) error {]
ff=是否存活
timeout=超时时长

[func (c *Client) Database(name string, options ...*options.DatabaseOptions) *Database {]
ff=设置数据库
options=可选选项
name=数据库名称

[func (c *Client) Session(opt ...*options.SessionOptions) (*Session, error) {]
ff=创建Session
opt=可选选项

[func (c *Client) DoTransaction(ctx context.Context, callback func(sessCtx context.Context) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {]
ff=事务
opts=可选选项
callback=回调函数
ctx=上下文

[func (c *Client) ServerVersion() string {]
ff=取版本号
