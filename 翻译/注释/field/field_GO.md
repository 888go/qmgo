
<原文开始>
// filedHandler defines the relations between field type and handler
<原文结束>

# <翻译开始>
// filedHandler 定义了字段类型与处理器之间的关系
# <翻译结束>


<原文开始>
//func init() {
//	middleware.Register(Do)
//}
<原文结束>

# <翻译开始>
// 注册函数
//func init() {
//	// 将Do函数注册到中间件
//	middleware.Register(Do)
//}
# <翻译结束>


<原文开始>
// Do call the specific method to handle field based on fType
// Don't use opts here
<原文结束>

# <翻译开始>
// 根据fType调用特定方法来处理字段
// 在此处不要使用opts
# <翻译结束>


<原文开始>
//fmt.Println("not support type")
<原文结束>

# <翻译开始>
// 输出："不支持的类型"
# <翻译结束>


<原文开始>
// sliceHandle handles the slice docs
<原文结束>

# <翻译开始>
// sliceHandle 处理切片文档
# <翻译结束>


<原文开始>
// []interface{}{UserType{}...}
<原文结束>

# <翻译开始>
// []interface{}{UserType{}...} 
// 创建一个接口类型切片，其中包含零个或多个UserType结构体实例。这里的"..."表示可变数量的参数，表示可以传入任意数量的UserType实例到切片中。
# <翻译结束>


<原文开始>
// beforeInsert handles field before insert
// If value of field createAt is valid in doc, upsert doesn't change it
// If value of field id is valid in doc, upsert doesn't change it
// Change the value of field updateAt anyway
<原文结束>

# <翻译开始>
// beforeInsert 在插入前处理字段
// 如果doc中createAt字段的值有效，upsert不会改变它
// 如果doc中id字段的值有效，upsert不会改变它
// 无论如何都会更新updateAt字段的值
# <翻译结束>


<原文开始>
// beforeUpdate handles field before update
<原文结束>

# <翻译开始>
// beforeUpdate 在更新字段前进行处理
# <翻译结束>


<原文开始>
// beforeUpsert handles field before upsert
// If value of field createAt is valid in doc, upsert doesn't change it
// If value of field id is valid in doc, upsert doesn't change it
// Change the value of field updateAt anyway
<原文结束>

# <翻译开始>
// beforeUpsert 在执行upsert操作前处理字段
// 如果doc中createAt字段的值有效，upsert操作不会改变它
// 如果doc中id字段的值有效，upsert操作也不会改变它
// 无论如何都会更新updateAt字段的值
# <翻译结束>


<原文开始>
// do check if opType is supported and call fieldHandler
<原文结束>

# <翻译开始>
// 检查opType是否支持，并调用fieldHandler
# <翻译结束>

