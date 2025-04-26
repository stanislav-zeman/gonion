package dto

type Entity struct {
	Name   string
	Fields []Field

	Import    Import `json:"-"`
	HasValues bool   `json:"-"`
}
