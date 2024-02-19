
<原文开始>
// Next gets the next document for this cursor. It returns true if there were no errors and the cursor has not been
// exhausted.
<原文结束>

# <翻译开始>
// Next 获取此游标的下一个文档。如果未发生错误且游标未被耗尽，则返回true。
# <翻译结束>


<原文开始>
// All iterates the cursor and decodes each document into results. The results parameter must be a pointer to a slice.
// recommend to use All() in struct Query or Aggregate
<原文结束>

# <翻译开始>
// All 方法遍历游标并对每个文档进行解码，将结果存入 results 参数。results 参数必须是指向切片的指针。
// 推荐在结构体 Query 或 Aggregate 中使用 All() 方法。
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
// ID返回该游标的ID，如果游标已关闭或耗尽，则返回0。
//func (c *Cursor) ID() int64 {
//	// 如果存在错误（即游标已关闭或耗尽），则返回0
//	if c.err != nil {
//		return 0
//	}
//	// 返回当前游标的ID
//	return c.cursor.ID()
//}
# <翻译结束>


<原文开始>
// Close closes this cursor. Next and TryNext must not be called after Close has been called.
// When the cursor object is no longer in use, it should be actively closed
<原文结束>

# <翻译开始>
// Close 关闭此游标。在调用 Close 后，不得再调用 Next 和 TryNext。
// 当游标对象不再使用时，应主动关闭它。
# <翻译结束>


<原文开始>
// Err return the last error of Cursor, if no error occurs, return nil
<原文结束>

# <翻译开始>
// Err 返回Cursor的最后一个错误，如果没有发生错误，则返回nil
# <翻译结束>

