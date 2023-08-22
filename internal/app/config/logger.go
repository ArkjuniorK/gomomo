package config

import "go.uber.org/zap"

// Logger wrapper, the Core could be replaced with any kind of log package
// but for now the default core is *zap.Logger.
// This dependency is optional.
type Logger struct {
	Core *zap.Logger
}

func NewLogger() *Logger {
	conf := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:          "console",
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig:     zap.NewProductionConfig().EncoderConfig,
	}

	core, err := conf.Build()
	if err != nil {
		panic(err)
	}

	defer core.Sugar().Infoln("Logger initialized!")
	return &Logger{Core: core}
}
