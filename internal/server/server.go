package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"webServer/internal/recipes/dto"

	"webServer/internal/recipes/repo"

	"github.com/go-chi/chi/v5"
)

func Server(ctx context.Context, r *repo.Repo) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/tags", func(w http.ResponseWriter, router *http.Request) {
		tags, err := r.TagsAll(ctx)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		var response []ResponseTag
		for _, tag := range tags {
			response = append(response, ResponseTag{
				Id:    tag.Id,
				Name:  tag.Name,
				Emoji: tag.Emoji,
			})
		}

		data, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(err.Error()))
		}

		w.Write(data)
	})

	router.Post("/tags", func(w http.ResponseWriter, router *http.Request) {

		newTag := &RequestTag{}

		json.NewDecoder(router.Body).Decode(newTag)

		response, err := r.TagCreate(ctx, TagCreateConverter(newTag))
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		} else {
			data, _ := json.Marshal(response)
			w.Write(data)
		}
	})

	router.Get("/previews", func(w http.ResponseWriter, router *http.Request) {

		name := router.URL.Query().Get("name")

		tags := router.URL.Query()["tag"]

		budget := router.URL.Query().Get("budget")

		tagIds := make([]int, len(tags))
		if len(tagIds) != 0 {
			for i, tag := range tags {
				tagIds[i], _ = strconv.Atoi(tag)
			}
		}
		budgetInt, _ := strconv.Atoi(budget)

		recipeQuery := RequestSearchQuery{
			Name:   name,
			Tags:   tagIds,
			Budget: budgetInt,
		}

		// log
		fmt.Println(recipeQuery)

		recipes, err := r.RecipesSearch(ctx, RecipeSearchConverter(&recipeQuery))
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(err.Error()))
		}

		// log
		fmt.Println(len(recipes))

		response := make([]ResponsePreview, 0, len(recipes))
		for _, recipe := range recipes {
			response = append(response, ResponsePreview{
				Id:     recipe.Id,
				Name:   recipe.Name,
				Time:   recipe.Time,
				Budget: recipe.Budget,
				Tags:   recipe.Tags,
				ImgSrc: recipe.ImgSrc,
			})
		}

		data, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(err.Error()))
		}

		w.Write(data)

	})

	router.Get("/previews/{id}", func(w http.ResponseWriter, router *http.Request) {
		idParam, err := strconv.Atoi(chi.URLParam(router, "id"))

		recipe, err := r.RecipeById(ctx, idParam)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(err.Error()))
		}

		response := ResponseDetail{
			Id:          recipe.Id,
			Name:        recipe.Name,
			Calories:    recipe.Calories,
			Time:        recipe.Time,
			Budget:      recipe.Budget,
			Tags:        recipe.Tags,
			Ingredients: recipe.Ingredients,
			Steps:       recipe.Steps,
			ImgSrc:      recipe.ImgSrc,
		}

		data, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(err.Error()))
		}

		w.Write(data)
	})

	router.Post("/previews", func(w http.ResponseWriter, router *http.Request) {
		recipeInput := RequestRecipe{}

		json.NewDecoder(router.Body).Decode(&recipeInput)

		newRecipe := dto.NewRecipe{
			Name:        recipeInput.Name,
			Tags:        make([]int, 0, len(recipeInput.Tags)),
			Calories:    recipeInput.Calories,
			Time:        recipeInput.Time,
			Budget:      recipeInput.Budget,
			Ingredients: recipeInput.Ingredients,
			Steps:       recipeInput.Steps,
			ImgSrc:      recipeInput.ImgSrc,
		}

		for _, tag := range recipeInput.Tags {
			newRecipe.Tags = append(newRecipe.Tags, tag.Id)
		}

		recipeId, err := r.RecipeCreate(ctx, &newRecipe)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		} else {
			data, _ := json.Marshal(recipeId)
			w.Write(data)
		}

	})

	http.ListenAndServe(":8080", router)
}
