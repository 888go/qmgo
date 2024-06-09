package main

import (
	"fmt"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_测试(t *testing.T) {
	one := X记账{}
	err := cli.Find(ctx, bson.M{"名称": "d4"}).One(&one)
	fmt.Println(one)
	fmt.Println(err)

}
func Test_查找一个文档(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#literal-values
	var userInfo = X记账{
		X名称: "a1",
		X年龄: 7,
		X重量: 40,
	}
	one := X记账{}
	_ = cli.Find(ctx, bson.M{"名称": userInfo.X名称}).One(&one)
	fmt.Println(one)
}

func Test_查找所有_排序和限制(t *testing.T) {
	batch := []X记账{}
	//用bson.M{}作为条件, 等同
	cli.Find(ctx, bson.M{"年龄": 6}).Sort("重量").Limit(7).All(&batch)
	fmt.Println(batch)

	batch = []X记账{}
	//用bson.D{}作为条件, 等同
	cli.Find(ctx, bson.D{{"年龄", 6}}).Sort("重量").Limit(7).All(&batch)
	fmt.Println(batch)
}

func Test_对比查询操作符(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#comparison
	//$eq匹配等于指定值的值。
	//$gt匹配大于指定值的值。
	//$gte匹配大于等于指定值的值。
	//$in匹配数组中指定的任何值。
	//$lt匹配小于指定值的值。
	//$lte匹配小于等于指定值的值。
	//$ne匹配所有不等于指定值的值。
	//$nin不匹配数组中指定的任何值。

	//用MongoDB自带比较符号,年龄大于6
	batch := []X记账{}
	cli.Find(ctx, bson.M{"年龄": bson.M{"$gt": 6}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄大于6-->", batch)

	//用MongoDB自带比较符号,年龄大于6且小于8
	batch = []X记账{}
	cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6}}, {"年龄", bson.M{"$lt": 8}}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄大于6且小于8-->", batch)

	//用MongoDB自带比较符号,年龄大于6且小于8
	batch = []X记账{}
	cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6, "$lt": 8}}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄大于6且小于8-->", batch)

	//用MongoDB自带比较符号,年龄等于6,且名称等于b2
	batch = []X记账{}
	cli.Find(ctx, bson.D{{"年龄", bson.M{"$eq": 6}}, {"名称", bson.M{"$eq": "b2"}}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄等于6,且名称等于b2-->", batch)

	//用qmgo包装后的常量,年龄大于6且小于8
	batch = []X记账{}
	cli.Find(ctx, bson.D{{"年龄", bson.M{operator.Gt: 6, operator.Lt: 8}}}).All(&batch)
	fmt.Println("用qmgo包装后的常量,年龄大于6且小于8-->", batch)
}

func Test_逻辑查询操作符(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#logical
	batch := []X记账{}
	//$and使用逻辑 AND 连接查询子句将返回与两个子句的条件匹配的所有文档。
	//$not反转查询表达式的效果，并返回与查询表达式不匹配的文档。
	//$nor使用逻辑 NOR 的联接查询子句会返回无法匹配这两个子句的所有文档。
	//$or使用逻辑 OR 连接多个查询子句会返回符合任一子句条件的所有文档。
	//用MongoDB自带比较符号,年龄等于6,且名称等于b2
	逻辑比较 := bson.D{
		{
			"$and",
			bson.A{
				bson.D{{"年龄", bson.M{"$eq": 6}}},
				bson.D{{"名称", bson.M{"$eq": "b2"}}},
			},
		},
	}

	cli.Find(ctx, 逻辑比较).All(&batch)
	fmt.Println("用MongoDB自带逻辑符号,年龄等于6,且名称等于b2-->", batch)
}

func Test_字段是否存在(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#element
	//备注,在go里面,似乎没办法插入一个空数据,所以,以下很难生效.

	//查找名称为空数据的文档
	batch := []X记账{}
	filter := bson.D{{"名称", bson.D{{"$exists", false}}}}
	cli.Find(ctx, filter).All(&batch)
	fmt.Println(batch)

	//以下是另外一个用途,名称必须存在,且不等于d4
	batch = []X记账{}
	cli.Find(ctx, bson.D{{"名称", bson.M{"$exists": true, "$ne": "d4"}}}).All(&batch)
	fmt.Println(batch)
}

func Test_求值(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#evaluation
	//$regex操作符, 支持正则表达式, 匹配"名称"以b2开头的.
	batch := []X记账{}
	cli.Find(ctx, bson.D{{"名称", bson.M{operator.Regex: "^b2.*"}}}).All(&batch)
	fmt.Println("$regex操作符-->", batch)

	//$expr操作符, 支持把查询字段作为变量来比较.
	//查出"重量"比"年龄"字段大的数据.
	batch = []X记账{}
	query := bson.D{
		{"$expr", bson.D{
			{"$gt", bson.A{"$重量", "$年龄"}},
		}},
	}
	cli.Find(ctx, query).All(&batch)
	fmt.Println("$expr操作符--->", batch)
}

func Test_取文档数量(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/count/#std-label-golang-estimated-count
	文档数量, _ := cli.Find(ctx, bson.D{{"名称", bson.M{operator.Regex: "^b.*"}}}).Count()
	fmt.Println(文档数量)
}

func Test_去重(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/count/#std-label-golang-estimated-count
	去重返回 := []string{}
	cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6}}}).Distinct("姓名", &去重返回)
	fmt.Println(去重返回)
}
func Test_排序(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/sort/
	//按照年龄升排序
	返回 := []X记账{}
	_ = cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6}}}).Sort("年龄").All(&返回)
	fmt.Println(返回)
	//按照年龄降排序
	返回 = []X记账{}
	_ = cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6}}}).Sort("-年龄").All(&返回)
	fmt.Println(返回)
}

func Test_字段(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/project/

	//bson.M{"姓名": 1} 表示只显示"姓名"字段
	返回 := []X记账{}
	_ = cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6}}}).Select(bson.M{"姓名": 1}).All(&返回)
	fmt.Println(返回)

	//bson.M{"姓名": 0} 表示除了"姓名"以外的其他字段都显示
	返回 = []X记账{}
	_ = cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6}}}).Select(bson.M{"姓名": 0}).All(&返回)
	fmt.Println(返回)
}

func Test_分页(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/sort/
	//用qmgo自带的分页
	第几页 := 2
	每页 := 3
	返回 := []X记账{}
	_ = cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 0}}}).Limit(int64(每页)).Skip(int64((每页 * (第几页 - 1)))).All(&返回)
	fmt.Println(返回)

	//用额外追加的分页功能操作
	返回 = []X记账{}
	_ = cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 0}}}).X分页(第几页, 每页).All(&返回)
	fmt.Println(返回)

	//用额外追加的分页功能取分页数
	fmt.Println("总分页数:", cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 0}}}).X取分页数(3))
}

func Test_全文搜索(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/text/

	//注意,需要提前创建索引,否则报错: (IndexNotFound) text index required for $text query
	//"$search"是MongoDB中的一个特殊操作符，用于在$text查询中执行全文搜索
	返回 := []X记账{}
	全文搜索条件 := bson.D{{"$text", bson.D{{"$search", "张三"}}}}
	err := cli.Find(ctx, 全文搜索条件).Select(bson.M{"姓名": 1}).All(&返回)
	fmt.Println(err)
	fmt.Println(返回)
}
