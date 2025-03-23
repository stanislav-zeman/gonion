package config

type RESTAPIConfig struct {
	Framework string                  `json:"framework,omitempty"`
	Endpoints []RESTAPIEndpointConfig `json:"endpoints,omitempty"`
}

type RESTAPIEndpointConfig struct {
	Name       string   `json:"name,omitempty"`
	Controller string   `json:"controller,omitempty"`
	Methods    []string `json:"methods,omitempty"`
}
