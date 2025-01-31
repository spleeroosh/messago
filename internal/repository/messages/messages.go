package messages

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spleeroosh/messago/internal/entity"
	"github.com/spleeroosh/messago/internal/valueobject"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type Repository struct {
	db      *pgxpool.Pool
	tracer  trace.Tracer
	builder goqu.DialectWrapper
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:      db,
		tracer:  otel.Tracer("MessagesRepository"),
		builder: goqu.Dialect("postgres"),
	}
}

func (r *Repository) GetAllMessages(ctx context.Context) ([]entity.Message, error) {
	ctx, span := r.tracer.Start(ctx, "Repository.GetAllMessages()")
	defer span.End()

	sql, args, err := r.builder.
		Select("messages.id", "messages.sender", "messages.content", "messages.created_at").
		From("messages").
		ToSQL()

	if err != nil {
		return nil, fmt.Errorf("build sql: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query rows: %w", err)
	}
	defer rows.Close()

	var items []entity.Message
	for rows.Next() {
		var item entity.Message
		if err := rows.Scan(&item.ID, &item.Sender, &item.Content, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return items, nil
}

func (r *Repository) SaveMessage(ctx context.Context, msg valueobject.Message) error {
	ctx, span := r.tracer.Start(ctx, "Repository.SaveMessage()")
	defer span.End()

	now := time.Now()

	record := goqu.Record{
		"type":       msg.Type,
		"sender":     msg.Sender,
		"content":    msg.Content,
		"created_at": now,
	}

	sql, args, err := r.builder.Insert("messages").Rows(record).ToSQL()
	if err != nil {
		return fmt.Errorf("build sql: %w", err)
	}

	if _, err := r.db.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("exec sql: %w", err)
	}

	return nil
}
