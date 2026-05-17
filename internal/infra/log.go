package infra

import (
	"log"

	logger "github.com/ebisuG/search-all-user-bookmark/internal/logger"
)

type Logger struct {
	level logger.LogLevel
}

func (l *Logger) Error(msg any) {
	if logger.LevelError <= l.level {
		log.Println("[ERROR]", msg)
	}
}

func (l *Logger) Debug(msg any) {
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

func NewLogger(level string) Logger {
	return Logger{level: Parse(level)}
}
