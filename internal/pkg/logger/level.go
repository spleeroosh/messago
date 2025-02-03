package logger

import "github.com/rs/zerolog"

type Level string

const (
	LevelTrace    Level = "trace"
	LevelDebug    Level = "debug"
	LevelInfo     Level = "info"
	LevelWarn     Level = "warn"
	LevelError    Level = "error"
	LevelFatal    Level = "fatal"
	LevelPanic    Level = "panic"
	LevelDisabled Level = "disabled"
)

func (l Level) ZerologLevel() zerolog.Level {
	switch l {
	case LevelTrace:
		return zerolog.TraceLevel
	case LevelDebug:
		return zerolog.DebugLevel
	case LevelInfo:
		return zerolog.InfoLevel
	case LevelWarn:
		return zerolog.WarnLevel
	case LevelError:
		return zerolog.ErrorLevel
	case LevelFatal:
		return zerolog.FatalLevel
	case LevelPanic:
		return zerolog.PanicLevel
	case LevelDisabled:
		return zerolog.Disabled
	default:
		return zerolog.InfoLevel
	}
}
