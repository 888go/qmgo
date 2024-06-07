package main

import (
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_聚合_分组(t *testing.T) {
	//分组, 以"姓名"分组, 并加一个列表示"重量合计"
	//其中"_id", 是固定表达式, 表示分组的键, 不可更改.
	分组 := bson.D{
		{operator.Group,
			bson.D{{"_id", "$姓名"},
				{"重量合计", bson.D{{operator.Sum, "$重量"}}},
				{"年龄平均", bson.D{{operator.Avg, "$年龄"}}},
			},
		},
	}
	var showsWithInfo []bson.M
	err := cli.Aggregate(ctx, qmgo.Pipeline{分组}).All(&showsWithInfo)
	fmt.Println(err)
	fmt.Println(showsWithInfo)
}

func Test_聚合_管道(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/read-operations/retrieve/#aggregation
	//条件,过滤条件,"姓名"字段="张三"
	过滤条件 := bson.D{{operator.Match, []bson.E{{"姓名", bson.D{{operator.Eq, "张三"}}}}}}

	//分组, 以"姓名"分组, 并加一个列表示"重量合计"
	//其中"_id", 是固定表达式, 表示分组的键, 不可更改.
	分组 := bson.D{
		{operator.Group,
			bson.D{{"_id", "$姓名"},
				{"重量合计", bson.D{{operator.Sum, "$重量"}}},
				{"年龄平均", bson.D{{operator.Avg, "$年龄"}}},
			},
		},
	}
	var 显示信息 []bson.M
	err := cli.Aggregate(ctx, qmgo.Pipeline{过滤条件, 分组}).All(&显示信息)
	fmt.Println(err)
	fmt.Println(显示信息)
}
