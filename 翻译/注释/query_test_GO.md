
<原文开始>
// filter is bson.M{}，match all and return one
<原文结束>

# <翻译开始>
// filter 是 bson.M 类型的空映射，表示匹配所有文档并返回一个结果。 md5:5a0dc74674539e4e
# <翻译结束>


<原文开始>
// res is nil or can't parse
<原文结束>

# <翻译开始>
// res 为 nil 或者无法解析 md5:970a874db5a3d5c0
# <翻译结束>


<原文开始>
// res is a parseable object, but the bson tag is inconsistent with the mongodb record, no error is reported, res is the initialization state of the data structure
<原文结束>

# <翻译开始>
// res是一个解析的对象，但是bson标签与mongodb记录不一致，没有报告错误，res的数据结构处于初始化状态。 md5:60d100e8fd5c135d
# <翻译结束>


<原文开始>
// filter is bson.M{}, which means to match all, will return all records in the collection
<原文结束>

# <翻译开始>
// filter 是 bson.M{}，这意味着匹配所有，会返回集合中的所有记录 md5:c0c66af96a433502
# <翻译结束>


<原文开始>
	// res is a parseable object, but the bson tag is inconsistent with the mongodb record, and no error is reported
	// The corresponding value will be mapped according to the bson tag of the res data structure, and the tag without the value will be the default value of the corresponding type
	// The length of res is the number of records filtered by the filter condition
<原文结束>

# <翻译开始>
// res 是一个可解析的对象，但其 bson 标签与 mongodb 记录不一致，且不会报告错误
// 将根据 res 数据结构的 bson 标签映射相应的值，没有值的标签将使用对应类型的默认值
// res 的长度表示过滤条件筛选出的记录数
// md5:fa2c9312a213eab9
# <翻译结束>


<原文开始>
// filter can match records, skip 1 record, and return the remaining records
<原文结束>

# <翻译开始>
// filter 可以匹配记录，跳过一条记录，并返回剩余的记录。 md5:b966e759fac20d97
# <翻译结束>


<原文开始>
// filter can match the records, the number of skips is greater than the total number of existing records, res returns empty
<原文结束>

# <翻译开始>
// filter 可以匹配记录，跳过的数量大于现有记录的总数时，res 返回空 md5:d4411346be877b9e
# <翻译结束>


<原文开始>
// Sort a single field in ascending order
<原文结束>

# <翻译开始>
// 按升序对单个字段进行排序 md5:cb85098a3b639ea3
# <翻译结束>


<原文开始>
// Sort a single field in descending order
<原文结束>

# <翻译开始>
// 以降序对单个字段进行排序 md5:e53fe948db01b8ef
# <翻译结束>


<原文开始>
// Sort a single field in descending order, and sort the other field in ascending order
<原文结束>

# <翻译开始>
// 以降序对单个字段进行排序，然后按升序对其他字段进行排序 md5:8dec6f6f1880f356
# <翻译结束>


<原文开始>
// fields is empty, does not panic or error (#128)
<原文结束>

# <翻译开始>
// fields为空，不会引发恐慌或错误（#128） md5:65471cfbb3cddea4
# <翻译结束>


<原文开始>
	// different behavior with different version of mongod, v4.4.0 return err and v4.0.19 return nil
	//var res6 []int32
	//err = cli.Find(context.Background(), filter2).Distinct("", &res6)
	//ast.Error(err) // (Location40352) FieldPath cannot be constructed with empty string
	//ast.Equal(0, len(res6))
<原文结束>

# <翻译开始>
// 对于不同版本的mongod（如v4.4.0和v4.0.19），行为有所不同：v4.4.0会返回错误，而v4.0.19则可能返回nil
// 不使用res6
// _, err = cli.Find(context.Background(), filter2).Distinct("", &res6)
// 如果err非nil，则打印错误信息：(Location40352) FieldPath不能使用空字符串构建
// 验证res6的长度为0
// md5:db8b4089027d21a0
# <翻译结束>

