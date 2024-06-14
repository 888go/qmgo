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
	"context"
	"testing"
	"time"

	"github.com/888go/qmgo/operator"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	X默认字段名称 `bson:",inline"`

	Name         string             `bson:"name"`
	Age          int                `bson:"age"`
	CreateTimeAt time.Time          `bson:"createTimeAt"`
	UpdateTimeAt int64              `bson:"updateTimeAt"`
	MyId         primitive.ObjectID `bson:"myId"`
}

func (u *User) CustomFields() CustomFieldsBuilder {
	return NewCustom().X设置创建时间字段名("CreateTimeAt").X设置更新时间字段名("UpdateTimeAt").X设置ID字段名("MyId")
}

func TestBeforeInsert(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()

	u := &User{Name: "Lucas", Age: 7}
	err := Do(ctx, u, mgo常量.X钩子_插入前)
	ast.NoError(err)
	// default fields
	ast.NotEqual(time.Time{}, u.X创建时间)
	ast.NotEqual(time.Time{}, u.X更新时间)
	ast.NotEqual(primitive.NilObjectID, u.Id)
	// custom fields
	ast.NotEqual(time.Time{}, u.CreateTimeAt)
	ast.NotEqual(time.Time{}, u.UpdateTimeAt)
	ast.NotEqual(primitive.NilObjectID, u.MyId)

	u1, u2 := &User{Name: "Lucas", Age: 7}, &User{Name: "Alice", Age: 8}
	us := []*User{u1, u2}
	err = Do(ctx, us, mgo常量.X钩子_插入前)
	ast.NoError(err)

	for _, v := range us {
		ast.NotEqual(time.Time{}, v.X创建时间)
		ast.NotEqual(time.Time{}, v.X更新时间)
		ast.NotEqual(primitive.NilObjectID, v.Id)
	}

	u3 := User{Name: "Lucas", Age: 7}
	err = Do(ctx, u3, mgo常量.X钩子_插入前)
	ast.NoError(err)

	// 使用有效值插入 md5:c68783e3102ab2a6
	tBefore3s := time.Now().Add(-3 * time.Second)
	id := primitive.NewObjectID()
	u = &User{Name: "Lucas", Age: 7}
	u.X创建时间 = tBefore3s
	u.X更新时间 = tBefore3s
	u.Id = id
	u.MyId = id
	u.CreateTimeAt = tBefore3s
	u.UpdateTimeAt = tBefore3s.Unix()

	err = Do(ctx, u, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

	ast.Equal(tBefore3s, u.X创建时间)
	ast.Equal(id, u.Id)
	ast.NotEqual(tBefore3s, u.X更新时间)

	ast.Equal(tBefore3s, u.CreateTimeAt)
	ast.Equal(id, u.MyId)
	ast.NotEqual(tBefore3s.Unix(), u.UpdateTimeAt)
}

func TestBeforeUpdate(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()

	u := &User{Name: "Lucas", Age: 7}
	err := Do(ctx, u, mgo常量.X钩子_更新前)
	ast.NoError(err)
	// default field
	ast.NotEqual(time.Time{}, u.X更新时间)

	// custom fields
	ast.NotEqual(time.Time{}, u.UpdateTimeAt)

	u1, u2 := &User{Name: "Lucas", Age: 7}, &User{Name: "Alice", Age: 8}
	us := []*User{u1, u2}
	err = Do(ctx, us, mgo常量.X钩子_更新前)
	ast.NoError(err)
	for _, v := range us {
		// default field
		ast.NotEqual(time.Time{}, v.X更新时间)

		// custom fields
		ast.NotEqual(time.Time{}, v.UpdateTimeAt)
	}

	us1 := []interface{}{u1, u2}
	err = Do(ctx, us1, mgo常量.X钩子_更新前)
	ast.NoError(err)
	for _, v := range us {
		// default field
		ast.NotEqual(time.Time{}, v.X更新时间)

		// custom fields
		ast.NotEqual(time.Time{}, v.UpdateTimeAt)
	}

}

type UserField struct {
	X默认字段名称 `bson:",inline"`

	Name         string             `bson:"name"`
	Age          int                `bson:"age"`
	CreateTimeAt int64              `bson:"createTimeAt"`
	UpdateTimeAt time.Time          `bson:"updateTimeAt"`
	MyId         primitive.ObjectID `bson:"myId"`
}

func (u *UserField) CustomFields() CustomFieldsBuilder {
	return NewCustom().X设置创建时间字段名("CreateTimeAt").X设置更新时间字段名("UpdateTimeAt").X设置ID字段名("MyId")
}

func TestBeforeUpsert(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()

	// with empty fields
	u := &User{Name: "Lucas", Age: 7}
	err := Do(ctx, u, mgo常量.X钩子_替换插入前)
	ast.NoError(err)
	// default fields
	ast.NotEqual(time.Time{}, u.X创建时间)
	ast.NotEqual(time.Time{}, u.X更新时间)
	ast.NotEqual(primitive.NilObjectID, u.Id)
	// custom fields
	ast.NotEqual(time.Time{}, u.CreateTimeAt)
	ast.NotEqual(0, u.UpdateTimeAt)
	ast.NotEqual(primitive.NilObjectID, u.MyId)

	u1, u2 := &User{Name: "Lucas", Age: 7}, &User{Name: "Alice", Age: 8}
	us := []*User{u1, u2}
	err = Do(ctx, us, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

	for _, v := range us {
		ast.NotEqual(time.Time{}, v.X创建时间)
		ast.NotEqual(time.Time{}, v.X更新时间)
		ast.NotEqual(time.Time{}, u.CreateTimeAt)
		ast.NotEqual(0, u.UpdateTimeAt)
		ast.NotEqual(primitive.NilObjectID, v.Id)
	}

	u3 := User{Name: "Lucas", Age: 7}
	err = Do(ctx, u3, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

	// 使用有效值进行更新插入操作 md5:6dcc17bfca098279
	tBefore3s := time.Now().Add(-3 * time.Second)
	id := primitive.NewObjectID()
	u = &User{Name: "Lucas", Age: 7}
	u.X创建时间 = tBefore3s
	u.X更新时间 = tBefore3s
	u.Id = id
	u.MyId = id
	u.CreateTimeAt = tBefore3s
	u.UpdateTimeAt = tBefore3s.Unix()

	err = Do(ctx, u, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

	ast.Equal(tBefore3s, u.X创建时间)
	ast.Equal(id, u.Id)
	ast.NotEqual(tBefore3s, u.X更新时间)

	ast.Equal(tBefore3s, u.CreateTimeAt)
	ast.Equal(id, u.MyId)
	ast.NotEqual(tBefore3s.Unix(), u.UpdateTimeAt)

}

// 与 TestBeforeUpsert 相同，只是将 CreateTimeAt 和 UpdateTimeAt 的类型进行了切换 md5:90f4c0b23d59b403
func TestBeforeUpsertUserFiled(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()

	// with empty fileds
	u := &UserField{Name: "Lucas", Age: 7}
	err := Do(ctx, u, mgo常量.X钩子_替换插入前)
	ast.NoError(err)
	// default fields
	ast.NotEqual(time.Time{}, u.X创建时间)
	ast.NotEqual(time.Time{}, u.X更新时间)
	ast.NotEqual(primitive.NilObjectID, u.Id)
	// custom fields
	ast.NotEqual(0, u.CreateTimeAt)
	ast.NotEqual(time.Time{}, u.UpdateTimeAt)
	ast.NotEqual(primitive.NilObjectID, u.MyId)

	u1, u2 := &UserField{Name: "Lucas", Age: 7}, &UserField{Name: "Alice", Age: 8}
	us := []*UserField{u1, u2}
	err = Do(ctx, us, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

	for _, v := range us {
		ast.NotEqual(time.Time{}, v.X创建时间)
		ast.NotEqual(time.Time{}, v.X更新时间)
		ast.NotEqual(0, u.CreateTimeAt)
		ast.NotEqual(time.Time{}, u.UpdateTimeAt)
		ast.NotEqual(primitive.NilObjectID, v.Id)
	}

	u3 := User{Name: "Lucas", Age: 7}
	err = Do(ctx, u3, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

	// 使用有效值进行更新插入操作 md5:6dcc17bfca098279
	tBefore3s := time.Now().Add(-3 * time.Second)
	id := primitive.NewObjectID()
	u = &UserField{Name: "Lucas", Age: 7}
	u.X创建时间 = tBefore3s
	u.X更新时间 = tBefore3s
	u.Id = id
	u.MyId = id
	u.CreateTimeAt = tBefore3s.Unix()
	u.UpdateTimeAt = tBefore3s

	err = Do(ctx, u, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

	ast.Equal(tBefore3s, u.X创建时间)
	ast.Equal(id, u.Id)
	ast.NotEqual(tBefore3s, u.X更新时间)

	ast.NotEqual(tBefore3s, u.UpdateTimeAt)
	ast.Equal(id, u.MyId)
	ast.Equal(tBefore3s.Unix(), u.CreateTimeAt)

}

func TestNilError(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()

	err := Do(ctx, nil, mgo常量.X钩子_替换插入前)
	ast.NoError(err)

}
