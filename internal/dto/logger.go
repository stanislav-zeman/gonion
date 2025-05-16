package dto

import (
	"errors"
)

var ErrUnknownLogger = errors.New("unknown logger")

type Logger struct {
	Struct  string
	Package string
}

func ParseLogger(loggerName string) (logger Logger, err error) {
	switch loggerName {
	case "slog":
		logger = Logger{
			Struct:  "*slog.Logger",
			Package: "slog",
		}
		return
	case "zap":
		logger = Logger{
			Struct:  "*zap.Logger",
			Package: "go.uber.org/zap",
		}
		return
	case "zerolog":
		logger = Logger{
			Struct:  "*zerolog.Logger",
			Package: "github.com/rs/zerolog",
		}
		return
	default:
		err = ErrUnknownLogger
		return
	}
}
