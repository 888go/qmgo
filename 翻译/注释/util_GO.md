
<原文开始>
// Now return Millisecond current time
<原文结束>

# <翻译开始>
// 现在返回当前时间的毫秒级别
# <翻译结束>


<原文开始>
// NewObjectID generates a new ObjectID.
// Watch out: the way it generates objectID is different from mgo
<原文结束>

# <翻译开始>
// NewObjectID 生成一个新的 ObjectID。
// 注意：它生成 objectID 的方式与 mgo 不同。
# <翻译结束>


<原文开始>
// CompareVersions compares two version number strings (i.e. positive integers separated by
// periods). Comparisons are done to the lesser precision of the two versions. For example, 3.2 is
// considered equal to 3.2.11, whereas 3.2.0 is considered less than 3.2.11.
//
// Returns a positive int if version1 is greater than version2, a negative int if version1 is less
// than version2, and 0 if version1 is equal to version2.
<原文结束>

# <翻译开始>
// CompareVersions 比较两个版本号字符串（即由点分隔的正整数）。比较操作会以两者中精度较低的那个为准。例如，3.2 被视为等于 3.2.11，而 3.2.0 则被视为小于 3.2.11。
//
// 如果 version1 大于 version2，则返回一个正整数；如果 version1 小于 version2，则返回一个负整数；如果 version1 等于 version2，则返回 0。
# <翻译结束>

