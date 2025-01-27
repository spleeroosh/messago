package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func (r *Routes) WebsocketHandler(c *gin.Context) {
	// Обновляем соединение до WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Ошибка при апгрейде до WebSocket:", err)
		return
	}
	defer conn.Close()

	messagesHandler(conn)
}

func messagesHandler(conn *websocket.Conn) {
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
