package Redis

import (
    "go.uber.org/zap"
    "github.com/stanislav-zeman/gonion/example/internal/management/domain/repository"
    "github.com/stanislav-zeman/gonion/example/internal/management/domain/entity"
)

var _ repository.StreamRepository = &StreamRepository{}

type StreamRepository struct{
    log *zap.Logger
}

func NewStreamRepository(
    log *zap.Logger,
) *StreamRepository {
    return &StreamRepository{
         log: log,
    }
}

func (r *StreamRepository) GetAllStreams() ([]entity.Stream, error) {
    panic("unimplemented")
}

func (r *StreamRepository) GetStream(id string) (entity.Stream, error) {
    _ = id
    panic("unimplemented")
}

func (r *StreamRepository) CreateStream(e entity.Stream) (entity.Stream, error) {
    _ = e
    panic("unimplemented")
}

func (r *StreamRepository) UpdateStream(e entity.Stream) (entity.Stream, error) {
    _ = e
    panic("unimplemented")
}

func (r *StreamRepository) DeleteStream(id string) (entity.Stream, error){
    _ = id
    panic("unimplemented")
}

