
<原文开始>
// check if unique indexs is working
<原文结束>

# <翻译开始>
// 检查唯一索引是否生效
# <翻译结束>


<原文开始>
// filter is nil or wrong BSON Document format
<原文结束>

# <翻译开始>
// filter 为空或其格式错误，非正确的 BSON 文档格式
# <翻译结束>


<原文开始>
// replacement is nil or wrong BSON Document format
<原文结束>

# <翻译开始>
// 如果replacement为nil，或者其格式不符合BSON文档规范，则出现错误
# <翻译结束>


<原文开始>
// id3 will insert into the inserted document
<原文结束>

# <翻译开始>
// id3 将会插入到待插入的文档中
# <翻译结束>


<原文开始>
// filter with id different from id in document, error
<原文结束>

# <翻译开始>
// 对具有与文档中id不同的id进行过滤，错误
# <翻译结束>


<原文开始>
// update already exist record
<原文结束>

# <翻译开始>
// 更新已存在的记录
# <翻译结束>


<原文开始>
// update is nil or wrong BSON Document format
<原文结束>

# <翻译开始>
// update 为空或其格式错误的 BSON 文档
# <翻译结束>


<原文开始>
// delete record: name = "Alice" , after that, expect one name = "Alice" record
<原文结束>

# <翻译开始>
// 删除记录：姓名 = "Alice"，之后预期还存在一条姓名为 "Alice" 的记录
// 请注意，根据这段注释描述的操作与实际代码逻辑可能存在不符的情况。从字面意思理解，这段注释表达的是删除一个名为"Alice"的记录，但操作后仍然期望存在一条姓名为"Alice"的记录，这在通常情况下是矛盾的。若要准确翻译并符合代码逻辑，请提供更多上下文或检查代码实现。
# <翻译结束>


<原文开始>
// delete not match  record , report err
<原文结束>

# <翻译开始>
// 删除不匹配的记录，并报告错误
# <翻译结束>


<原文开始>
// delete record: name = "Alice" ,after that, expect - record : name = "Alice"
<原文结束>

# <翻译开始>
// 删除记录：姓名 = "Alice"，之后期望 - 记录：姓名 = "Alice"
// 这段代码注释的含义是：
// ```go
// 删除名为"Alice"的记录，删除后预期该记录（姓名 = "Alice"）将不存在
# <翻译结束>

