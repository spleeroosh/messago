-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages
(
    id         BIGSERIAL PRIMARY KEY,
    content    TEXT        NOT NULL,
    sender     TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
