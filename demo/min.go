package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	ctx := context.Background()
	//mongodb://mongo_DibRHZ@124.220.182.235:27017
	client, _ := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://mongo_DibRHZ:chengnuandi1@124.220.182.235:27017"})
	db := client.Database("测试")
	集合 := db.Collection("测试")
	var one bson.M
	集合.Find(ctx, bson.M{"名称": "菜鸟教程"}).One(&one)
	fmt.Println(one)
	///替换文档不能包含以'$'开头的键
	//替换文档不能包含以'$'开头的键
	//_, err = 集合.Upsert(ctx, bson.M{"名称": "菜鸟教程2"}, bson.M{"name": "菜鸟教程", "网址": "http://www.runoob.com"})
	_, err := 集合.UpdateAll(ctx, bson.M{"名称": "菜鸟教程22"}, bson.M{"$set": bson.M{"by": "菜鸟教程1111", "网址": "http://www.runoob.com"}})
	if err != nil {
		return
	}

}
