提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

[func IsErrNoDocuments(err error) bool {]
ff=是否为无文档错误
err=错误

[func IsDup(err error) bool {]
ff=是否为重复键错误
err=错误
