package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 替换选项 struct {
//     更新钩子 interface{}
//     *options.替换选项
// }
// [结束]
type ReplaceOptions struct {
	UpdateHook interface{}
	*options.ReplaceOptions
}
