package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

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

	//查找一个文档
	one := UserInfo{}
	err := cli.Find(ctx, bson.M{"age": 10}).Select(bson.M{"age": 1}).One(&one)
	fmt.Println(one)
	fmt.Println(err)
}
