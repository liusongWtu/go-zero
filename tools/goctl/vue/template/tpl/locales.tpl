{{.lowerStartCamelObject}}:
  create{{.upperStartCamelObject}}: 创建{{.tableComment}}
  edit{{.upperStartCamelObject}}: 编辑{{.tableComment}}
  edit{{.upperStartCamelObject}}ById: 编辑{{.tableComment}} - {id}
  {{.lowerStartCamelObject}}List: {{.tableComment}}列表
  itemFields:
{{.localesItemFields}}
  {{.lowerStartCamelObject}}Editor:
    nameReg: 必须输入名称
    remarkReg: 请输入备注
    statusReg: 必须选择状态
  search:
    dateRange: 日期范围
    placeholderId: id, id, id, ...
    placeholderName: 搜索名称
    idsRuleReg: Id 列表的格式应该为 id, id, id, ...