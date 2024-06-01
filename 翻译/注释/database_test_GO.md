
<原文开始>
//func TestCreateCollection(t *testing.T) {
//	ast := require.New(t)
//
//	cli := initClient("test")
//
//	timeSeriesOpt := options.TimeSeriesOptions{
//		TimeField:"timestamp",
//	}
//	timeSeriesOpt.SetMetaField("metadata")
//	ctx := context.Background()
//	createCollectionOpts := opts.CreateCollectionOptions{CreateCollectionOptions: options.CreateCollection().SetTimeSeriesOptions(&timeSeriesOpt)}
//	if err := cli.CreateCollection(ctx, "syslog", createCollectionOpts); err != nil {
//		ast.NoError(err)
//	}
//	cli.DropCollection(ctx)
//	cli.DropDatabase(ctx)
//}
<原文结束>

# <翻译开始>
// ```go
// 测试创建集合
// func TestCreateCollection(t *testing.T) {
// 初始化断言工具
// 	ast := require.New(t)
// 
// 初始化客户端，连接名为"test"的数据库
// 	cli := initClient("test")
// 
// 设置时间序列选项
// 	timeSeriesOpt := options.TimeSeriesOptions{
// 		TimeField: "timestamp",
// 	}
// 设置元数据字段
// 	timeSeriesOpt.SetMetaField("metadata")
// 
// 创建上下文
// 	ctx := context.Background()
// 创建集合选项，设置时间序列相关选项
// 	createCollectionOpts := opts.CreateCollectionOptions{CreateCollectionOptions: options.CreateCollection().SetTimeSeriesOptions(&timeSeriesOpt)}
// 创建名为"syslog"的集合，检查是否出错
// 	if err := cli.CreateCollection(ctx, "syslog", createCollectionOpts); err != nil {
// 		ast.NoError(err)
// 	}
// 删除集合
// 	cli.DropCollection(ctx)
// 删除数据库
// 	cli.DropDatabase(ctx)
// }
// ```
// 
// 这段代码是一个测试函数，用于测试在ArangoDB中创建带有时间序列选项的集合。它首先初始化测试所需的工具和客户端，然后定义时间序列的配置，接着在上下文中创建集合。如果在创建过程中没有错误，会删除创建的集合和数据库。
// md5:79faec56c35696a6
# <翻译结束>

