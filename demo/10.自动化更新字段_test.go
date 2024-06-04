package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/field"
	"testing"
	"time"
)

type User struct {
	Name string `bson:"名字"`
	Age  int    `bson:"年龄"`

	MyId         string    `bson:"ID"`
	CreateTimeAt time.Time `bson:"创建时间"`
	UpdateTimeAt int64     `bson:"更新时间"`
}

func Test_默认字段(t *testing.T) {
	ctx := context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "学校数据库", Coll: "学生"})

	u := &User{Name: "Lucas", Age: 7}
	_, err := cli.InsertOne(context.Background(), u) // tag为createAt、updateAt 和 _id 的字段会自动更新插入
	fmt.Println(err)
}

// 自定义字段,指定自定义field的field名
func (u *User) CustomFields() field.CustomFieldsBuilder {
	return field.NewCustom().SetCreateAt("CreateTimeAt").SetUpdateAt("UpdateTimeAt").SetId("MyId")
}

func Test_自定义字段(t *testing.T) {
	ctx := context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "学校数据库", Coll: "学生"})

	u := &User{Name: "Lucas", Age: 7}
	_, err := cli.InsertOne(context.Background(), u)
	fmt.Println(err)
}
