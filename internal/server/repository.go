package server

import (
	"webServer/internal/recipes/dto"
)

func TagCreateConverter(query *RequestTag) *dto.NewTag {
	newTag := &dto.NewTag{
		Name:  query.Name,
		Emoji: query.Emoji,
	}

	return newTag
}

func RecipeSearchConverter(query *RequestSearchQuery) *dto.RecipeFilter {
	filter := &dto.RecipeFilter{
		Name:     query.Name,
		Tags:     query.Tags,
		Budget:   query.Budget,
		Time:     query.Time,
		Calories: query.Calories,
	}

	return filter
}
