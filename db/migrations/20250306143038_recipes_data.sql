-- +goose Up
-- +goose StatementBegin
INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES (
           'Яичница', 350, 15, 200,
           '[{"name": "Яйца", "quantity": "2 шт."}, {"name": "Хлеб", "quantity": "2 ломтика"}]'::jsonb,
           '[{"step": 1, "text": "Поджарить хлеб до золотистой корочки."}, {"step": 2, "text": "Приготовить яичницу из яиц."}]'::jsonb,
           'https://topfood.club/assets/components/phpthumbof/cache/2024-08-12-qubadr-2022-06-15-y4wj2n-glazunya-1.2ceee6ffad271e3be1f7409fa7a41c20.jpg'
       );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Паста Карбонара',
        550,
        30,
        250,
        '[{"name": "Спагетти", "quantity": "200 г"}, {"name": "Яйца", "quantity": "2 шт"}, {"name": "Пармезан", "quantity": "50 г"}, {"name": "Бекон", "quantity": "100 г"}, {"name": "Чёрный перец", "quantity": "по вкусу"}, {"name": "Соль", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Отварить спагетти в подсоленной воде до состояния аль денте."}, {"step": 2, "text": "Обжарить бекон до хрустящей корочки."}, {"step": 3, "text": "Взбить яйца с тертым пармезаном и перцем."}, {"step": 4, "text": "Соединить пасту с беконом и яично-сырной смесью."}]'::jsonb,
        'https://www.patee.ru/r/x6/11/70/8e/960m.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Овощной рататуй',
        300,
        45,
        200,
        '[{"name": "Баклажан", "quantity": "1 шт"}, {"name": "Кабачок", "quantity": "1 шт"}, {"name": "Помидоры", "quantity": "3 шт"}, {"name": "Перец", "quantity": "1 шт"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Чеснок", "quantity": "2 зубчика"}, {"name": "Оливковое масло", "quantity": "2 ст. ложки"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Чёрный перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Нарезать овощи крупными кусками."}, {"step": 2, "text": "Обжарить лук и чеснок до мягкости."}, {"step": 3, "text": "Добавить баклажан и кабачок, тушить несколько минут."}, {"step": 4, "text": "Ввести помидоры и перец, приправить солью и перцем."}, {"step": 5, "text": "Тушить до полной готовности овощей."}]'::jsonb,
        'https://eco-buffet.com/wp-content/uploads/2023/02/ratatuj.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Куриный суп с лапшой',
        400,
        60,
        180,
        '[{"name": "Курица", "quantity": "500 г"}, {"name": "Лапша", "quantity": "150 г"}, {"name": "Морковь", "quantity": "1 шт"}, {"name": "Сельдерей", "quantity": "2 стебля"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Чеснок", "quantity": "2 зубчика"}, {"name": "Лавровый лист", "quantity": "2 шт"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Чёрный перец", "quantity": "по вкусу"}, {"name": "Петрушка", "quantity": "для украшения"}]'::jsonb,
        '[{"step": 1, "text": "Варить курицу в большом объеме воды до готовности."}, {"step": 2, "text": "Добавить лук, морковь и сельдерей, варить до мягкости овощей."}, {"step": 3, "text": "Ввести лапшу и лавровый лист, продолжать варить до готовности лапши."}, {"step": 4, "text": "Приправить солью и перцем, украсить свежей петрушкой."}]'::jsonb,
        'https://img.iamcook.ru/old/upl/recipes/cat/u1378-5ab5e94a53aabb5f55709f23c3521a2d.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Салат Цезарь с курицей',
        300,
        20,
        150,
        '[{"name": "Ромэн", "quantity": "200 г"}, {"name": "Куриное филе", "quantity": "150 г"}, {"name": "Пармезан", "quantity": "50 г"}, {"name": "Крутоны", "quantity": "100 г"}, {"name": "Цезарь соус", "quantity": "50 мл"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить куриное филе до готовности."}, {"step": 2, "text": "Нарезать ромэн и выложить в миску."}, {"step": 3, "text": "Добавить курицу, крутоны и тертый пармезан."}, {"step": 4, "text": "Полить соусом и перемешать."}]'::jsonb,
        'https://img.iamcook.ru/old/upl/recipes/misc/mobile/d6ae16b2d755bb0c5e4384f46ea71d0d.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Том Ям с креветками',
        350,
        40,
        300,
        '[{"name": "Креветки", "quantity": "200 г"}, {"name": "Грибы", "quantity": "100 г"}, {"name": "Лемongрасс", "quantity": "1 стебель"}, {"name": "Лайм", "quantity": "1 шт"}, {"name": "Кафр лайм", "quantity": "3 листа"}, {"name": "Чили", "quantity": "2 шт"}, {"name": "Рыбный соус", "quantity": "30 мл"}]'::jsonb,
        '[{"step": 1, "text": "Довести до кипения бульон с лемongрасс и кафр лайм."}, {"step": 2, "text": "Добавить грибы и чили, варить 5 минут."}, {"step": 3, "text": "Добавить креветки и сок лайма, варить до готовности."}, {"step": 4, "text": "Приправить рыбным соусом и подать горячим."}]'::jsonb,
        'https://i.ytimg.com/vi/AICQdRi_8tM/maxresdefault.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Вегетарианский борщ',
        250,
        90,
        200,
        '[{"name": "Свекла", "quantity": "2 шт"}, {"name": "Капуста", "quantity": "300 г"}, {"name": "Картофель", "quantity": "3 шт"}, {"name": "Морковь", "quantity": "1 шт"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Томатная паста", "quantity": "2 ст. ложки"}, {"name": "Чеснок", "quantity": "2 зубчика"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить лук и морковь до мягкости."}, {"step": 2, "text": "Добавить свеклу и тушить 10 минут."}, {"step": 3, "text": "Добавить картофель и капусту, залить водой."}, {"step": 4, "text": "Ввести томатную пасту, чеснок, соль и перец, варить до готовности овощей."}]'::jsonb,
        'https://bottva.ru/image/cache/catalog/2619_max-660x425.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Пельмени с мясом',
        500,
        60,
        220,
        '[{"name": "Мука", "quantity": "300 г"}, {"name": "Вода", "quantity": "150 мл"}, {"name": "Мясной фарш", "quantity": "400 г"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Масло", "quantity": "для подачи"}]'::jsonb,
        '[{"step": 1, "text": "Замесить тесто из муки и воды."}, {"step": 2, "text": "Приготовить начинку из фарша с мелко нарезанным луком и солью."}, {"step": 3, "text": "Раскатать тесто и вырезать кружки."}, {"step": 4, "text": "Заполнить кружки начинкой и защипнуть края."}, {"step": 5, "text": "Отварить пельмени до готовности и подать с маслом."}]'::jsonb,
        'https://img.7dach.ru/image/600/02/65/28/2013/11/05/59be9f.png'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Запеченная рыба с лимоном',
        400,
        35,
        280,
        '[{"name": "Филе рыбы", "quantity": "400 г"}, {"name": "Лимон", "quantity": "1 шт"}, {"name": "Оливковое масло", "quantity": "2 ст. ложки"}, {"name": "Чеснок", "quantity": "2 зубчика"}, {"name": "Тимьян", "quantity": "1 ч. ложка"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Замариновать рыбу с лимонным соком, чесноком, тимьяном, солью и перцем."}, {"step": 2, "text": "Выложить рыбу в форму для запекания, полить оливковым маслом."}, {"step": 3, "text": "Запекать в духовке при 200°C 25-30 минут."}, {"step": 4, "text": "Подавать с ломтиками лимона и зеленью."}]'::jsonb,
        'https://static.1000.menu/img/content-v2/98/de/17649/krasnaya-ryba-v-folge-v-duxovke_1586057835_12_max.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Овсяная каша с ягодами',
        350,
        10,
        100,
        '[{"name": "Овсяные хлопья", "quantity": "100 г"}, {"name": "Молоко", "quantity": "200 мл"}, {"name": "Ягоды", "quantity": "100 г"}, {"name": "Мёд", "quantity": "1 ст. ложка"}, {"name": "Орехи", "quantity": "30 г"}]'::jsonb,
        '[{"step": 1, "text": "Смешать овсяные хлопья с молоком и довести до кипения."}, {"step": 2, "text": "Добавить ягоды и варить 5 минут на медленном огне."}, {"step": 3, "text": "Подавать, украсив мёдом и орехами."}]'::jsonb,
        'https://www.patee.ru/r/x6/05/0f/ce/960m.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Стейк из говядины с овощами',
        650,
        50,
        350,
        '[{"name": "Говяжий стейк", "quantity": "300 г"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Чёрный перец", "quantity": "по вкусу"}, {"name": "Оливковое масло", "quantity": "2 ст. ложки"}, {"name": "Овощи (перец, цукини, баклажан)", "quantity": "300 г"}]'::jsonb,
        '[{"step": 1, "text": "Приправить стейк солью и перцем."}, {"step": 2, "text": "Обжарить стейк на раскаленной сковороде до желаемой степени готовности."}, {"step": 3, "text": "Обжарить овощи до мягкости."}, {"step": 4, "text": "Выложить стейк и овощи на тарелку, полить оливковым маслом."}]'::jsonb,
        'https://steakandgrill.ru/wp-content/uploads/2020/03/2-stejk-s-ovoshami-i-sousom-1024x682.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Крем-суп из тыквы',
        300,
        40,
        180,
        '[{"name": "Тыква", "quantity": "500 г"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Морковь", "quantity": "1 шт"}, {"name": "Чеснок", "quantity": "2 зубчика"}, {"name": "Куриный бульон", "quantity": "500 мл"}, {"name": "Сливки", "quantity": "100 мл"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить лук, морковь и чеснок до мягкости."}, {"step": 2, "text": "Добавить тыкву и тушить 10 минут."}, {"step": 3, "text": "Залить бульоном и варить до мягкости тыквы."}, {"step": 4, "text": "Взблендеровать суп до однородной массы."}, {"step": 5, "text": "Добавить сливки, посолить и поперчить."}]'::jsonb,
        'https://e52e3ee2-628b-49a9-9e26-e5a61fd72b20.selcdn.net/upload/resize_cache/iblock/005/1920_1080_1/1%20%D0%A2%D1%8B%D0%BA%D0%B2%D0%B5%D0%BD%D0%BD%D1%8B%D0%B9%20%D0%BA%D1%80%D0%B5%D0%BC-%D1%81%D1%83%D0%BF.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Сырная лазанья',
        700,
        90,
        400,
        '[{"name": "Листы лазаньи", "quantity": "250 г"}, {"name": "Моцарелла", "quantity": "200 г"}, {"name": "Пармезан", "quantity": "50 г"}, {"name": "Томатный соус", "quantity": "300 мл"}, {"name": "Базилик", "quantity": "несколько листьев"}, {"name": "Чеснок", "quantity": "2 зубчика"}, {"name": "Оливковое масло", "quantity": "2 ст. ложки"}, {"name": "Соль", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Смешать томатный соус с измельченным чесноком и базиликом."}, {"step": 2, "text": "Выложить слой соуса на дно формы для запекания."}, {"step": 3, "text": "Чередовать слои лазаньи и сыра, поливая соусом."}, {"step": 4, "text": "Запекать в духовке при 180°C около 45 минут."}]'::jsonb,
        'https://www.patee.ru/r/x6/11/8a/9d/960m.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Ризотто с грибами',
        550,
        50,
        300,
        '[{"name": "Арборио рис", "quantity": "200 г"}, {"name": "Грибы", "quantity": "300 г"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Белое вино", "quantity": "100 мл"}, {"name": "Овощной бульон", "quantity": "600 мл"}, {"name": "Пармезан", "quantity": "50 г"}, {"name": "Сливочное масло", "quantity": "2 ст. ложки"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить лук до прозрачности."}, {"step": 2, "text": "Добавить рис и обжарить до легкой золотистости."}, {"step": 3, "text": "Влить белое вино и дать ему испариться."}, {"step": 4, "text": "Постепенно добавить горячий бульон, постоянно помешивая."}, {"step": 5, "text": "Добавить грибы и варить до мягкости риса."}, {"step": 6, "text": "Вмешать сливочное масло и тертый пармезан, приправить солью и перцем."}]'::jsonb,
        'https://cdn.food.ru/unsigned/fit/640/480/ce/0/czM6Ly9tZWRpYS9waWN0dXJlcy9yZWNpcGVzLzQzOC9jb3ZlcnMvM1YycXBULkpQRw.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Мексиканские тако с говядиной',
        600,
        35,
        220,
        '[{"name": "Говяжий фарш", "quantity": "300 г"}, {"name": "Тако лепешки", "quantity": "6 шт"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Помидоры", "quantity": "2 шт"}, {"name": "Кукуруза", "quantity": "100 г"}, {"name": "Чили", "quantity": "1 шт"}, {"name": "Кинза", "quantity": "несколько веточек"}, {"name": "Лайм", "quantity": "1 шт"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить говяжий фарш с луком до золотистости."}, {"step": 2, "text": "Добавить нарезанные помидоры и кукурузу, тушить несколько минут."}, {"step": 3, "text": "Приправить чили, солью и перцем."}, {"step": 4, "text": "Подогреть лепешки, выложить начинку и украсить кинзой и дольками лайма."}]'::jsonb,
        'https://primebeef.ru/images/cms/data/tex-mex-beef-fajitas-61.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Сливочный грибной суп',
        320,
        45,
        200,
        '[{"name": "Грибы", "quantity": "300 г"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Чеснок", "quantity": "2 зубчика"}, {"name": "Картофель", "quantity": "2 шт"}, {"name": "Сливки", "quantity": "150 мл"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить нарезанные грибы, лук и чеснок до золотистого цвета."}, {"step": 2, "text": "Добавить нарезанный картофель и залить водой, варить до мягкости."}, {"step": 3, "text": "Измельчить суп блендером до пюреобразного состояния."}, {"step": 4, "text": "Вернуть суп на огонь, добавить сливки, приправить солью и перцем."}]'::jsonb,
        'https://static.1000.menu/img/content/30703/syrnyi-sup-s-gribami-v-multivarke_1545198478_1_max.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Фалафель с соусом тахини',
        450,
        60,
        180,
        '[{"name": "Нут", "quantity": "250 г"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Чеснок", "quantity": "3 зубчика"}, {"name": "Петрушка", "quantity": "пучок"}, {"name": "Кинза", "quantity": "пучок"}, {"name": "Специи (кориандр, тмин)", "quantity": "по вкусу"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Соус тахини", "quantity": "50 мл"}, {"name": "Оливковое масло", "quantity": "для жарки"}]'::jsonb,
        '[{"step": 1, "text": "Замочить нут на ночь."}, {"step": 2, "text": "Смешать нут с луком, чесноком, петрушкой, кинзой и специями в блендере."}, {"step": 3, "text": "Сформировать шарики и обжарить до золотистой корочки."}, {"step": 4, "text": "Подавать фалафель с соусом тахини."}]'::jsonb,
        'https://img.povar.ru/mobile/39/34/4f/c9/falafel_s_tahini-814677.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Блины с творогом',
        400,
        35,
        150,
        '[{"name": "Мука", "quantity": "200 г"}, {"name": "Молоко", "quantity": "300 мл"}, {"name": "Яйца", "quantity": "2 шт"}, {"name": "Сахар", "quantity": "1 ст. ложка"}, {"name": "Соль", "quantity": "щепотка"}, {"name": "Творог", "quantity": "200 г"}, {"name": "Ваниль", "quantity": "по вкусу"}, {"name": "Масло", "quantity": "для жарки"}]'::jsonb,
        '[{"step": 1, "text": "Приготовить тесто, смешав муку, молоко, яйца, сахар и соль."}, {"step": 2, "text": "Жарить блины до золотистого цвета."}, {"step": 3, "text": "Смешать творог с ванилью."}, {"step": 4, "text": "Подавать блины с творогом, свернув в рулет."}]'::jsonb,
        'https://img.iamcook.ru/2021/upl/recipes/cat/u-1cd85a3ebf3937b326b51a3870582298.JPG'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Плов с бараниной',
        700,
        90,
        350,
        '[{"name": "Баранина", "quantity": "500 г"}, {"name": "Рис", "quantity": "300 г"}, {"name": "Морковь", "quantity": "2 шт"}, {"name": "Лук", "quantity": "2 шт"}, {"name": "Чеснок", "quantity": "3 зубчика"}, {"name": "Специи (зира, барбарис)", "quantity": "по вкусу"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Оливковое масло", "quantity": "2 ст. ложки"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить баранину до золотистой корочки."}, {"step": 2, "text": "Добавить лук и морковь, тушить до мягкости."}, {"step": 3, "text": "Добавить рис, чеснок и специи, залить водой."}, {"step": 4, "text": "Тушить на медленном огне до готовности риса."}]'::jsonb,
        'https://i0.wp.com/justcookvrn.ru/wp-content/uploads/2022/10/plov-iz-baraniny.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Чизкейк Нью-Йоркский',
        500,
        120,
        300,
        '[{"name": "Сливочный сыр", "quantity": "500 г"}, {"name": "Печенье", "quantity": "200 г"}, {"name": "Сливочное масло", "quantity": "100 г"}, {"name": "Сахар", "quantity": "150 г"}, {"name": "Яйца", "quantity": "3 шт"}, {"name": "Ванильный экстракт", "quantity": "1 ч. ложка"}]'::jsonb,
        '[{"step": 1, "text": "Измельчить печенье и смешать с растопленным маслом, выложить основу в форму."}, {"step": 2, "text": "Взбить сливочный сыр с сахаром, добавить яйца и ванильный экстракт."}, {"step": 3, "text": "Вылить сырную массу на основу."}, {"step": 4, "text": "Запекать в духовке при 160°C около 70 минут, затем охладить."}]'::jsonb,
        'https://art-lunch.ru/content/uploads/2014/08/cheesecake-new-york-001x2-1.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Рамен с курицей',
        550,
        40,
        300,
        '[{"name": "Куриное филе", "quantity": "200 г"}, {"name": "Рамен лапша", "quantity": "1 упаковка"}, {"name": "Бульон", "quantity": "500 мл"}, {"name": "Соевый соус", "quantity": "2 ст. ложки"}, {"name": "Имбирь", "quantity": "1 кусочек"}, {"name": "Зелёный лук", "quantity": "пучок"}, {"name": "Яйцо", "quantity": "1 шт"}]'::jsonb,
        '[{"step": 1, "text": "Приготовить курицу в бульоне с соевым соусом и имбирем."}, {"step": 2, "text": "Отварить рамен лапшу согласно инструкции."}, {"step": 3, "text": "Собрать блюдо, добавив нарезанный зелёный лук и варёное яйцо."}]'::jsonb,
        'https://static.insales-cdn.com/images/articles/1/1729/6629057/delicious-ramen-on-dark-surface.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Шакшука',
        400,
        30,
        150,
        '[{"name": "Помидоры", "quantity": "4 шт"}, {"name": "Перец", "quantity": "1 шт"}, {"name": "Лук", "quantity": "1 шт"}, {"name": "Чеснок", "quantity": "3 зубчика"}, {"name": "Яйца", "quantity": "3 шт"}, {"name": "Оливковое масло", "quantity": "2 ст. ложки"}, {"name": "Паприка", "quantity": "1 ч. ложка"}, {"name": "Кумин", "quantity": "1 ч. ложка"}, {"name": "Соль", "quantity": "по вкусу"}, {"name": "Перец", "quantity": "по вкусу"}]'::jsonb,
        '[{"step": 1, "text": "Обжарить лук и чеснок до мягкости."}, {"step": 2, "text": "Добавить нарезанные помидоры и перец, тушить до загустения."}, {"step": 3, "text": "Добавить специи и аккуратно вбить яйца, готовить до желаемой консистенции."}]'::jsonb,
        'https://www.patee.ru/r/x6/01/39/cf/960m.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Бургер с картофелем фри',
        800,
        35,
        250,
        '[{"name": "Булочка для бургера", "quantity": "1 шт"}, {"name": "Говяжий котлет", "quantity": "150 г"}, {"name": "Лист салата", "quantity": "2 шт"}, {"name": "Помидор", "quantity": "1 шт"}, {"name": "Сыр", "quantity": "1 ломтик"}, {"name": "Картофель", "quantity": "200 г"}, {"name": "Майонез", "quantity": "30 г"}, {"name": "Кетчуп", "quantity": "30 г"}]'::jsonb,
        '[{"step": 1, "text": "Приготовить говяжий котлет до готовности."}, {"step": 2, "text": "Обжарить картофель фри до хрустящей корочки."}, {"step": 3, "text": "Собрать бургер: булочка, салат, котлет, сыр, ломтики помидора, соус."}, {"step": 4, "text": "Подавать бургер с картофелем фри."}]'::jsonb,
        'https://static.tildacdn.com/tild3561-3731-4732-b662-636461363134/image.png'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Крем-брюле',
        450,
        90,
        200,
        '[{"name": "Сливки", "quantity": "500 мл"}, {"name": "Сахар", "quantity": "100 г"}, {"name": "Яичные желтки", "quantity": "5 шт"}, {"name": "Ваниль", "quantity": "1 стручок"}]'::jsonb,
        '[{"step": 1, "text": "Нагреть сливки с ванилью до кипения, затем остудить."}, {"step": 2, "text": "Смешать желтки с сахаром и влить сливки."}, {"step": 3, "text": "Разлить смесь по формочкам и запекать на водяной бане."}, {"step": 4, "text": "Охладить, посыпать сахаром и карамелизовать сверху."}]'::jsonb,
        'https://static.1000.menu/img/content/22226/-krem-brule-francuzskii-krem-brule_1505309028_2_max.jpg'
    );

INSERT INTO recipes (name, calories, time, budget, ingredients, steps, imgsrc)
VALUES
    (
        'Фруктовый салат',
        250,
        15,
        120,
        '[{"name": "Яблоки", "quantity": "2 шт"}, {"name": "Груши", "quantity": "2 шт"}, {"name": "Киви", "quantity": "2 шт"}, {"name": "Апельсин", "quantity": "1 шт"}, {"name": "Мята", "quantity": "несколько веточек"}, {"name": "Мед", "quantity": "1 ст. ложка"}, {"name": "Лимонный сок", "quantity": "1 ст. ложка"}]'::jsonb,
        '[{"step": 1, "text": "Нарезать фрукты кубиками."}, {"step": 2, "text": "Смешать фрукты с медом и лимонным соком."}, {"step": 3, "text": "Добавить измельченную мяту и аккуратно перемешать."}]'::jsonb,
        'https://cdn.nur.kz/images/1200x675/4c7de76e38c02b7e.jpeg'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM recipes;
-- +goose StatementEnd
