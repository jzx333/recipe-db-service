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
	Id   int    `json:"id"`
	Step string `json:"step"`
}

type SmallInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	Budget string `json:"budget"`
	Tags   []Tag  `json:"tags"`
	ImgSrc string `json:"imgsrc"`
}

type FullInfo struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Calories    string       `json:"calories"`
	Time        string       `json:"time"`
	Budget      string       `json:"budget"`
	Tags        []Tag        `json:"tags"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []Step       `json:"steps"`
	Imgsrc      string       `json:"imgsrc"`
}
