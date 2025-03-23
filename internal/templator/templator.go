package templator

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/stanislav-zeman/gonion/internal/dto"
	"github.com/stanislav-zeman/gonion/internal/layers"
)

type Templator struct {
	service *template.Template
	command *template.Template
	query   *template.Template
}

func New(assetsDirector string) (t Templator, err error) {
	fp := filepath.Join(assetsDirector, layers.ApplicationLayer, "service_template.tmpl")
	service, err := template.ParseFiles(fp)
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

	t = Templator{
		service: service,
		command: command,
		query:   query,
	}

	return
}

func (t *Templator) TemplateService(s dto.Service) (data []byte, err error) {
	return templateObject(t.service, s)
}

func (t *Templator) TemplateCommand(c dto.Command) (data []byte, err error) {
	return templateObject(t.command, c)
}

func (t *Templator) TemplateQuery(q dto.Query) (data []byte, err error) {
	return templateObject(t.query, q)
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
