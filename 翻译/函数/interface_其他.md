提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

[Update    interface{}]
hs=更新替换

[Replace   bool]
hs=是否替换

[Remove    bool]
hs=是否删除

[Upsert    bool]
hs=未找到是否插入

[ReturnNew bool]
hs=是否返回新文档

[Next(result interface{}) bool]
hs=下一个

[Close() error]
hs=关闭

[Err() error]
hs=取错误

[All(results interface{}) error]
hs=取全部
cf=2

[Collation(collation *options.Collation) QueryI]
hs=设置排序规则

[SetArrayFilters(*options.ArrayFilters) QueryI]
hs=设置数组过滤

[Sort(fields ...string) QueryI]
hs=排序

[Select(selector interface{}) QueryI]
hs=字段

[Skip(n int64) QueryI]
hs=跳过

[BatchSize(n int64) QueryI]
hs=设置批量处理数量

[NoCursorTimeout(n bool) QueryI]
hs=设置不超时

[Limit(n int64) QueryI]
hs=设置最大返回数

[One(result interface{}) error]
hs=取一条
cf=2

[All(result interface{}) error]
hs=取全部


[Count() (n int64, err error)]
hs=取数量

[EstimatedCount() (n int64, err error)]
hs=取预估数量

[Distinct(key string, result interface{}) error]
hs=去重

[Cursor() CursorI]
hs=取结果集
cf=2

[Apply(change Change, result interface{}) error]
hs=执行命令

[Hint(hint interface{}) QueryI]
hs=指定索引字段

[Iter() CursorI]
hs=Iter弃用
