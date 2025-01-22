-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT        NOT NULL,
    age        INT         NOT NULL,
    social     TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
