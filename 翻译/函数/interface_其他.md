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

提示:
ff= 方法,重命名方法名称
qm= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

[Update interface{}]
qm=更新替换
zz=Update +interface{} +/*

[Replace bool]
qm=是否替换
cz=Replace bool

[Remove bool]
qm=是否删除
cz=Remove bool

[Upsert bool]
qm=是否未找到时插入
cz=Upsert bool

[ReturnNew bool]
qm=是否返回新文档
cz=ReturnNew bool

[Next(result interface{}) bool]
qm=下一个
cz=Next(result interface{}) bool

[Close() error]
qm=关闭
cz=Close() error

[Err() error]
qm=取错误
cz=Err() error

[Collation(collation *options.Collation) QueryI]
qm=设置排序规则
cz=Collation(collation *options.Collation) QueryI

[SetArrayFilters(*options.ArrayFilters) QueryI]
qm=设置切片过滤
cz=SetArrayFilters(*options.ArrayFilters) QueryI

[Sort(fields ...string) QueryI]
qm=排序
cz=Sort(fields ...string) QueryI

[Select(selector interface{}) QueryI]
qm=字段
cz=Select(selector interface{}) QueryI

[Skip(n int64) QueryI]
qm=跳过
cz=Skip(n int64) QueryI

[BatchSize(n int64) QueryI]
qm=设置批量处理数量
cz=BatchSize(n int64) QueryI

[NoCursorTimeout(n bool) QueryI]
qm=设置不超时
cz=NoCursorTimeout(n bool) QueryI

[Limit(n int64) QueryI]
qm=设置最大返回数
cz=Limit(n int64) QueryI

[All(result interface{}) error]
qm=取全部
cz=All(result interface{}) error

[Count() (n int64, err error)]
qm=取数量
cz=Count() (n int64, err error)

[EstimatedCount() (n int64, err error)]
qm=取预估数量
cz=EstimatedCount() (n int64, err error)

[Distinct(key string, result interface{}) error]
qm=去重
cz=Distinct(key string, result interface{}) error

[Apply(change Change, result interface{}) error]
qm=执行命令
cz=Apply(change Change, result interface{}) error

[Hint(hint interface{}) QueryI]
qm=指定索引字段
cz=Hint(hint interface{}) QueryI

[All(results interface{}) error]
qm=取全部
cz=All(results interface{}) error
cf=2

[One(result interface{}) error]
qm=取一条
cz=One(result interface{}) error
cf=2

[Iter() CursorI]
qm=Iter弃用
cz=Iter() CursorI

[Cursor() CursorI]
qm=取结果集
cz=Cursor() CursorI
cf=2
