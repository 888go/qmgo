package middleware

import (
	"context"
	"fmt"
	"github.com/888go/qmgo/operator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMiddleware(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()
	// not register
	ast.NoError(Do(ctx, "success", mgo常量.X钩子_插入前))

	// valid register
	Register(callbackTest)
	ast.NoError(Do(ctx, "success", mgo常量.X钩子_插入前))
	ast.Error(Do(ctx, "failure", mgo常量.X钩子_替换插入前))
	ast.NoError(Do(ctx, "failure", mgo常量.X钩子_更新前, "success"))
}

func callbackTest(ctx context.Context, doc interface{}, opType mgo常量.OpType, opts ...interface{}) error {
	if doc.(string) == "success" && opType == mgo常量.X钩子_插入前 {
		return nil
	}
	if len(opts) > 0 && opts[0].(string) == "success" {
		return nil
	}
	if doc.(string) == "failure" && opType == mgo常量.X钩子_替换插入前 {
		return fmt.Errorf("this is error")
	}
	return nil
}
