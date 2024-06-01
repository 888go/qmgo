package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 变更流选项 struct {
//     *options.ChangeStreamOptions
// }
// [结束]
type ChangeStreamOptions struct {
	*options.ChangeStreamOptions
}