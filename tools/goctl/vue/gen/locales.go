package gen

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

func (g *defaultGenerator) genLocales(in parser.Table) (codeFile, error) {

	text, err := pathx.LoadTemplate(category, localesItemFieldTemplateFile, template.Locales)
	if err != nil {
		return codeFile{}, err
	}

	primaryKey, uniqueKey := genCacheKeys(in)

	var table Table
	table.Table = in
	table.PrimaryCacheKey = primaryKey
	table.UniqueCacheKey = uniqueKey
	table.ContainsUniqueCacheKey = len(uniqueKey) > 0
	table.ignoreColumns = g.ignoreColumns

	itemFieldCode, err := genLocalesItemFields(table, in.Fields)
	if err != nil {
		return codeFile{}, err
	}

	t := util.With("locales").
		Parse(text).
		GoFmt(false)
	output, err := t.Execute(map[string]any{
		"pkg":                   g.pkg,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
		"kebabObject":           in.Name.ToKebab(),
		"localesItemFields":     itemFieldCode,
		"tableComment":          in.TableComment.Lower(),
	})
	if err != nil {
		return codeFile{}, err
	}

	return codeFile{
		filename: "locales/" + in.Name.ToSnake() + ".yaml",
		content:  output.String(),
	}, nil

}
