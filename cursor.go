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

// X下一个 gets the next document for this cursor. It returns true if there were no errors and the cursor has not been
// exhausted.
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

// X取全部 iterates the cursor and decodes each document into results. The results parameter must be a pointer to a slice.
// recommend to use X取全部() in struct Query or Aggregate
func (c *Cursor) X取全部(results interface{}) error {
	if c.err != nil {
		return c.err
	}
	return c.cursor.All(c.ctx, results)
}

// ID returns the ID of this cursor, or 0 if the cursor has been closed or exhausted.
//func (c *Cursor) ID() int64 {
//	if c.err != nil {
//		return 0
//	}
//	return c.cursor.ID()
//}

// X关闭 closes this cursor. Next and TryNext must not be called after X关闭 has been called.
// When the cursor object is no longer in use, it should be actively closed
func (c *Cursor) X关闭() error {
	if c.err != nil {
		return c.err
	}
	return c.cursor.Close(c.ctx)
}

// X取错误 return the last error of Cursor, if no error occurs, return nil
func (c *Cursor) X取错误() error {
	if c.err != nil {
		return c.err
	}
	return c.cursor.Err()
}
