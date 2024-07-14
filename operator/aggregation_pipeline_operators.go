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

// Aggregation Pipeline Operators
const (
	// 算术运算符 md5:862d53c43dccda12
	Abs      = "$abs" //qm:聚合绝对值  cz:Abs = "$abs"
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
	Hour           = "$hour" //qm:小时  cz:Hour = "$hour"
	IsoDayOfWeek   = "$isoDayOfWeek"
	IsoWeek        = "$isoWeek"
	IsoWeekYear    = "$isoWeekYear"
	Millisecond    = "$millisecond"
	Minute         = "$minute"
	Month          = "$month"  //qm:月  cz:Month = "$month"
	Second         = "$second" //qm:秒  cz:Second = "$second"
	ToDate         = "$toDate"
	Week           = "$week" //qm:周  cz:Week = "$week"
	Year           = "$year" //qm:年  cz:Year = "$year"

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
	Avg   = "$avg" //qm:平均值  cz:Avg = "$avg"
	First = "$first"
	Last  = "$last"

	StdDevPop  = "$stdDevPop"
	StdDevSamp = "$stdDevSamp"
	Sum        = "$sum" //qm:求和  cz:Sum = "$sum"

	// 变量表达式操作符 md5:b84290815043dc99
	Let = "$let"
)
