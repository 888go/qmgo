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
	BeforeInsert  OpType = "beforeInsert"  //qm:插入前
	AfterInsert   OpType = "afterInsert"   //qm:插入后
	BeforeUpdate  OpType = "beforeUpdate"  //qm:更新前
	AfterUpdate   OpType = "afterUpdate"   //qm:更新后
	BeforeQuery   OpType = "beforeQuery"   //qm:查询前
	AfterQuery    OpType = "afterQuery"    //qm:查询后
	BeforeRemove  OpType = "beforeRemove"  //qm:删除前
	AfterRemove   OpType = "afterRemove"   //qm:删除后
	BeforeUpsert  OpType = "beforeUpsert"  //qm:更新或插入前
	AfterUpsert   OpType = "afterUpsert"   //qm:更新或插入后
	BeforeReplace OpType = "beforeReplace" //qm:替换前
	AfterReplace  OpType = "afterReplace"  //qm:替换后
)
