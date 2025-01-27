package bootstrap

import (
	wsapi "github.com/spleeroosh/messago/internal/api/ws"
	"github.com/spleeroosh/messago/internal/config"
	routerfx "github.com/spleeroosh/messago/internal/infrastructure/http/routerfx"
	serverfx "github.com/spleeroosh/messago/internal/infrastructure/http/serverfx"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(providers...),
		fx.Invoke(func(*serverfx.ServerFX) {}),
	)
}

var providers = []any{
	config.GetConfig,
	newHTTPRouter,
	newHTTPServer,
	newPostgresClient,
	// Регистрация роутов
	fx.Annotate(wsapi.NewRoutes, fx.As(new(routerfx.Provider)), fx.ResultTags(`group:"providers"`)),
	// Коллектор роутов
	fx.Annotate(routerfx.NewRouter, fx.ParamTags(`group:"providers"`)),
}
