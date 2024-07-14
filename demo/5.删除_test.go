package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_删除一条文档(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/delete/
	err := cli.Remove(ctx, bson.M{"年龄": 7})
	fmt.Println(err)
}

func Test_删除多条文档(t *testing.T) {
	//https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/crud/write-operations/delete/
	返回, _ := cli.RemoveAll(ctx, bson.M{"年龄": 7})
	fmt.Println(返回.DeletedCount)
}
