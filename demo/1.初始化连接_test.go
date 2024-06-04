package main

import (
	"context"
	"github.com/qiniu/qmgo"
)

var cli *qmgo.QmgoClient
var ctx context.Context

type UserInfo struct {
	Id     string `bson:"_id,omitempty"` // 注意使用omitempty标签，使得在插入时如果未指定ID则MongoDB会自动生成
	Name   string `bson:"名称"`
	Age    uint16 `bson:"年龄"`
	Weight uint32 `bson:"重量"`
}

func init() {
	ctx = context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "学校数据库", Coll: "用户"})
}
