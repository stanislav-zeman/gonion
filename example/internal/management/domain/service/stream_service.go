package services

import (
    "go.uber.org/zap"
    "github.com/stanislav-zeman/gonion/example/internal/management/domain/interface"
)

var _ interfaces.StreamService = &StreamService{}

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
