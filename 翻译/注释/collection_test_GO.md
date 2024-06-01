
<原文开始>
// check if unique indexs is working
<原文结束>

# <翻译开始>
// 检查唯一索引是否正常工作 md5:9b2257b60d7b5998
# <翻译结束>


<原文开始>
// filter is nil or wrong BSON Document format
<原文结束>

# <翻译开始>
// filter 是空或者不符合正确的BSON文档格式 md5:a55c6ef20a253667
# <翻译结束>


<原文开始>
// replacement is nil or wrong BSON Document format
<原文结束>

# <翻译开始>
// replacement 是空的或者不符合正确的BSON文档格式 md5:7b0ecb01590a648b
# <翻译结束>


<原文开始>
// id3 will insert into the inserted document
<原文结束>

# <翻译开始>
// id3 将会插入到已插入的文档中 md5:4cdfbeaa6a4c59ce
# <翻译结束>


<原文开始>
// filter with id different from id in document, error
<原文结束>

# <翻译开始>
// 使用与文档中id不同的过滤器，错误 md5:1864a41611ea40ba
# <翻译结束>


<原文开始>
// update already exist record
<原文结束>

# <翻译开始>
// 更新已存在的记录 md5:cc4fac8615b8fc8a
# <翻译结束>


<原文开始>
// update is nil or wrong BSON Document format
<原文结束>

# <翻译开始>
// update 是 nil 或者格式错误的 BSON 文档 md5:8f6e8bd5cf0af638
# <翻译结束>


<原文开始>
// if record is not exist，err is nil， MatchedCount in res is 0
<原文结束>

# <翻译开始>
// 如果记录不存在，err为nil，res中的MatchedCount为0 md5:ffbbcabc3c0f02fe
# <翻译结束>


<原文开始>
// delete record: name = "Alice" , after that, expect one name = "Alice" record
<原文结束>

# <翻译开始>
// 删除记录：名称为 "Alice"，之后预期存在一条名称为 "Alice" 的记录。 md5:274874b30e4288bb
# <翻译结束>


<原文开始>
// delete not match  record , report err
<原文结束>

# <翻译开始>
// 删除不匹配的记录，如果发生错误则报告错误 md5:46e3eb8e95abdfcc
# <翻译结束>


<原文开始>
// filter is bson.M{}，delete one document
<原文结束>

# <翻译开始>
// filter 是 bson.M{}，删除一个文档 md5:dc8fa3aa9522cd67
# <翻译结束>


<原文开始>
// delete record: name = "Alice" ,after that, expect - record : name = "Alice"
<原文结束>

# <翻译开始>
// 删除记录：名称为 "Alice"，之后，预期 - 记录：名称为 "Alice" md5:e6ccda4a8c588184
# <翻译结束>


<原文开始>
// delete with not match filter， DeletedCount in res is 0
<原文结束>

# <翻译开始>
// 使用不匹配的过滤器删除，结果中的DeletedCount为0 md5:61f4e36a0742d763
# <翻译结束>


<原文开始>
// filter is bson.M{}，delete all docs
<原文结束>

# <翻译开始>
// filter 是 bson.M{}，删除所有文档 md5:bbe0ca02d153a930
# <翻译结束>

