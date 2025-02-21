package data

type TagStore struct {
	list []Tag
}

func NewTagStore() *TagStore {
	var list = []Tag{
		{Name: "Жареное", Emoji: "🔥"},
		{Name: "Яйца", Emoji: "🥚"},
		{Name: "Бекон", Emoji: "🥓"},
	}
	return &TagStore{
		list,
	}
}

func (t TagStore) Get() ([]Tag, error) {
	return t.list, nil
}
