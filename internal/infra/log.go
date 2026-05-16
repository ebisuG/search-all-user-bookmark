package infra

import (
	"log"

	logger "github.com/ebisuG/search-all-user-bookmark/internal/logger"
)

type Logger struct {
	level logger.LogLevel
}

func (l *Logger) Error(msg string) {
	if logger.LevelError <= logger.LevelDebug {
		log.Println("[ERROR]", msg)
	}
}

func (l *Logger) Debug(msg string) {
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
