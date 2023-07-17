package gen

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

// genStorePinia 生成Store Pinia 文件
func (g *defaultGenerator) genStorePinia(in parser.Table) (codeFile, error) {
	text, err := pathx.LoadTemplate(category, storePiniaTemplateFile, template.StorePinia)
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

	fieldCode, err := genFields(table, in.Fields)
	if err != nil {
		return codeFile{}, err
	}

	t := util.With("store-pinia").
		Parse(text).
		GoFmt(false)
	output, err := t.Execute(map[string]any{
		"pkg":                   g.pkg,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
		"kebabObject":           in.Name.ToKebab(),
		"fields":                fieldCode,
	})
	if err != nil {
		return codeFile{}, err
	}

	return codeFile{
		filename: "src/store/modules/" + stringx.From(in.Name.ToCamel()).Untitle() + "PageOpts.ts",
		content:  output.String(),
	}, nil
}
