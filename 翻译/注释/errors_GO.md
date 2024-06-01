
<原文开始>
// ErrQueryNotSlicePointer return if result argument is not a pointer to a slice
<原文结束>

# <翻译开始>
// ErrQueryNotSlicePointer 如果结果参数不是一个切片的指针，返回此错误 md5:99bf2bfe686c166d
# <翻译结束>


<原文开始>
// ErrQueryNotSliceType return if result argument is not slice address
<原文结束>

# <翻译开始>
// ErrQueryNotSliceType 如果结果参数不是切片的地址时返回错误 md5:a70365d8a017acd7
# <翻译结束>


<原文开始>
// ErrQueryResultTypeInconsistent return if result type is not equal mongodb value type
<原文结束>

# <翻译开始>
// ErrQueryResultTypeInconsistent 如果查询结果类型与MongoDB值类型不一致时返回 md5:940d09b0f6052a9f
# <翻译结束>


<原文开始>
// ErrQueryResultValCanNotChange return if the value of result can not be changed
<原文结束>

# <翻译开始>
// ErrQueryResultValCanNotChange 如果结果值不能更改，返回这个错误 md5:95f73ebb72c4463a
# <翻译结束>


<原文开始>
// ErrNoSuchDocuments return if no document found
<原文结束>

# <翻译开始>
// ErrNoSuchDocuments 如果未找到文档，则返回此错误 md5:381e7dce5c56bc42
# <翻译结束>


<原文开始>
// ErrTransactionRetry return if transaction need to retry
<原文结束>

# <翻译开始>
// ErrTransactionRetry 如果事务需要重试时返回该错误 md5:82e274f71f7c0175
# <翻译结束>


<原文开始>
// ErrTransactionNotSupported return if transaction not supported
<原文结束>

# <翻译开始>
// ErrTransactionNotSupported 如果不支持事务，则返回该错误 md5:97ad2e1fbe2e7207
# <翻译结束>


<原文开始>
// ErrNotSupportedUsername return if username is invalid
<原文结束>

# <翻译开始>
// ErrNotSupportedUsername 如果用户名无效，则返回此错误 md5:c9df5d462362cad6
# <翻译结束>


<原文开始>
// ErrNotSupportedPassword return if password is invalid
<原文结束>

# <翻译开始>
// ErrNotSupportedPassword 如果密码无效时返回 md5:e9df4f7f8304cc70
# <翻译结束>


<原文开始>
// ErrNotValidSliceToInsert return if insert argument is not valid slice
<原文结束>

# <翻译开始>
// ErrNotValidSliceToInsert 如果插入的参数不是一个有效的切片，返回该错误 md5:2b940e5853972183
# <翻译结束>


<原文开始>
// ErrReplacementContainUpdateOperators return if replacement document contain update operators
<原文结束>

# <翻译开始>
// ErrReplacementContainUpdateOperators 如果替换文档中包含更新操作符，返回错误 md5:4f127578930f91fa
# <翻译结束>


<原文开始>
// IsErrNoDocuments check if err is no documents, both mongo-go-driver error and qmgo custom error
// Deprecated, simply call if err == ErrNoSuchDocuments or if err == mongo.ErrNoDocuments
<原文结束>

# <翻译开始>
// IsErrNoDocuments 检查错误是否表示没有找到文档，既包括 mongo-go-driver 的错误也包括 qmgo 自定义的错误
// 已弃用，直接判断 err == ErrNoSuchDocuments 或者 err == mongo.ErrNoDocuments 即可
// md5:a9bcbf0c80c5509c
# <翻译结束>


<原文开始>
// IsDup check if err is mongo E11000 (duplicate err)。
<原文结束>

# <翻译开始>
// IsDup 检查错误是否为MongoDB的E11000（重复错误）。 md5:4a3435e9efa67970
# <翻译结束>

