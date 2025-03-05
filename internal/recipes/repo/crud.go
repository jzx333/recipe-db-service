package repo

import (
	"context"
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
	tagsAllQuery = sq.Select("t.id", "t.name", "t.emoji", "t.created_at").
			From("tags t").
			GroupBy("t.id", "t.name", "t.emoji", "t.created_at").
			PlaceholderFormat(sq.Dollar)

	tagCreateQuery = sq.Insert("tags").Columns("name", "emoji").
			PlaceholderFormat(sq.Dollar)

	recipesAllQuery = sq.Select("r.id", "r.name", "r.calories", "r.time", "r.budget",
		"json_agg(json_build_object('id', t.id, 'name', t.name, 'emoji', t.emoji)) as tags",
		"r.ingredients", "r.steps", "r.imgsrc", "r.created_at").
		From("recipes r").
		LeftJoin("recipe_tags rt on r.id = rt.recipe_id").
		LeftJoin("tags t on rt.tag_id = t.id").
		GroupBy("r.id", "r.name", "r.calories", "r.time", "r.budget", "r.ingredients",
			"r.steps", "r.imgsrc", "r.created_at").
		PlaceholderFormat(sq.Dollar)

	recipeCreateQuery = sq.Insert("recipes").Columns("name", "calories",
		"time", "budget", "ingredients", "steps", "imgsrc").
		PlaceholderFormat(sq.Dollar)

	recipeRelQuery = sq.Insert("recipe_tags").
			Columns("recipe_id", "tag_id").PlaceholderFormat(sq.Dollar)
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

	builder := recipeRelQuery
	for _, tagId := range recipe.Tags {
		builder = builder.Values(recipeId, tagId)
	}

	q, args, err = builder.ToSql()
	if err != nil {
		return 0, err
	}

	if _, err := r.db.QueryxContext(ctx, q, args...); err != nil {
		return 0, err
	}

	return recipeId, nil
}

func (r *Repo) RecipeById(ctx context.Context, id int) (*dto.Recipe, error) {
	q, args := recipesAllQuery.Where(sq.Eq{"r.id": id}).Limit(1).MustSql()

	var recipe dto.Recipe

	if err := r.db.GetContext(ctx, &recipe, q, args...); err != nil {
		return nil, ErrNotExist
	}

	return &recipe, nil

}

func (r *Repo) RecipesSearch(ctx context.Context, filter *dto.RecipeFilter) ([]dto.Recipe, error) {
	var conditions []sq.Sqlizer
	recipesSearchQuery := recipesAllQuery

	if filter.Name != "" {
		conditions = append(conditions, sq.Eq{"r.name": filter.Name})
	}

	if filter.Budget != 0 {
		conditions = append(conditions, sq.Eq{"r.budget": filter.Budget})
	}

	if len(filter.Tags) > 0 {
		conditions = append(conditions, sq.Eq{"rt.tag_id": filter.Tags})
		recipesSearchQuery = recipesSearchQuery.
			Where(sq.And(conditions)).
			Having("count(distinct rt.tag_id) = ?", len(filter.Tags))
	} else {
		recipesSearchQuery = recipesSearchQuery.
			Where(sq.And(conditions))
	}

	q, args, err := recipesSearchQuery.ToSql()
	if err != nil {
		return nil, err
	}

	var recipes []dto.Recipe
	if err := r.db.SelectContext(ctx, &recipes, q, args...); err != nil {
		return nil, ErrNotExist
	}

	return recipes, nil
}
