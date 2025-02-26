-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tags
(
    id         SERIAL PRIMARY KEY,
    name       TEXT      NOT NULL,
    emoji      TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tags;
-- +goose StatementEnd
