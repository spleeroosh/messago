package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/pkg/logger"
)

type Routes struct {
	logger    logger.Logger
	messages  Messages
	wsService WebsocketService
}

func NewRoutes(logger logger.Logger, messages Messages, wsService WebsocketService) *Routes {
	return &Routes{
		logger:    logger,
		messages:  messages,
		wsService: wsService,
	}
}

func (r *Routes) Apply(e *gin.Engine) {
	g := e.Group("/ws")
	g.GET("/chat", r.WebsocketHandler)
	g.GET("/messages", r.GetMessagesHandler)
	r.logger.Info().Msg("ws routers are registered")
}
