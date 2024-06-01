package options

import "go.mongodb.org/mongo-driver/mongo/options"

// [提示]
//type 运行命令选项 struct {
//     *基础运行命令选项
// }
// [结束]
type RunCommandOptions struct {
	*options.RunCmdOptions
}
