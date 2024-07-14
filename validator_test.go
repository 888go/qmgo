package qmgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// User 包含用户信息 md5:0449710cca9a8191
type User struct {
	FirstName string            `bson:"fname"`
	LastName  string            `bson:"lname"`
	Age       uint8             `bson:"age" validate:"gte=0,lte=130" `    // Age must in [0,130]
	Email     string            `bson:"e-mail" validate:"required,email"` //  Email can't be empty string, and must has email format
	CreateAt  time.Time         `bson:"createAt" validate:"lte"`          // CreateAt 必须小于或等于当前时间 md5:588f116766addee2
	Relations map[string]string `bson:"relations" validate:"max=2"`       // Relations can't has more than 2 elements
}

func TestValidator(t *testing.T) {
	ast := require.New(t)
	cli := initClient("test")
	ctx := context.Background()
	defer cli.Close(ctx)
	defer cli.DropCollection(ctx)

	user := &User{
		FirstName: "",
		LastName:  "",
		Age:       45,
		Email:     "1234@gmail.com",
	}
	_, err := cli.InsertOne(ctx, user)
	ast.NoError(err)

	user.Age = 200 // invalid age
	_, err = cli.InsertOne(ctx, user)
	ast.Error(err)

	users := []*User{user, user, user}
	_, err = cli.InsertMany(ctx, users)
	ast.Error(err)

	user.Age = 20
	user.Email = "1234@gmail" // 邮件标签，无效邮件 md5:5e57b4e04096fa8c
	err = cli.ReplaceOne(ctx, bson.M{"age": 45}, user)
	ast.Error(err)

	user.Email = "" // 必要的标签，无效的空字符串 md5:1d307d1c696f38f6
	_, err = cli.Upsert(ctx, bson.M{"age": 45}, user)
	ast.Error(err)

	user.Email = "1234@gmail.com"
	user.CreateAt = time.Now().Add(1 * time.Hour) // lte 标签用于时间，时间必须小于或等于当前时间 md5:d8aebd8f3f7b532d
	_, err = cli.Upsert(ctx, bson.M{"age": 45}, user)
	ast.Error(err)

	user.CreateAt = time.Now()
	user.Relations = map[string]string{"Alex": "friend", "Joe": "friend"}
	_, err = cli.Upsert(ctx, bson.M{"age": 45}, user)
	ast.NoError(err)

	user.Relations = map[string]string{"Alex": "friend", "Joe": "friend", "Bob": "sister"} // 最大标签，映射中的数字数量 md5:82e032c216b8c99d
	_, err = cli.Upsert(ctx, bson.M{"age": 45}, user)
	ast.Error(err)
}
