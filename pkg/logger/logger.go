package logger

import (
	"fullstackhero/golang/pkg/core/utils"
	"fullstackhero/golang/pkg/logger/zap"
)

const (
	Zap = "zap"
)

// package level variable that should be used by every application for logging.
var Instance Logger

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Println(args ...interface{})
}

func init() {
	loggerType := utils.Getenv("LOGGER_TYPE")
	if loggerType == "" {
		loggerType = Zap
	}
	switch loggerType {
	case Zap:
		Instance = zap.NewZapLogger()
	default:
		Instance = zap.NewZapLogger()
	}
}
