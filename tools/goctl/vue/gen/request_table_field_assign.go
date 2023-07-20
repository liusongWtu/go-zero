package gen

import (
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

func genRequestTableFieldAssigns(table Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genRequestTableFieldAssign(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genRequestTableFieldAssign(table Table, field *parser.Field) (string, error) {
	text, err := pathx.LoadTemplate(category, requestTableFieldAssignTemplateFile, template.RequestTableFieldAssign)
	if err != nil {
		return "", err
	}

	name := util.SafeString(field.Name.ToSnake())
	output, err := util.With("request-table-field-assign").
		Parse(text).
		Execute(map[string]any{
			"name": name,
			"type": tsFieldType(field),
			// "tag":        tag,
			"upperStartCamelObject": table.Table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Table.Name.ToCamel()).Untitle(),
			"comment":               field.Comment,
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
