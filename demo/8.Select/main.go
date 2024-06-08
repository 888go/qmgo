package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

// [提示]
//type 用户信息 struct {
//     姓名   string `bson:"name"`
//     年龄    uint16 `bson:"age"`
//     体重   uint32 `bson:"weight"`
// }
// [结束]
type UserInfo struct {
	Name   string `bson:"名称"`
	Age    uint16 `bson:"年龄"`
	Weight uint32 `bson:"重量"`
}

var userInfo = UserInfo{
	Name:   "xm",
	Age:    7,
	Weight: 40,
}

func main() {
	ctx := context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "数据库_demo", Coll: "记账"})

	//关闭连接
	defer func() {
		if err := cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	//查找一个文档
	one := UserInfo{}
	cli.Find(ctx, bson.M{"年龄": 7}).Select(bson.M{"年龄": 7}).One(&one)
	fmt.Println(one)
}
