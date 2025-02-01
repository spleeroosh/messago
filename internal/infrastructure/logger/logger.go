package logger

import (
	"context"
	"io"
	stdlog "log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

type Logger interface {
	Debug() *zerolog.Event
	Err(error) *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	GetLevel() zerolog.Level
	Hook(...zerolog.Hook) zerolog.Logger
	Info() *zerolog.Event
	Level(zerolog.Level) zerolog.Logger
	Log() *zerolog.Event
	Output(io.Writer) zerolog.Logger
	Panic() *zerolog.Event
	Print(...interface{})
	Printf(string, ...interface{})
	Sample(zerolog.Sampler) zerolog.Logger
	Trace() *zerolog.Event
	UpdateContext(func(c zerolog.Context) zerolog.Context)
	Warn() *zerolog.Event
	With() zerolog.Context
	WithContext(context.Context) context.Context
	WithLevel(zerolog.Level) *zerolog.Event
	Write([]byte) (int, error)
}

type LoggerImpl struct {
	zerolog.Logger
}

var mtx sync.Mutex

// NewLogger creates logger that should be used by default.
func NewLogger(appName string, opts ...Option) *LoggerImpl {
	registerMetrics(strings.ToLower(strings.ReplaceAll(appName, "-", "_")))

	o := &options{
		writer:            os.Stdout,
		overrideStdLogOut: false,
	}
	for _, opt := range opts {
		opt(o)
	}

	if o.prettify {
		o.writer = zerolog.NewConsoleWriter()
	}

	l := zerolog.New(o.writer).
		Level(o.level.ZerologLevel()).
		Hook(errorCounterHook{})

	if !o.noTimestamp {
		mtx.Lock()
		zerolog.TimeFieldFormat = time.RFC3339Nano
		mtx.Unlock()
		l = l.Hook(timestampHook{})
	}

	c := l.With()

	if o.env != "" {
		c = c.Str("env", o.env)
	}
	if o.buildCommit != "" {
		c = c.Str("build_commit", o.buildCommit)
	}
	if !o.buildTime.IsZero() {
		c = c.Str("build_time", o.buildTime.Format(time.RFC3339))
	}

	logger := &LoggerImpl{c.Logger()}

	if o.overrideStdLogOut {
		// Use the logger as an output for stdlog.
		// https://github.com/rs/zerolog#set-as-standard-logger-output
		stdlog.SetFlags(0)
		stdlog.SetOutput(logger)
	}

	return logger
}

// NewDiscardLogger creates logger that writes nothing.
// Useful in testing and as default value for loggers.
func NewDiscardLogger() *LoggerImpl {
	return &LoggerImpl{zerolog.Nop()}
}

// NewNonBlockingWriter writes logs quickly but with no guarantee.
func NewNonBlockingWriter(w io.Writer, size int, pollInterval time.Duration, fallback Logger) io.Writer {
	return diode.NewWriter(w, size, pollInterval, func(missed int) {
		fallback.Warn().Msgf("non-blocking log writer has dropped %v message(s)", missed)
	})
}

func WithName(logger Logger, name string) *LoggerImpl {
	return &LoggerImpl{logger.With().Str("logger", name).Logger()}
}
