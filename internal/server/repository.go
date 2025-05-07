package server

import (
	"webServer/internal/recipes/dto"
)

//func RecipeCreateConverter(query *RequestRecipe) *dto.NewRecipe {
//	newRecipe := &dto.NewRecipe{
//		Name:        query.Name,
//		Tags:        query.Tags,
//		Calories:    query.Calories,
//		Time:        query.Time,
//		Budget:      query.Budget,
//		Ingredients: query.Ingredients,
//		Steps:       query.Steps,
//		ImgSrc:      query.ImgSrc,
//	}
//
//	return newRecipe
//}

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
