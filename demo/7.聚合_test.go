package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

/*
更多聚合查看 https://www.mongodb.com/zh-cn/docs/manual/reference/operator/aggregation/abs/
*/

import (
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
)

func Test_聚合分组(t *testing.T) {
	分组 := bson.D{
		{operator.Group,
			bson.D{
				{"_id", "$姓名"}, //分组, 以"姓名"分组, 其中"_id", 是固定表达式, 表示分组的键, 不可更改.
				{"重量合计", bson.D{{operator.Sum, "$重量"}}}, //并加一个列表示"重量合计"
				{"年龄平均", bson.D{{operator.Avg, "$年龄"}}}, //并加一个列表示"平均年龄"
			},
		},
	}
	var showsWithInfo []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{分组}).All(&showsWithInfo)
	fmt.Println(showsWithInfo)
}

func Test_聚合过滤条件(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/retrieve/#aggregation
	//条件,过滤条件,"姓名"字段="张三"
	过滤条件 := bson.D{{operator.Match, []bson.E{{"姓名", bson.D{{operator.Eq, "张三"}}}}}}

	//分组, 以"姓名"分组, 并加一个列表示"重量合计"
	//其中"_id", 是固定表达式, 表示分组的键, 不可更改.
	分组 := bson.D{
		{operator.Group,
			bson.D{
				{"_id", "$姓名"},
				{"重量合计", bson.D{{operator.Sum, "$重量"}}}, //如果不是求和, 只是要统计文档数量, 可以传入参数1, {operator.Sum, 1}
				{"年龄平均", bson.D{{operator.Avg, "$年龄"}}},
			},
		},
	}
	var 显示信息 []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{过滤条件, 分组}).All(&显示信息)
	fmt.Println(显示信息)
}

func Test_聚合文档数量(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/retrieve/#aggregation
	//条件,过滤条件,"姓名"字段="张三"
	过滤条件 := bson.D{{operator.Match, []bson.E{{"姓名", bson.D{{operator.Eq, "张三"}}}}}}

	//分组, 以"姓名"分组, 并加一个列表示"重量合计"
	//其中"_id", 是固定表达式, 表示分组的键, 不可更改.
	分组 := bson.D{
		{operator.Group,
			bson.D{
				{"_id", "$姓名"},
				{"文档数量", bson.D{{operator.Sum, 1}}}, //如果不是求和, 只是要统计文档数量, 可以传入参数1, {operator.Sum, 1}
				{"年龄平均", bson.D{{operator.Avg, "$年龄"}}},
			},
		},
	}
	var 显示信息 []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{过滤条件, 分组}).All(&显示信息)
	fmt.Println(显示信息)
}

func Test_聚合排序(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/sort/
	排序 := bson.D{
		{operator.Sort,
			bson.D{
				{"年龄", 1},  //按年龄升序
				{"重量", -1}, //按重量降序
			},
		},
	}
	var 显示信息 []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{排序}).All(&显示信息)
	fmt.Println(显示信息)
}

func Test_聚合限制返回数(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/limit/

	//限制返回数量3条.
	限制返回 := bson.D{{"$limit", 3}}

	var 显示信息 []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{限制返回}).All(&显示信息)
	fmt.Println(显示信息)
}

func Test_聚合字段(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/project/#aggregation
	//字段,只返回"姓名"和"年龄"字段
	字段 := bson.D{{"$project", bson.D{{"姓名", 1}, {"年龄", 1}}}}
	显示信息 := []bson.M{}
	cli.Aggregate(ctx, qmgo.Pipeline{字段}).All(&显示信息)
	fmt.Println(显示信息)

	//bson.D{{"姓名", 0}, {"年龄",0}}表示除了"姓名"和"年龄"以外的其他字段都显示
	字段 = bson.D{{"$project", bson.D{{"姓名", 0}, {"年龄", 0}}}}
	显示信息 = []bson.M{}
	cli.Aggregate(ctx, qmgo.Pipeline{字段}).All(&显示信息)
	fmt.Println(显示信息)
}

func Test_聚合关联(t *testing.T) {
	//https://www.bilibili.com/video/BV1i14y1H7gK?p=37&vd_source=5ea92576e0a34ead0bc129a4eeb1fa45
	//https://www.mongodb.com/zh-cn/docs/manual/reference/operator/aggregation/lookup/
	聚合关联初始化() //初始化数据库的价格集合
	聚合 := bson.D{
		{"$lookup", bson.M{
			"from":         "价格",   //要关联的集合名称
			"localField":   "购买产品", //当前集合的关联字段,可以是数组字段,也可以是字符串字段
			"foreignField": "产品名称", //关联集合的关联字段
			"as":           "关联集合", //显示字段名称
		}},
	}

	显示信息 := []bson.M{}
	cli.Aggregate(ctx, qmgo.Pipeline{聚合}).All(&显示信息)
	fmt.Println(显示信息)
}

func 聚合关联初始化() {
	//插入"价格"集合, 聚合关联查询用.
	cli_价格, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "数据库_demo", Coll: "价格"})
	文档数量, _ := cli_价格.Find(ctx, bson.D{}).EstimatedCount()
	if 文档数量 == 0 {
		userInfo := []X产品价格{
			{X产品名称: "西瓜", X产品价格: 10},
			{X产品名称: "老虎钳", X产品价格: 20},
			{X产品名称: "番茄", X产品价格: 30},
			{X产品名称: "草鱼", X产品价格: 40},
		}
		result, _ := cli_价格.InsertMany(ctx, userInfo)
		fmt.Println(result.InsertedIDs)
	}
}
