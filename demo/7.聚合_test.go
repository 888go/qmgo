package main

import (
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_聚合_分组(t *testing.T) {
	分组 := bson.D{
		{operator.Group,
			bson.D{
				{"_id", "$姓名"},                              //分组, 以"姓名"分组, 其中"_id", 是固定表达式, 表示分组的键, 不可更改.
				{"重量合计", bson.D{{operator.Sum, "$重量"}}}, //并加一个列表示"重量合计"
				{"年龄平均", bson.D{{operator.Avg, "$年龄"}}}, //并加一个列表示"平均年龄"
			},
		},
	}
	var showsWithInfo []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{分组}).All(&showsWithInfo)
	fmt.Println(showsWithInfo)
}

func Test_聚合_过滤条件(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/retrieve/#aggregation
	//条件,过滤条件,"姓名"字段="张三"
	过滤条件 := bson.D{{operator.Match, []bson.E{{"姓名", bson.D{{operator.Eq, "张三"}}}}}}

	//分组, 以"姓名"分组, 并加一个列表示"重量合计"
	//其中"_id", 是固定表达式, 表示分组的键, 不可更改.
	分组 := bson.D{
		{operator.Group,
			bson.D{
				{"_id", "$姓名"},
				{"重量合计", bson.D{{operator.Sum, "$重量"}}},
				{"年龄平均", bson.D{{operator.Avg, "$年龄"}}},
			},
		},
	}
	var 显示信息 []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{过滤条件, 分组}).All(&显示信息)
	fmt.Println(显示信息)
}

func Test_聚合_排序(t *testing.T) {
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

func Test_聚合_限制返回数(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/limit/
	限制返回 := bson.D{{"$limit", 3}}

	var 显示信息 []bson.M
	cli.Aggregate(ctx, qmgo.Pipeline{限制返回}).All(&显示信息)
	fmt.Println(显示信息)
}

func Test_聚合_字段(t *testing.T) {
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
