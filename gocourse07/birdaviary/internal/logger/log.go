package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New(writer *os.File) *Logger {
	log := &Logger{
		Logger: slog.New(slog.NewJSONHandler(writer, nil)),
	}
	slog.SetDefault(log.Logger)
	return log
}
