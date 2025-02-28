package repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"webServer/internal/recipes/dto"
)

import sq "github.com/Masterminds/squirrel"

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db}
}

var (
	recipeRelQuery = sq.Insert("recipe_tags").
		Columns("recipe_id", "tag_id").PlaceholderFormat(sq.Dollar)

	tagsAllQuery = sq.Select("t.id", "t.name", "t.emoji", "t.created_at").
		From("tags t").
		GroupBy("t.id", "t.name", "t.emoji", "t.created_at").
		PlaceholderFormat(sq.Dollar)

	tagCreateQuery = sq.Insert("tags").Columns("name", "emoji").
		PlaceholderFormat(sq.Dollar)

	recipesAllQuery = sq.Select("r.id", "r.name", "r.calories", "r.time", "r.budget",
		"string_agg(distinct t.name, ', ') as tags", "r.ingredients",
		"r.steps", "r.imgsrc", "r.created_at").
		From("recipes r").
		LeftJoin("recipe_tags rt on r.id = rt.recipe_id").
		LeftJoin("tags t on rt.tag_id = t.id").
		GroupBy("r.id", "r.name", "r.calories", "r.time", "r.budget", "r.ingredients",
			"r.steps", "r.imgsrc", "r.created_at").
		PlaceholderFormat(sq.Dollar)

	recipeCreateQuery = sq.Insert("recipes").Columns("name", "calories",
		"time", "budget", "ingredients", "steps", "imgsrc").
		PlaceholderFormat(sq.Dollar)
)

func (r *Repo) TagsAll(ctx context.Context) ([]dto.Tag, error) {
	q, args := tagsAllQuery.MustSql()

	var tags []dto.Tag

	if err := r.db.SelectContext(ctx, &tags, q, args...); err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *Repo) TagCreate(ctx context.Context, tag *dto.NewTag) (int, error) {
	q, args, err := tagCreateQuery.
		Values(tag.Name, tag.Emoji).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	if err := r.db.QueryRowxContext(ctx, q, args...).Scan(&id); err != nil {
		return 0, ErrDuplicate
	}

	return id, nil
}

func (r *Repo) RecipesAll(ctx context.Context) ([]dto.Recipe, error) {
	q, args := recipesAllQuery.MustSql()
	fmt.Println(q, args)

	var recipes []dto.Recipe
	if err := r.db.SelectContext(ctx, &recipes, q, args...); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (r *Repo) RecipeCreate(ctx context.Context, recipe *dto.NewRecipe) (int, error) {

	q, args, err := recipeCreateQuery.
		Values(recipe.Name, recipe.Calories, recipe.Time, recipe.Budget, recipe.Ingredients,
			recipe.Steps, recipe.ImgSrc).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var recipeId int
	if err := r.db.QueryRowxContext(ctx, q, args...).Scan(&recipeId); err != nil {
		return 0, ErrDuplicate
	}
	fmt.Println(recipeId)

	builder := recipeRelQuery
	for _, tagId := range recipe.Tags {
		builder = builder.Values(recipeId, tagId)
	}
	fmt.Println(builder)

	q, args, err = builder.ToSql()
	if err != nil {
		return 0, err
	}

	fmt.Println(q, args)
	if _, err := r.db.QueryxContext(ctx, q, args...); err != nil {
		return 0, err
	}

	return recipeId, nil
}
