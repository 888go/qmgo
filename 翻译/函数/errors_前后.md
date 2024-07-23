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

[ErrQueryNotSlicePointer = errors.New("result argument must be a pointer to a slice")]
qm=错误_结果参数_必须切片指针
cz=ErrQueryNotSlicePointer #等号# errors.New

[ErrQueryNotSliceType = errors.New("result argument must be a slice address")]
qm=错误_结果参数_必须切片地址
cz=ErrQueryNotSliceType #等号# errors.New

[ErrQueryResultTypeInconsistent = errors.New("result type is not equal mongodb value type")]
qm=错误_查询结果_类型不一致
cz=ErrQueryResultTypeInconsistent #等号# errors.New

[ErrQueryResultValCanNotChange = errors.New("the value of result can not be changed")]
qm=错误_结果值不能更改
cz=ErrQueryResultValCanNotChange #等号# errors.New

[ErrNoSuchDocuments = mongo.ErrNoDocuments]
qm=错误_未找到文档
cz=ErrNoSuchDocuments #等号#

[ErrTransactionRetry = errors.New("retry transaction")]
qm=错误_事务_重试
cz=ErrTransactionRetry #等号# errors.New

[ErrTransactionNotSupported = errors.New("transaction not supported")]
qm=错误_事务_不支持
cz=ErrTransactionNotSupported #等号# errors.New

[ErrNotSupportedUsername = errors.New("username not supported")]
qm=错误_不支持用户名
cz=ErrNotSupportedUsername #等号# errors.New

[ErrNotSupportedPassword = errors.New("password not supported")]
qm=错误_不支持密码
cz=ErrNotSupportedPassword #等号# errors.New

[ErrNotValidSliceToInsert = errors.New("must be valid slice to insert")]
qm=错误_插入_无效切片
cz=ErrNotValidSliceToInsert #等号# errors.New

[ErrReplacementContainUpdateOperators = errors.New("replacement document cannot contain keys beginning with '$'")]
qm=错误_替换_文档含更新操作符
cz=ErrReplacementContainUpdateOperators #等号# errors.New
