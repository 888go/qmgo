提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

[func (c *Collection) Find(ctx context.Context, filter interface{}, opts ...opts.FindOptions) QueryI {]
ff=查询
opts=可选选项
filter=查询条件
ctx=上下文

[func (c *Collection) InsertOne(ctx context.Context, doc interface{}, opts ...opts.InsertOneOptions) (result *InsertOneResult, err error) {]
ff=插入
err=错误
result=插入结果
opts=可选选项
doc=待插入文档
ctx=上下文

[func (c *Collection) InsertMany(ctx context.Context, docs interface{}, opts ...opts.InsertManyOptions) (result *InsertManyResult, err error) {]
ff=插入多个
err=错误
result=插入结果
opts=可选选项
docs=待插入文档
ctx=上下文

[func (c *Collection) Upsert(ctx context.Context, filter interface{}, replacement interface{}, opts ...opts.UpsertOptions) (result *UpdateResult, err error) {]
ff=更新或插入
err=错误
result=更新结果
opts=可选选项
replacement=更新内容
filter=更新条件
ctx=上下文

[func (c *Collection) UpsertId(ctx context.Context, id interface{}, replacement interface{}, opts ...opts.UpsertOptions) (result *UpdateResult, err error) {]
ff=更新或插入并按ID
err=错误
result=更新结果
opts=可选选项
replacement=更新内容
id=更新ID
ctx=上下文

[func (c *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...opts.UpdateOptions) (err error) {]
ff=更新一条
err=错误
opts=可选选项
update=更新内容
filter=更新条件
ctx=上下文

[func (c *Collection) UpdateId(ctx context.Context, id interface{}, update interface{}, opts ...opts.UpdateOptions) (err error) {]
ff=更新并按ID
err=错误
opts=可选选项
update=更新内容
id=更新ID
ctx=上下文

[func (c *Collection) UpdateAll(ctx context.Context, filter interface{}, update interface{}, opts ...opts.UpdateOptions) (result *UpdateResult, err error) {]
ff=更新
err=错误
result=更新结果
opts=可选选项
update=更新内容
filter=更新条件
ctx=上下文

[func (c *Collection) ReplaceOne(ctx context.Context, filter interface{}, doc interface{}, opts ...opts.ReplaceOptions) (err error) {]
ff=替换一条
err=错误
opts=可选选项
doc=替换内容
filter=替换条件
ctx=上下文

[func (c *Collection) Remove(ctx context.Context, filter interface{}, opts ...opts.RemoveOptions) (err error) {]
ff=删除一条
err=错误
opts=可选选项
filter=删除条件
ctx=上下文

[func (c *Collection) RemoveId(ctx context.Context, id interface{}, opts ...opts.RemoveOptions) (err error) {]
ff=删除并按ID
err=错误
opts=可选选项
id=删除ID
ctx=上下文

[func (c *Collection) RemoveAll(ctx context.Context, filter interface{}, opts ...opts.RemoveOptions) (result *DeleteResult, err error) {]
ff=删除
err=错误
result=删除结果
opts=可选选项
filter=删除条件
ctx=上下文

[func (c *Collection) Aggregate(ctx context.Context, pipeline interface{}, opts ...opts.AggregateOptions) AggregateI {]
ff=聚合
opts=可选选项
pipeline=聚合管道
ctx=上下文

[func (c *Collection) EnsureIndexes(ctx context.Context, uniques #左中括号##右中括号#string, indexes #左中括号##右中括号#string) (err error) {]
ff=EnsureIndexes弃用

[func (c *Collection) CreateIndexes(ctx context.Context, indexes #左中括号##右中括号#opts.IndexModel) (err error) {]
ff=索引多条
err=错误
indexes=索引s
ctx=上下文

[func (c *Collection) CreateOneIndex(ctx context.Context, index opts.IndexModel) error {]
ff=索引一条
index=索引
ctx=上下文

[func (c *Collection) DropAllIndexes(ctx context.Context) (err error) {]
ff=删除全部索引
err=错误
ctx=上下文

[func (c *Collection) DropIndex(ctx context.Context, indexes #左中括号##右中括号#string) error {]
ff=删除索引
indexes=索引s
ctx=上下文

[func (c *Collection) DropCollection(ctx context.Context) error {]
ff=删除集合
ctx=上下文

[func (c *Collection) CloneCollection() (*mongo.Collection, error) {]
ff=取副本

[func (c *Collection) GetCollectionName() string {]
ff=取集合名

[func (c *Collection) Watch(ctx context.Context, pipeline interface{}, opts ...*opts.ChangeStreamOptions) (*mongo.ChangeStream, error) {]
ff=取变更流
opts=可选选项
pipeline=管道
ctx=上下文
