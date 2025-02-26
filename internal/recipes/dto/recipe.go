package dto

import (
	"gorm.io/gorm"
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

type Recipe struct {
	gorm.Model
	Id          int
	Name        string
	Calories    int
	Time        int
	Budget      int
	Ingredients []Ingredient
	Steps       []Step
	ImgSrc      string
	CreatedAt   time.Time
}
