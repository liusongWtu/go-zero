package gen

import (
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

func genListHooksTableFields(table Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genListHooksTableField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genListHooksTableField(table Table, field *parser.Field) (string, error) {
	text, err := pathx.LoadTemplate(category, listHooksTableFieldTemplateFile, template.ListHooksTableField)
	if err != nil {
		return "", err
	}

	name := util.SafeString(field.Name.ToSnake())
	output, err := util.With("list-hooks-table-field").
		Parse(text).
		Execute(map[string]any{
			"name": name,
			"type": tsFieldType(field),
			// "tag":        tag,
			"upperStartCamelObject": table.Table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Table.Name.ToCamel()).Untitle(),
			"hasSuffixTime":         strings.HasSuffix(name, "time"),
			"hasSuffixStatus":       strings.HasSuffix(name, "status"),
			"comment":               field.Comment,
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
