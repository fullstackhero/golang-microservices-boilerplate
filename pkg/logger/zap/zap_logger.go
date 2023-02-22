package zap

import "go.uber.org/zap"

type ZapLogger struct {
	sugarLogger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	logger, _ := zap.NewProduction()
	return &ZapLogger{
		sugarLogger: logger.Sugar(),
	}
}

func (logger *ZapLogger) Debug(args ...interface{}) {
	logger.sugarLogger.Debug(args...)
}

func (logger *ZapLogger) Info(args ...interface{}) {
	logger.sugarLogger.Info(args...)
}

func (logger *ZapLogger) Println(args ...interface{}) {
	logger.sugarLogger.Info(args, "\n")
}
