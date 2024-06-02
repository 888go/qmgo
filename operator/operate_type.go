package operator

type OpType string

// [提示]
// const (
//
//	插入前 OpType = "beforeInsert"
//	插入后   OpType = "afterInsert"
//	更新前  OpType = "beforeUpdate"
//	更新后   OpType = "afterUpdate"
//	查询前   OpType = "beforeQuery"
//	查询后    OpType = "afterQuery"
//	删除前  OpType = "beforeRemove"
//	删除后   OpType = "afterRemove"
//	Upsert前  OpType = "beforeUpsert"
//	Upsert后   OpType = "afterUpsert"
//	替换前  OpType = "beforeReplace"
//	替换后   OpType = "afterReplace"
//
// )
// [结束]
const (
	BeforeInsert  OpType = "beforeInsert"  //qm:常量_钩子_插入前  cz:BeforeInsert OpType = "beforeInsert"
	AfterInsert   OpType = "afterInsert"   //qm:常量_钩子_插入后  cz:AfterInsert OpType = "afterInsert"
	BeforeUpdate  OpType = "beforeUpdate"  //qm:常量_钩子_更新前  cz:BeforeUpdate OpType = "beforeUpdate"
	AfterUpdate   OpType = "afterUpdate"   //qm:常量_钩子_更新后  cz:AfterUpdate OpType = "afterUpdate"
	BeforeQuery   OpType = "beforeQuery"   //qm:常量_钩子_查询前  cz:BeforeQuery OpType = "beforeQuery"
	AfterQuery    OpType = "afterQuery"    //qm:常量_钩子_查询后  cz:AfterQuery OpType = "afterQuery"
	BeforeRemove  OpType = "beforeRemove"  //qm:常量_钩子_删除前  cz:BeforeRemove OpType = "beforeRemove"
	AfterRemove   OpType = "afterRemove"   //qm:常量_钩子_删除后  cz:AfterRemove OpType = "afterRemove"
	BeforeUpsert  OpType = "beforeUpsert"  //qm:常量_钩子_更新插入前  cz:BeforeUpsert OpType = "beforeUpsert"
	AfterUpsert   OpType = "afterUpsert"   //qm:常量_钩子_更新插入后  cz:AfterUpsert OpType = "afterUpsert"
	BeforeReplace OpType = "beforeReplace" //qm:常量_钩子_替换前  cz:BeforeReplace OpType = "beforeReplace"
	AfterReplace  OpType = "afterReplace"  //qm:常量_钩子_替换后  cz:AfterReplace OpType = "afterReplace"
)
