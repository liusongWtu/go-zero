package gen

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/model"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const pwd = "."

type (
	defaultGenerator struct {
		console.Console
		// source string
		dir           string
		pkg           string
		cfg           *config.Config
		isPostgreSql  bool
		ignoreColumns []string
		requestPath   string
	}

	// Option defines a function with argument defaultGenerator
	Option func(generator *defaultGenerator)

	code struct {
		importsCode string
		varsCode    string
		typesCode   string
		newCode     string
		insertCode  string
		findCode    []string
		updateCode  string
		deleteCode  string
		cacheExtra  string
		tableName   string
	}

	codeFile struct {
		filename string
		content  string
	}

	codeTuple struct {
		codes []codeFile
	}
)

// NewDefaultGenerator creates an instance for defaultGenerator
func NewDefaultGenerator(dir string, cfg *config.Config, opt ...Option) (*defaultGenerator, error) {
	if dir == "" {
		dir = pwd
	}
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	dir = dirAbs
	pkg := util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dir)
	if err != nil {
		return nil, err
	}

	generator := &defaultGenerator{dir: dir, cfg: cfg, pkg: pkg}
	var optionList []Option
	optionList = append(optionList, newDefaultOption())
	optionList = append(optionList, opt...)
	for _, fn := range optionList {
		fn(generator)
	}

	return generator, nil
}

// WithConsoleOption creates a console option.
func WithConsoleOption(c console.Console) Option {
	return func(generator *defaultGenerator) {
		generator.Console = c
	}
}

// WithIgnoreColumns ignores the columns while insert or update rows.
func WithIgnoreColumns(ignoreColumns []string) Option {
	return func(generator *defaultGenerator) {
		generator.ignoreColumns = ignoreColumns
	}
}

func WithRequestPath(requestPath string) Option {
	return func(generator *defaultGenerator) {
		generator.requestPath = requestPath
	}
}

// WithPostgreSql marks  defaultGenerator.isPostgreSql true.
func WithPostgreSql() Option {
	return func(generator *defaultGenerator) {
		generator.isPostgreSql = true
	}
}

func newDefaultOption() Option {
	return func(generator *defaultGenerator) {
		generator.Console = console.NewColorConsole()
	}
}

func (g *defaultGenerator) StartFromInformationSchema(tables map[string]*model.Table, withCache, strict bool) error {
	m := make(map[string]*codeTuple)
	for _, each := range tables {
		table, err := parser.ConvertDataType(each, strict)
		if err != nil {
			return err
		}

		requestApiCode, err := g.genRequestApi(*table, g.requestPath)
		if err != nil {
			return err
		}

		storePiniaCode, err := g.genStorePinia(*table)
		if err != nil {
			return err
		}

		localesCode, err := g.genLocales(*table)
		if err != nil {
			return err
		}

		//editor
		editorIndexVueCode, err := g.genEditorIndexVue(*table)
		if err != nil {
			return err
		}

		editorFormVueCode, err := g.genEditorForm(*table)
		if err != nil {
			return err
		}

		editorFormHooksCode, err := g.genEditorFormHooks(*table)
		if err != nil {
			return err
		}

		editorFormRulesCode, err := g.genEditorFormRules(*table)
		if err != nil {
			return err
		}

		editorFormIndexCode, err := g.genEditorFormIndex(*table)
		if err != nil {
			return err
		}

		//search
		searchFormCode, err := g.genSearchForm(*table)
		if err != nil {
			return err
		}

		searchFormDataCode, err := g.genSearchFormData(*table)
		if err != nil {
			return err
		}

		searchFormHooksCode, err := g.genSearchFormHooks(*table)
		if err != nil {
			return err
		}

		searchFormRulesCode, err := g.genSearchFormRules(*table)
		if err != nil {
			return err
		}

		//index
		indexCode, err := g.genIndex(*table)
		if err != nil {
			return err
		}

		m[table.Name.Source()] = &codeTuple{
			codes: []codeFile{
				requestApiCode,
				storePiniaCode,
				editorIndexVueCode,
				localesCode,
				editorFormVueCode,
				editorFormHooksCode,
				editorFormRulesCode,
				editorFormIndexCode,
				searchFormCode,
				searchFormDataCode,
				searchFormHooksCode,
				searchFormRulesCode,
				indexCode,
			}}
	}

	return g.createFile(m)
}

func (g *defaultGenerator) createFile(modelList map[string]*codeTuple) error {
	dirAbs, err := filepath.Abs(g.dir)
	if err != nil {
		return err
	}

	g.dir = dirAbs
	g.pkg = util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	for _, codes := range modelList {
		for _, code := range codes.codes {
			filename := filepath.Join(dirAbs, code.filename)
			err := pathx.MkdirIfNotExistByFile(filename)
			if err != nil {
				return err
			}
			err = os.WriteFile(filename, []byte(code.content), os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	g.Success("Done.")
	return nil
}

// Table defines mysql table
type Table struct {
	parser.Table
	PrimaryCacheKey        Key
	UniqueCacheKey         []Key
	ContainsUniqueCacheKey bool
	ignoreColumns          []string
}

func (t Table) isIgnoreColumns(columnName string) bool {
	for _, v := range t.ignoreColumns {
		if v == columnName {
			return true
		}
	}
	return false
}

func wrapWithRawString(v string, postgreSql bool) string {
	if postgreSql {
		return v
	}

	if v == "`" {
		return v
	}

	if !strings.HasPrefix(v, "`") {
		v = "`" + v
	}

	if !strings.HasSuffix(v, "`") {
		v = v + "`"
	} else if len(v) == 1 {
		v = v + "`"
	}

	return v
}
