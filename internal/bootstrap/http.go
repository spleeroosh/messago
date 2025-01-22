package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/config"
	"go.uber.org/fx"
	"golang.org/x/net/context"
	"net/http"
)

// Создание нового HTTP-сервера
func newHTTPServer(lc fx.Lifecycle, conf config.Config) *http.Server {
	// Установите режим Gin (например, ReleaseMode или DebugMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Пример использования переменной из конфигурации
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("pong from %s", conf.App.Name),
		})
	})

	// Создание HTTP-сервера с указанным портом
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.App.Port), // Использование порта из конфигурации
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("Starting HTTP server %s at :%d\n", conf.App.Name, conf.App.Port)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					fmt.Printf("HTTP server error: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server")
			return server.Shutdown(ctx)
		},
	})

	return server
}
