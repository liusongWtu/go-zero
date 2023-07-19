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

// genSearchForm 生成/search/form.vue文件
func (g *defaultGenerator) genSearchForm(in parser.Table) (codeFile, error) {
	text, err := pathx.LoadTemplate(category, searchFormTemplateFile, template.SearchForm)
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

	// fieldCode, err := genEditorForeItemFields(table, in.Fields)
	// if err != nil {
	// 	return codeFile{}, err
	// }

	t := util.With("search-form").
		Parse(text).
		GoFmt(false)
	output, err := t.Execute(map[string]any{
		"pkg":                   g.pkg,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
		"kebabObject":           in.Name.ToKebab(),
		// "editorFormVueFields":   fieldCode,
	})
	if err != nil {
		return codeFile{}, err
	}

	return codeFile{
		filename: "src/views/" + strings.ReplaceAll(in.Name.Lower(), "_", "") + "/search/form.vue",
		content:  output.String(),
	}, nil
}
