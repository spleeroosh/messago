package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/config"
	"github.com/spleeroosh/messago/internal/infrastructure/http/application"
	"github.com/spleeroosh/messago/internal/infrastructure/http/routerfx"
	"github.com/spleeroosh/messago/internal/infrastructure/http/serverfx"
	"go.uber.org/fx"
)

// Создание нового HTTP-сервера
func newHTTPServer(lc fx.Lifecycle, sh fx.Shutdowner, engine *gin.Engine, conf config.Config, router *routerfx.AppRoute) *serverfx.ServerFX {
	// Установите режим Gin (например, ReleaseMode или DebugMode)
	gin.SetMode(gin.ReleaseMode)

	// Настройка роутера
	router.SetupRouter(engine)

	// Создание HTTP сервера
	srv := serverfx.New(
		conf.App.Name,
		serverfx.Handler(engine.Handler()),
		serverfx.Port(conf.App.Port),
	)

	lc.Append(application.ServerHooks(sh, srv))

	return srv
}
