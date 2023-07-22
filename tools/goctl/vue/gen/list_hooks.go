package gen

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

// genListHooks 生成/listHooks.ts文件
func (g *defaultGenerator) genListHooks(in parser.Table) (codeFile, error) {
	text, err := pathx.LoadTemplate(category, listHooksTemplateFile, template.ListHooks)
	if err != nil {
		return codeFile{}, err
	}

	if len(in.PrimaryKey.Name.Source()) == 0 {
		return codeFile{}, fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}

	primaryKey, uniqueKey := genCacheKeys(in)

	var table Table
	table.Table = in
	table.PrimaryCacheKey = primaryKey
	table.UniqueCacheKey = uniqueKey
	table.ContainsUniqueCacheKey = len(uniqueKey) > 0
	table.ignoreColumns = g.ignoreColumns

	tableListFieldsCode, err := genListHooksTableFields(table, in.Fields)
	if err != nil {
		return codeFile{}, err
	}

	requestTableFieldAssignCode, err := genRequestTableFieldAssigns(table, in.Fields)
	if err != nil {
		return codeFile{}, err
	}

	templateParams := map[string]any{
		"pkg":                   g.pkg,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
		"kebabObject":           in.Name.ToKebab(),
		"lowerObject":           strings.ReplaceAll(in.Name.Lower(), "_", ""),
		"tableListFields":       tableListFieldsCode,
		"requestTableField":     requestTableFieldAssignCode,
	}

	t := util.With("list-hooks").
		Parse(text).
		GoFmt(false)
	output, err := t.Execute(templateParams)
	if err != nil {
		return codeFile{}, err
	}

	return codeFile{
		filename: "src/views/" + strings.ReplaceAll(in.Name.Lower(), "_", "") + "/listHooks.ts",
		content:  output.String(),
	}, nil
}
