package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spleeroosh/messago/internal/entity"
	"github.com/spleeroosh/messago/internal/pkg/logger"
	messagesService "github.com/spleeroosh/messago/internal/usecases/messages"
	"github.com/spleeroosh/messago/internal/utils"
	"github.com/spleeroosh/messago/internal/valueobject"
)

type Message struct {
	Client *websocket.Conn
	Data   []byte
}

type Service struct {
	logger          logger.Logger
	messageService  messagesService.Repository
	clients         map[*websocket.Conn]string
	incomingChannel chan Message
	outgoingChannel chan Message
}

func NewService(messageService messagesService.Repository, logger logger.Logger) *Service {
	return &Service{
		logger:          logger,
		messageService:  messageService,
		clients:         make(map[*websocket.Conn]string),
		incomingChannel: make(chan Message),
		outgoingChannel: make(chan Message),
	}
}

func (s *Service) HandleConnection(ctx context.Context, conn *websocket.Conn) error {
	nickname := utils.GenerateNickname()
	s.clients[conn] = nickname

	defer func() {
		delete(s.clients, conn)
		conn.Close()
	}()

	if err := s.SendLastMessages(ctx, conn); err != nil {
		s.logger.Err(err).Msg("failed to send last messages")
		return fmt.Errorf("failed to send last messages: %w", err)
	}

	// Запуск горутин для обработки сообщений
	go s.ReadMessages(conn)
	go s.BroadcastMessages()

	// Обработка новых сообщений через канал
	for msg := range s.incomingChannel {
		s.HandleIncomingMessage(ctx, msg.Client, msg.Data, nickname)
	}

	return nil
}

func (s *Service) ReadMessages(conn *websocket.Conn) {
	for {
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			s.logger.Err(err).Msg("Ошибка чтения сообщения")
			delete(s.clients, conn)
			return
		}

		// Отправляем сообщение в канал
		s.incomingChannel <- Message{Client: conn, Data: rawMessage}

		// Также пишем сообщение для других клиентов через исходящий канал
		s.outgoingChannel <- Message{Client: conn, Data: rawMessage}
	}
}

func (s *Service) HandleIncomingMessage(ctx context.Context, conn *websocket.Conn, rawMessage []byte, nickname string) error {
	var msg valueobject.Message
	if err := json.Unmarshal(rawMessage, &msg); err != nil {
		s.logger.Err(err).Msgf("failed to unmarshal raw message: %v", err)
		return fmt.Errorf("invalid message: %w", err)
	}

	message := valueobject.Message{
		Content: msg.Content,
		Sender:  nickname,
	}
	s.logger.Info().Msgf("received message %v", message)

	if err := s.messageService.SaveMessage(ctx, message); err != nil {
		s.logger.Error().Msgf("failed to save message: %v", err)
		return fmt.Errorf("save message failed: %w", err)
	}

	response, _ := json.Marshal(message)
	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		s.logger.Error().Msgf("failed to send response: %v", err)
		return fmt.Errorf("send confirmation failed: %w", err)
	}

	return nil
}

func (s *Service) GetAllMessages(ctx context.Context) ([]entity.Message, error) {
	return s.messageService.GetAllMessages(ctx)
}

func (s *Service) SendLastMessages(ctx context.Context, conn *websocket.Conn) error {
	messages, err := s.messageService.GetLatestMessages(ctx, 10)
	if err != nil {
		s.logger.Err(err).Msg("failed to get messages")
		return err
	}

	for _, message := range messages {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			s.logger.Err(err).Msg("failed to marshal message")
			return fmt.Errorf("serialization error: %w", err)
		}

		if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
			s.logger.Err(err).Msg("failed to send message")
			return fmt.Errorf("send message error: %w", err)
		}
	}
	return nil
}

func (s *Service) BroadcastMessages() {
	for msg := range s.outgoingChannel {
		for client := range s.clients {
			if client != msg.Client {
				err := client.WriteMessage(websocket.TextMessage, msg.Data)
				if err != nil {
					s.logger.Err(err).Msg("Ошибка отправки сообщения клиенту")
					client.Close()
					delete(s.clients, client)
				}
			}
		}
	}
}
