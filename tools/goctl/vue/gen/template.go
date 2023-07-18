package gen

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	category = "vue"

	requestApiTemplateFile       = "request_api.tpl"
	fieldTsTemplateFile          = "field.tpl"
	storePiniaTemplateFile       = "store_pinia.tpl"
	editorIndexVueTemplateFile   = "editor_index_vue.tpl"
	localesItemFieldTemplateFile = "locales_item_field.tpl"
	localesTemplateFile          = "locales.tpl"
)

var templates = map[string]string{}

// Category returns model const value
func Category() string {
	return category
}

// Clean deletes all template files
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates creates template files if not exists
func GenTemplates() error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the deleted template files
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}

	return pathx.CreateTemplate(category, name, content)
}

// Update provides template clean and init
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}
