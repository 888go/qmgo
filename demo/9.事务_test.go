package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_简单事务(t *testing.T) {
	ctx := context.Background()
	callback := func(sessCtx context.Context) (interface{}, error) {
		// 重要提示：确保在整个事务的每个操作中都使用了sessCtx
		if _, err := cli.X插入(sessCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sessCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		return nil, nil
	}
	result, err := cli.X事务(ctx, callback)
	fmt.Println(result, err)
}

func Test_会话事务(t *testing.T) {
	ctx := context.Background()
	//同时，您可以创建会话并使用会话启动交易： （如果不再使用会话，请不要忘记调用 EndSession）
	s, err := cli.X创建Session事务()
	defer s.X结束Session(ctx)

	callback := func(sessCtx context.Context) (interface{}, error) {
		// 重要提示：确保在整个事务的每个操作中都使用了sessCtx
		if _, err := cli.X插入(sessCtx, bson.D{{"abc", int32(1)}}); err != nil {
			return nil, err
		}
		if _, err := cli.X插入(sessCtx, bson.D{{"xyz", int32(999)}}); err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err = s.X开始事务(ctx, callback)
	fmt.Println(err)
}
