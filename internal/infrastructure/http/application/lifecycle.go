package application

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

func ServerHooks(sh fx.Shutdowner, srv Server) fx.Hook {
	return fx.Hook{
		OnStart: ServerStartHook(sh, srv),
		OnStop:  ServerStopHook(srv),
	}
}

// ServerStartHook returns start hook for fx lifecycle. Calls Server.Start and shutdowns application upon receiving an error. Ignores http.ErrServerClosed
func ServerStartHook(sh fx.Shutdowner, srv Server) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		go func() {
			if err := srv.Start(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("server down %v\n", srv.Name())
				_ = sh.Shutdown()
			}
		}()
		return nil
	}
}

// ServerStopHook returns stop hook for fx lifecycle. Calls Server.Stop
func ServerStopHook(srv Server) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return srv.Stop(ctx)
	}
}
