package main

import (
	"fmt"
	"github.com/qiniu/qmgo/options"
	"testing"
)

func Test_创建索引(t *testing.T) {
	cli.CreateOneIndex(ctx, options.IndexModel{Key: []string{"名称"}})            //单个索引
	cli.CreateIndexes(ctx, []options.IndexModel{{Key: []string{"id2", "id3"}}}) //复合索引
}

func Test_删除所有索引(t *testing.T) {
	err := cli.DropAllIndexes(ctx)
	fmt.Println(err)
}

func Test_删除索引(t *testing.T) {
	err := cli.DropIndex(ctx, []string{"名称"}) //删除字段"名称" 索引
	fmt.Println(err)
}
