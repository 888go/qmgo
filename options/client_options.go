package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 客户端选项 struct {
//     *options.客户端选项
// }
// [结束]
type ClientOptions struct {
	*options.ClientOptions
}
