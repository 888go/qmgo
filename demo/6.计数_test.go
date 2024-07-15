package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_求和(t *testing.T) {
	count, err := cli.X查询(ctx, bson.M{"年龄": 6}).X取数量()
	fmt.Println(count, err)
}
