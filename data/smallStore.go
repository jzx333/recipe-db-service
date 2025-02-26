package data

type SmallStore struct {
	list []RecipePreview
}

func NewSmallStore() *SmallStore {
	var list = []RecipePreview{
		{Id: 1, Name: "Блины", Time: 30, Budget: 175, Tags: []Tag{
			{Name: "Из муки", Emoji: "🌾"}, {Name: "Жареное", Emoji: "🔥"}},
			ImgSrc: "path",
		},
		{Id: 2, Name: "Яичница", Time: 15, Budget: 50, Tags: []Tag{
			{Name: "Яйца", Emoji: "🥚"}, {Name: "Жареное", Emoji: "🔥"}},
			ImgSrc: "path",
		},
	}
	return &SmallStore{list: list}
}

func (s SmallStore) Get() ([]RecipePreview, error) {
	return s.list, nil
}
