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
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/qmgo/field"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserField struct {
	field.DefaultField `bson:",inline"`

	Name string `bson:"name"`
	Age  int    `bson:"age"`

	MyId         string    `bson:"myId"`
	CreateTimeAt time.Time `bson:"createTimeAt"`
	UpdateTimeAt int64     `bson:"updateTimeAt"`
}

func (u *UserField) CustomFields() field.CustomFieldsBuilder {
	return field.NewCustom().X设置创建时间字段名("CreateTimeAt").X设置更新时间字段名("UpdateTimeAt").X设置ID字段名("MyId")
}

func TestFieldInsert(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	ctx := context.Background()
	defer cli.X关闭(ctx)
	defer cli.X删除集合(ctx)

	u := &UserField{Name: "Lucas", Age: 7}
	_, err := cli.X插入(context.Background(), u)
	ast.NoError(err)

	uc := bson.M{"name": "Lucas"}
	ur := &UserField{}
	err = cli.X查询(ctx, uc).X取一条(ur)
	ast.NoError(err)

	// default fields
	ast.NotEqual(time.Time{}, ur.CreateAt)
	ast.NotEqual(time.Time{}, ur.UpdateAt)
	ast.NotEqual(primitive.NilObjectID, ur.Id)
	// custom fields
	ast.NotEqual(time.Time{}, ur.CreateTimeAt)
	ast.NotEqual(int64(0), ur.UpdateTimeAt)
	ast.NotEqual("", ur.MyId)

}

func TestFieldInsertMany(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	ctx := context.Background()
	defer cli.X关闭(ctx)
	defer cli.X删除集合(ctx)

	u1 := &UserField{Name: "Lucas", Age: 7}
	u2 := &UserField{Name: "Alice", Age: 7}
	us := []*UserField{u1, u2}
	_, err := cli.X插入多个(ctx, us)
	ast.NoError(err)

	uc := bson.M{"age": 7}
	ur := []UserField{}
	err = cli.X查询(ctx, uc).X取全部(&ur)
	ast.NoError(err)

	// default fields
	ast.NotEqual(time.Time{}, ur[0].CreateAt)
	ast.NotEqual(time.Time{}, ur[0].UpdateAt)
	ast.NotEqual(primitive.NilObjectID, ur[0].Id)
	// default fields
	ast.NotEqual(time.Time{}, ur[1].CreateAt)
	ast.NotEqual(time.Time{}, ur[1].UpdateAt)
	ast.NotEqual(primitive.NilObjectID, ur[1].Id)

	// custom fields
	ast.NotEqual(time.Time{}, ur[0].CreateTimeAt)
	ast.NotEqual(int64(0), ur[0].UpdateTimeAt)
}

func TestFieldUpdate(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	ui := &UserField{Name: "Lucas", Age: 17}
	_, err := cli.X插入(context.Background(), ui)
	ast.NoError(err)

	err = cli.X更新一条(context.Background(), bson.M{"name": "Lucas"}, bson.M{"$set": bson.M{"updateTimeAt": 0, "updateAt": time.Time{}}})
	ast.NoError(err)

	findUi := UserField{}
	err = cli.X查询(context.Background(), bson.M{"name": "Lucas"}).X取一条(&findUi)
	ast.Equal(int64(0), findUi.UpdateTimeAt)
	ast.Equal(time.Time{}, findUi.UpdateAt)

	ast.NoError(err)
	ui.Id = findUi.Id
	err = cli.X替换一条(context.Background(), bson.M{"_id": findUi.Id}, &ui)
	ast.NoError(err)
	err = cli.X查询(context.Background(), bson.M{"name": "Lucas"}).X取一条(&findUi)
	ast.NotEqual(int64(0), findUi.UpdateTimeAt)
	ast.NotEqual(time.Time{}, findUi.UpdateAt)

}

