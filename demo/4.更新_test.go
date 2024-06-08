package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func Test_更新一个(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	err := cli.UpdateOne(ctx, bson.M{"名称": "d4"}, bson.M{"$set": bson.M{"age": 7}})
	fmt.Println(err)
}
func Test_按id更新(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	Update := bson.M{
		"姓名": "Alice12",
		"年龄": 121,
	}
	//注意,id是十六进制, 不是文本型,所以需要转换.
	id, _ := primitive.ObjectIDFromHex("666484a8bc42fc4e667f82f7")
	result, _ := cli.UpsertId(ctx, id, Update) //注意,id是十六进制, 不是文本型
	fmt.Println(result.ModifiedCount)
}

func Test_更新全部(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	result, err := cli.UpdateAll(ctx, bson.M{"年龄": 6}, bson.M{"$set": bson.M{"年龄": 10}})
	fmt.Println(result.MatchedCount, err)
	fmt.Println(result.ModifiedCount, err)
}

func Test_更新字段名(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/#return-values
	// 全部更新
	result, err := cli.UpdateAll(ctx, bson.M{"年龄": 6}, bson.M{"$rename": bson.M{"年龄": "年龄2"}})
	fmt.Println(result.MatchedCount, err)
}

func Test_更新数值递增(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/#return-values
	// 把全部的年龄等于7的年龄加100
	result, err := cli.UpdateAll(ctx, bson.M{"年龄": 7}, bson.M{"$inc": bson.M{"年龄": 100}})
	fmt.Println(result.MatchedCount, err)
}
func Test_替换一条(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/#replace
	var userInfo = X记账{
		X姓名: "张三", X名称: "a1", X年龄: 99, X重量: 20, X够买产品: []string{"西瓜", "香蕉", "老虎钳"}, X够买时间: time.Now(), X支付方式: X支付方式{
			X支付方式: "支付宝", X联系方式: "177777777", X可选支付方式: []string{"支付宝", "微信", "银行卡"},
		}}

	err := cli.ReplaceOne(ctx, bson.M{"年龄": 8}, userInfo)
	fmt.Println(err)
}

///替换文档不能包含以'$'开头的键
//替换文档不能包含以'$'开头的键
//_, err = 集合.Upsert(ctx, bson.M{"名称": "菜鸟教程2"}, bson.M{"name": "菜鸟教程", "网址": "http://www.runoob.com"})
//_, err := 集合.UpdateAll(ctx, bson.M{"名称": "菜鸟教程22"}, bson.M{"$set": bson.M{"by": "菜鸟教程1111", "网址": "http://www.runoob.com"}})
