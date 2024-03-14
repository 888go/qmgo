package mgo类

import (
	"context"
	
	"testing"
	
	"github.com/888go/qmgo/operator"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBulk(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭(context.Background())
	defer cli.X删除集合(context.Background())

	id := primitive.NewObjectID()
	lucas := UserInfo{Id: primitive.NewObjectID(), Name: "Lucas", Age: 12}
	alias := UserInfo{Id: id, Name: "Alias", Age: 21}
	jess := UserInfo{Id: primitive.NewObjectID(), Name: "Jess", Age: 22}
	joe := UserInfo{Id: primitive.NewObjectID(), Name: "Joe", Age: 22}
	ethanId := primitive.NewObjectID()
	ethan := UserInfo{Id: ethanId, Name: "Ethan", Age: 8}

	result, err := cli.X创建批量执行().
		X插入(lucas).X插入(alias).X插入(jess).
		X更新一条(bson.M{"name": "Jess"}, bson.M{operator.Set: bson.M{"age": 23}}).X更新并按ID(id, bson.M{operator.Set: bson.M{"age": 23}}).
		X更新(bson.M{"age": 23}, bson.M{operator.Set: bson.M{"age": 18}}).
		X更新或插入(bson.M{"age": 17}, joe).X更新或插入并按ID(ethanId, ethan).
		X删除一条(bson.M{"name": "Joe"}).X删除并按ID(ethanId).X删除(bson.M{"age": 18}).
		X执行(context.Background())
	ast.NoError(err)
	ast.Equal(int64(3), result.InsertedCount)
	ast.Equal(int64(4), result.ModifiedCount)
	ast.Equal(int64(4), result.DeletedCount)
	ast.Equal(int64(2), result.UpsertedCount)
	ast.Equal(2, len(result.UpsertedIDs))
	ast.Equal(int64(4), result.MatchedCount)

}

func TestBulkUpsertOne(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭(context.Background())
	defer cli.X删除集合(context.Background())

	result, err := cli.X创建批量执行().
		X更新或插入一条(bson.M{"name": "Jess"}, bson.M{operator.Set: bson.M{"age": 20}, operator.SetOnInsert: bson.M{"weight": 40}}).
		X更新或插入一条(bson.M{"name": "Jess"}, bson.M{operator.Set: bson.M{"age": 30}, operator.SetOnInsert: bson.M{"weight": 40}}).
		X执行(context.Background())

	ast.NoError(err)
	ast.Equal(int64(0), result.InsertedCount)
	ast.Equal(int64(1), result.ModifiedCount)
	ast.Equal(int64(0), result.DeletedCount)
	ast.Equal(int64(1), result.UpsertedCount)
	ast.Equal(1, len(result.UpsertedIDs))
	ast.Equal(int64(1), result.MatchedCount)
}
