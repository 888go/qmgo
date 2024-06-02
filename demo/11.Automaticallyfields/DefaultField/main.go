package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/field"
)

type User struct {
	field.DefaultField `bson:",inline"`

	Name string `bson:"名称"`
	Age  int    `bson:"年龄"`
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

	u := &User{Name: "Lucas", Age: 7}
	_, err := cli.InsertOne(context.Background(), u) // tag为createAt、updateAt 和 _id 的字段会自动更新插入
	fmt.Println(err)
}
