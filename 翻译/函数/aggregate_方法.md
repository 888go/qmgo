提示:
ff= 是方法名称
sx= 是属性或者结构体重命名,默认会跳转到行首进行重命名.
    文档内如果有多个相同的,会一起重命名.
bm= 包名称
th= 替换文本

[func (a *Aggregate) All(results interface{}) error {]
results=结果指针
ff=取全部

[func (a *Aggregate) One(result interface{}) error {]
ff=取一条
result=结果指针

[func (a *Aggregate) Iter() CursorI {]
ff=Iter弃用

[func (a *Aggregate) Cursor() CursorI {]
ff=取结果集
