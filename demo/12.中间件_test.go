package main

import (
	"context"
	"fmt"
	"testing"
)

func (u *X记账) BeforeInsert(ctx context.Context) error {
	u.X名称 = "修改后-爱丽丝" //此处可以直接修改数据.
	fmt.Println("在调用插入之前")
	return nil
}

func (u *X记账) AfterInsert(ctx context.Context) error {
	fmt.Println("在调用插入之后")
	return nil
}

func Test_hook(t *testing.T) {
	u := &X记账{X名称: "爱丽丝", X年龄: 7}
	_, err := cli.X插入(context.Background(), u)
	fmt.Println(err)
}
