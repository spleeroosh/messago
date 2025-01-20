package main

import (
	messagoHttp "github.com/spleeroosh/messago/internal/infrastructure/http"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(messagoHttp.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
