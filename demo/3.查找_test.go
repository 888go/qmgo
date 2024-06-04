package main

import (
	"fmt"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_查找一个文档(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#literal-values
	var userInfo = UserInfo{
		Name:   "a1",
		Age:    7,
		Weight: 40,
	}
	one := UserInfo{}
	_ = cli.Find(ctx, bson.M{"名称": userInfo.Name}).One(&one)
	fmt.Println(one)
}

func Test_查找所有_排序和限制(t *testing.T) {
	batch := []UserInfo{}
	//用bson.M{}作为条件, 等同
	cli.Find(ctx, bson.M{"年龄": 6}).Sort("重量").Limit(7).All(&batch)
	fmt.Println(batch)

	//用bson.D{}作为条件, 等同
	cli.Find(ctx, bson.D{{"年龄", 6}}).Sort("重量").Limit(7).All(&batch)
	fmt.Println(batch)
}

func Test_对比查询操作符(t *testing.T) {
	batch := []UserInfo{}
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
	cli.Find(ctx, bson.M{"年龄": bson.M{"$gt": 6}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄大于6-->", batch)

	//用MongoDB自带比较符号,年龄大于6且小于8
	cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6}}, {"年龄", bson.M{"$lt": 8}}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄大于6且小于8-->", batch)

	//用MongoDB自带比较符号,年龄大于6且小于8
	cli.Find(ctx, bson.D{{"年龄", bson.M{"$gt": 6, "$lt": 8}}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄大于6且小于8-->", batch)

	//用MongoDB自带比较符号,年龄等于6,且名称等于b2
	cli.Find(ctx, bson.D{{"年龄", bson.M{"$eq": 6}}, {"名称", bson.M{"$eq": "b2"}}}).All(&batch)
	fmt.Println("用MongoDB自带比较符号,年龄等于6,且名称等于b2-->", batch)

	//用qmgo包装后的常量,年龄大于6且小于8
	cli.Find(ctx, bson.D{{"年龄", bson.M{operator.Gt: 6, operator.Lt: 8}}}).All(&batch)
	fmt.Println("用qmgo包装后的常量,年龄大于6且小于8-->", batch)
}

func Test_逻辑查询操作符(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#logical
	batch := []UserInfo{}
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
	batch := []UserInfo{}
	//备注,在go里面,似乎没办法插入一个空数据,所以,以下很难生效.

	//查找名称为空数据的文档
	filter := bson.D{{"名称", bson.D{{"$exists", false}}}}
	cli.Find(ctx, filter).All(&batch)
	fmt.Println(batch)

	//以下是另外一个用途,名称必须存在,且不等于d4
	cli.Find(ctx, bson.D{{"名称", bson.M{"$exists": true, "$ne": "d4"}}}).All(&batch)
	fmt.Println(batch)
}

func Test_求值(t *testing.T) {
	//www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/#evaluation
	batch := []UserInfo{}
	//$regex操作符, 支持正则表达式, 匹配"名称"以b2开头的.
	cli.Find(ctx, bson.D{{"名称", bson.M{operator.Regex: "^b2.*"}}}).All(&batch)
	fmt.Println(batch)
}
