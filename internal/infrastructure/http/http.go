package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"golang.org/x/net/context"
	"net/http"
)

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	// Установите режим Gin (например, ReleaseMode или DebugMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting HTTP server at", "8080")
			go r.Run() // listen and serve on 0.0.0.0:8080
			return nil
		},
		//OnStop: func(ctx context.Context) error {
		//	return r.R
		//},
	})

	return nil
}
