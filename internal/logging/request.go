package logging

import "go.uber.org/zap"

var requestLoggerConfig = zap.Config{
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

var requestLogger *zap.Logger

func initRequestLogger() error {
	if (requestLogger == nil) {
		logger, err := requestLoggerConfig.Build()
		if err != nil {
			return err
		}
		requestLogger = logger
	}
	return nil
}

// get the request logger
func Request() (*zap.Logger) {
	return requestLogger
}