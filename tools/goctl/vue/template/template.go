package template

import (
	_ "embed"
)

// Field defines a filed template for types
//
//go:embed tpl/field.tpl
var Field string

//go:embed tpl/request_api.tpl
var RequestApi string

//go:embed tpl/field.tpl
var FieldTs string

//go:embed tpl/store_pinia.tpl
var StorePinia string

//go:embed tpl/editor_index.tpl
var EditorIndex string

//go:embed tpl/locales_item_field.tpl
var LocalesItemField string

//go:embed tpl/locales.tpl
var Locales string

//go:embed tpl/editor_form_field.tpl
var EditorFormField string

//go:embed tpl/editor_form.tpl
var EditorForm string

//go:embed tpl/editor_form_hooks.tpl
var EditorFormHooks string

//go:embed tpl/editor_form_rules.tpl
var EditorFormRules string

//go:embed tpl/editor_form_index.tpl
var EditorFormIndex string

//go:embed tpl/search_form.tpl
var SearchForm string

//go:embed tpl/search_form_data.tpl
var SearchFormData string

//go:embed tpl/search_form_hooks.tpl
var SearchFormHooks string

//go:embed tpl/search_form_rules.tpl
var SearchFormRules string

//go:embed tpl/index.tpl
var Index string
