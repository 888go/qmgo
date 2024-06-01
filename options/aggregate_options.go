package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 分组统计选项 struct {
//     *options.AggregateOptions
// }
// [结束]
type AggregateOptions struct {
	*options.AggregateOptions
}
