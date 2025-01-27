package routerfx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

const CIHealthcheckHeader = "X-CI-Healthcheck"
const CIHealthcheckValue = "0c207c7b-e8f4-4332-b7a1-2fdddb3fd75b"

func New(appName string, opts ...Option) *gin.Engine {
	o := options{
		middlewares:            []gin.HandlerFunc{},
		handleMethodNotAllowed: true,
		enableContextFallback:  true,
		pprof:                  true,
		pprofPrefix:            "debug/pprof",
	}

	for _, opt := range opts {
		opt(&o)
	}

	engine := gin.New()
	engine.HandleMethodNotAllowed = o.handleMethodNotAllowed
	engine.ContextWithFallback = o.enableContextFallback

	engine.Use(o.middlewares...)

	if !o.disableDefaultRoutes {
		engine.GET("/", func(c *gin.Context) {
			c.Status(http.StatusNoContent)
		})
		engine.GET("/internal/version", func(c *gin.Context) {
			hostname, _ := os.Hostname()
			c.JSON(http.StatusOK, map[string]any{"data": struct {
				Hostname string `json:"hostname"`
				Commit   string `json:"commit"`
				Time     string `json:"time"`
			}{
				Hostname: hostname,
				Commit:   o.buildCommit,
				Time:     o.buildTime.In(time.UTC).Format(time.RFC3339),
			}})
		})
		engine.GET("/external/api/healthcheck", func(c *gin.Context) {
			if c.GetHeader(CIHealthcheckHeader) == CIHealthcheckValue {
				c.Status(http.StatusNoContent)
			} else {
				c.Status(http.StatusNotFound)
			}
		})
	}

	return engine
}
