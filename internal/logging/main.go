package logging

import (
	goLog "log"

	"go.uber.org/zap"
)

func Init() {
	// init default logger
	if err := initDefaultLogger(); err != nil {
		goLog.Print("Failed to init default logger")
		panic(err)
	}

	// use default logger for secondary init errors
	log := defaultLogger
	failed := false
	defer log.Sync()

	// request logger
	if err := initRequestLogger(); err != nil {
		log.Fatal("Failed to init request logger", err)
		failed = true
	}

	// task logger
	if err := initTaskLogger(); err != nil {
		log.Fatal("Failed to init task logger", err)
		failed = true
	}

	// panic if any of the secondary loggers failed to init
	if failed {
		panic("One or more loggers failed to init")
	}
}

var defaultLoggerConfig = zap.Config{
	Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
	Development: false,
	Sampling: &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	},
	Encoding:         "json",
	EncoderConfig:    zap.NewProductionEncoderConfig(),
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
}

var defaultLogger *zap.SugaredLogger

func initDefaultLogger() error {
	if (defaultLogger == nil) {
		logger, err := defaultLoggerConfig.Build()
		if err != nil {
			return err
		}
		defaultLogger = logger.Sugar()
	}
	return nil
}


// get the default logger
func Default() (*zap.SugaredLogger) {
	return defaultLogger
}
