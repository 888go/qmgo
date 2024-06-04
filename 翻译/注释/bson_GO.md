
<原文开始>
// alias mongo drive bson primitives
// thus user don't need to import go.mongodb.org/mongo-driver/mongo, it's all in qmgo
<原文结束>

# <翻译开始>
// 别名mongo驱动的bson原语
// 因此用户不需要导入go.mongodb.org/mongo-driver/mongo，所有内容都在qmgo中可用
// md5:2f6e3ba77edc7a63
# <翻译结束>


<原文开始>
// M is an alias of bson.M
<原文结束>

# <翻译开始>
// map[string]interface{} , 如:bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}, M是 bson.M 的别名 md5:66b7bee0d7904542
# <翻译结束>


<原文开始>
// A is an alias of bson.A
<原文结束>

# <翻译开始>
// []interface{},如:bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}} , A是bson.A的别名 md5:7a6f09b99ea36324
# <翻译结束>


<原文开始>
// D is an alias of bson.D
<原文结束>

# <翻译开始>
// Key/Value结构体数组, 如:bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}} ,D是 bson.D 的别名  md5:a2fd7b05e87775b6
# <翻译结束>


<原文开始>
// E is an alias of bson.E
<原文结束>

# <翻译开始>
// Key/Value结构体, E 内部的单个元素,是bson.E的别名 md5:d1a800789b88ac58
# <翻译结束>

