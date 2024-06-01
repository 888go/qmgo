package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
)

// [提示]
//type 用户 struct {
//     姓名 string `bson:"name"`
//     年龄 int    `bson:"age"`
// }
// [结束]
type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
// [提示:] func (u *用户) 在插入前处理(ctx 上下文.Context) 错误 {}
// ff:
// ctx:
func (u *User) BeforeInsert(ctx context.Context) error {
	fmt.Println("在调用插入之前")
	return nil
}
// [提示:] func (u *用户) 插入后处理(ctx 上下文.Context) 错误 {} 
// ff:
// ctx:
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
