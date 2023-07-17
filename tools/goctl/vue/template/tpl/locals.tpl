{{.lowerStartCamelObject}}:
  create{{.upperStartCamelObject}}: 创建配置
  edit{{.upperStartCamelObject}}: 编辑配置
  edit{{.upperStartCamelObject}}ById: 编辑配置 - {id}
  {{.lowerStartCamelObject}}List: 配置列表
  itemFields:
    id: 配置Id
    name: 名称
    code: Code
    data: 配置数据
    type_id: Schema Id
    type_name: Schema 名称
    type_code: Schema Code
    schema: Schema 内容
    remark: 备注
    status: 状态
    create_time: 创建时间
    update_time: 更新时间
  {{.lowerStartCamelObject}}Editor:
    typeId: 配置类型
    nameReg: 必须输入名称
    codeReg: 必须输入Code
    dataReg: 必须填写配置数据
    typeIdReg: 必须选择 Schema
    remarkReg: 请输入备注
    statusReg: 必须选择状态
    schemaRuleReg: Schema 格式错误
    dataRuleReg: 配置数据格式错误
  search:
    dateRange: 日期范围
    placeholderId: id, id, id, ...
    placeholderName: 搜索配置名称
    placeholderCode: 搜索Code
    idsRuleReg: Id 列表的格式应该为 id, id, id, ...
