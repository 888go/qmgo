提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

[func (d *Database) Collection(name string, opts ...*options.CollectionOptions) *Collection {]
ff=取集合
opts=可选选项
name=名称

[func (d *Database) GetDatabaseName() string {]
ff=取数据库名称

[func (d *Database) DropDatabase(ctx context.Context) error {]
ff=删除数据库
ctx=上下文

[func (d *Database) RunCommand(ctx context.Context, runCommand interface{}, opts ...options.RunCommandOptions) *mongo.SingleResult {]
ff=执行命令
opts=可选选项
ctx=上下文

[func (db *Database) CreateCollection(ctx context.Context, name string, opts ...options.CreateCollectionOptions) error {]
ff=创建集合
opts=可选选项
name=集合名称
ctx=上下文
