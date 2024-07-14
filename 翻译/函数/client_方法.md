# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

 

[func Open(ctx context.Context, conf *Config, o ...options.ClientOptions) (cli *QmgoClient, err error) {]
ff=连接
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
ff=关闭连接
ctx=上下文

[func (c *Client) Ping(timeout int64) error {]
ff=是否存活
timeout=超时时长

[func (c *Client) Database(name string, options ...*options.DatabaseOptions) *Database {]
ff=设置数据库
options=可选选项
name=数据库名称

[func (c *Client) Session(opt ...*options.SessionOptions) (*Session, error) {]
ff=创建Session事务
opt=可选选项

[func (c *Client) DoTransaction(ctx context.Context, callback func(sessCtx context.Context) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {]
ff=事务
opts=可选选项
callback=回调函数
ctx=上下文
sessCtx=事务上下文

[func (c *Client) ServerVersion() string {]
ff=取版本号
