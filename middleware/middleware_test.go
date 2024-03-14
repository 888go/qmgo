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
	ast.NoError(Do(ctx, "success", 操作符.X插入前))

	// valid register
	Register(callbackTest)
	ast.NoError(Do(ctx, "success", 操作符.X插入前))
	ast.Error(Do(ctx, "failure", 操作符.X更新或插入前))
	ast.NoError(Do(ctx, "failure", 操作符.X更新前, "success"))
}

func callbackTest(ctx context.Context, doc interface{}, opType 操作符.OpType, opts ...interface{}) error {
	if doc.(string) == "success" && opType == 操作符.X插入前 {
		return nil
	}
	if len(opts) > 0 && opts[0].(string) == "success" {
		return nil
	}
	if doc.(string) == "failure" && opType == 操作符.X更新或插入前 {
		return fmt.Errorf("this is error")
	}
	return nil
}
