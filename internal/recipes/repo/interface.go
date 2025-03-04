package repo

import (
	"context"
	"fmt"
	"webServer/internal/recipes/dto"
)

var (
	ErrDuplicate = fmt.Errorf("duplicate")
	ErrNotExist  = fmt.Errorf("not exist")
)

type RecipeRepo interface {
	RecipesAll(ctx context.Context) ([]dto.Recipe, error)
	RecipeCreate(ctx context.Context, recipe *dto.Recipe) (int, error)
	RecipeById(ctx context.Context, id int) (*dto.Recipe, error)
	RecipeByName(ctx context.Context, name string) (*dto.Recipe, error)
	RecipesByTags(ctx context.Context, tags ...int) ([]dto.Recipe, error)
	RecipesByBudget(ctx context.Context, budget int) ([]dto.Recipe, error)
	// TODO: implement this
	RecipesByTagsAndBudget(ctx context.Context, budget int, tags ...int) ([]dto.Recipe, error)
}

type TagsRepo interface {
	TagsAll(ctx context.Context) ([]dto.Tag, error)
	TagCreate(ctx context.Context, tag *dto.Tag) (int, error)
}
