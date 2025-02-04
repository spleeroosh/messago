package bootstrap

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	wsapi "github.com/spleeroosh/messago/internal/api/ws"
	"github.com/spleeroosh/messago/internal/config"
	"github.com/spleeroosh/messago/internal/pkg/application"
	"github.com/spleeroosh/messago/internal/pkg/routerfx"
	"github.com/spleeroosh/messago/internal/pkg/serverfx"
	"github.com/spleeroosh/messago/internal/repository/messages"
	usemessages "github.com/spleeroosh/messago/internal/usecases/messages"
	usews "github.com/spleeroosh/messago/internal/usecases/websocket"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(providers...),
		fx.Invoke(registerHooks),
	)
}

var providers = []any{
	fx.Annotate(usemessages.NewService, fx.As(new(wsapi.Messages))),
	fx.Annotate(usews.NewService, fx.As(new(wsapi.WebsocketService))),
	fx.Annotate(messages.NewRepository, fx.As(new(usemessages.Repository))),
	newHTTPRouter,
	config.GetConfig,
	application.GetBuildVersion,
	newLogger,
	newHTTPServer,
	newPostgresClient,
	// Регистрация роутов
	fx.Annotate(wsapi.NewRoutes, fx.As(new(routerfx.Provider)), fx.ResultTags(`group:"providers"`)),
	// Коллектор роутов
	fx.Annotate(routerfx.NewRouter, fx.ParamTags(`group:"providers"`)),
}

func registerHooks(pool *pgxpool.Pool, server *serverfx.ServerFX) {
	fmt.Println("Postgres pool successfully initialized")
}
