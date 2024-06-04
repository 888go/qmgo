package main

import (
	"fmt"
	"testing"
)

func Test_插入一条数据(t *testing.T) {
	var userInfo = UserInfo{
		Name:   "xm",
		Age:    7,
		Weight: 40,
	}
	result, _ := cli.InsertOne(ctx, userInfo)
	fmt.Println(result.InsertedID)
}

func Test_插入多条数据(t *testing.T) {
	var userInfos = []UserInfo{
		{Name: "a1", Age: 6, Weight: 20},
		{Name: "b2", Age: 6, Weight: 25},
		{Name: "c3", Age: 6, Weight: 30},
		{Name: "d4", Age: 6, Weight: 35},
		{Name: "a1", Age: 7, Weight: 40},
		{Name: "a1", Age: 8, Weight: 45},
		{Name: "d4", Age: 6, Weight: 35},
		{Name: "a1", Age: 3, Weight: 40},
		{Name: "a1", Age: 33, Weight: 45},
		{Name: "d4", Age: 33, Weight: 35},
		{Name: "a1", Age: 22, Weight: 40},
		{Name: "a1", Age: 22, Weight: 45},
		{Name: "d4", Age: 2, Weight: 35},
		{Age: 7, Weight: 40},
		{Name: "", Age: 8, Weight: 45},
	}
	result, err := cli.Collection.InsertMany(ctx, userInfos)
	fmt.Println(result.InsertedIDs)
	fmt.Println(err)
}
