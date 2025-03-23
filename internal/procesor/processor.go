package processor

import (
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
		for _, appService := range service.Application.Service {
			appService.Logger = service.Logger
			appService.Import = dto.Import{
				Module:  p.config.Module,
				Service: serviceName,
			}

			log.Printf("service: %v\n", appService)

			data, err := p.templator.TemplateService(appService)
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
