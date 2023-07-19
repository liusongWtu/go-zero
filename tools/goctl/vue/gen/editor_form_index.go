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

// genEditorFormRules 生成/editor/form/index.ts文件
func (g *defaultGenerator) genEditorFormIndex(in parser.Table) (codeFile, error) {
	text, err := pathx.LoadTemplate(category, editorFormIndexTemplateFile, template.EditorFormIndex)
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

	t := util.With("editor-form-index").
		Parse(text).
		GoFmt(false)
	output, err := t.Execute(map[string]any{
		"pkg":                   g.pkg,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
		"kebabObject":           in.Name.ToKebab(),
	})
	if err != nil {
		return codeFile{}, err
	}

	return codeFile{
		filename: "src/views/" + strings.ReplaceAll(in.Name.Lower(), "_", "") + "/editor/form/index.ts",
		content:  output.String(),
	}, nil
}
