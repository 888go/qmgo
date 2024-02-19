
<原文开始>
// hookHandler defines the relations between hook type and handler
<原文结束>

# <翻译开始>
// hookHandler 定义了钩子类型与处理函数之间的关联关系
# <翻译结束>


<原文开始>
//
//func init() {
//	middleware.Register(Do)
//}
<原文结束>

# <翻译开始>
// 
// // 注册Do函数到中间件
// func init() {
// 	middleware.Register(Do)
// }
# <翻译结束>


<原文开始>
// Do call the specific method to handle hook based on hType
// If opts has valid value, use it instead of original hook
<原文结束>

# <翻译开始>
// 根据hType调用特定方法来处理钩子
// 如果opts中有有效值，则使用它替代原始钩子
# <翻译结束>


<原文开始>
// sliceHandle handles the slice hooks
<原文结束>

# <翻译开始>
// sliceHandle 处理切片钩子
# <翻译结束>


<原文开始>
// []interface{}{UserType{}...}
<原文结束>

# <翻译开始>
// []interface{}{UserType{}...} 
// 创建一个接口类型切片，其中包含零个或多个UserType结构体实例。这里的"..."表示可变数量的参数，表示可以传入任意数量的UserType实例到切片中。
# <翻译结束>


<原文开始>
// BeforeInsertHook InsertHook defines the insert hook interface
<原文结束>

# <翻译开始>
// BeforeInsertHook InsertHook 定义了插入钩子接口
# <翻译结束>


<原文开始>
// beforeInsert calls custom BeforeInsert
<原文结束>

# <翻译开始>
// beforeInsert 在插入前调用自定义的 BeforeInsert
# <翻译结束>


<原文开始>
// afterInsert calls custom AfterInsert
<原文结束>

# <翻译开始>
// afterInsert 在插入后调用自定义的 AfterInsert
# <翻译结束>


<原文开始>
// BeforeUpdateHook defines the Update hook interface
<原文结束>

# <翻译开始>
// BeforeUpdateHook 定义了 Update 钩子接口
# <翻译结束>


<原文开始>
// beforeUpdate calls custom BeforeUpdate
<原文结束>

# <翻译开始>
// beforeUpdate 调用自定义的 BeforeUpdate
# <翻译结束>


<原文开始>
// afterUpdate calls custom AfterUpdate
<原文结束>

# <翻译开始>
// afterUpdate 调用自定义的 AfterUpdate
# <翻译结束>


<原文开始>
// BeforeQueryHook QueryHook defines the query hook interface
<原文结束>

# <翻译开始>
// BeforeQueryHook QueryHook 定义了查询钩子接口
# <翻译结束>


<原文开始>
// beforeQuery calls custom BeforeQuery
<原文结束>

# <翻译开始>
// beforeQuery 调用自定义的 BeforeQuery 方法
# <翻译结束>


<原文开始>
// afterQuery calls custom AfterQuery
<原文结束>

# <翻译开始>
// afterQuery 调用自定义的 AfterQuery
# <翻译结束>


<原文开始>
// BeforeRemoveHook RemoveHook defines the remove hook interface
<原文结束>

# <翻译开始>
// BeforeRemoveHook RemoveHook 定义了删除钩子接口
# <翻译结束>


<原文开始>
// beforeRemove calls custom BeforeRemove
<原文结束>

# <翻译开始>
// beforeRemove 调用自定义的 BeforeRemove
# <翻译结束>


<原文开始>
// afterRemove calls custom AfterRemove
<原文结束>

# <翻译开始>
// afterRemove 调用自定义的 AfterRemove 函数
# <翻译结束>


<原文开始>
// BeforeUpsertHook UpsertHook defines the upsert hook interface
<原文结束>

# <翻译开始>
// BeforeUpsertHook UpsertHook 定义了 upsert 钩子接口
# <翻译结束>


<原文开始>
// beforeUpsert calls custom BeforeUpsert
<原文结束>

# <翻译开始>
// beforeUpsert调用自定义的BeforeUpsert函数
# <翻译结束>


<原文开始>
// afterUpsert calls custom AfterUpsert
<原文结束>

# <翻译开始>
// afterUpsert 调用自定义的 AfterUpsert 方法
# <翻译结束>


<原文开始>
// do check if opType is supported and call hookHandler
<原文结束>

# <翻译开始>
// 检查opType是否支持，并调用hookHandler
# <翻译结束>

