提示:
ff= 方法,重命名方法名称
qm= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

[Update    interface{}]
qm=更新替换

[Replace   bool]
qm=是否替换

[Remove    bool]
qm=是否删除

[Upsert    bool]
qm=未找到是否插入

[ReturnNew bool]
qm=是否返回新文档

[Next(result interface{}) bool]
qm=下一个

[Close() error]
qm=关闭

[Err() error]
qm=取错误

[All(results interface{}) error]
qm=取全部
cf=2

[Collation(collation *options.Collation) QueryI]
qm=设置排序规则

[SetArrayFilters(*options.ArrayFilters) QueryI]
qm=设置数组过滤

[Sort(fields ...string) QueryI]
qm=排序

[Select(selector interface{}) QueryI]
qm=字段

[Skip(n int64) QueryI]
qm=跳过

[BatchSize(n int64) QueryI]
qm=设置批量处理数量

[NoCursorTimeout(n bool) QueryI]
qm=设置不超时

[Limit(n int64) QueryI]
qm=设置最大返回数

[One(result interface{}) error]
qm=取一条
cf=2

[All(result interface{}) error]
qm=取全部


[Count() (n int64, err error)]
qm=取数量

[EstimatedCount() (n int64, err error)]
qm=取预估数量

[Distinct(key string, result interface{}) error]
qm=去重

[Cursor() CursorI]
qm=取结果集
cf=2

[Apply(change Change, result interface{}) error]
qm=执行命令

[Hint(hint interface{}) QueryI]
qm=指定索引字段

[Iter() CursorI]
qm=Iter弃用
