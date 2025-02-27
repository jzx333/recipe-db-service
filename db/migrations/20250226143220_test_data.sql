-- +goose Up
-- +goose StatementBegin
INSERT INTO tags (name, emoji) VALUES ('Жареное', '🔥');
INSERT INTO tags (name, emoji) VALUES ('Из муки', '🌾');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tags;
-- +goose StatementEnd
