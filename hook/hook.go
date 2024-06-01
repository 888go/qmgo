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
	// []interface{}{UserType{}...} 的中文翻译为：
// []interface{}{UserType{}...} 的中文翻译为：
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
type BeforeInsertHook interface {
	BeforeInsert(ctx context.Context) error
}
type AfterInsertHook interface {
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
type BeforeUpdateHook interface {
	BeforeUpdate(ctx context.Context) error
}
type AfterUpdateHook interface {
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
type BeforeQueryHook interface {
	BeforeQuery(ctx context.Context) error
}
type AfterQueryHook interface {
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
type BeforeRemoveHook interface {
	BeforeRemove(ctx context.Context) error
}
type AfterRemoveHook interface {
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
type BeforeUpsertHook interface {
	BeforeUpsert(ctx context.Context) error
}
type AfterUpsertHook interface {
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
