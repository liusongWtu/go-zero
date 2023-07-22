package gen

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	category = "vue"

	requestApiTemplateFile              = "request_api.tpl"
	requestApiFieldTemplateFile         = "request_api_field.tpl"
	storePiniaTemplateFile              = "store_pinia.tpl"
	localesItemFieldTemplateFile        = "locales_item_field.tpl"
	localesTemplateFile                 = "locales.tpl"
	editorIndexTemplateFile             = "editor_index.tpl"
	editorFormFieldTemplateFile         = "editor_form_field.tpl"
	editorFormTemplateFile              = "editor_form.tpl"
	editorFormHooksTemplateFile         = "editor_form_hooks.tpl"
	editorFormRulesTemplateFile         = "editor_form_rules.tpl"
	editorFormIndexTemplateFile         = "editor_form_index.tpl"
	searchFormTemplateFile              = "search_form.tpl"
	searchFormDataTemplateFile          = "search_form_data.tpl"
	searchFormHooksTemplateFile         = "search_form_hooks.tpl"
	searchFormRulesTemplateFile         = "search_form_rules.tpl"
	indexTemplateFile                   = "index.tpl"
	listHooksTableFieldTemplateFile     = "list_hooks_table_field.tpl"
	requestTableFieldAssignTemplateFile = "request_table_field_assign.tpl"
	listHooksTemplateFile               = "list_hooks.tpl"
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