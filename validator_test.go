package mgo类

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
	
	"github.com/stretchr/testify/require"
)

// User 包含用户信息
type User struct {
	FirstName string            `bson:"fname"`
	LastName  string            `bson:"lname"`
	Age       uint8             `bson:"age" validate:"gte=0,lte=130" `    // Age must in [0,130]
	Email     string            `bson:"e-mail" validate:"required,email"` //  Email can't be empty string, and must has email format
	CreateAt  time.Time         `bson:"createAt" validate:"lte"`          // CreateAt 必须小于等于当前时间
	Relations map[string]string `bson:"relations" validate:"max=2"`       // Relations can't has more than 2 elements
}

func TestValidator(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	ctx := context.Background()
	defer cli.X关闭(ctx)
	defer cli.X删除集合(ctx)

	user := &User{
		FirstName: "",
		LastName:  "",
		Age:       45,
		Email:     "1234@gmail.com",
	}
	_, err := cli.X插入(ctx, user)
	ast.NoError(err)

	user.Age = 200 // invalid age
	_, err = cli.X插入(ctx, user)
	ast.Error(err)

	users := []*User{user, user, user}
	_, err = cli.X插入多个(ctx, users)
	ast.Error(err)

	user.Age = 20
	user.Email = "1234@gmail" // email标签，无效邮箱地址
	err = cli.X替换一条(ctx, bson.M{"age": 45}, user)
	ast.Error(err)

	user.Email = "" // 必需的标签，空字符串无效
	_, err = cli.X更新或插入(ctx, bson.M{"age": 45}, user)
	ast.Error(err)

	user.Email = "1234@gmail.com"
	user.CreateAt = time.Now().Add(1 * time.Hour) // lte标签用于时间，时间值必须小于等于当前时间
	_, err = cli.X更新或插入(ctx, bson.M{"age": 45}, user)
	ast.Error(err)

	user.CreateAt = time.Now()
	user.Relations = map[string]string{"Alex": "friend", "Joe": "friend"}
	_, err = cli.X更新或插入(ctx, bson.M{"age": 45}, user)
	ast.NoError(err)

	user.Relations = map[string]string{"Alex": "friend", "Joe": "friend", "Bob": "sister"} // 最大标签数，映射的数量
	_, err = cli.X更新或插入(ctx, bson.M{"age": 45}, user)
	ast.Error(err)
}
