package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func (r *Routes) WebsocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	err = r.wsService.HandleConnection(c.Request.Context(), conn)
	if err != nil {
		r.logger.Err(err).Msg("Failed to handle connection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection error"})
	}
}

func (r *Routes) GetMessagesHandler(c *gin.Context) {
	messages, err := r.wsService.GetAllMessages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": messages})
}
