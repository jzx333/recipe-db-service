package server

type RequestTag struct {
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type RequestRecipe struct {
	Name        string
	Tags        []int
	Calories    int
	Time        int
	Budget      int
	Ingredients []byte
	Steps       []byte
	ImgSrc      string
}

type RequestSearchQuery struct {
	Name   string
	Tags   []int
	Budget int
}
