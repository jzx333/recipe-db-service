package dto

import (
	"encoding/json"
	"time"
)

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Step struct {
	Step int    `json:"step"`
	Text string `json:"text"`
}

type SelectedTag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type Ingredients []Ingredient

func (ing *Ingredients) Scan(src interface{}) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, ing)
}

type Steps []Step

func (steps *Steps) Scan(src interface{}) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, steps)
}

type SelectedTags []SelectedTag

func (t *SelectedTags) Scan(src interface{}) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, t)
}

type Recipe struct {
	Id          int          `db:"id"`
	Name        string       `db:"name"`
	Calories    int          `db:"calories"`
	Time        int          `db:"time"`
	Budget      int          `db:"budget"`
	Tags        SelectedTags `json:"tags"`
	Ingredients Ingredients  `db:"ingredients"`
	Steps       Steps        `db:"steps"`
	ImgSrc      string       `db:"imgsrc"`
	CreatedAt   time.Time    `db:"created_at"`
}

type NewRecipe struct {
	Name        string       `db:"name"`
	Tags        []int        `json:"tags"`
	Calories    int          `db:"calories"`
	Time        int          `db:"time"`
	Budget      int          `db:"budget"`
	Ingredients []Ingredient `db:"ingredients"`
	Steps       []Step       `db:"steps"`
	ImgSrc      string       `db:"imgsrc"`
}

type RecipeFilter struct {
	Name   string `db:"name"`
	Tags   []int
	Budget int `db:"budget"`
}
