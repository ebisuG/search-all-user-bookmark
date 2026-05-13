package log

import "log"

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelError
)

type Logger struct {
	level LogLevel
}

func (l *Logger) Error(msg string) {
	if LevelError <= l.level {
		log.Println("[ERROR]", msg)
	}
}

func (l *Logger) Debug(msg string) {
	if l.level <= LevelDebug {
		log.Println("[DEBUG]", msg)
	}
}

func ParseLevel(level string) LogLevel {
	switch level {
	case "debug":
		return LevelDebug
	case "error":
		return LevelError
	default:
		return LevelDebug
	}
}
