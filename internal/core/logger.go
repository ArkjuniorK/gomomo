package core

import (
	"log/slog"
	"os"
	"path/filepath"
)

// Logger act application logging using slog, standard Go library for
// structured log. This dependency could also be used by
// the package by injecting it at initialization.
type Logger struct {
	core *slog.Logger
}

func NewLogger() *Logger {
	opts := &slog.HandlerOptions{AddSource: true, ReplaceAttr: replaceAttr}
	hdlr := slog.NewTextHandler(os.Stdout, opts)
	core := slog.New(hdlr)

	return &Logger{core: core}
}

func (l *Logger) GetCore() *slog.Logger {
	return l.core
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	// Replace full path of source file to workdir path
	if a.Key == slog.SourceKey {
		source := a.Value.Any().(*slog.Source)

		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		file, err := filepath.Rel(wd, source.File)
		if err != nil {
			panic(err)
		}

		source.File = file
	}
	return a
}
