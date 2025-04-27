package dto

type Repository struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	Import Import `json:"-"`
	Logger Logger `json:"-"`
	Entity Entity `json:"-"`
}
