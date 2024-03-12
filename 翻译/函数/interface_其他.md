提示:
ff= 是方法名称
sx= 是属性或者结构体重命名,默认会跳转到行首进行重命名.
    文档内如果有多个相同的,会一起重命名.
bm= 包名称
th= 替换文本

[Update    interface{}]
sx=更新替换

[Replace   bool]
sx=是否替换

[Remove    bool]
sx=是否删除

[Upsert    bool]
sx=未找到是否插入

[ReturnNew bool]
sx=是否返回新文档

[Next(result interface{}) bool]
sx=下一个

[Close() error]
sx=关闭

[Err() error]
sx=取错误

[All(results interface{}) error]
sx=取全部

[Collation(collation *options.Collation) QueryI]
sx=设置排序规则

[SetArrayFilters(*options.ArrayFilters) QueryI]
sx=设置数组过滤

[Sort(fields ...string) QueryI]
sx=排序

[Select(selector interface{}) QueryI]
sx=字段

[Skip(n int64) QueryI]
sx=跳过

[BatchSize(n int64) QueryI]
sx=设置批量处理数量

[NoCursorTimeout(n bool) QueryI]
sx=设置不超时

[Limit(n int64) QueryI]
sx=设置最大返回数

[One(result interface{}) error]
sx=取一条

[All(result interface{}) error]
sx=取全部

[Count() (n int64, err error)]
sx=取数量

[EstimatedCount() (n int64, err error)]
sx=取预估数量

[Distinct(key string, result interface{}) error]
sx=去重

[Cursor() CursorI]
sx=取结果集

[Apply(change Change, result interface{}) error]
sx=执行命令

[Hint(hint interface{}) QueryI]
sx=指定索引字段

[Iter() CursorI]
sx=Iter弃用
