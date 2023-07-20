package gen

import (
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/converter"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

func genRequestApiFields(table Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genRequestApiField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genRequestApiField(table Table, field *parser.Field) (string, error) {
	// tag, err := genTag(table, field.NameOriginal)
	// if err != nil {
	// 	return "", err
	// }

	text, err := pathx.LoadTemplate(category, requestApiFieldTemplateFile, template.RequestApiField)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
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

func tsFieldType(field *parser.Field) string {
	name := util.SafeString(field.Name.ToSnake())
	switch name {
	case "status":
		return "CommonStatus"
	}
	val, ok := converter.CommonGoDataTypeMapTs[field.DataType]
	if !ok {
		panic("unknown field type: " + field.DataType)
	}
	return val
}
