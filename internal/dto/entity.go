package dto

type Entity struct {
	Name   string  `json:"name,omitempty"`
	Fields []Field `json:"fields,omitzero"`

	Import    Import `json:"-"`
	HasValues bool   `json:"-"`
}
