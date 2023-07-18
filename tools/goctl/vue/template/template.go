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

//go:embed tpl/editor_index_vue.tpl
var EditorIndexVue string

//go:embed tpl/locales_item_field.tpl
var LocalesItemField string

//go:embed tpl/locales.tpl
var Locales string
