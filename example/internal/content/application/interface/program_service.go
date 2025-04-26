package interfaces

import (

    "github.com/stanislav-zeman/gonion/example/internal/content/application/query"
)

type ProgramService interface{
    Program(q query.ProgramQuery) (r query.ProgramQueryResult, err error)
    Programs(q query.ProgramsQuery) (r query.ProgramsQueryResult, err error)
}
