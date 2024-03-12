提示:
ff= 是方法名称
hs= 是属性或者结构体重命名,默认会跳转到行首进行重命名.
    文档内如果有多个相同的,会一起重命名.
bm= 包名称
th= 替换文本

[func (q *Query) Collation(collation *options.Collation) QueryI {]
ff=设置排序规则
collation=规则

[func (q *Query) NoCursorTimeout(n bool) QueryI {]
ff=设置不超时
n=是否不超时

[func (q *Query) BatchSize(n int64) QueryI {]
ff=设置批量处理数量
n=数量

[func (q *Query) Sort(fields ...string) QueryI {]
ff=排序
fields=排序字段

[func (q *Query) SetArrayFilters(filter *options.ArrayFilters) QueryI {]
ff=设置数组过滤
filter=过滤条件

[func (q *Query) Select(projection interface{}) QueryI {]
ff=字段
projection=字段Map

[func (q *Query) Skip(n int64) QueryI {]
ff=跳过
n=跳过数量

[func (q *Query) Hint(hint interface{}) QueryI {]
ff=指定索引字段
hint=索引字段

[func (q *Query) Limit(n int64) QueryI {]
ff=设置最大返回数
n=数量

[func (q *Query) One(result interface{}) error {]
ff=取一条
result=结果指针

[func (q *Query) All(result interface{}) error {]
ff=取全部
result=结果指针

[func (q *Query) Count() (n int64, err error) {]
ff=取数量
err=错误
n=数量

[func (q *Query) EstimatedCount() (n int64, err error) {]
ff=取预估数量
err=错误
n=数量

[func (q *Query) Distinct(key string, result interface{}) error {]
ff=去重
result=数组指针
key=字段名

[func (q *Query) Cursor() CursorI {]
ff=取结果集

[func (q *Query) Apply(change Change, result interface{}) error {]
ff=执行命令
