package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
)

// [提示]
//type 用户信息 struct {
//     姓名   string `bson:"name"`
//     年龄    uint16 `bson:"age"`
//     体重   uint32 `bson:"weight"`
// }
// [结束]
type UserInfo struct {
	Name   string `bson:"name"`
	Age    uint16 `bson:"age"`
	Weight uint32 `bson:"weight"`
}

var userInfo = UserInfo{
	Name:   "xm",
	Age:    7,
	Weight: 40,
}

func main() {
	ctx := context.Background()
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "class", Coll: "user"})

	//关闭连接
	defer func() {
		if err := cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	//插入一条数据
	result, _ := cli.InsertOne(ctx, userInfo)
	fmt.Println(result.InsertedID)

}
