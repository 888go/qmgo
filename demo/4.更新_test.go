package main

import (
	"fmt"
	"github.com/888go/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func Test_更新一个(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	err := cli.X更新一条(ctx, bson.M{"名称": "d4"}, bson.M{"$set": bson.M{"age": 7}})
	fmt.Println(err)
}

func Test_按id更新(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	Update := bson.M{
		"姓名": "Alice12",
		"年龄": 121,
	}
	//注意,id是十六进制, 不是文本型,所以需要转换.
	id, _ := primitive.ObjectIDFromHex("666484a8bc42fc4e667f82f7")
	result, _ := cli.X替换插入并按ID(ctx, id, Update) //注意,id是十六进制, 不是文本型
	fmt.Println(result.X修改数)
}

func Test_更新子文档(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	err := cli.X更新一条(ctx, bson.M{"名称": "d4"}, bson.M{"$set": bson.M{"支付方式.联系方式": "18059710086"}})
	fmt.Println(err)
}

func Test_更新文档数组(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/embedded-arrays/#first-array-element
	//更新要特别注意同原类型, 否则会导致后期读取失败.

	//".$[]" 代表数组每一个成员
	//把数组"购买产品" 每个成员都改成"精品西瓜",
	err := cli.X更新一条(ctx, bson.M{"名称": "d4"}, bson.M{"$set": bson.M{"购买产品.$[]": "精品西瓜1"}})
	fmt.Println(err)

	//".$" 代表数组第一个成员, 注意!!!,必须将数组字段包含在 查询 条件中
	//把数组"购买产品"第一个成员改成"精品西瓜",
	err = cli.X更新一条(ctx, bson.M{"购买产品": bson.M{"$eq": "西瓜"}}, bson.M{"$set": bson.M{"购买产品.$": "精品西瓜2"}})
	fmt.Println(err)
}

func Test_更新子文档数组(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	err := cli.X更新一条(ctx, bson.M{"名称": "d4"}, bson.M{"$set": bson.M{"支付方式.联系方式": "18059710086"}})
	fmt.Println(err)
}

func Test_更新全部(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	result, err := cli.X更新(ctx, bson.M{"年龄": 6}, bson.M{"$set": bson.M{"年龄": 10}})
	fmt.Println(result.X匹配数, err)
	fmt.Println(result.X修改数, err)
}

func Test_更新字段名(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/#return-values
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	result, err := cli.X更新(ctx, bson.M{"年龄": 6}, bson.M{"$rename": bson.M{"年龄": "年龄2"}})
	fmt.Println(result.X匹配数, err)
}

func Test_更新数值递增(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/#return-values
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	// 把全部的年龄等于7的年龄加100
	result, err := cli.X更新(ctx, bson.M{"年龄": 7}, bson.M{"$inc": bson.M{"年龄": 100}})
	fmt.Println(result.X匹配数, err)
}

func Test_更新数值相乘(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/manual/reference/operator/update/mul/#examples
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	//把姓名为张三的,年龄乘以100后更新
	result, err := cli.X更新(ctx, bson.M{"姓名": "张三"}, bson.M{mgo常量.X更新相乘: bson.M{"年龄": 100}})
	fmt.Println(result.X匹配数, err)
}

func Test_替换插入(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/manual/reference/operator/update/mul/#examples
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	//qmgo已经封装此方法, MongoDB自带的"$setOnInsert"没用了,
	//Upsert(filter interface{}, replacement interface{})

	//如果没查找到,则插入
	//如果查找到,则更新
	id, _ := primitive.ObjectIDFromHex("66648f742b1cca43c095f992") //id是十六进制, 不是文本型
	result, err := cli.X替换插入(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"年龄": 10}})
	fmt.Println(result.X匹配数, err)
}

func Test_替换一条(t *testing.T) {
	// https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/modify/#replace
	//更新要特别注意同原类型, 否则会导致后期读取失败.
	var userInfo = X记账{
		X姓名: "张三", X名称: "a1", X年龄: 99, X重量: 20, X购买产品: []string{"西瓜", "香蕉", "老虎钳"}, X购买时间: time.Now(), X支付方式: X支付方式{
			X支付方式: "支付宝", X联系方式: "177777777", X可选支付方式: []string{"支付宝", "微信", "银行卡"},
		}}

	err := cli.X替换一条(ctx, bson.M{"年龄": 8}, userInfo)
	fmt.Println(err)
}
