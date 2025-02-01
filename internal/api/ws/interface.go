package ws

import (
	"context"

	"github.com/spleeroosh/messago/internal/entity"
	"github.com/spleeroosh/messago/internal/valueobject"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -destination=repository_mock_test.go -package=dictionary DictionaryRepository
type Messages interface {
	GetAllMessages(ctx context.Context) ([]entity.Message, error)
	GetLatestMessages(ctx context.Context, limit int) ([]entity.Message, error)
	SaveMessage(ctx context.Context, msg valueobject.Message) error
}
