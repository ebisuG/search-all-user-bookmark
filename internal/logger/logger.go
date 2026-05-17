package logger

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelProduction
)

type Logger interface {
	Production(msg any)
	Debug(msg any)
}
