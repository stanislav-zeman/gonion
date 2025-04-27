package processor

import (
	"fmt"
	"log"

	"github.com/stanislav-zeman/gonion/internal/dto"
)

func (p *Processor) processInfrastructureRepositories(
	serviceName string,
	logger dto.Logger,
	repositories []dto.Repository,
	entities []dto.Entity,
) error {
	for _, repository := range repositories {
		repository.Logger = logger
		repository.Import = dto.Import{
			Module:  p.config.Module,
			Service: serviceName,
		}

		for _, entity := range entities {
			if entity.Name == repository.Name {
				repository.Entity = entity
				break
			}
		}

		log.Printf("repository: %v\n", repository)

		data, err := p.templator.TemplateInfrastructureRepository(repository)
		if err != nil {
			return fmt.Errorf("failed templating infrastructure repository: %w", err)
		}

		err = p.writer.WriteInfrastructureRepository(serviceName, repository.Name, repository.Type, data)
		if err != nil {
			return fmt.Errorf("failed writing infrastructure repository: %w", err)
		}
	}

	return nil
}
