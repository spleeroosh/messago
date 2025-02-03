package application

import "context"

type Server interface {
	Name() string
	Start() error
	Stop(ctx context.Context) error
}
