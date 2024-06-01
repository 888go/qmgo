package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func (u *User) BeforeInsert(ctx context.Context) error {
	fmt.Println("在调用插入之前")
	return nil
}
func (u *User) AfterInsert(ctx context.Context) error {
	fmt.Println("在调用插入之前")
	return nil
}

func main() {
	ctx := context.Background()
	cli, _ := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://mongo_tdBG3A:mongo_RSmrcT@121.89.206.172:27017", Database: "class", Coll: "user"})

	////关闭连接
	//defer func() {
	//	if err := cli.Close(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	u := &User{Name: "Alice", Age: 7}
	_, err := cli.InsertOne(context.Background(), u)
	fmt.Println(err)
}
