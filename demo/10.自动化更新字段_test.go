package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/field"
	"testing"
	"time"
)

type X学生 struct {
	X名字 string `bson:"名字"`
	X年龄 int    `bson:"年龄"`

	MyId         string    `bson:"ID"`
	CreateTimeAt time.Time `bson:"创建时间"`
	UpdateTimeAt int64     `bson:"更新时间"`
}

func Test_默认字段(t *testing.T) {
	ctx := context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "数据库_demo", Coll: "学生"})

	u := &X学生{X名字: "Lucas", X年龄: 7}
	_, err := cli.InsertOne(context.Background(), u) // tag为createAt、updateAt 和 _id 的字段会自动更新插入
	fmt.Println(err)
}

// 自定义字段,指定自定义field的field名
func (u *X学生) CustomFields() field.CustomFieldsBuilder {
	return field.NewCustom().SetCreateAt("CreateTimeAt").SetUpdateAt("UpdateTimeAt").SetId("MyId")
}

func Test_自定义字段(t *testing.T) {
	ctx := context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "数据库_demo", Coll: "学生"})

	u := &X学生{X名字: "Lucas", X年龄: 7}
	_, err := cli.InsertOne(context.Background(), u)
	fmt.Println(err)
}
