package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 事务Options struct {
//     *options.事务Options
// }
// [结束]
type TransactionOptions struct {
	*options.TransactionOptions
}
