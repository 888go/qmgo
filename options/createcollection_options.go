package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 创建集合选项 struct {
//     *options.CreateCollectionOptions
// }
// [结束]
type CreateCollectionOptions struct {
	*options.CreateCollectionOptions
}