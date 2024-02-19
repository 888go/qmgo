
<原文开始>
// ErrQueryNotSlicePointer return if result argument is not a pointer to a slice
<原文结束>

# <翻译开始>
// ErrQueryNotSlicePointer 当查询结果参数不是指向切片的指针时返回该错误
# <翻译结束>


<原文开始>
// ErrQueryNotSliceType return if result argument is not slice address
<原文结束>

# <翻译开始>
// ErrQueryNotSliceType 当结果参数不是切片地址时返回该错误
# <翻译结束>


<原文开始>
// ErrQueryResultTypeInconsistent return if result type is not equal mongodb value type
<原文结束>

# <翻译开始>
// ErrQueryResultTypeInconsistent 当查询结果类型与mongodb的值类型不相等时返回该错误
# <翻译结束>


<原文开始>
// ErrQueryResultValCanNotChange return if the value of result can not be changed
<原文结束>

# <翻译开始>
// ErrQueryResultValCanNotChange 当查询结果的值不能被更改时返回该错误
# <翻译结束>


<原文开始>
// ErrNoSuchDocuments return if no document found
<原文结束>

# <翻译开始>
// ErrNoSuchDocuments 当没有找到任何文档时返回这个错误
# <翻译结束>


<原文开始>
// ErrTransactionRetry return if transaction need to retry
<原文结束>

# <翻译开始>
// ErrTransactionRetry：如果事务需要重试则返回该错误
# <翻译结束>


<原文开始>
// ErrTransactionNotSupported return if transaction not supported
<原文结束>

# <翻译开始>
// ErrTransactionNotSupported 当事务不被支持时返回该错误
# <翻译结束>


<原文开始>
// ErrNotSupportedUsername return if username is invalid
<原文结束>

# <翻译开始>
// ErrNotSupportedUsername 当用户名无效时返回该错误
# <翻译结束>


<原文开始>
// ErrNotSupportedPassword return if password is invalid
<原文结束>

# <翻译开始>
// ErrNotSupportedPassword 当密码无效时返回该错误
# <翻译结束>


<原文开始>
// ErrNotValidSliceToInsert return if insert argument is not valid slice
<原文结束>

# <翻译开始>
// ErrNotValidSliceToInsert 当插入参数不是一个有效的切片时返回该错误
# <翻译结束>


<原文开始>
// ErrReplacementContainUpdateOperators return if replacement document contain update operators
<原文结束>

# <翻译开始>
// ErrReplacementContainUpdateOperators 当替换文档包含更新操作符时返回错误
# <翻译结束>


<原文开始>
// IsErrNoDocuments check if err is no documents, both mongo-go-driver error and qmgo custom error
// Deprecated, simply call if err == ErrNoSuchDocuments or if err == mongo.ErrNoDocuments
<原文结束>

# <翻译开始>
// IsErrNoDocuments 检查 err 是否表示没有文档，包括mongo-go-driver库的错误和qmgo自定义错误
// 已弃用，直接调用 if err == ErrNoSuchDocuments 或 if err == mongo.ErrNoDocuments 即可
# <翻译结束>

