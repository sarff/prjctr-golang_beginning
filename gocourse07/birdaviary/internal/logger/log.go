package logger

import (
	"io"
	"log/slog"
)

type Logger struct {
	*slog.Logger
}

func New(writer io.Writer) *Logger {
	return &Logger{
		Logger: slog.New(slog.NewJSONHandler(writer, nil)),
	}
}
