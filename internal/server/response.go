package server

import "webServer/internal/recipes/dto"

type ResponseTag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type ResponsePreview struct {
	Id     int              `json:"id"`
	Name   string           `json:"name"`
	Time   int              `json:"time"`
	Budget int              `json:"budget"`
	Tags   dto.SelectedTags `json:"tags"`
	ImgSrc string           `json:"imgsrc"`
}

type ResponseDetail struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	Calories    int              `json:"calories"`
	Time        int              `json:"time"`
	Budget      int              `json:"budget"`
	Tags        dto.SelectedTags `json:"tags"`
	Ingredients dto.Ingredients  `json:"ingredients"`
	Steps       dto.Steps        `json:"steps"`
	ImgSrc      string           `json:"imgsrc"`
}
