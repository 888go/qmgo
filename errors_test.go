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
	"errors"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestIsErrNoDocuments(t *testing.T) {
	ast := require.New(t)
	ast.False(X是否为无文档错误(errors.New("dont match")))
	ast.True(X是否为无文档错误(ErrNoSuchDocuments))
	ast.True(X是否为无文档错误(mongo.ErrNoDocuments))
}

func TestIsDup(t *testing.T) {
	ast := require.New(t)
	ast.False(X是否为重复键错误(nil))
	ast.False(X是否为重复键错误(errors.New("invaliderror")))
	ast.True(X是否为重复键错误(errors.New("E11000")))
}
