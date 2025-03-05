package server

import (
	"webServer/internal/recipes/dto"
)

func Converter(query *RequestQuery) *dto.RecipeFilter {
	filter := &dto.RecipeFilter{
		Name:   query.Name,
		Tags:   query.Tags,
		Budget: query.Budget,
	}

	return filter
}
