package messages

import (
	"context"
	"fmt"
	"github.com/spleeroosh/messago/internal/entity"
	"github.com/spleeroosh/messago/internal/valueobject"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	repo   Repository
	tracer trace.Tracer
}

func NewService(repo Repository) *Service {
	return &Service{
		repo:   repo,
		tracer: otel.Tracer("DictionaryService"),
	}
}

func (s *Service) GetAllMessages(ctx context.Context) ([]entity.Message, error) {
	ctx, span := s.tracer.Start(ctx, "MessagesService:GetMessages()")
	defer span.End()

	items, err := s.repo.GetAllMessages(ctx)
	if err != nil {
		return nil, fmt.Errorf("get messages: %w", err) // Возвращаем nil в случае ошибки
	}

	return items, nil
}

func (s *Service) GetLatestMessages(ctx context.Context, limit int) ([]entity.Message, error) {
	ctx, span := s.tracer.Start(ctx, "MessagesService:GetLatestMessages()")
	defer span.End()

	items, err := s.repo.GetLatestMessages(ctx, limit)
	if err != nil {
		return nil, fmt.Errorf("get last messages: %w", err) // Возвращаем nil в случае ошибки
	}

	return items, nil
}

func (s *Service) SaveMessage(ctx context.Context, msg valueobject.Message) error {
	ctx, span := s.tracer.Start(ctx, "MessagesService:SaveMessage()")
	defer span.End()

	err := s.repo.SaveMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("save message: %w", err) // Возвращаем nil в случае ошибки
	}

	return nil
}
