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
	entity          *template.Template
	value           *template.Template
	domainInterface *template.Template
	domainService   *template.Template

	// Application templates.
	command      *template.Template
	query        *template.Template
	appInterface *template.Template
	appService   *template.Template
}

func New(assetsDirector string) (t Templator, err error) {
	fp := filepath.Join(assetsDirector, layers.DomainLayer, "entity.tmpl")
	entity, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "value.tmpl")
	value, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "interface.tmpl")
	domainInterface, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.DomainLayer, "service.tmpl")
	domainService, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "command.tmpl")
	command, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "query.tmpl")
	query, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "interface.tmpl")
	appInterface, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	fp = filepath.Join(assetsDirector, layers.ApplicationLayer, "service.tmpl")
	appService, err := template.ParseFiles(fp)
	if err != nil {
		return
	}

	t = Templator{
		entity:          entity,
		value:           value,
		domainInterface: domainInterface,
		domainService:   domainService,

		command:      command,
		query:        query,
		appInterface: appInterface,
		appService:   appService,
	}

	return
}

// ----------------------------------------------------------------------------

func (t *Templator) TemplateEntity(e dto.Entity) (data []byte, err error) {
	return templateObject(t.entity, e)
}

func (t *Templator) TemplateValue(e dto.Value) (data []byte, err error) {
	return templateObject(t.value, e)
}

func (t *Templator) TemplateDomainInterface(s dto.Service) (data []byte, err error) {
	return templateObject(t.domainInterface, s)
}

func (t *Templator) TemplateDomainService(e dto.Service) (data []byte, err error) {
	return templateObject(t.domainService, e)
}

// ----------------------------------------------------------------------------

func (t *Templator) TemplateCommand(c dto.Command) (data []byte, err error) {
	return templateObject(t.command, c)
}

func (t *Templator) TemplateQuery(q dto.Query) (data []byte, err error) {
	return templateObject(t.query, q)
}

func (t *Templator) TemplateApplicationInterface(s dto.Service) (data []byte, err error) {
	return templateObject(t.appInterface, s)
}

func (t *Templator) TemplateApplicationService(s dto.Service) (data []byte, err error) {
	return templateObject(t.appService, s)
}

// ----------------------------------------------------------------------------

func templateObject(t *template.Template, object any) (data []byte, err error) {
	b := bytes.NewBuffer(make([]byte, 0))

	err = t.Execute(b, object)
	if err != nil {
		return
	}

	data = b.Bytes()
	return
}
