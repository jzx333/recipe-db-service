-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recipes
(
    id          SERIAL PRIMARY KEY,
    name        TEXT      NOT NULL,
    calories    INTEGER   NOT NULL,
    time        INTEGER   NOT NULL,
    budget      INTEGER   NOT NULL,
    ingredients JSONB     NOT NULL,
    steps       JSONB     NOT NULL,
    imgsrc      TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recipes;
-- +goose StatementEnd
