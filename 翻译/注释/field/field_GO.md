
<原文开始>
// filedHandler defines the relations between field type and handler
<原文结束>

# <翻译开始>
// filedHandler 定义字段类型和处理器之间的关系 md5:c7cd659bd6a053b2
# <翻译结束>


<原文开始>
//func init() {
//	middleware.Register(Do)
//}
<原文结束>

# <翻译开始>
// 函数 init() {
// 注册 middleware，参数为 Do
//}
// md5:4bdefdddb5ec33c1
# <翻译结束>


<原文开始>
// Do call the specific method to handle field based on fType
// Don't use opts here
<原文结束>

# <翻译开始>
// Do 调用特定方法根据 fType 处理字段
// 不在这里使用 opts
// md5:01967b5b64a19adb
# <翻译结束>


<原文开始>
//fmt.Println("not support type")
<原文结束>

# <翻译开始>
//fmt.Println("不支持此类类型") md5:2ba1fad322480d74
# <翻译结束>


<原文开始>
// sliceHandle handles the slice docs
<原文结束>

# <翻译开始>
// sliceHandle处理切片文档 md5:92800dd5899836ce
# <翻译结束>


<原文开始>
// []interface{}{UserType{}...}
<原文结束>

# <翻译开始>
// []interface{}{UserType实例...} md5:bda81608072dd1ad
# <翻译结束>


<原文开始>
// beforeInsert handles field before insert
// If value of field createAt is valid in doc, upsert doesn't change it
// If value of field id is valid in doc, upsert doesn't change it
// Change the value of field updateAt anyway
<原文结束>

# <翻译开始>
// beforeInsert 在插入前处理字段
// 如果文档中的createAt字段的值有效，upsert 不会改变它
// 如果文档中的id字段的值有效，upsert 不会改变它
// 无论如何，改变updateAt字段的值
// md5:f49d81597c8212f6
# <翻译结束>


<原文开始>
// beforeUpdate handles field before update
<原文结束>

# <翻译开始>
// beforeUpdate处理更新前的字段 md5:a783a1aa99fba490
# <翻译结束>


<原文开始>
// beforeUpsert handles field before upsert
// If value of field createAt is valid in doc, upsert doesn't change it
// If value of field id is valid in doc, upsert doesn't change it
// Change the value of field updateAt anyway
<原文结束>

# <翻译开始>
// beforeUpsert 处理字段的before upsert操作
// 如果doc中field createAt的值有效，upsert操作不会改变它
// 如果doc中field id的值有效，upsert操作也不会改变它
// 无论如何都会更新field updateAt的值
// md5:d286cfb6c0a1f1da
# <翻译结束>


<原文开始>
// do check if opType is supported and call fieldHandler
<原文结束>

# <翻译开始>
// 检查opType是否被支持，并调用fieldHandler方法 md5:3bb8cbff6cb4f5e3
# <翻译结束>

