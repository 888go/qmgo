
<原文开始>
// Next gets the next document for this cursor. It returns true if there were no errors and the cursor has not been
// exhausted.
<原文结束>

# <翻译开始>
// Next 获取此游标下的下一个文档。如果未发生错误且游标未耗尽，它将返回true。
// md5:29446221269baaee
# <翻译结束>


<原文开始>
// All iterates the cursor and decodes each document into results. The results parameter must be a pointer to a slice.
// recommend to use All() in struct Query or Aggregate
<原文结束>

# <翻译开始>
// All 使用游标遍历每个文档，并将其解码到结果中。results 参数必须是指向切片的指针。
// 建议在 struct Query 或 Aggregate 中使用 All() 方法。
// md5:283225edc771266b
# <翻译结束>


<原文开始>
// ID returns the ID of this cursor, or 0 if the cursor has been closed or exhausted.
//func (c *Cursor) ID() int64 {
//	if c.err != nil {
//		return 0
//	}
//	return c.cursor.ID()
//}
<原文结束>

# <翻译开始>
// ID 返回游标ID，如果游标已关闭或耗尽，则返回0。
//func (c *Cursor) ID() int64 {
// 如果c.err不为nil，则返回0
// 否则返回游标c.cursor的ID
//}
// md5:bfd41b068bf5e581
# <翻译结束>


<原文开始>
// Close closes this cursor. Next and TryNext must not be called after Close has been called.
// When the cursor object is no longer in use, it should be actively closed
<原文结束>

# <翻译开始>
// Close 关闭这个游标。在调用 Close 之后，不应再调用 Next 或 TryNext。
// 当游标对象不再使用时，应主动关闭它。
// md5:7c67b9468038ed61
# <翻译结束>


<原文开始>
// Err return the last error of Cursor, if no error occurs, return nil
<原文结束>

# <翻译开始>
// Err 返回Cursor的最后一个错误，如果没有发生错误，则返回nil md5:2ebbf5e5b4796f72
# <翻译结束>

