package data

type FullStore struct {
	recipe RecipeDetails
}

func NewFullStore() *FullStore {
	var recipe = RecipeDetails{
		Id: 1, Name: "Блины", Calories: 150, Time: 175, Budget: 175, Tags: []Tag{
			{Name: "Из муки", Emoji: "🌾"}, {Name: "Жареное", Emoji: "🔥"}},
		Ingredients: []Ingredient{
			{Name: "Мука", Quantity: "300 г"}, {Name: "Яйцо", Quantity: "2 шт"},
			{Name: "Молоко", Quantity: "500 г"}, {Name: "Сахар", Quantity: "1 ст. л. = 25 г"},
			{Name: "Соль", Quantity: "0,5 ч. л. = 1 г"},
		},
		Steps: []Step{
			{Step: 1, Text: "В глубокую миску наливаем молоко, разбиваем яйцо"},
			{Step: 2, Text: "Добавляем муку, частями, тщательно перемешивая"},
		},
		ImgSrc: "path",
	}
	return &FullStore{recipe}
}

func (f FullStore) Get() (RecipeDetails, error) {
	return f.recipe, nil
}
