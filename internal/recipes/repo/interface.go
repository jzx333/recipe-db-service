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
	RecipeCreate(ctx context.Context, recipe *dto.NewRecipe) (int, error)
	RecipeById(ctx context.Context, id int) (*dto.Recipe, error)
	RecipesSearch(ctx context.Context, filter *dto.RecipeFilter) ([]dto.Recipe, error)
}

type TagsRepo interface {
	TagsAll(ctx context.Context) ([]dto.Tag, error)
	TagCreate(ctx context.Context, tag *dto.NewTag) (int, error)
}
