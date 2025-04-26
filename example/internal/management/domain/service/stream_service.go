package services

import (
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
