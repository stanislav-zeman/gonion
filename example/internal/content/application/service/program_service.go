package services

import (

    "github.com/stanislav-zeman/gonion/example/internal/content/application/query"
    "go.uber.org/zap"
)

type ProgramService struct{
    log *zap.Logger
}

func NewProgramService(
    log *zap.Logger,
) *ProgramService {
    return &ProgramService{
         log: log,
    }
}

func (s *ProgramService) Program(q query.ProgramQuery) (r query.ProgramQueryResult, err error) {
    _ = q
    panic("unimplemented")
}

func (s *ProgramService) Programs(q query.ProgramsQuery) (r query.ProgramsQueryResult, err error) {
    _ = q
    panic("unimplemented")
}
