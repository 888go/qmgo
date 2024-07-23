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

[Eq = "$eq"]
qm=条件等于
cz=Eq #等号# "$eq"

[Gt = "$gt"]
qm=条件大于
cz=Gt #等号# "$gt"

[Gte = "$gte"]
qm=条件大于等于
cz=Gte #等号# "$gte"

[In = "$in"]
qm=条件包含
cz=In #等号# "$in"

[Lt = "$lt"]
qm=条件小于
cz=Lt #等号# "$lt"

[Lte = "$lte"]
qm=条件小于等于
cz=Lte #等号# "$lte"

[Ne = "$ne"]
qm=条件不等于
cz=Ne #等号# "$ne"

[Nin = "$nin"]
qm=条件不包含
cz=Nin #等号# "$nin"

[And = "$and"]
qm=条件且
cz=And #等号# "$and"

[Not = "$not"]
qm=条件非
cz=Not #等号# "$not"

[Nor = "$nor"]
qm=条件或非
cz=Nor #等号# "$nor"

[Or = "$or"]
qm=条件或
cz=Or #等号# "$or"

[Exists = "$exists"]
qm=条件字段存在
cz=Exists #等号# "$exists"

[Type = "$type"]
qm=条件类型
cz=Type #等号# "$type"

[Expr = "$expr"]
qm=条件表达式
cz=Expr #等号# "$expr"

[JsonSchema = "$jsonSchema"]
qm=Json效验
cz=JsonSchema #等号# "$jsonSchema"

[Mod = "$mod"]
qm=取模
cz=Mod #等号# "$mod"

[Regex = "$regex"]
qm=条件正则
cz=Regex #等号# "$regex"

[Text = "$text"]
qm=条件全文搜索
cz=Text #等号# "$text"

[Where = "$where"]
qm=条件Js
cz=Where #等号# "$where"

[All = "$all"]
qm=数组全部
cz=All #等号# "$all"

[ElemMatch = "$elemMatch"]
qm=数组匹配条件
cz=ElemMatch #等号# "$elemMatch"

[Size = "$size"]
qm=数组数量
cz=Size #等号# "$size"
