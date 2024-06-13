package main

import (
	"fmt"
	"github.com/888go/qmgo/options"
	"testing"
)

func Test_创建单个索引(t *testing.T) {
	cli.X创建索引(ctx, options.X索引选项{X索引字段: []string{"名称"}}) //单个索引
}
func Test_创建复合索引(t *testing.T) {
	cli.X创建多条索引(ctx, []options.X索引选项{{X索引字段: []string{"id2", "id3"}}}) //复合索引
}

func Test_删除所有索引(t *testing.T) {
	err := cli.X删除全部索引(ctx)
	fmt.Println(err)
}

func Test_删除索引(t *testing.T) {
	err := cli.X删除索引(ctx, []string{"名称"}) //删除字段"名称" 索引
	fmt.Println(err)
}

/*
备注, qmgo没有实现索引列出,以及全文索引创建. 需要自己实现
*/
