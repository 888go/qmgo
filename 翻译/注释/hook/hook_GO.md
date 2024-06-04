
<原文开始>
// hookHandler defines the relations between hook type and handler
<原文结束>

# <翻译开始>
// hookHandler 定义钩子类型和处理器之间的关系 md5:bce577bc34fd8393
# <翻译结束>


<原文开始>
//
//func init() {
//	middleware.Register(Do)
//}
<原文结束>

# <翻译开始>
// ```go
// 函数init() {
// 中间件注册(Do)
// }
// ```
// 
// 这段Go代码的注释是描述`init()`函数的作用，它用于在程序启动时注册一个名为`Do`的中间件。
// md5:a0604c723a346113
# <翻译结束>


<原文开始>
// Do call the specific method to handle hook based on hType
// If opts has valid value, use it instead of original hook
<原文结束>

# <翻译开始>
// 根据hType调用特定的方法来处理钩子
// 如果opts有有效的值，将使用它替换原始的钩子
// md5:8a28d86282a2f1cb
# <翻译结束>


<原文开始>
// sliceHandle handles the slice hooks
<原文结束>

# <翻译开始>
// sliceHandle 处理切片钩子 md5:c688842b5e68c3d2
# <翻译结束>


<原文开始>
// []interface{}{UserType{}...}
<原文结束>

# <翻译开始>
// []interface{}{UserType实例...} md5:bda81608072dd1ad
# <翻译结束>


<原文开始>
// BeforeInsertHook InsertHook defines the insert hook interface
<原文结束>

# <翻译开始>
// BeforeInsertHook 插入钩子接口定义了插入操作前的钩子函数 md5:d21219ecf0626005
# <翻译结束>


<原文开始>
// beforeInsert calls custom BeforeInsert
<原文结束>

# <翻译开始>
// beforeInsert 调用自定义的 BeforeInsert md5:615b3c8fedf08917
# <翻译结束>


<原文开始>
// afterInsert calls custom AfterInsert
<原文结束>

# <翻译开始>
// afterInsert 调用自定义的 AfterInsert md5:2c328449f2524700
# <翻译结束>


<原文开始>
// BeforeUpdateHook defines the Update hook interface
<原文结束>

# <翻译开始>
// BeforeUpdateHook 定义了 Update 钩子接口 md5:5b0bf7445582bfc4
# <翻译结束>


<原文开始>
// beforeUpdate calls custom BeforeUpdate
<原文结束>

# <翻译开始>
// beforeUpdate 调用自定义的 BeforeUpdate md5:4241dc99bc7049cb
# <翻译结束>


<原文开始>
// afterUpdate calls custom AfterUpdate
<原文结束>

# <翻译开始>
// afterUpdate 调用自定义的 AfterUpdate md5:e97728f60d7fb285
# <翻译结束>


<原文开始>
// BeforeQueryHook QueryHook defines the query hook interface
<原文结束>

# <翻译开始>
// BeforeQueryHook QueryHook 定义了查询钩子接口 md5:7190d574d8ba3bb9
# <翻译结束>


<原文开始>
// beforeQuery calls custom BeforeQuery
<原文结束>

# <翻译开始>
// beforeQuery 调用自定义的 BeforeQuery md5:269716e251327a4b
# <翻译结束>


<原文开始>
// afterQuery calls custom AfterQuery
<原文结束>

# <翻译开始>
// afterQuery 调用自定义的 AfterQuery md5:3975e33a3442aa92
# <翻译结束>


<原文开始>
// BeforeRemoveHook RemoveHook defines the remove hook interface
<原文结束>

# <翻译开始>
// BeforeRemoveHook RemoveHook 定义了移除钩子接口 md5:9c4d45d4f016c9cc
# <翻译结束>


<原文开始>
// beforeRemove calls custom BeforeRemove
<原文结束>

# <翻译开始>
// beforeRemove 调用自定义的 BeforeRemove md5:28aee6c76322664d
# <翻译结束>


<原文开始>
// afterRemove calls custom AfterRemove
<原文结束>

# <翻译开始>
// afterRemove 调用自定义的 AfterRemove 方法 md5:76432724d5d50929
# <翻译结束>


<原文开始>
// BeforeUpsertHook UpsertHook defines the upsert hook interface
<原文结束>

# <翻译开始>
// BeforeUpsertHook UpsertHook 定义了 Upsert 操作前的钩子接口 md5:745e467bebed93fc
# <翻译结束>


<原文开始>
// beforeUpsert calls custom BeforeUpsert
<原文结束>

# <翻译开始>
// beforeUpsert 调用自定义的 BeforeUpsert md5:c4bfb36f702295c2
# <翻译结束>


<原文开始>
// afterUpsert calls custom AfterUpsert
<原文结束>

# <翻译开始>
// afterUpsert 调用自定义的 AfterUpsert md5:2bcc20678061e065
# <翻译结束>


<原文开始>
// do check if opType is supported and call hookHandler
<原文结束>

# <翻译开始>
// 检查opType是否支持，并调用hookHandler md5:1b5144f1d5dc2b78
# <翻译结束>

