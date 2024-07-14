package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_求和(t *testing.T) {
	count, err := cli.Find(ctx, bson.M{"年龄": 6}).Count()
	fmt.Println(count, err)
}
