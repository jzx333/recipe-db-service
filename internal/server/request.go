package server

import "webServer/internal/recipes/dto"

type RequestTag struct {
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type RequestTagsRecipe struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type RequestRecipe struct {
	Name        string              `json:"name"`
	Tags        []RequestTagsRecipe `json:"tags"`
	Calories    int                 `json:"calories"`
	Time        int                 `json:"time"`
	Budget      int                 `json:"budget"`
	Ingredients []dto.Ingredient    `json:"ingredients"`
	Steps       []dto.Step          `json:"steps"`
	ImgSrc      string              `json:"imgSrc"`
}

type RequestSearchQuery struct {
	Name   string
	Tags   []int
	Budget int
}
