package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_更新(t *testing.T) {
	// 更新一个
	err := cli.UpdateOne(ctx, bson.M{"名称": "d4"}, bson.M{"$set": bson.M{"age": 7}})
	fmt.Println(err)

	// 全部更新
	result, err := cli.UpdateAll(ctx, bson.M{"年龄": 6}, bson.M{"$set": bson.M{"年龄": 10}})
	fmt.Println(result.MatchedCount, err)
	fmt.Println(result.ModifiedCount, err)
}
