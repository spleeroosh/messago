package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spleeroosh/messago/internal/config"
	"go.uber.org/fx"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

// Настройка апгрейдера WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Добавьте обработку CORS (по умолчанию запросы от других источников блокируются)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(c *gin.Context) {
	// Обновляем соединение до WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Ошибка при апгрейде до WebSocket:", err)
		return
	}
	defer conn.Close()

	// Обработка сообщений
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}

		log.Printf("Получено сообщение: %s\n", message)

		// Отправка обратно (эхо)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Ошибка отправки сообщения:", err)
			break
		}
	}
}

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
	r.GET("/ws", websocketHandler)

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
