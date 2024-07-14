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



[BeforeInsert OpType = "beforeInsert"]
qm=钩子_插入前
cz=BeforeInsert OpType #等号# "beforeInsert"

[AfterInsert OpType = "afterInsert"]
qm=钩子_插入后
cz=AfterInsert OpType #等号# "afterInsert"

[BeforeUpdate OpType = "beforeUpdate"]
qm=钩子_更新前
cz=BeforeUpdate OpType #等号# "beforeUpdate"

[AfterUpdate OpType = "afterUpdate"]
qm=钩子_更新后
cz=AfterUpdate OpType #等号# "afterUpdate"

[BeforeQuery OpType = "beforeQuery"]
qm=钩子_查询前
cz=BeforeQuery OpType #等号# "beforeQuery"

[AfterQuery OpType = "afterQuery"]
qm=钩子_查询后
cz=AfterQuery OpType #等号# "afterQuery"

[BeforeRemove OpType = "beforeRemove"]
qm=钩子_删除前
cz=BeforeRemove OpType #等号# "beforeRemove"

[AfterRemove OpType = "afterRemove"]
qm=钩子_删除后
cz=AfterRemove OpType #等号# "afterRemove"

[BeforeUpsert OpType = "beforeUpsert"]
qm=钩子_替换插入前
cz=BeforeUpsert OpType #等号# "beforeUpsert"

[AfterUpsert OpType = "afterUpsert"]
qm=钩子_替换插入后
cz=AfterUpsert OpType #等号# "afterUpsert"

[BeforeReplace OpType = "beforeReplace"]
qm=钩子_替换前
cz=BeforeReplace OpType #等号# "beforeReplace"

[AfterReplace OpType = "afterReplace"]
qm=钩子_替换后
cz=AfterReplace OpType #等号# "afterReplace"
