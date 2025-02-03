package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/config"
	"github.com/spleeroosh/messago/internal/pkg/routerfx"
)

func newHTTPRouter(conf config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	return routerfx.New(
		conf.App.Name,
	)
}
