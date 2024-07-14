package main

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var cli *qmgo.QmgoClient
var ctx context.Context

type X记账 struct {
	//注意, _id是十六进制格式
	Id    primitive.ObjectID `bson:"_id,omitempty"` // 注意使用omitempty标签，使得在插入时如果未指定ID则MongoDB会自动生成
	X姓名   string             `bson:"姓名"`
	X名称   string             `bson:"名称"`
	X年龄   int                `bson:"年龄"`
	X重量   int                `bson:"重量"`
	X购买产品 []string           `bson:"购买产品"`
	X购买时间 time.Time          `bson:"购买时间"`
	X支付方式 X支付方式              `bson:"支付方式"`
}
type X支付方式 struct {
	X支付方式   string   `bson:"支付方式"`
	X联系方式   string   `bson:"联系方式"`
	X可选支付方式 []string `bson:"可选支付方式"`
}

type X产品价格 struct {
	//注意, _id是十六进制格式
	Id    primitive.ObjectID `bson:"_id,omitempty"` // 注意使用omitempty标签，使得在插入时如果未指定ID则MongoDB会自动生成
	X产品名称 string             `bson:"产品名称"`
	X产品价格 int                `bson:"产品价格"`
}

func init() {
	ctx = context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "数据库_demo", Coll: "记账"})
}
