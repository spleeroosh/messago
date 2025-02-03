package ws

import (
	"context"
	"github.com/gorilla/websocket"

	"github.com/spleeroosh/messago/internal/entity"
	"github.com/spleeroosh/messago/internal/valueobject"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -destination=repository_mock_test.go -package=dictionary DictionaryRepository
type Messages interface {
	GetAllMessages(ctx context.Context) ([]entity.Message, error)
	GetLatestMessages(ctx context.Context, limit int) ([]entity.Message, error)
	SaveMessage(ctx context.Context, msg valueobject.Message) error
}

// WebsocketService describes the interface for working with the WebSocket service.
type WebsocketService interface {
	// HandleConnection handles the client connection and message management.
	HandleConnection(ctx context.Context, conn *websocket.Conn) error

	// HandleIncomingMessage processes the incoming message from the client.
	HandleIncomingMessage(ctx context.Context, conn *websocket.Conn, rawMessage []byte, nickname string) error

	// GetAllMessages retrieves all messages.
	GetAllMessages(ctx context.Context) ([]entity.Message, error)

	// SendLastMessages sends the last messages to the user.
	SendLastMessages(ctx context.Context, conn *websocket.Conn) error

	// BroadcastMessage broadcasts a message to all clients.
	BroadcastMessage(sender *websocket.Conn, messageType int, jsonMessage []byte) error
}
