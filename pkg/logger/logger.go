package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger 全局日志实例
var Logger *logrus.Logger

// Init 初始化日志
func Init(level string) {
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别
	switch level {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

// Debug 调试日志
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Info 信息日志
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Warn 警告日志
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Error 错误日志
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Fatal 致命错误日志
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}