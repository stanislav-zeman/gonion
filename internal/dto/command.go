package dto

type Command struct {
	Name   string  `json:"name,omitempty"`
	Fields []Field `json:"fields,omitzero"`
}
