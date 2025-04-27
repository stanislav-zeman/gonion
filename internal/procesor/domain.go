package processor

import (
	"fmt"
	"log"
	"strings"

	"github.com/stanislav-zeman/gonion/internal/dto"
)

func (p *Processor) processEntities(serviceName string, entities []dto.Entity) error {
	for _, entity := range entities {
		for _, field := range entity.Fields {
			if strings.Contains(field.Type, "value") {
				entity.HasValues = true
				entity.Import = dto.Import{
					Module:  p.config.Module,
					Service: serviceName,
				}

				break
			}
		}

		log.Printf("entity: %v\n", entity)

		data, err := p.templator.TemplateEntity(entity)
		if err != nil {
			return fmt.Errorf("failed templating entity: %w", err)
		}

		err = p.writer.WriteDomainEntity(serviceName, entity.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing entity: %w", err)
		}
	}

	return nil
}

func (p *Processor) processValues(serviceName string, values []dto.Value) error {
	for _, value := range values {
		log.Printf("value: %v\n", value)

		data, err := p.templator.TemplateValue(value)
		if err != nil {
			return fmt.Errorf("failed templating entity: %w", err)
		}

		err = p.writer.WriteDomainValue(serviceName, value.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing value: %w", err)
		}
	}

	return nil
}

func (p *Processor) processDomainServices(serviceName string, logger dto.Logger, services []dto.Service) error {
	for _, domainService := range services {
		domainService.Logger = logger
		domainService.Import = dto.Import{
			Module:  p.config.Module,
			Service: serviceName,
		}

		log.Printf("Generating domain service: %v\n", domainService)

		data, err := p.templator.TemplateDomainService(domainService)
		if err != nil {
			return fmt.Errorf("failed templating service: %w", err)
		}

		err = p.writer.WriteDomainService(serviceName, domainService.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing service: %w", err)
		}

		log.Printf("Generating domain service interface: %v\n", domainService)

		data, err = p.templator.TemplateDomainInterface(domainService)
		if err != nil {
			return fmt.Errorf("failed templating service interface: %w", err)
		}

		err = p.writer.WriteDomainInterface(serviceName, domainService.Name+"_service", data)
		if err != nil {
			return fmt.Errorf("failed writing service interface: %w", err)
		}
	}

	return nil
}

func (p *Processor) processDomainRepositories(
	serviceName string,
	repositories []dto.Repository,
	entities []dto.Entity,
) error {
	for _, repository := range repositories {
		repository.Import = dto.Import{
			Module:  p.config.Module,
			Service: serviceName,
		}

		var foundEntity bool
		for _, entity := range entities {
			if entity.Name == repository.Name {
				repository.Entity = entity
				foundEntity = true
				break
			}
		}

		if !foundEntity {
			return errNoMatchingRepositoryEntity
		}

		log.Printf("Generating domain repository: %v\n", repository)

		data, err := p.templator.TemplateDomainRepository(repository)
		if err != nil {
			return fmt.Errorf("failed templating domain repository: %w", err)
		}

		err = p.writer.WriteDomainRepository(serviceName, repository.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing domain repository: %w", err)
		}
	}

	return nil
}
