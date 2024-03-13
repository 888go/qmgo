提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
    但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"


[func (c *CustomFields) SetUpdateAt(fieldName string) CustomFieldsBuilder {]
ff=设置更新时间字段名
fieldName=字段名称

[func (c *CustomFields) SetCreateAt(fieldName string) CustomFieldsBuilder {]
ff=设置创建时间字段名
fieldName=字段名称

[func (c *CustomFields) SetId(fieldName string) CustomFieldsBuilder {]
ff=设置ID字段名
fieldName=字段名称

[func (c CustomFields) CustomCreateTime(doc interface{}) {]
ff=自定义创建时间
doc=待插入文档

[func (c CustomFields) CustomUpdateTime(doc interface{}) {]
ff=自定义更新时间
doc=待插入文档

[func (c CustomFields) CustomId(doc interface{}) {]
ff=自定义ID
doc=待插入文档
