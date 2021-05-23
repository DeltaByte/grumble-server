package logging

import "go.uber.org/zap"

var taskLoggerConfig = zap.Config{
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

var taskLogger *zap.SugaredLogger

func initTaskLogger() error {
	if (taskLogger == nil) {
		logger, err := taskLoggerConfig.Build()
		if err != nil {
			return err
		}
		taskLogger = logger.Sugar()
	}
	return nil
}

// get the task logger
func Task() (*zap.SugaredLogger) {
	return taskLogger
}