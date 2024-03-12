提示:
ff= 是方法名称
hs= 是属性或者结构体重命名,默认会跳转到行首进行重命名.
    文档内如果有多个相同的,会一起重命名.
bm= 包名称
th= 替换文本

[func (c *Collection) Bulk() *Bulk {]
ff=创建批量执行

[func (b *Bulk) SetOrdered(ordered bool) *Bulk {]
ff=设置有序执行
ordered=开启有序

[func (b *Bulk) InsertOne(doc interface{}) *Bulk {]
ff=插入
doc=待插入文档

[func (b *Bulk) Remove(filter interface{}) *Bulk {]
ff=删除一条
filter=删除条件

[func (b *Bulk) RemoveId(id interface{}) *Bulk {]
ff=删除并按ID
id=删除ID

[func (b *Bulk) RemoveAll(filter interface{}) *Bulk {]
ff=删除
filter=删除条件

[func (b *Bulk) Upsert(filter interface{}, replacement interface{}) *Bulk {]
ff=更新或插入
replacement=更新内容
filter=更新条件

[func (b *Bulk) UpsertOne(filter interface{}, update interface{}) *Bulk {]
ff=更新或插入一条
update=更新内容
filter=更新条件

[func (b *Bulk) UpsertId(id interface{}, replacement interface{}) *Bulk {]
ff=更新或插入并按ID
replacement=更新内容
id=更新ID

[func (b *Bulk) UpdateOne(filter interface{}, update interface{}) *Bulk {]
ff=更新一条
update=更新内容
filter=更新条件

[func (b *Bulk) UpdateId(id interface{}, update interface{}) *Bulk {]
ff=更新并按ID
update=更新内容
id=更新ID

[func (b *Bulk) UpdateAll(filter interface{}, update interface{}) *Bulk {]
ff=更新
update=更新内容
filter=更新条件

[func (b *Bulk) Run(ctx context.Context) (*BulkResult, error) {]
ff=执行
ctx=上下文
