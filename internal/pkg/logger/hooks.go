package logger

import (
	"time"

	"github.com/rs/zerolog"
)

type timestampHook struct{}

func (timestampHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	e.Time(zerolog.TimestampFieldName, time.Now())
}

type errorCounterHook struct{}

func (errorCounterHook) Run(_ *zerolog.Event, level zerolog.Level, _ string) {
	switch level {
	case zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel:
		errorCounter.WithLabelValues().Inc()
	}
}
