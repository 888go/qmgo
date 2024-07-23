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
ff=设置切片过滤
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
result=切片指针
key=字段名

[func (q *Query) Cursor() CursorI {]
ff=取结果集

[func (q *Query) Apply(change Change, result interface{}) error {]
ff=执行命令
