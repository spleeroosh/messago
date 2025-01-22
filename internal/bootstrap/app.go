package bootstrap

import (
	"github.com/spleeroosh/messago/internal/config"
	"go.uber.org/fx"
	"net/http"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(providers...),
		fx.Invoke(func(*http.Server) {}),
	)
}

var providers = []any{
	config.GetConfig,
	newHTTPServer,
	newPostgresClient,
}
