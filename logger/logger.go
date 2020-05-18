package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// Initialize logging configurations
func init()  {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			MessageKey: "msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}


// Function for Info logs
func Info(msg string, tags ...zap.Field){
	log.Info(msg, tags...)
	log.Sync()
}

// Function for Error logs
func Error(msg string, err error, tags ...zap.Field)  {
	tags = append(tags, zap.NamedError("Error", err))
	log.Error(msg, tags...)
	log.Sync()
}
