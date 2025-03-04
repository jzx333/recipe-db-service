package server

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"

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

	router.Get("/previews", func(w http.ResponseWriter, router *http.Request) {

		name := router.URL.Query().Get("name")
		if name != "" {
			recipe, err := r.RecipeByName(ctx, name)
			if err != nil {
				w.WriteHeader(422)
				w.Write([]byte(err.Error()))
			}

			response := ResponsePreview{
				Id:     recipe.Id,
				Name:   recipe.Name,
				Time:   recipe.Time,
				Budget: recipe.Budget,
				Tags:   recipe.Tags,
				ImgSrc: recipe.ImgSrc,
			}

			data, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(422)
				w.Write([]byte(err.Error()))
			}
			w.Write(data)

		}

		tags := router.URL.Query()["tag"]
		if len(tags) > 0 {
			var tagIds []int

			for i := 0; i < len(tags); i++ {
				tagId, err := strconv.Atoi(tags[i])
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(err.Error()))
				}
				tagIds = append(tagIds, tagId)
			}

			recipes, err := r.RecipesByTags(ctx, tagIds...)
			if err != nil {
				w.WriteHeader(404)
				w.Write([]byte(err.Error()))
			}

			var response []ResponsePreview
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
		}

		budget := router.URL.Query().Get("budget")
		if budget != "" {
			budgetInt, err := strconv.Atoi(budget)

			recipes, err := r.RecipesByBudget(ctx, budgetInt)
			if err != nil {
				w.WriteHeader(404)
				w.Write([]byte(err.Error()))
			}

			var response []ResponsePreview
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
		}

		if name == "" && len(tags) == 0 && budget == "" {
			recipes, err := r.RecipesAll(ctx)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
			}

			var response []ResponsePreview
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
		}

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

	http.ListenAndServe(":8080", router)
}
