package genvue

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vue/template"
)

const (
	category = "model"

	apiTsTemplateFile   = "api.tpl"
	fieldTsTemplateFile = "field.tpl"
)

var templates = map[string]string{
	deleteTemplateFile:                    template.Delete,
	deleteMethodTemplateFile:              template.DeleteMethod,
	fieldTemplateFile:                     template.Field,
	findOneTemplateFile:                   template.FindOne,
	findOneMethodTemplateFile:             template.FindOneMethod,
	findOneByFieldTemplateFile:            template.FindOneByField,
	findOneByFieldMethodTemplateFile:      template.FindOneByFieldMethod,
	findOneByFieldExtraMethodTemplateFile: template.FindOneByFieldExtraMethod,
	importsTemplateFile:                   template.Imports,
	importsWithNoCacheTemplateFile:        template.ImportsNoCache,
	insertTemplateFile:                    template.Insert,
	insertTemplateMethodFile:              template.InsertMethod,
	modelGenTemplateFile:                  template.ModelGen,
	modelCustomTemplateFile:               template.ModelCustom,
	modelNewTemplateFile:                  template.New,
	tableNameTemplateFile:                 template.TableName,
	tagTemplateFile:                       template.Tag,
	typesTemplateFile:                     template.Types,
	updateTemplateFile:                    template.Update,
	updateMethodTemplateFile:              template.UpdateMethod,
	varTemplateFile:                       template.Vars,
	errTemplateFile:                       template.Error,
}

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
