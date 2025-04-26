package dto

type Query struct {
	Name   string  `json:"name,omitempty"`
	Fields []Field `json:"fields,omitzero"`
}
