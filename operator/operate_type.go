package operator

type OpType string

const (
	BeforeInsert  OpType = "beforeInsert"  //hs:插入前
	AfterInsert   OpType = "afterInsert"   //hs:插入后
	BeforeUpdate  OpType = "beforeUpdate"  //hs:更新前
	AfterUpdate   OpType = "afterUpdate"   //hs:更新后
	BeforeQuery   OpType = "beforeQuery"   //hs:查询前
	AfterQuery    OpType = "afterQuery"    //hs:查询后
	BeforeRemove  OpType = "beforeRemove"  //hs:删除前
	AfterRemove   OpType = "afterRemove"   //hs:删除后
	BeforeUpsert  OpType = "beforeUpsert"  //hs:更新或插入前
	AfterUpsert   OpType = "afterUpsert"   //hs:更新或插入后
	BeforeReplace OpType = "beforeReplace" //hs:替换前
	AfterReplace  OpType = "afterReplace"  //hs:替换后
)
