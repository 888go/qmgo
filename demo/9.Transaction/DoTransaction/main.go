package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

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
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "class", Coll: "user"})

	//关闭连接
	defer func() {
		if err := cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	callback := func(sessCtx context.Context) (interface{}, error) {
		// 重要提示：确保在整个事务的每个操作中都使用了sessCtx
		if _, err := cli.InsertOne(sessCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.InsertOne(sessCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		return nil, nil
	}
	result, err := cli.DoTransaction(ctx, callback)
	fmt.Println(result, err)
}