func TestFieldUpsert(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	ctx := context.Background()
	defer cli.X关闭(ctx)
	defer cli.X删除集合(ctx)

	u := &UserField{Name: "Lucas", Age: 7}
	id := primitive.NewObjectID()
	u.Id = id
	id_1 := primitive.NewObjectID()
	u.MyId = id_1.String()
	_, err := cli.X插入(context.Background(), u)
	ast.NoError(err)

	time.Sleep(2 * time.Second)
	u.Age = 17
	tBefore3s := time.Now().Add(-3 * time.Second).Local()
	u.CreateAt = tBefore3s
	u.UpdateAt = tBefore3s
	u.CreateTimeAt = tBefore3s
	u.UpdateTimeAt = tBefore3s.Unix()
	result, err := cli.X更新或插入(ctx, bson.M{"_id": id}, u)
	ast.NoError(err)
	fmt.Println(result)

	ui := UserField{}
	err = cli.X查询(ctx, bson.M{"_id": id}).X取一条(&ui)

	ast.NoError(err)
	ast.Equal(u.Age, ui.Age)
	ast.Equal(id, ui.Id)
	ast.Equal(id_1.String(), ui.MyId)
	fmt.Println(tBefore3s.Unix(), ui.CreateAt.Unix())
	ast.Equal(tBefore3s.Unix(), ui.CreateAt.Unix())
	ast.Equal(tBefore3s.Unix(), ui.CreateTimeAt.Unix())
	ast.NotEqual(tBefore3s.Unix(), ui.UpdateAt.Unix())
	ast.NotEqual(tBefore3s.Unix(), ui.UpdateTimeAt)

}

func TestFieldUpsertId(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	ctx := context.Background()
	defer cli.X关闭(ctx)
	defer cli.X删除集合(ctx)

	u := &UserField{Name: "Lucas", Age: 7}
	id := primitive.NewObjectID()
	u.Id = id
	id_1 := primitive.NewObjectID()
	u.MyId = id_1.String()
	_, err := cli.X插入(context.Background(), u)
	ast.NoError(err)

	time.Sleep(2 * time.Second)
	u.Age = 17
	tBefore3s := time.Now().Add(-3 * time.Second).Local()
	u.CreateAt = tBefore3s
	u.UpdateAt = tBefore3s
	u.CreateTimeAt = tBefore3s
	u.UpdateTimeAt = tBefore3s.Unix()
	_, err = cli.X更新或插入并按ID(ctx, id, u)
	ast.NoError(err)

	ui := UserField{}
	err = cli.X查询(ctx, bson.M{"_id": id}).X取一条(&ui)

	ast.NoError(err)
	ast.Equal(u.Age, ui.Age)
	ast.Equal(id, ui.Id)
	ast.Equal(id_1.String(), ui.MyId)
	ast.Equal(tBefore3s.Unix(), ui.CreateAt.Unix())
	ast.Equal(tBefore3s.Unix(), ui.CreateTimeAt.Unix())
	ast.NotEqual(tBefore3s.Unix(), ui.UpdateAt.Unix())
	ast.NotEqual(tBefore3s.Unix(), ui.UpdateTimeAt)
}

func TestFieldUpdateId(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭(context.Background())
	defer cli.X删除集合(context.Background())
	cli.EnsureIndexes弃用(context.Background(), []string{"name"}, nil)

	ui := &UserField{Name: "Lucas", Age: 17}
	res, err := cli.X插入(context.Background(), ui)
	ast.NoError(err)

	err = cli.X更新并按ID(context.Background(), res.InsertedID, bson.M{"$set": bson.M{"updateTimeAt": 0, "updateAt": time.Time{}}})
	ast.NoError(err)

	findUi := UserField{}
	err = cli.X查询(context.Background(), bson.M{"name": "Lucas"}).X取一条(&findUi)
	ast.NoError(err)
	ast.Equal(int64(0), findUi.UpdateTimeAt)
	ast.Equal(time.Time{}, findUi.UpdateAt)
}
