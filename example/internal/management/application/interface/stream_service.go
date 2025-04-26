package interfaces

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/command"
    "github.com/stanislav-zeman/gonion/example/internal/management/application/query"
)

type StreamService interface{
    Stream(q query.StreamQuery) (r query.StreamQueryResult, err error)
    Configuration(q query.ConfigurationQuery) (r query.ConfigurationQueryResult, err error)
    StreamCreate(c command.StreamCreateCommand) (r command.StreamCreateCommandResult, err error)
    ConfigurationCreate(c command.ConfigurationCreateCommand) (r command.ConfigurationCreateCommandResult, err error)
    ConfigurationUpdate(c command.ConfigurationUpdateCommand) (r command.ConfigurationUpdateCommandResult, err error)
}
