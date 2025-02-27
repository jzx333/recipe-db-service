package repo

import (
	"context"
	"webServer/internal/recipes/dto"
)

type RecipeRepo interface {
	RecipeByName(ctx context.Context, name string) (*dto.Recipe, error)
}

type TagsRepo interface {
	TagsAll(ctx context.Context) ([]*dto.Tag, error)
}
