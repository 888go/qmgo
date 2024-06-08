/*
 Copyright 2020 The Qmgo Authors.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package hook

import (
	"context"
	"github.com/qiniu/qmgo/operator"
	"reflect"
)

// hookHandler 定义钩子类型和处理器之间的关系 md5:bce577bc34fd8393
var hookHandler = map[operator.OpType]func(ctx context.Context, hook interface{}) error{
	operator.BeforeInsert:  beforeInsert,
	operator.AfterInsert:   afterInsert,
	operator.BeforeUpdate:  beforeUpdate,
	operator.AfterUpdate:   afterUpdate,
	operator.BeforeQuery:   beforeQuery,
	operator.AfterQuery:    afterQuery,
	operator.BeforeRemove:  beforeRemove,
	operator.AfterRemove:   afterRemove,
	operator.BeforeUpsert:  beforeUpsert,
	operator.AfterUpsert:   afterUpsert,
	operator.BeforeReplace: beforeUpdate,
	operator.AfterReplace:  afterUpdate,
}

// ```go
// 函数init() {
// 中间件注册(Do)
// }
// ```
// 
// 这段Go代码的注释是描述`init()`函数的作用，它用于在程序启动时注册一个名为`Do`的中间件。
// md5:a0604c723a346113

// 根据hType调用特定的方法来处理钩子
// 如果opts有有效的值，将使用它替换原始的钩子
// md5:8a28d86282a2f1cb
// [提示:] func 执行Hook(ctx 上下文.Context, 钩子 interface{})
// ff:
// ctx:
// hook:
// opType:
// opts:
func Do(ctx context.Context, hook interface{}, opType operator.OpType, opts ...interface{}) error {
	if len(opts) > 0 {
		hook = opts[0]
	}

	to := reflect.TypeOf(hook)
	if to == nil {
		return nil
	}
	switch to.Kind() {
	case reflect.Slice:
		return sliceHandle(ctx, hook, opType)
	case reflect.Ptr:
		v := reflect.ValueOf(hook).Elem()
		switch v.Kind() {
		case reflect.Slice:
			return sliceHandle(ctx, v.Interface(), opType)
		default:
			return do(ctx, hook, opType)
		}
	default:
		return do(ctx, hook, opType)
	}
}

// sliceHandle 处理切片钩子 md5:c688842b5e68c3d2
func sliceHandle(ctx context.Context, hook interface{}, opType operator.OpType) error {
	// []interface{}{UserType实例...} md5:bda81608072dd1ad
	if h, ok := hook.([]interface{}); ok {
		for _, v := range h {
			if err := do(ctx, v, opType); err != nil {
				return err
			}
		}
		return nil
	}
	// []UserType{}
	s := reflect.ValueOf(hook)
	for i := 0; i < s.Len(); i++ {
		if err := do(ctx, s.Index(i).Interface(), opType); err != nil {
			return err
		}
	}
	return nil
}

// BeforeInsertHook 插入钩子接口定义了插入操作前的钩子函数 md5:d21219ecf0626005
// [提示]
//type BeforeInsertHook interface {
//     // 方法名：处理BeforeInsert
//     ProcessBeforeInsert(document interface{}) (interface{}, error)
// }
// [结束]
type BeforeInsertHook interface {
	// [提示]
//type BeforeInsertHook interface {
//     在插入前处理(ctx 上下文) error
// }
// [结束]
BeforeInsert(ctx context.Context) error
}
// [提示:] type 插入后钩子接口 interface {}
type AfterInsertHook interface {
	// [提示]
//AfterInsert(ctx 上下文环境) 错误
// 
// BeforeInsert(ctx 上下文环境) (interface{}, error)
// 
// AfterUpdate(ctx 上下文环境, info *UpdateInfo) error
// 
// BeforeUpdate(ctx 上下文环境, info *UpdateInfo) (interface{}, error)
// 
// AfterDelete(ctx 上下文环境, info *DeleteInfo) error
// 
// BeforeDelete(ctx 上下文环境, info *DeleteInfo) (interface{}, error)
// 
// AfterFind(ctx 上下文环境, doc interface{}) error
// 
// BeforeFind(ctx 上下文环境, query interface{}) (interface{}, error)
// [结束]
AfterInsert(ctx context.Context) error
}

// beforeInsert 调用自定义的 BeforeInsert md5:615b3c8fedf08917
func beforeInsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeInsertHook); ok {
		return ih.BeforeInsert(ctx)
	}
	return nil
}

// afterInsert 调用自定义的 AfterInsert md5:2c328449f2524700
func afterInsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterInsertHook); ok {
		return ih.AfterInsert(ctx)
	}
	return nil
}

// BeforeUpdateHook 定义了 Update 钩子接口 md5:5b0bf7445582bfc4
// [提示]
//type BeforeUpdateHook interface {
//     // 在更新操作执行前调用的方法
//     OnBeforeUpdate(context.Context, *mongo.UpdateOptions, interface{}) error
// }
// [结束]
type BeforeUpdateHook interface {
	// [提示]
//在更新前操作(BeforeUpdate) 
// 参数：
// - 环境上下文(Context): 用于处理请求的上下文对象
// - 错误(Error): 如果发生错误，返回错误信息
// 
// 返回值：
// - 错误(Error): 执行过程中可能出现的错误
// [结束]
BeforeUpdate(ctx context.Context) error
}
// [提示]
//类型名：更新后钩子接口
// 
// 接口方法：无（因为这是一个空接口，没有定义任何方法）
// 
// 参数和返回值：无（接口本身没有指定任何方法，所以没有参数和返回值需要翻译）
// [结束]
type AfterUpdateHook interface {
	// [提示:] AfterUpdate(ctx 上下文环境) 错误
AfterUpdate(ctx context.Context) error
}

// beforeUpdate 调用自定义的 BeforeUpdate md5:4241dc99bc7049cb
func beforeUpdate(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeUpdateHook); ok {
		return ih.BeforeUpdate(ctx)
	}
	return nil
}

// afterUpdate 调用自定义的 AfterUpdate md5:e97728f60d7fb285
func afterUpdate(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterUpdateHook); ok {
		return ih.AfterUpdate(ctx)
	}
	return nil
}

// BeforeQueryHook QueryHook 定义了查询钩子接口 md5:7190d574d8ba3bb9
// [提示]
//type 查询前Hook接口 interface {
// }
// [结束]
type BeforeQueryHook interface {
	// [提示]
//在查询前执行(ctx上下文环境) (错误)
// 
// BeforeInsert(ctx context.Context, docs ...interface{}) error
// 
// 在插入前执行(ctx上下文环境, 文档... interface{}) (错误)
// 
// BeforeUpdate(ctx context.Context, filter, update interface{}) error
// 
// 在更新前执行(ctx上下文环境, 过滤器, 更新条件 interface{}) (错误)
// 
// BeforeDelete(ctx context.Context, filter interface{}) error
// 
// 在删除前执行(ctx上下文环境, 删除过滤器 interface{}) (错误)
// 
// AfterQuery(ctx context.Context, cursor *mongo.Cursor, err error) error
// 
// 查询后执行(ctx上下文环境, 游标 *mongo.Cursor, 错误 error) (错误)
// 
// AfterInsert(ctx context.Context, insertedID interface{}, err error) error
// 
// 插入后执行(ctx上下文环境, 插入的ID interface{}, 错误 error) (错误)
// 
// AfterUpdate(ctx context.Context, result *mongo.UpdateResult, err error) error
// 
// 更新后执行(ctx上下文环境, 更新结果 *mongo.UpdateResult, 错误 error) (错误)
// 
// AfterDelete(ctx context.Context, result *mongo.DeleteResult, err error) error
// 
// 删除后执行(ctx上下文环境, 删除结果 *mongo.DeleteResult, 错误 error) (错误)
// [结束]
BeforeQuery(ctx context.Context) error
}
// [提示]
//类型名称：查询后钩子接口
// 
// 接口方法：无（因为这是一个空接口，没有具体的方法定义）
// 
// 在Go语言中，接口类型定义了一组方法签名。在这个例子中，`AfterQueryHook`是一个接口，表示实现了该接口的类型需要提供一个或多个在查询操作之后执行的钩子函数。但由于接口本身没有定义任何方法，所以实现这个接口的类型可以自由地定义它们自己的后查询处理逻辑。
// [结束]
type AfterQueryHook interface {
	// [提示:] After查询(ctx 上下文环境) 错误
AfterQuery(ctx context.Context) error
}

// beforeQuery 调用自定义的 BeforeQuery md5:269716e251327a4b
func beforeQuery(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeQueryHook); ok {
		return ih.BeforeQuery(ctx)
	}
	return nil
}

// afterQuery 调用自定义的 AfterQuery md5:3975e33a3442aa92
func afterQuery(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterQueryHook); ok {
		return ih.AfterQuery(ctx)
	}
	return nil
}

// BeforeRemoveHook RemoveHook 定义了移除钩子接口 md5:9c4d45d4f016c9cc
// [提示]
//类型名称：删除前钩子接口
// 
// 接口方法：无（因为这是一个空接口，通常用于表示一个在删除操作前需要执行的钩子函数。实际的实现可能会包含具体的钩子函数，但这里没有定义具体的方法。）
// 
// 参数和返回值：无（由于接口本身没有定义方法，所以没有参数和返回值。在实现这个接口时，会根据实际需求定义钩子函数的参数和返回值。）
// [结束]
type BeforeRemoveHook interface {
	// [提示]
//在删除前执行(ctx上下文环境)错误处理
// 
// BeforeInsert(ctx context.Context) error
// 
// 在插入前执行(ctx上下文环境)错误处理
// 
// BeforeUpdate(ctx context.Context) error
// 
// 在更新前执行(ctx上下文环境)错误处理
// 
// AfterRemove(ctx context.Context) error
// 
// 在删除后执行(ctx上下文环境)错误处理
// 
// AfterInsert(ctx context.Context) error
// 
// 在插入后执行(ctx上下文环境)错误处理
// 
// AfterUpdate(ctx context.Context) error
// 
// 在更新后执行(ctx上下文环境)错误处理
// [结束]
BeforeRemove(ctx context.Context) error
}
// [提示]
//类型名称：移除后钩子接口
// 
// 参数和返回值名称：无（因为这是一个接口定义，没有具体的函数方法，所以没有参数和返回值名称。在实现这个接口的方法中，会根据具体需求定义参数和返回值）
// [结束]
type AfterRemoveHook interface {
	// [提示:] AfterRemove(ctx 上下文环境) 错误
AfterRemove(ctx context.Context) error
}

// beforeRemove 调用自定义的 BeforeRemove md5:28aee6c76322664d
func beforeRemove(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeRemoveHook); ok {
		return ih.BeforeRemove(ctx)
	}
	return nil
}

// afterRemove 调用自定义的 AfterRemove 方法 md5:76432724d5d50929
func afterRemove(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterRemoveHook); ok {
		return ih.AfterRemove(ctx)
	}
	return nil
}

// BeforeUpsertHook UpsertHook 定义了 Upsert 操作前的钩子接口 md5:745e467bebed93fc
// [提示]
//类型名称：BeforeUpsert钩子接口
// 
// 参数和返回值名称：无，因为这是一个接口类型定义，具体的实现可能会有不同的参数和返回值，需要在实现该接口时定义。
// [结束]
type BeforeUpsertHook interface {
	// [提示]
//在插入更新前操作 (BeforeUpsert) 
// 参数：
// - ctx：上下文 (Context)
// 返回值：
// - 错误 (error)
// [结束]
BeforeUpsert(ctx context.Context) error
}
// [提示]
//类型名称：AfterUpsertHook 翻译为：更新后钩子接口
// 
// 接口方法无，因此无需翻译参数和返回值。
// [结束]
type AfterUpsertHook interface {
	// [提示:] AfterUpsert(ctx 上下文环境) 错误
AfterUpsert(ctx context.Context) error
}

// beforeUpsert 调用自定义的 BeforeUpsert md5:c4bfb36f702295c2
func beforeUpsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeUpsertHook); ok {
		return ih.BeforeUpsert(ctx)
	}
	return nil
}

// afterUpsert 调用自定义的 AfterUpsert md5:2bcc20678061e065
func afterUpsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterUpsertHook); ok {
		return ih.AfterUpsert(ctx)
	}
	return nil
}

// 检查opType是否支持，并调用hookHandler md5:1b5144f1d5dc2b78
func do(ctx context.Context, hook interface{}, opType operator.OpType) error {
	if f, ok := hookHandler[opType]; !ok {
		return nil
	} else {
		return f(ctx, hook)
	}
}
