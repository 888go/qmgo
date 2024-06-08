/*
 Copyright 2020 The Qmgo Authors.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package operator

// 聚合管道运算符
// 参考：https://docs.mongodb.com/manual/reference/operator/aggregation/
// md5:b743948494162fab
// [提示]
//const (
// 	绝对值      = "$abs"
// 	加法        = "$add"
// 	天花板      = "$ceil"
// 	除法        = "$divide"
// 	指数        = "$exp"
// 	地板        = "$floor"
// 	自然对数     = "$ln"
// 	对数        = "$log"
// 	以10为底的对数 = "$log10"
// 	乘法        = "$multiply"
// 	次方        = "$pow"
// 	四舍五入     = "$round"
// 	平方根      = "$sqrt"
// 	减法        = "$subtract"
// 	截断        = "$trunc"
// 
// 	数组元素At   = "$arrayElemAt"
// 	数组转对象   = "$arrayToObject"
// 	连接数组    = "$concatArrays"
// 	过滤        = "$filter"
// 	数组索引Of   = "$indexOfArray"
// 	是否为数组   = "$isArray"
// 	映射         = "$map"
// 	对象转数组   = "$objectToArray"
// 	范围         = "$range"
// 	归约        = "$reduce"
// 	反转数组     = "$reverseArray"
// 	合并         = "$zip"
// 
// 	比较        = "$cmp"
// 
// 	条件        = "$cond"
// 	如果为Null   = "$ifNull"
// 	开关        = "$switch"
// 
// 	累加器     = "$accumulator"
// 	函数        = "$function"
// 
// 	二进制大小   = "$binarySize"
// 	BSON大小     = "$bsonSize"
// 
// 	日期从部件   = "$dateFromParts"
// 	日期从字符串 = "$dateFromString"
// 	日期转部件   = "$dateToParts"
// 	日期转字符串 = "$dateToString"
// 	月份中的天数   = "$dayOfMonth"
// 	星期中的天数   = "$dayOfWeek"
// 	年中的天数     = "$dayOfYear"
// 	小时         = "$hour"
// 	ISO星期中的天数 = "$isoDayOfWeek"
// 	ISO周        = "$isoWeek"
// 	ISO周年的年份  = "$isoWeekYear"
// 	毫秒         = "$millisecond"
// 	分钟         = "$minute"
// 	月份         = "$month"
// 	秒           = "$second"
// 	日期转日期   = "$toDate"
// 	周           = "$week"
// 	年           = "$year"
// 
// 	字面量       = "$literal"
// 
// 	合并对象     = "$mergeObjects"
// 
// 	所有元素为真 = "$allElementsTrue"
// 	任何元素为真   = "$anyElementTrue"
// 	集合差集     = "$setDifference"
// 	集合相等      = "$setEquals"
// 	集合交集     = "$setIntersection"
// 	集合是子集    = "$setIsSubset"
// 	集合并集     = "$setUnion"
// 
// 	连接        = "$concat"
// 	字节索引Of   = "$indexOfBytes"
// 	字符索引Of   = "$indexOfCP"
// 	左截断      = "$ltrim"
// 	正则查找     = "$regexFind"
// 	正则查找所有 = "$regexFindAll"
// 	正则匹配     = "$regexMatch"
// 	右截断      = "$rtrim"
// 	分割        = "$split"
// 	字节长度     = "$strLenBytes"
// 	字符长度     = "$strLenCP"
// 	不区分大小写比较 = "$strcasecmp"
// 	子串         = "$substr"
// 	字节子串     = "$substrBytes"
// 	字符子串     = "$substrCP"
// 	转换为小写   = "$toLower"
// 	转换为字符串 = "$toString"
// 	修剪        = "$trim"
// 	转换为大写   = "$toUpper"
// 	替换一次     = "$replaceOne"
// 	全部替换     = "$replaceAll"
// 
// 	正弦        = "$sin"
// 	余弦        = "$cos"
// 	正切        = "$tan"
// 	反正弦      = "$asin"
// 	反余弦      = "$acos"
// 	反正切      = "$atan"
// 	反正切2      = "$atan2"
// 	双曲正弦     = "$asinh"
// 	双曲余弦     = "$acosh"
// 	双曲正切     = "$atanh"
// 	度转弧度     = "$degreesToRadians"
// 	弧度转度     = "$radiansToDegrees"
// 
// 	转换       = "$convert"
// 	转换为布尔   = "$toBool"
// 	转换为十进制 = "$toDecimal"
// 	转换为浮点数 = "$toDouble"
// 	转换为整数   = "$toInt"
// 	转换为长整数 = "$toLong"
// 	转换为ObjectId = "$toObjectId"
// 	是否为数字   = "$isNumber"
// 
// 	平均值       = "$avg"
// 	第一个       = "$first"
// 	最后一个     = "$last"
// 
// 	样本标准差   = "$stdDevPop"
// 	总体标准差   = "$stdDevSamp"
// 	求和         = "$sum"
// 
// 	局部变量     = "$let"
// )
// [结束]
const (
	// 算术运算符 md5:862d53c43dccda12
	Abs      = "$abs"
	Add      = "$add"
	Ceil     = "$ceil"
	Divide   = "$divide"
	Exp      = "$exp"
	Floor    = "$floor"
	Ln       = "$ln"
	Log      = "$log"
	Log10    = "$log10"
	Multiply = "$multiply"
	Pow      = "$pow"
	Round    = "$round"
	Sqrt     = "$sqrt"
	Subtract = "$subtract"
	Trunc    = "$trunc"

	// 数组表达式运算符 md5:ec78a0b8bc2385f3
	ArrayElemAt   = "$arrayElemAt"
	ArrayToObject = "$arrayToObject"
	ConcatArrays  = "$concatArrays"
	Filter        = "$filter"
	IndexOfArray  = "$indexOfArray"
	IsArray       = "$isArray"
	Map           = "$map"
	ObjectToArray = "$objectToArray"
	Range         = "$range"
	Reduce        = "$reduce"
	ReverseArray  = "$reverseArray"
	Zip           = "$zip"

	// 比较表达式运算符 md5:c44b496f05f62dec
	Cmp = "$cmp"

	// 条件表达式运算符 md5:5dd78ba4ae108ba1
	Cond   = "$cond"
	IfNull = "$ifNull"
	Switch = "$switch"

	// 自定义聚合表达式运算符 md5:c2e3d7a49a64bbbf
	Accumulator = "$accumulator"
	Function    = "$function"

	// Data Size Operators
	BinarySize = "$binarySize"
	BsonSize   = "$bsonSize"

	// 日期表达式运算符 md5:9109a9260ffcf4b1
	DateFromParts  = "$dateFromParts"
	DateFromString = "$dateFromString"
	DateToParts    = "$dateToParts"
	DateToString   = "$dateToString"
	DayOfMonth     = "$dayOfMonth"
	DayOfWeek      = "$dayOfWeek"
	DayOfYear      = "$dayOfYear"
	Hour           = "$hour"//qm:小时  cz:Hour = "$hour"  
	IsoDayOfWeek   = "$isoDayOfWeek"
	IsoWeek        = "$isoWeek"
	IsoWeekYear    = "$isoWeekYear"
	Millisecond    = "$millisecond"
	Minute         = "$minute"
	Month          = "$month"//qm:月  cz:Month = "$month"  
	Second         = "$second"//qm:秒  cz:Second = "$second"  
	ToDate         = "$toDate"
	Week           = "$week"//qm:周  cz:Week = "$week"  
	Year           = "$year"//qm:年  cz:Year = "$year"  

	// 字面量表达式操作符 md5:8501f5c82ee2c883
	Literal = "$literal"

	// 对象表达式运算符 md5:b6f2383d804984d9
	MergeObjects = "$mergeObjects"

	// 设置表达式运算符 md5:e745772d17491d89
	AllElementsTrue = "$allElementsTrue"
	AnyElementTrue  = "$anyElementTrue"
	SetDifference   = "$setDifference"
	SetEquals       = "$setEquals"
	SetIntersection = "$setIntersection"
	SetIsSubset     = "$setIsSubset"
	SetUnion        = "$setUnion"

	// 字符串表达式操作符 md5:eb3461a712d14fac
	Concat       = "$concat"
	IndexOfBytes = "$indexOfBytes"
	IndexOfCP    = "$indexOfCP"
	Ltrim        = "$ltrim"
	RegexFind    = "$regexFind"
	RegexFindAll = "$regexFindAll"
	RegexMatch   = "$regexMatch"
	Rtrim        = "$rtrim"
	Split        = "$split"
	StrLenBytes  = "$strLenBytes"
	StrLenCP     = "$strLenCP"
	Strcasecmp   = "$strcasecmp"
	Substr       = "$substr"
	SubstrBytes  = "$substrBytes"
	SubstrCP     = "$substrCP"
	ToLower      = "$toLower"
	ToString     = "$toString"
	Trim         = "$trim"
	ToUpper      = "$toUpper"
	ReplaceOne   = "$replaceOne"
	ReplaceAll   = "$replaceAll"

	// 三角函数表达式运算符 md5:c83fcb6cf067b355
	Sin              = "$sin"
	Cos              = "$cos"
	Tan              = "$tan"
	Asin             = "$asin"
	Acos             = "$acos"
	Atan             = "$atan"
	Atan2            = "$atan2"
	Asinh            = "$asinh"
	Acosh            = "$acosh"
	Atanh            = "$atanh"
	DegreesToRadians = "$degreesToRadians"
	RadiansToDegrees = "$radiansToDegrees"

	// 类型表达式运算符 md5:021bcd14e2b085e1
	Convert    = "$convert"
	ToBool     = "$toBool"
	ToDecimal  = "$toDecimal"
	ToDouble   = "$toDouble"
	ToInt      = "$toInt"
	ToLong     = "$toLong"
	ToObjectID = "$toObjectId"
	IsNumber   = "$isNumber"

	// Accumulators ($group)
	Avg   = "$avg"//qm:平均值  cz:Avg = "$avg"  
	First = "$first"
	Last  = "$last"

	StdDevPop  = "$stdDevPop"
	StdDevSamp = "$stdDevSamp"
	Sum        = "$sum"//qm:求和  cz:Sum = "$sum"  

	// 变量表达式操作符 md5:b84290815043dc99
	Let = "$let"
)
