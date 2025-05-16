package processor

import (
	"errors"
	"fmt"

	"github.com/stanislav-zeman/gonion/internal/config"
	"github.com/stanislav-zeman/gonion/internal/dto"
	"github.com/stanislav-zeman/gonion/internal/templator"
	"github.com/stanislav-zeman/gonion/internal/writer"
)

var errNoMatchingRepositoryEntity = errors.New("no matching entity defined for repository")

// Processor creates project structure using the supplied configuration
// and processes all of the structure to Go code.
type Processor struct {
	config    config.Config
	templator templator.Templator
	writer    writer.Writer
}

func New(config config.Config, templator templator.Templator, writer writer.Writer) Processor {
	return Processor{
		config:    config,
		templator: templator,
		writer:    writer,
	}
}

func (p *Processor) Run() error {
	for serviceName, service := range p.config.Services {
		logger, err := dto.ParseLogger(service.Logger)
		if err != nil {
			return fmt.Errorf("failed parsing logger: %w", err)
		}

		// --------------------------------------------------------------------------

		err = p.processEntities(serviceName, service.Domain.Entity)
		if err != nil {
			return fmt.Errorf("failed processing entities: %w", err)
		}

		err = p.processValues(serviceName, service.Domain.Value)
		if err != nil {
			return fmt.Errorf("failed processing values: %w", err)
		}

		err = p.processDomainServices(serviceName, logger, service.Domain.Service)
		if err != nil {
			return fmt.Errorf("failed processing domain services: %w", err)
		}

		err = p.processDomainRepositories(serviceName, service.Domain.Repository, service.Domain.Entity)
		if err != nil {
			return fmt.Errorf("failed processing domain repositories: %w", err)
		}

		// --------------------------------------------------------------------------

		err = p.processApplicationServices(serviceName, logger, service.Application.Service)
		if err != nil {
			return fmt.Errorf("failed processing application services: %w", err)
		}

		// --------------------------------------------------------------------------

		err = p.processInfrastructureRepositories(
			serviceName,
			logger,
			service.Infrastructure.Repository,
			service.Domain.Entity,
		)
		if err != nil {
			return fmt.Errorf("failed processing infrastructure repositories: %w", err)
		}
	}

	return nil
}
