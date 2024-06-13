package main

import (
	"context"
	"fmt"
	"github.com/888go/qmgo"
	"github.com/888go/qmgo/field"
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
	cli, _ := mgo类.X连接(ctx, &mgo类.X配置{X连接URI: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", X数据库名: "数据库_demo", X集合名: "学生"})

	u := &X学生{X名字: "Lucas", X年龄: 7}
	_, err := cli.X插入(context.Background(), u) // tag为createAt、updateAt 和 _id 的字段会自动更新插入
	fmt.Println(err)
}

// 自定义字段,指定自定义field的field名
func (u *X学生) CustomFields() field.CustomFieldsBuilder {
	return field.NewCustom().X设置创建时间字段名("CreateTimeAt").X设置更新时间字段名("UpdateTimeAt").X设置ID字段名("MyId")
}

func Test_自定义字段(t *testing.T) {
	ctx := context.Background()
	//注意要替换连接 mongodb://账号:密码@ip:端口
	cli, _ := mgo类.X连接(ctx, &mgo类.X配置{X连接URI: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", X数据库名: "数据库_demo", X集合名: "学生"})

	u := &X学生{X名字: "Lucas", X年龄: 7}
	_, err := cli.X插入(context.Background(), u)
	fmt.Println(err)
}
