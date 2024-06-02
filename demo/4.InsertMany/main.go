package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
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

	// 多重插入
	var userInfos = []UserInfo{
		{Name: "a1", Age: 6, Weight: 20},
		{Name: "b2", Age: 6, Weight: 25},
		{Name: "c3", Age: 6, Weight: 30},
		{Name: "d4", Age: 6, Weight: 35},
		{Name: "a1", Age: 7, Weight: 40},
		{Name: "a1", Age: 8, Weight: 45},
	}
	result, err := cli.Collection.InsertMany(ctx, userInfos)
	fmt.Println(result.InsertedIDs)
	fmt.Println(err)
}
