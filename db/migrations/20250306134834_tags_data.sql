-- +goose Up
-- +goose StatementBegin
INSERT INTO tags (name, emoji) VALUES ('Завтрак', '🌅');
INSERT INTO tags (name, emoji) VALUES ('Обед', '🍽️');
INSERT INTO tags (name, emoji) VALUES ('Ужин', '🌙');
INSERT INTO tags (name, emoji) VALUES ('Жареное', '🔥');
INSERT INTO tags (name, emoji) VALUES ('Из муки', '🌾');
INSERT INTO tags (name, emoji) VALUES ('Сладкое', '🍭');
INSERT INTO tags (name, emoji) VALUES ('Острое', '🌶️');
INSERT INTO tags (name, emoji) VALUES ('Диетическое', '🥗');
INSERT INTO tags (name, emoji) VALUES ('Вегетарианское', '🥦');
INSERT INTO tags (name, emoji) VALUES ('Веганское', '🌱');
INSERT INTO tags (name, emoji) VALUES ('Мясное', '🍖');
INSERT INTO tags (name, emoji) VALUES ('Рыбное', '🐟');
INSERT INTO tags (name, emoji) VALUES ('Морепродукты', '🦐');
INSERT INTO tags (name, emoji) VALUES ('Гриль', '🥘');
INSERT INTO tags (name, emoji) VALUES ('Суп', '🥣');
INSERT INTO tags (name, emoji) VALUES ('Напитки', '🥤');
INSERT INTO tags (name, emoji) VALUES ('Десерт', '🍰');
INSERT INTO tags (name, emoji) VALUES ('Фастфуд', '🍔');
INSERT INTO tags (name, emoji) VALUES ('Пицца', '🍕');
INSERT INTO tags (name, emoji) VALUES ('Паста', '🍝');
INSERT INTO tags (name, emoji) VALUES ('Овощное', '🥕');
INSERT INTO tags (name, emoji) VALUES ('Фруктовое', '🍎');
INSERT INTO tags (name, emoji) VALUES ('Здоровое питание', '💪');
INSERT INTO tags (name, emoji) VALUES ('Без глютена', '🚫');
INSERT INTO tags (name, emoji) VALUES ('Без сахара', '🥛');
INSERT INTO tags (name, emoji) VALUES ('Легкое', '🍃');
INSERT INTO tags (name, emoji) VALUES ('Тяжелая еда', '🥩');
INSERT INTO tags (name, emoji) VALUES ('На компанию', '👨‍👩‍👧‍👦');
INSERT INTO tags (name, emoji) VALUES ('Экспресс', '⏳');
INSERT INTO tags (name, emoji) VALUES ('Для детей', '🧒');
INSERT INTO tags (name, emoji) VALUES ('Домашняя кухня', '🏠');
INSERT INTO tags (name, emoji) VALUES ('Экзотическое', '🌴');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tags;
-- +goose StatementEnd
