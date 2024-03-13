提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

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
