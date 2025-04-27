package processor

import (
	"errors"
	"fmt"
	"log"
	"strings"

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
		logger, err := p.parseLogger(service.Logger)
		if err != nil {
			return fmt.Errorf("failed parsing logger: %w", err)
		}

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

		err = p.processApplicationServices(serviceName, logger, service.Application.Service)
		if err != nil {
			return fmt.Errorf("failed processing application services: %w", err)
		}
	}

	return nil
}

func (p *Processor) parseLogger(loggerName string) (logger dto.Logger, err error) {
	switch loggerName {
	case "slog":
		logger = dto.Logger{
			Struct:  "*slog.Logger",
			Package: "slog",
		}
		return
	case "zap":
		logger = dto.Logger{
			Struct:  "*zap.Logger",
			Package: "go.uber.org/zap",
		}
		return
	case "zerolog":
		logger = dto.Logger{
			Struct:  "*zerolog.Logger",
			Package: "github.com/rs/zerolog",
		}
		return
	default:
		err = errors.New("unknown logger")
		return
	}
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

func (p *Processor) processApplicationServices(serviceName string, logger dto.Logger, services []dto.Service) error {
	for _, appService := range services {
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

		log.Printf("Generating application service interface: %v\n", appService)

		data, err = p.templator.TemplateApplicationInterface(appService)
		if err != nil {
			return fmt.Errorf("failed templating service interface: %w", err)
		}

		err = p.writer.WriteApplicationInterface(serviceName, appService.Name+"_service", data)
		if err != nil {
			return fmt.Errorf("failed writing service interface: %w", err)
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

	return nil
}

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
