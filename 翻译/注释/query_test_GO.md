
<原文开始>
// res is nil or can't parse
<原文结束>

# <翻译开始>
// res 为空或者无法解析
# <翻译结束>


<原文开始>
// res is a parseable object, but the bson tag is inconsistent with the mongodb record, no error is reported, res is the initialization state of the data structure
<原文结束>

# <翻译开始>
// res 是一个可解析的对象，但其 bson 标签与 MongoDB 中的记录不一致，此时不会报告错误，res 保持着数据结构的初始状态
# <翻译结束>


<原文开始>
// filter is bson.M{}, which means to match all, will return all records in the collection
<原文结束>

# <翻译开始>
// filter 是 bson.M{} 类型，这意味着匹配所有条件，将会返回集合中的所有记录
# <翻译结束>


<原文开始>
	// res is a parseable object, but the bson tag is inconsistent with the mongodb record, and no error is reported
	// The corresponding value will be mapped according to the bson tag of the res data structure, and the tag without the value will be the default value of the corresponding type
	// The length of res is the number of records filtered by the filter condition
<原文结束>

# <翻译开始>
// res 是一个可解析的对象，但其 bson 标签与 MongoDB 中的记录不一致，并且不会报告错误
// 对应的值将根据 res 数据结构的 bson 标签进行映射，没有值的标签将会使用对应类型的默认值
// res 的长度是经过 filter 条件筛选出的记录数量
# <翻译结束>


<原文开始>
// filter can match records, skip 1 record, and return the remaining records
<原文结束>

# <翻译开始>
// filter 可以匹配记录，跳过 1 条记录，并返回剩余的记录
# <翻译结束>


<原文开始>
// filter can match the records, the number of skips is greater than the total number of existing records, res returns empty
<原文结束>

# <翻译开始>
// 如果filter能够匹配记录，且跳过的数量大于现有记录的总数，则res返回空结果集
# <翻译结束>


<原文开始>
// Sort a single field in ascending order
<原文结束>

# <翻译开始>
// 按升序对单个字段进行排序
# <翻译结束>


<原文开始>
// Sort a single field in descending order
<原文结束>

# <翻译开始>
// 按降序对单个字段进行排序
# <翻译结束>


<原文开始>
// Sort a single field in descending order, and sort the other field in ascending order
<原文结束>

# <翻译开始>
// 按照一个字段降序排序，同时其他字段按升序排序
# <翻译结束>


<原文开始>
// fields is empty, does not panic or error (#128)
<原文结束>

# <翻译开始>
// 当fields为空时，不会导致panic或error（#128）
# <翻译结束>


<原文开始>
	// different behavior with different version of mongod, v4.4.0 return err and v4.0.19 return nil
	//var res6 []int32
	//err = cli.Find(context.Background(), filter2).Distinct("", &res6)
	//ast.Error(err) // (Location40352) FieldPath cannot be constructed with empty string
	//ast.Equal(0, len(res6))
<原文结束>

# <翻译开始>
// 不同版本的mongod表现出不同的行为，v4.4.0返回错误，v4.0.19返回nil
// 定义一个int32类型的切片变量res6
// 使用cli调用Find方法并执行Distinct操作，在context.Background()环境下，根据filter2过滤条件，将查询结果存入res6中
// ast.Error(err) // (Location40352) 当字段路径为空字符串时，无法构建FieldPath，此处捕获到该错误
// ast.Equal(0, len(res6)) // 断言res6的结果长度为0
# <翻译结束>

