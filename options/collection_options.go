package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//
//	type 集合选项 struct {
//	    *options.集合选项
//	}
//
// [结束]
type CollectionOptions struct {
	*options.CollectionOptions
}
