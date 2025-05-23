package config

import "github.com/stanislav-zeman/gonion/internal/dto"

type Config struct {
	Version  int                      `yaml:"version,omitempty"`
	Module   string                   `yaml:"module,omitempty"`
	Services map[string]ServiceConfig `yaml:"services,omitempty"`
	Misc     MiscConfig               `yaml:"misc,omitempty"`
}

type ServiceConfig struct {
	Logger         string               `yaml:"logger,omitempty"`
	Config         bool                 `yaml:"config,omitempty"`
	API            map[string]any       `yaml:"api,omitempty"`
	Application    ApplicationConfig    `yaml:"application,omitempty"`
	Domain         DomainConfig         `yaml:"domain,omitempty"`
	Infrastructure InfrastructureConfig `yaml:"infrastructure,omitempty"`
}

type ApplicationConfig struct {
	Service []dto.Service `yaml:"service,omitempty"`
}

type ApplicationServiceConfig struct {
	Name    string        `yaml:"name,omitempty"`
	Query   []dto.Query   `yaml:"query,omitempty"`
	Command []dto.Command `yaml:"command,omitempty"`
}

type DomainConfig struct {
	Entity     []dto.Entity     `yaml:"entity,omitempty"`
	Value      []dto.Value      `yaml:"value,omitempty"`
	Repository []dto.Repository `yaml:"repository,omitempty"`
	Service    []dto.Service    `yaml:"service,omitempty"`
}

type StructConfig struct {
	Name   string              `yaml:"name,omitempty"`
	Fields []StructFieldConfig `yaml:"fields,omitempty"`
}

type StructFieldConfig struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type,omitempty"`
}

type InfrastructureConfig struct {
	Clients    []ClientConfig   `yaml:"clients,omitempty"`
	Repository []dto.Repository `yaml:"repository,omitempty"`
}

type ClientConfig struct {
	Name string `yaml:"name,omitempty"`
}

type MiscConfig struct {
	Makefile  string `yaml:"makefile,omitempty"`
	Gitignore string `yaml:"gitignore,omitempty"`
	GolangCI  string `yaml:"golangci,omitempty"`
}
