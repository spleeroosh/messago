package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/messago/internal/config"
)

type Routes struct {
}

func NewRoutes(conf config.Config) *Routes {
	return &Routes{}
}

func (r *Routes) Apply(e *gin.Engine) {
	g := e.Group("/ws")
	g.GET("/chat", r.WebsocketHandler)
	fmt.Println("routers are registered")
}
