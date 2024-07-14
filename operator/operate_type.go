package operator

type OpType string

const (
	BeforeInsert  OpType = "beforeInsert"  //qm:钩子_插入前  cz:BeforeInsert OpType = "beforeInsert"
	AfterInsert   OpType = "afterInsert"   //qm:钩子_插入后  cz:AfterInsert OpType = "afterInsert"
	BeforeUpdate  OpType = "beforeUpdate"  //qm:钩子_更新前  cz:BeforeUpdate OpType = "beforeUpdate"
	AfterUpdate   OpType = "afterUpdate"   //qm:钩子_更新后  cz:AfterUpdate OpType = "afterUpdate"
	BeforeQuery   OpType = "beforeQuery"   //qm:钩子_查询前  cz:BeforeQuery OpType = "beforeQuery"
	AfterQuery    OpType = "afterQuery"    //qm:钩子_查询后  cz:AfterQuery OpType = "afterQuery"
	BeforeRemove  OpType = "beforeRemove"  //qm:钩子_删除前  cz:BeforeRemove OpType = "beforeRemove"
	AfterRemove   OpType = "afterRemove"   //qm:钩子_删除后  cz:AfterRemove OpType = "afterRemove"
	BeforeUpsert  OpType = "beforeUpsert"  //qm:钩子_替换插入前  cz:BeforeUpsert OpType = "beforeUpsert"
	AfterUpsert   OpType = "afterUpsert"   //qm:钩子_替换插入后  cz:AfterUpsert OpType = "afterUpsert"
	BeforeReplace OpType = "beforeReplace" //qm:钩子_替换前  cz:BeforeReplace OpType = "beforeReplace"
	AfterReplace  OpType = "afterReplace"  //qm:钩子_替换后  cz:AfterReplace OpType = "afterReplace"
)
