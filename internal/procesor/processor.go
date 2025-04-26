package processor

import (
	"errors"
	"fmt"
	"log"

	"github.com/stanislav-zeman/gonion/internal/config"
	"github.com/stanislav-zeman/gonion/internal/dto"
	"github.com/stanislav-zeman/gonion/internal/templator"
	"github.com/stanislav-zeman/gonion/internal/writer"
)

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
		var logger dto.Logger
		switch service.Logger {
		case "zap":
			logger.Struct = "*zap.Logger"
			logger.Package = "go.uber.org/zap"
		default:
			return errors.New("unknown logger")
		}

		err := p.processEntities(serviceName, service.Domain.Entity)
		if err != nil {
			return fmt.Errorf("failed processing entities: %w", err)
		}

		for _, domainService := range service.Domain.Service {
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
		}

		for _, appService := range service.Application.Service {
			appService.Logger = logger
			appService.Import = dto.Import{
				Module:  p.config.Module,
				Service: serviceName,
			}

			log.Printf("Generating application service: %v\n", appService)

			data, err := p.templator.TemplateApplicationService(appService)
			if err != nil {
				return fmt.Errorf("failed templating service: %w", err)
			}

			err = p.writer.WriteApplicationService(serviceName, appService.Name, data)
			if err != nil {
				return fmt.Errorf("failed writing service: %w", err)
			}

			err = p.processQueries(serviceName, appService.Queries)
			if err != nil {
				return fmt.Errorf("failed processing queries: %w", err)
			}

			err = p.processCommands(serviceName, appService.Commands)
			if err != nil {
				return fmt.Errorf("failed processing commands: %w", err)
			}
		}
	}

	return nil
}

func (p *Processor) processEntities(serviceName string, entities []dto.Entity) error {
	for _, entity := range entities {
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

func (p *Processor) processQueries(serviceName string, queries []dto.Query) error {
	for _, query := range queries {
		log.Printf("query: %v\n", query)

		data, err := p.templator.TemplateQuery(query)
		if err != nil {
			return fmt.Errorf("failed templating query: %w", err)
		}

		err = p.writer.WriteApplicationQuery(serviceName, query.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing query: %w", err)
		}
	}

	return nil
}

func (p *Processor) processCommands(serviceName string, commands []dto.Command) error {
	for _, command := range commands {
		log.Printf("command: %v\n", command)

		data, err := p.templator.TemplateCommand(command)
		if err != nil {
			return fmt.Errorf("failed templating command: %w", err)
		}

		err = p.writer.WriteApplicationCommand(serviceName, command.Name, data)
		if err != nil {
			return fmt.Errorf("failed writing command: %w", err)
		}
	}

	return nil
}
