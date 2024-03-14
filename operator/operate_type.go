package 操作符

type OpType string

const (
	X插入前  OpType = "beforeInsert"
	X插入后   OpType = "afterInsert"
	X更新前  OpType = "beforeUpdate"
	X更新后   OpType = "afterUpdate"
	X查询前   OpType = "beforeQuery"
	X查询后    OpType = "afterQuery"
	X删除前  OpType = "beforeRemove"
	X删除后   OpType = "afterRemove"
	X更新或插入前  OpType = "beforeUpsert"
	X更新或插入后   OpType = "afterUpsert"
	X替换前 OpType = "beforeReplace"
	X替换后  OpType = "afterReplace"
)
