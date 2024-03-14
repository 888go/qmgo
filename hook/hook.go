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
	"github.com/888go/qmgo/operator"
	"reflect"
)

// hookHandler 定义了钩子类型与处理函数之间的关联关系
var hookHandler = map[操作符.OpType]func(ctx context.Context, hook interface{}) error{
	操作符.X插入前:  beforeInsert,
	操作符.X插入后:   afterInsert,
	操作符.X更新前:  beforeUpdate,
	操作符.X更新后:   afterUpdate,
	操作符.X查询前:   beforeQuery,
	操作符.X查询后:    afterQuery,
	操作符.X删除前:  beforeRemove,
	操作符.X删除后:   afterRemove,
	操作符.X更新或插入前:  beforeUpsert,
	操作符.X更新或插入后:   afterUpsert,
	操作符.X替换前: beforeUpdate,
	操作符.X替换后:  afterUpdate,
}

// 
// // 注册Do函数到中间件
// func init() {
// 	middleware.Register(Do)
// }

// 根据hType调用特定方法来处理钩子
// 如果opts中有有效值，则使用它替代原始钩子
func Do(ctx context.Context, hook interface{}, opType 操作符.OpType, opts ...interface{}) error {
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

// sliceHandle 处理切片钩子
func sliceHandle(ctx context.Context, hook interface{}, opType 操作符.OpType) error {
	// []interface{}{UserType{}...} 
// 创建一个接口类型切片，其中包含零个或多个UserType结构体实例。这里的"..."表示可变数量的参数，表示可以传入任意数量的UserType实例到切片中。
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

// BeforeInsertHook InsertHook 定义了插入钩子接口
type BeforeInsertHook interface {
	BeforeInsert(ctx context.Context) error
}
type AfterInsertHook interface {
	AfterInsert(ctx context.Context) error
}

// beforeInsert 在插入前调用自定义的 BeforeInsert
func beforeInsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeInsertHook); ok {
		return ih.BeforeInsert(ctx)
	}
	return nil
}

// afterInsert 在插入后调用自定义的 AfterInsert
func afterInsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterInsertHook); ok {
		return ih.AfterInsert(ctx)
	}
	return nil
}

// BeforeUpdateHook 定义了 Update 钩子接口
type BeforeUpdateHook interface {
	BeforeUpdate(ctx context.Context) error
}
type AfterUpdateHook interface {
	AfterUpdate(ctx context.Context) error
}

// beforeUpdate 调用自定义的 BeforeUpdate
func beforeUpdate(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeUpdateHook); ok {
		return ih.BeforeUpdate(ctx)
	}
	return nil
}

// afterUpdate 调用自定义的 AfterUpdate
func afterUpdate(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterUpdateHook); ok {
		return ih.AfterUpdate(ctx)
	}
	return nil
}

// BeforeQueryHook QueryHook 定义了查询钩子接口
type BeforeQueryHook interface {
	BeforeQuery(ctx context.Context) error
}
type AfterQueryHook interface {
	AfterQuery(ctx context.Context) error
}

// beforeQuery 调用自定义的 BeforeQuery 方法
func beforeQuery(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeQueryHook); ok {
		return ih.BeforeQuery(ctx)
	}
	return nil
}

// afterQuery 调用自定义的 AfterQuery
func afterQuery(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterQueryHook); ok {
		return ih.AfterQuery(ctx)
	}
	return nil
}

// BeforeRemoveHook RemoveHook 定义了删除钩子接口
type BeforeRemoveHook interface {
	BeforeRemove(ctx context.Context) error
}
type AfterRemoveHook interface {
	AfterRemove(ctx context.Context) error
}

// beforeRemove 调用自定义的 BeforeRemove
func beforeRemove(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeRemoveHook); ok {
		return ih.BeforeRemove(ctx)
	}
	return nil
}

// afterRemove 调用自定义的 AfterRemove 函数
func afterRemove(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterRemoveHook); ok {
		return ih.AfterRemove(ctx)
	}
	return nil
}

// BeforeUpsertHook UpsertHook 定义了 upsert 钩子接口
type BeforeUpsertHook interface {
	BeforeUpsert(ctx context.Context) error
}
type AfterUpsertHook interface {
	AfterUpsert(ctx context.Context) error
}

// beforeUpsert调用自定义的BeforeUpsert函数
func beforeUpsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(BeforeUpsertHook); ok {
		return ih.BeforeUpsert(ctx)
	}
	return nil
}

// afterUpsert 调用自定义的 AfterUpsert 方法
func afterUpsert(ctx context.Context, hook interface{}) error {
	if ih, ok := hook.(AfterUpsertHook); ok {
		return ih.AfterUpsert(ctx)
	}
	return nil
}

// 检查opType是否支持，并调用hookHandler
func do(ctx context.Context, hook interface{}, opType 操作符.OpType) error {
	if f, ok := hookHandler[opType]; !ok {
		return nil
	} else {
		return f(ctx, hook)
	}
}
