package config

import (
	"log/slog"
	"os"
)

type Log struct {
}

func NewLogger() *Log {
	return &Log{}
}

func (l *Log) Setup() {

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
