package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]string)

// Настройка апгрейдера WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Добавьте обработку CORS (по умолчанию запросы от других источников блокируются)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type    string `json:"type"`
	Sender  string `json:"sender"`
	Content string `json:"content"`
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
	clients[conn] = generateNickname()

	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}

		// Парсим JSON в структуру
		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Invalid JSON:", err)
			continue
		}

		log.Printf("Получено сообщение: %s\n", msg)
		response := Message{
			Type:    "icoming",
			Content: msg.Content,
			Sender:  clients[conn],
		}

		jsonMessage, err := json.Marshal(response)

		// Рассылка всем подключенным клиентам
		for client := range clients {
			if client != conn { // Не отправлять самому себе
				err := client.WriteMessage(messageType, jsonMessage)
				if err != nil {
					log.Println("Ошибка отправки сообщения:", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
