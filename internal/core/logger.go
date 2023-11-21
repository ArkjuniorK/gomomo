package core

import (
	"log/slog"
	"os"
)

// Logger act application logging using slog, standard Go library for
// structured log. This dependency could also be used by
// the package by injecting it at initialization.
type Logger struct {
	core *slog.Logger
}

func NewLogger() *Logger {
	opts := &slog.HandlerOptions{AddSource: true}
	hdlr := slog.NewTextHandler(os.Stdout, opts)
	core := slog.New(hdlr)

	return &Logger{core: core}
}

func (l *Logger) GetCore() *slog.Logger {
	return l.core
}
