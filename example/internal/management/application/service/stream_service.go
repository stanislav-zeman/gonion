package services

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/command"
    "github.com/stanislav-zeman/gonion/example/internal/management/application/query"
    "go.uber.org/zap"
)

type StreamService struct{
    log *zap.Logger
}

func NewStreamService(
    log *zap.Logger,
) *StreamService {
    return &StreamService{
         log: log,
    }
}

func (s *StreamService) Stream(q query.StreamQuery) (r query.StreamQueryResult, err error) {
    _ = q
    panic("unimplemented")
}

func (s *StreamService) Configuration(q query.ConfigurationQuery) (r query.ConfigurationQueryResult, err error) {
    _ = q
    panic("unimplemented")
}

func (s *StreamService) StreamCreate(c command.StreamCreateCommand) (r command.StreamCreateCommandResult, err error) {
    _ = c
    panic("unimplemented")
}

func (s *StreamService) ConfigurationCreate(c command.ConfigurationCreateCommand) (r command.ConfigurationCreateCommandResult, err error) {
    _ = c
    panic("unimplemented")
}

func (s *StreamService) ConfigurationUpdate(c command.ConfigurationUpdateCommand) (r command.ConfigurationUpdateCommandResult, err error) {
    _ = c
    panic("unimplemented")
}
