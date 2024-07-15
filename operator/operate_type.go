package mgo常量

type OpType string

const (
	X钩子_插入前  OpType = "beforeInsert"
	X钩子_插入后   OpType = "afterInsert"
	X钩子_更新前  OpType = "beforeUpdate"
	X钩子_更新后   OpType = "afterUpdate"
	X钩子_查询前   OpType = "beforeQuery"
	X钩子_查询后    OpType = "afterQuery"
	X钩子_删除前  OpType = "beforeRemove"
	X钩子_删除后   OpType = "afterRemove"
	X钩子_替换插入前  OpType = "beforeUpsert"
	X钩子_替换插入后   OpType = "afterUpsert"
	X钩子_替换前 OpType = "beforeReplace"
	X钩子_替换后  OpType = "afterReplace"
)
