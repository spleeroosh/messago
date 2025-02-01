package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/config"
	"github.com/spleeroosh/messago/internal/infrastructure/logger"
)

type Routes struct {
	logger   logger.Logger
	messages Messages
}

func NewRoutes(conf config.Config, logger logger.Logger, messages Messages) *Routes {
	return &Routes{
		logger:   logger,
		messages: messages,
	}
}

func (r *Routes) Apply(e *gin.Engine) {
	g := e.Group("/ws")
	g.GET("/chat", r.WebsocketHandler)
	g.GET("/messages", r.GetMessagesHandler)
	fmt.Println("routers are registered")
}
