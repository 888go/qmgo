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

package field

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDefaultField(t *testing.T) {
	ast := require.New(t)

	df := &DefaultField{}
	df.X默认创建时间()
	df.X默认更新时间()
	df.X默认ID()
	ast.NotEqual(time.Time{}, df.X更新时间)
	ast.NotEqual(time.Time{}, df.X创建时间)
	ast.NotEqual(primitive.NilObjectID, df.Id)
}
