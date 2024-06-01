package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//
//	type 数据库Options struct {
//		*options.数据库Options
//	}
//
// [结束]
type DatabaseOptions struct {
	*options.DatabaseOptions
}
