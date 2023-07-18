package gen

import (
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

func genLocalesItemFields(table Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genLocalesItemField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genLocalesItemField(table Table, field *parser.Field) (string, error) {
	text, err := pathx.LoadTemplate(category, localesItemFieldTemplateFile, template.LocalesItemField)
	if err != nil {
		return "", err
	}

	output, err := util.With("locales-item-field").
		Parse(text).
		Execute(map[string]any{
			"name": util.SafeString(field.Name.ToSnake()),
			"type": tsFieldType(field),
			// "tag":        tag,
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
			"data":       table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
