package data

type Tag struct {
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Step struct {
	Step int    `json:"step"`
	Text string `json:"text"`
}

type RecipePreview struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Time   int    `json:"time"`
	Budget int    `json:"budget"`
	Tags   []Tag  `json:"tags"`
	ImgSrc string `json:"imgsrc"`
}

type RecipeDetails struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Calories    int          `json:"calories"`
	Time        int          `json:"time"`
	Budget      int          `json:"budget"`
	Tags        []Tag        `json:"tags"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []Step       `json:"steps"`
	ImgSrc      string       `json:"imgsrc"`
}
