package dto

// Service represent both Application and Domain services.
type Service struct {
	Import   Import    `json:"-"`
	Logger   string    `json:"-"`
	Name     string    `json:"name,omitempty"`
	Queries  []Query   `json:"queries,omitempty"`
	Commands []Command `json:"commands,omitempty"`
}
