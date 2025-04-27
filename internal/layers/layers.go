package layers

// Layer specifies all possible project layers.
type Layer = string

const (
	APILayer            Layer = "api"
	ApplicationLayer    Layer = "application"
	DomainLayer         Layer = "domain"
	InfrastructureLayer Layer = "infrastructure"
)
