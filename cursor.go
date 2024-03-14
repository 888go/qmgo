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

package mgo类

import (
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
)

// Cursor struct define
type Cursor struct {
	ctx    context.Context
	cursor *mongo.Cursor
	err    error
}

// Next 获取此游标的下一个文档。如果未发生错误且游标未被耗尽，则返回true。
func (c *Cursor) X下一个(result interface{}) bool {
	if c.err != nil {
		return false
	}
	if c.cursor.Next(c.ctx) {
		err := c.cursor.Decode(result)
		if err != nil {
			c.err = err
			return false
		}
		return true
	}
	return false
}

// All 方法遍历游标并对每个文档进行解码，将结果存入 results 参数。results 参数必须是指向切片的指针。
// 推荐在结构体 Query 或 Aggregate 中使用 All() 方法。
func (c *Cursor) X取全部(results interface{}) error {
	if c.err != nil {
		return c.err
	}
	return c.cursor.All(c.ctx, results)
}

// ID返回该游标的ID，如果游标已关闭或耗尽，则返回0。
//func (c *Cursor) ID() int64 {
//	// 如果存在错误（即游标已关闭或耗尽），则返回0
//	if c.err != nil {
//		return 0
//	}
//	// 返回当前游标的ID
//	return c.cursor.ID()
//}

// Close 关闭此游标。在调用 Close 后，不得再调用 Next 和 TryNext。
// 当游标对象不再使用时，应主动关闭它。
func (c *Cursor) X关闭() error {
	if c.err != nil {
		return c.err
	}
	return c.cursor.Close(c.ctx)
}

// Err 返回Cursor的最后一个错误，如果没有发生错误，则返回nil
func (c *Cursor) X取错误() error {
	if c.err != nil {
		return c.err
	}
	return c.cursor.Err()
}
