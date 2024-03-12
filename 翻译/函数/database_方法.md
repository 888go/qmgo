提示:
ff= 是方法名称
sx= 是属性或者结构体重命名,默认会跳转到行首进行重命名.
    文档内如果有多个相同的,会一起重命名.
bm= 包名称
th= 替换文本

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
