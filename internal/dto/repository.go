package dto

type Repository struct {
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Entity string `json:"entity,omitempty"`

	Import Import `json:"-"`
	Logger Logger `json:"-"`
}
