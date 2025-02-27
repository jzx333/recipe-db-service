package dto

import "time"

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Step struct {
	Step int    `json:"step"`
	Text string `json:"text"`
}

type Recipe struct {
	Id          int          `db:"id"`
	Name        string       `db:"name"`
	Tags        []Tag        `json:"tags"`
	Calories    int          `db:"calories"`
	Time        int          `db:"time"`
	Budget      int          `db:"budget"`
	Ingredients []Ingredient `db:"ingredients"`
	Steps       []Step       `db:"steps"`
	ImgSrc      string       `db:"imgsrc"`
	CreatedAt   time.Time    `db:"created_at"`
}
