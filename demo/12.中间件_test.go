package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func (u *UserInfo) BeforeInsert(ctx context.Context) error {
	u.Name = "修改后-爱丽丝" //此处可以直接修改数据.
	fmt.Println("在调用插入之前")
	return errors.New("你他吗没权限啊, 出错了")
	//return nil
}

func (u *UserInfo) AfterInsert(ctx context.Context) error {
	fmt.Println("在调用插入之后")
	return nil
}
func Test_hook(t *testing.T) {
	u := &UserInfo{Name: "爱丽丝", Age: 7}
	_, err := cli.InsertOne(context.Background(), u)
	fmt.Println(err)
}
