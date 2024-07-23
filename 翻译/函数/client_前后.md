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

[Uri string `json:"uri"`]
qm=连接URI
cz=Uri string `json:"uri"`

[Database string `json:"database"`]
qm=数据库名
cz=Database string `json:"database"`

[Coll string `json:"coll"`]
qm=集合名
cz=Coll string `json:"coll"`

[ConnectTimeoutMS *int64 `json:"connectTimeoutMS"`]
qm=连接超时毫秒
cz=ConnectTimeoutMS *int64 `json:"connectTimeoutMS"`

[MaxPoolSize *uint64 `json:"maxPoolSize"`]
qm=最大连接池大小
cz=MaxPoolSize *uint64 `json:"maxPoolSize"`

[MinPoolSize *uint64 `json:"minPoolSize"`]
qm=最小连接池大小
cz=MinPoolSize *uint64 `json:"minPoolSize"`

[SocketTimeoutMS *int64 `json:"socketTimeoutMS"`]
qm=套接字超时毫秒
cz=SocketTimeoutMS *int64 `json:"socketTimeoutMS"`

[ReadPreference *ReadPref `json:"readPreference"`]
qm=读取偏好
zz=ReadPreference \*.*`json:"readPreference"`

[Auth *Credential `json:"auth"`]
qm=身份凭证
zz=Auth \*.+ `json:"auth"`

[AuthMechanism string `json:"authMechanism"`]
qm=认证机制
cz=AuthMechanism string `json:"authMechanism"`

[AuthSource string `json:"authSource"`]
qm=认证源
cz=AuthSource string `json:"authSource"`

[Username string `json:"username"`]
qm=用户名
cz=Username string `json:"username"`

[Password string `json:"password"`]
qm=密码
cz=Password string `json:"password"`

[MaxStalenessMS int64 `json:"maxStalenessMS"`]
qm=最大延迟毫秒
cz=MaxStalenessMS int64 `json:"maxStalenessMS"`
