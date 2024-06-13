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
	defer cli.X关闭连接(context.Background())
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
		X更新一条(bson.M{"name": "Jess"}, bson.M{操作符.X更新值: bson.M{"age": 23}}).X更新并按ID(id, bson.M{操作符.X更新值: bson.M{"age": 23}}).
		X更新(bson.M{"age": 23}, bson.M{操作符.X更新值: bson.M{"age": 18}}).
		X替换插入(bson.M{"age": 17}, joe).X替换插入并按ID(ethanId, ethan).
		X删除一条(bson.M{"name": "Joe"}).X删除并按ID(ethanId).X删除(bson.M{"age": 18}).
		X执行(context.Background())
	ast.NoError(err)
	ast.Equal(int64(3), result.X插入数)
	ast.Equal(int64(4), result.X修改数)
	ast.Equal(int64(4), result.X删除数)
	ast.Equal(int64(2), result.X替换插入数)
	ast.Equal(2, len(result.X替换插入IDs))
	ast.Equal(int64(4), result.X匹配数)

}

func TestBulkUpsertOne(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	defer cli.X关闭连接(context.Background())
	defer cli.X删除集合(context.Background())

	result, err := cli.X创建批量执行().
		X替换插入一条(bson.M{"name": "Jess"}, bson.M{操作符.X更新值: bson.M{"age": 20}, 操作符.X更新插入时: bson.M{"weight": 40}}).
		X替换插入一条(bson.M{"name": "Jess"}, bson.M{操作符.X更新值: bson.M{"age": 30}, 操作符.X更新插入时: bson.M{"weight": 40}}).
		X执行(context.Background())

	ast.NoError(err)
	ast.Equal(int64(0), result.X插入数)
	ast.Equal(int64(1), result.X修改数)
	ast.Equal(int64(0), result.X删除数)
	ast.Equal(int64(1), result.X替换插入数)
	ast.Equal(1, len(result.X替换插入IDs))
	ast.Equal(int64(1), result.X匹配数)
}
