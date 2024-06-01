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

	// 更新一个一个
	err := cli.UpdateOne(ctx, bson.M{"name": "d4"}, bson.M{"$set": bson.M{"age": 7}})
	fmt.Println(err)

	// 全部更新
	result, err := cli.UpdateAll(ctx, bson.M{"age": 6}, bson.M{"$set": bson.M{"age": 10}})
	fmt.Println(result.MatchedCount, err)
	fmt.Println(result.ModifiedCount, err)
}
