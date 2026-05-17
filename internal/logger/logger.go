package logger

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelError
)

type Logger interface {
	Error(msg any)
	Debug(msg any)
}
