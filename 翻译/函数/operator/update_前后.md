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

[CurrentDate = "$currentDate"]
qm=更新为当前时间
cz=CurrentDate #等号# "$currentDate"

[Inc = "$inc"]
qm=更新数值递增
cz=Inc #等号# "$inc"

[Min = "$min"]
qm=更新最小
cz=Min #等号# "$min"

[Max = "$max"]
qm=更新最大
cz=Max #等号# "$max"

[Mul = "$mul"]
qm=更新相乘
cz=Mul #等号# "$mul"

[Rename = "$rename"]
qm=更新字段名
cz=Rename #等号# "$rename"

[Set = "$set"]
qm=更新值
cz=Set #等号# "$set"

[SetOnInsert = "$setOnInsert"]
qm=更新插入时
cz=SetOnInsert #等号# "$setOnInsert"

[Unset = "$unset"]
qm=聚合删除字段
cz=Unset #等号# "$unset"

[AddToSet = "$addToSet"]
qm=数组不存在追加
cz=AddToSet #等号# "$addToSet"

[Pop = "$pop"]
qm=数组删首尾
cz=Pop #等号# "$pop"

[Pull = "$pull"]
qm=数组删除
cz=Pull #等号# "$pull"

[Push = "$push"]
qm=数组追加
cz=Push #等号# "$push"

[PullAll = "$pullAll"]
qm=数组删除全部
cz=PullAll #等号# "$pullAll"

[Sort = "$sort"]
qm=聚合排序
cz=Sort #等号# "$sort"
