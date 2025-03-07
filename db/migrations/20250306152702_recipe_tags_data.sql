-- +goose Up
-- +goose StatementBegin
-- Яичница (id 1)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (1, 1);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (1, 4);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (1, 31);

-- Паста Карбонара (id 2)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (2, 20);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (2, 3);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (2, 11);

-- Овощной рататуй (id 3)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (3, 21);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (3, 9);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (3, 23);

-- Куриный суп с лапшой (id 4)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (4, 15);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (4, 2);

-- Салат Цезарь с курицей (id 5)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (5, 2);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (5, 11);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (5, 26);

-- Том Ям с креветками (id 6)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (6, 13);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (6, 15);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (6, 32);

-- Вегетарианский борщ (id 7)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (7, 9);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (7, 15);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (7, 23);

-- Пельмени с мясом (id 8)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (8, 11);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (8, 27);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (8, 31);

-- Запеченная рыба с лимоном (id 9)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (9, 12);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (9, 3);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (9, 23);

-- Овсяная каша с ягодами (id 10)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (10, 1);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (10, 22);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (10, 26);

-- Стейк из говядины с овощами (id 11)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (11, 11);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (11, 27);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (11, 14);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (11, 3);

-- Крем-суп из тыквы (id 12)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (12, 15);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (12, 23);

-- Сырная лазанья (id 13)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (13, 3);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (13, 27);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (13, 31);

-- Чизкейк Нью-Йоркский (id 20)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (20, 17);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (20, 6);

-- Ризотто с грибами (id 14)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (14, 3);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (14, 23);

-- Мексиканские тако с говядиной (id 15)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (15, 11);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (15, 18);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (15, 7);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (15, 3);

-- Сливочный грибной суп (id 16)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (16, 15);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (16, 23);

-- Фалафель с соусом тахини (id 17)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (17, 10);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (17, 9);

-- Блины с творогом (id 18)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (18, 1);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (18, 6);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (18, 31);

-- Плов с бараниной (id 19)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (19, 11);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (19, 27);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (19, 2);

-- Рамен с курицей (id 21)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (21, 15);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (21, 3);

-- Шакшука (id 22)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (22, 1);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (22, 29);

-- Бургер с картофелем фри (id 23)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (23, 18);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (23, 11);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (23, 2);

-- Крем-брюле (id 24)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (24, 17);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (24, 6);

-- Фруктовый салат (id 25)
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (25, 22);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (25, 26);
INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (25, 8);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM recipe_tags;
-- +goose StatementEnd
