package repository

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/domain/entity"
)

type StreamRepository interface {
    GetAllStreams() ([]entity.Stream, error)
    GetStream(id string) (entity.Stream, error)
    CreateStream(entity.Stream) (entity.Stream, error)
    UpdateStream(entity.Stream) (entity.Stream, error)
    DeleteStream(id string) (entity.Stream, error)
}

