package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/config"
)

type Routes struct {
	messages Messages
}

func NewRoutes(conf config.Config, messages Messages) *Routes {
	return &Routes{
		messages: messages,
	}
}

func (r *Routes) Apply(e *gin.Engine) {
	g := e.Group("/ws")
	g.GET("/chat", r.WebsocketHandler)
	g.GET("/messages", r.GetMessagesHandler)
	fmt.Println("routers are registered")
}
