package routerfx

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.xbet.lan/web-backend/go/pkg/log"
)

type Option func(o *options)

type options struct {
	logger                 log.Logger
	env                    string
	middlewares            []gin.HandlerFunc
	handleMethodNotAllowed bool
	enableContextFallback  bool
	pprof                  bool
	pprofSecret            string
	pprofPrefix            string
	buildCommit            string
	buildTime              time.Time
	disableDefaultRoutes   bool
	prettyLog              bool
}

func PprofPrefix(prefix string) Option {
	return func(o *options) {
		o.pprofPrefix = prefix
	}
}

func Pprof(enable bool) Option {
	return func(o *options) {
		o.pprof = enable
	}
}

func PprofSecret(secret string) Option {
	return func(o *options) {
		o.pprofSecret = secret
	}
}

func Env(env string) Option {
	return func(o *options) {
		o.env = env
	}
}

func Logger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func BuildCommit(buildCommit string) Option {
	return func(o *options) {
		o.buildCommit = buildCommit
	}
}

func BuildTime(t time.Time) Option {
	return func(o *options) {
		o.buildTime = t
	}
}

func Middlewares(middlewares ...gin.HandlerFunc) Option {
	return func(o *options) {
		o.middlewares = middlewares
	}
}

// HandleMethodNotAllowed refers to https://github.com/gin-gonic/gin/blob/v1.8.1/gin.go#L107
func HandleMethodNotAllowed(handle bool) Option {
	return func(o *options) {
		o.handleMethodNotAllowed = handle
	}
}

// EnableContextFallback refers to https://github.com/gin-gonic/gin/blob/v1.8.1/gin.go#L151
func EnableContextFallback(enable bool) Option {
	return func(o *options) {
		o.enableContextFallback = enable
	}
}

func DisableDefaultRoutes(disable bool) Option {
	return func(o *options) {
		o.disableDefaultRoutes = disable
	}
}

func PrettyLog(prettyLog bool) Option {
	return func(o *options) {
		o.prettyLog = prettyLog
	}
}
