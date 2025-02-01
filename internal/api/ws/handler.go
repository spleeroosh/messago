package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spleeroosh/messago/internal/valueobject"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]string)

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow cross-origin requests
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (r *Routes) GetMessagesHandler(c *gin.Context) {
	messages, err := r.messages.GetAllMessages(c)
	if err != nil {
		r.logger.Err(err).Msgf("Error retrieving messages: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

func (r *Routes) WebsocketHandler(c *gin.Context) {
	// Upgrade the connection to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		r.logger.Err(err).Msgf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	r.logger.Info().Msg("WS connection is upgraded")
	r.sendLastMessages(c, conn)
	r.messagesHandler(c, conn)
}

func (r *Routes) messagesHandler(c *gin.Context, conn *websocket.Conn) {
	clients[conn] = generateNickname()

	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			r.logger.Err(err).Msgf("Error reading message: %v", err)
			break
		}

		// Parse JSON into the message structure
		var msg valueobject.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			r.logger.Err(err).Msgf("Invalid JSON: %v", err)
			continue
		}

		log.Printf("Received message: %s\n", msg)
		response := valueobject.Message{
			Type:    "incoming",
			Content: msg.Content,
			Sender:  clients[conn],
		}

		err = r.messages.SaveMessage(c, response)
		if err != nil {
			r.logger.Err(err).Msgf("Error saving message in PostgreSQL: %v", err)
		}

		jsonMessage, err := json.Marshal(response)

		// Broadcast the message to all connected clients
		r.broadcastMessage(clients, conn, messageType, jsonMessage)
	}
}

func (r *Routes) sendLastMessages(c *gin.Context, conn *websocket.Conn) {
	// Retrieve the last 10 messages from the repository
	messages, err := r.messages.GetLatestMessages(c, 10)
	if err != nil {
		r.logger.Err(err).Msgf("Error retrieving last messages: %v", err)
		return
	}

	// Send the last messages to the newly connected client
	for _, message := range messages {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			r.logger.Err(err).Msgf("Error serializing message: %v", err)
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
			r.logger.Err(err).Msgf("Error sending message: %v", err)
			return
		}
	}
}

func (r *Routes) broadcastMessage(clients map[*websocket.Conn]string, sender *websocket.Conn, messageType int, jsonMessage []byte) {
	for client := range clients {
		if client != sender {
			if err := client.WriteMessage(messageType, jsonMessage); err != nil {
				r.logger.Err(err).Msgf("Error sending message to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
