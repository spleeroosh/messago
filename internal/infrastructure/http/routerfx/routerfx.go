package routerfx

import (
	"github.com/gin-gonic/gin"
)

type Provider interface {
	Apply(engine *gin.Engine)
}

type AppRoute struct {
	providers []Provider
}

func NewRouter(providers ...Provider) *AppRoute {
	return &AppRoute{providers: providers}
}

func (r AppRoute) SetupRouter(engine *gin.Engine) {
	for _, provider := range r.providers {
		provider.Apply(engine)
	}
}
