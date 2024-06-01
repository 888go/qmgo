package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//
//	type 会话选项 struct {
//	    *options.会话选项
//	}
//
// [结束]
type SessionOptions struct {
	*options.SessionOptions
}
