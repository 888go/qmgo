
<原文开始>
// use a single instance of Validate, it caches struct info
<原文结束>

# <翻译开始>
// 使用单一实例的Validate，它会缓存结构体信息
# <翻译结束>


<原文开始>
// SetValidate let validate can use custom rules
<原文结束>

# <翻译开始>
// SetValidate 设置验证器，使其可以使用自定义规则
# <翻译结束>


<原文开始>
// validatorNeeded checks if the validator is needed to opType
<原文结束>

# <翻译开始>
// validatorNeeded 检查对于 opType 是否需要验证器
# <翻译结束>


<原文开始>
// Do calls validator check
// Don't use opts here
<原文结束>

# <翻译开始>
// Do 调用验证器进行检查
// 在此处不要使用 opts
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
// do check if opType is supported and call fieldHandler
<原文结束>

# <翻译开始>
// 检查opType是否支持，并调用fieldHandler
# <翻译结束>


<原文开始>
// validatorStruct check if kind of doc is validator supported struct
// same implement as validator
<原文结束>

# <翻译开始>
// validatorStruct 检查doc的类型是否为validator支持的结构体
// 实现方式与validator相同
# <翻译结束>

