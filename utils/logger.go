package utils

import (
	"github.com/sirupsen/logrus"
)

// Logger is a struct that encapsulates the logrus logger.
type Logger struct {
	logger *logrus.Logger
}

// NewLogger initializes and returns a new Logger instance with default settings.
func NewLogger() *Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	l.SetLevel(logrus.InfoLevel)
	return &Logger{logger: l}
}

// LogInfo logs an informational message.
func (l *Logger) LogInfo(message string) {
	l.logger.Info(message)
}

// LogError logs an error message.
func (l *Logger) LogError(message string) {
	l.logger.Error(message)
}
