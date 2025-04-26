package entity

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/domain/value"
)

type Stream struct {
    ID string
    Name string
    State value.State
}
