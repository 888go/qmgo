package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/888go/qmgo/operator"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

// User 包含用户信息
type User struct {
	FirstName      string     `bson:"fname"`
	LastName       string     `bson:"lname"`
	Age            uint8      `bson:"age" validate:"gte=0,lte=130"`
	Email          string     `bson:"e-mail" validate:"required,email"`
	FavouriteColor string     `bson:"favouriteColor" validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `bson:"addresses" validate:"required,dive,required"` // 一个人可以拥有一个住宅和一间度假屋...
}

// Address 结构体用于存储用户地址信息
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// CustomRule 使用自定义规则
type CustomRule struct {
	Name string `validate:"required,foo"`
}

func TestValidator(t *testing.T) {
	ast := require.New(t)
	ctx := context.Background()

	user := &User{}
	// not need validator op
	ast.NoError(Do(ctx, user, 操作符.X删除前))
	ast.NoError(Do(ctx, user, 操作符.X插入后))
	// check success
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "Unknown",
	}

	user = &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "1234@gmail.com",
		FavouriteColor: "#000",
		Addresses:      []*Address{address, address},
	}
	ast.NoError(Do(ctx, user, 操作符.X插入前))
	ast.NoError(Do(ctx, user, 操作符.X更新或插入前))
	ast.NoError(Do(ctx, *user, 操作符.X更新或插入前))

	users := []*User{user, user, user}
	ast.NoError(Do(ctx, users, 操作符.X插入前))

	// check failure
	user.Age = 150
	ast.Error(Do(ctx, user, 操作符.X插入前))
	user.Age = 22
	user.Email = "1234@gmail" // invalid email
	ast.Error(Do(ctx, user, 操作符.X插入前))
	user.Email = "1234@gmail.com"
	user.Addresses[0].City = "" // 字符串标签使用默认值
	ast.Error(Do(ctx, user, 操作符.X插入前))

	// input slice
	users = []*User{user, user, user}
	ast.Error(Do(ctx, users, 操作符.X插入前))

	useris := []interface{}{user, user, user}
	ast.Error(Do(ctx, useris, 操作符.X插入前))

	user.Addresses[0].City = "shanghai"
	users = []*User{user, user, user}
	ast.NoError(Do(ctx, users, 操作符.X插入前))

	us := []User{*user, *user, *user}
	ast.NoError(Do(ctx, us, 操作符.X插入前))
	ast.NoError(Do(ctx, &us, 操作符.X插入前))

	// all bson type
	mdoc := []interface{}{bson.M{"name": "", "age": 12}, bson.M{"name": "", "age": 12}}
	ast.NoError(Do(ctx, mdoc, 操作符.X插入前))
	adoc := bson.A{"Alex", "12"}
	ast.NoError(Do(ctx, adoc, 操作符.X插入前))
	edoc := bson.E{"Alex", "12"}
	ast.NoError(Do(ctx, edoc, 操作符.X插入前))
	ddoc := bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
	ast.NoError(Do(ctx, ddoc, 操作符.X插入前))

	// nil ptr
	user = nil
	ast.NoError(Do(ctx, user, 操作符.X插入前))
	ast.NoError(Do(ctx, nil, 操作符.X插入前))

	// use custom rules
	customRule := &CustomRule{Name: "bar"}
	v := validator.New()
	_ = v.RegisterValidation("foo", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "bar"
	})
	SetValidate(v)
	ast.NoError(Do(ctx, customRule, 操作符.X插入前))
}
