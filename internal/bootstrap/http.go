package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/config"
	"github.com/spleeroosh/messago/internal/pkg/application"
	"github.com/spleeroosh/messago/internal/pkg/routerfx"
	serverfx2 "github.com/spleeroosh/messago/internal/pkg/serverfx"
	"go.uber.org/fx"
)

// Создание нового HTTP-сервера
func newHTTPServer(lc fx.Lifecycle, sh fx.Shutdowner, engine *gin.Engine, conf config.Config, router *routerfx.AppRoute) *serverfx2.ServerFX {
	// Установите режим Gin (например, ReleaseMode или DebugMode)
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("HTTP SERVER START")
	// Настройка роутера
	router.SetupRouter(engine)

	// Создание HTTP сервера
	srv := serverfx2.New(
		conf.App.Name,
		serverfx2.Handler(engine.Handler()),
		serverfx2.Port(conf.App.Port),
	)

	lc.Append(application.ServerHooks(sh, srv))
	fmt.Println("HTTP SERVER END")
	return srv
}
