package Postgres

import (
    "go.uber.org/zap"
    "github.com/stanislav-zeman/gonion/example/internal/management/domain/repository"
    "github.com/stanislav-zeman/gonion/example/internal/management/domain/entity"
)

var _ repository.ConfigurationRepository = &ConfigurationRepository{}

type ConfigurationRepository struct{
    log *zap.Logger
}

func NewConfigurationRepository(
    log *zap.Logger,
) *ConfigurationRepository {
    return &ConfigurationRepository{
         log: log,
    }
}

func (r *ConfigurationRepository) GetAllConfigurations() ([]entity.Configuration, error) {
    panic("unimplemented")
}

func (r *ConfigurationRepository) GetConfiguration(id string) (entity.Configuration, error) {
    _ = id
    panic("unimplemented")
}

func (r *ConfigurationRepository) CreateConfiguration(e entity.Configuration) (entity.Configuration, error) {
    _ = e
    panic("unimplemented")
}

func (r *ConfigurationRepository) UpdateConfiguration(e entity.Configuration) (entity.Configuration, error) {
    _ = e
    panic("unimplemented")
}

func (r *ConfigurationRepository) DeleteConfiguration(id string) (entity.Configuration, error){
    _ = id
    panic("unimplemented")
}

