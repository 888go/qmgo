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
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

[func (q *Query) Collation(collation *options.Collation) QueryI {]
collation=规则
ff=设置排序规则

[func (q *Query) NoCursorTimeout(n bool) QueryI {]
n=是否不超时
ff=设置不超时

[func (q *Query) BatchSize(n int64) QueryI {]
n=数量
ff=设置批量处理数量

[func (q *Query) Sort(fields ...string) QueryI {]
fields=排序字段
ff=排序

[func (q *Query) SetArrayFilters(filter *options.ArrayFilters) QueryI {]
filter=过滤条件
ff=设置切片过滤

[func (q *Query) Select(projection interface{}) QueryI {]
projection=字段Map
ff=字段

[func (q *Query) Skip(n int64) QueryI {]
n=跳过数量
ff=跳过

[func (q *Query) Hint(hint interface{}) QueryI {]
hint=索引字段
ff=指定索引字段

[func (q *Query) Limit(n int64) QueryI {]
n=数量
ff=设置最大返回数

[func (q *Query) One(result interface{}) error {]
result=结果指针
ff=取一条

[func (q *Query) All(result interface{}) error {]
result=结果指针
ff=取全部

[func (q *Query) Count() (n int64, err error) {]
err=错误
n=数量
ff=取数量

[func (q *Query) EstimatedCount() (n int64, err error) {]
err=错误
n=数量
ff=取预估数量

[func (q *Query) Distinct(key string, result interface{}) error {]
result=切片指针
key=字段名
ff=去重

[func (q *Query) Cursor() CursorI {]
ff=取结果集

[func (q *Query) Apply(change Change, result interface{}) error {]
ff=执行命令
