package templator

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/stanislav-zeman/gonion/internal/dto"
	"github.com/stanislav-zeman/gonion/internal/layers"
)

type Templator struct {
	// Domain templates.
	entity        *template.Template
	domainService *template.Template

	// Application templates.
	command    *template.Template
	query      *template.Template
	appService *template.Template
}

func New(assetsDirector string) (t Templator, err error) {
	fp := filepath.Join(assetsDirector, layers.DomainLayer, "entity_template.tmpl")
	entity, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "service_template.tmpl")
	domainService, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "command_template.tmpl")
	command, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "query_template.tmpl")
	query, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "service_template.tmpl")
	appService, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	t = Templator{
		entity:        entity,
		domainService: domainService,

		command:    command,
		query:      query,
		appService: appService,
	}

	return
}

func (t *Templator) TemplateEntity(e dto.Entity) (data []byte, err error) {
	return templateObject(t.entity, e)
}

func (t *Templator) TemplateDomainService(e dto.Service) (data []byte, err error) {
	return templateObject(t.domainService, e)
}

func (t *Templator) TemplateCommand(c dto.Command) (data []byte, err error) {
	return templateObject(t.command, c)
}

func (t *Templator) TemplateQuery(q dto.Query) (data []byte, err error) {
	return templateObject(t.query, q)
}

func (t *Templator) TemplateApplicationService(s dto.Service) (data []byte, err error) {
	return templateObject(t.appService, s)
}

func templateObject(t *template.Template, object any) (data []byte, err error) {
	b := bytes.NewBuffer(make([]byte, 0))

	err = t.Execute(b, object)
	if err != nil {
		return
	}

	data = b.Bytes()
	return
}
