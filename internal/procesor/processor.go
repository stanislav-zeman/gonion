package processor

import (
	"errors"
	"fmt"

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

		err = p.processDomainRepositories(serviceName, service.Domain.Repository)
		if err != nil {
			return fmt.Errorf("failed processing domain repositories: %w", err)
		}

		// --------------------------------------------------------------------------

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
