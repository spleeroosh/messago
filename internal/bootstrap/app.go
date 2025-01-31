package bootstrap

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	wsapi "github.com/spleeroosh/messago/internal/api/ws"
	"github.com/spleeroosh/messago/internal/config"
	routerfx "github.com/spleeroosh/messago/internal/infrastructure/http/routerfx"
	serverfx "github.com/spleeroosh/messago/internal/infrastructure/http/serverfx"
	"github.com/spleeroosh/messago/internal/repository/messages"
	usemessages "github.com/spleeroosh/messago/internal/usecases/messages"
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
	fx.Annotate(messages.NewRepository, fx.As(new(usemessages.Repository))),
	config.GetConfig,
	newHTTPRouter,
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
