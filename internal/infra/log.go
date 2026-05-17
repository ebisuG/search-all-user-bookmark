package infra

import (
	"log"

	logger "github.com/ebisuG/search-all-user-bookmark/internal/logger"
)

type loggerImpl struct {
	level logger.LogLevel
}

var _ logger.Logger = (*loggerImpl)(nil)

func (l *loggerImpl) Error(msg any) {
	if logger.LevelError <= l.level {
		log.Println("[ERROR]", msg)
	}
}

func (l *loggerImpl) Debug(msg any) {
	if l.level <= logger.LevelDebug {
		log.Println("[DEBUG]", msg)
	}
}

func Parse(level string) logger.LogLevel {
	switch level {
	case "debug":
		return logger.LevelDebug
	case "error":
		return logger.LevelError
	default:
		return logger.LevelDebug
	}
}

func NewLogger(level string) logger.Logger {
	return &loggerImpl{level: Parse(level)}
}
