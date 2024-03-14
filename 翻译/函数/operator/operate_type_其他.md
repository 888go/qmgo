提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
    但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"


[BeforeInsert  OpType = "beforeInsert"]
hs=插入前

[AfterInsert   OpType = "afterInsert"]
hs=插入后

[BeforeUpdate  OpType = "beforeUpdate"]
hs=更新前

[AfterUpdate   OpType = "afterUpdate"]
hs=更新后

[BeforeQuery   OpType = "beforeQuery"]
hs=查询前

[AfterQuery    OpType = "afterQuery"]
hs=查询后

[BeforeRemove  OpType = "beforeRemove"]
hs=删除前

[AfterRemove   OpType = "afterRemove"]
hs=删除后

[BeforeUpsert  OpType = "beforeUpsert"]
hs=更新或插入前

[AfterUpsert   OpType = "afterUpsert"]
hs=更新或插入后

[BeforeReplace OpType = "beforeReplace"]
hs=替换前

[AfterReplace  OpType = "afterReplace"]
hs=替换后
