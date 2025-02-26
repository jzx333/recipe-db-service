package repo

import (
	"context"
	"webServer/internal/recipes/dto"
)

type RecipeRepo interface {
	RecipeByName(ctx context.Context, name string) (*dto.Recipe, error)
}
