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

func (s *StreamService) Stream(q query.Stream) (r query.StreamResult, err error) {
    _ = q
    panic("unimplemented")
}

func (s *StreamService) Configuration(q query.Configuration) (r query.ConfigurationResult, err error) {
    _ = q
    panic("unimplemented")
}

func (s *StreamService) StreamCreate(c command.StreamCreate) (r command.StreamCreateResult, err error) {
    _ = c
    panic("unimplemented")
}

func (s *StreamService) ConfigurationCreate(c command.ConfigurationCreate) (r command.ConfigurationCreateResult, err error) {
    _ = c
    panic("unimplemented")
}

func (s *StreamService) ConfigurationUpdate(c command.ConfigurationUpdate) (r command.ConfigurationUpdateResult, err error) {
    _ = c
    panic("unimplemented")
}
