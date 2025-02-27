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
	tagQuery = sq.Select("t.name", "t.emoji").From("tags t").
		GroupBy("t.name", "t.emoji").PlaceholderFormat(sq.Dollar)
)

func (r *Repo) TagsAll(ctx context.Context) ([]*dto.Tag, error) {
	q, args := tagQuery.MustSql()

	var tags []dto.Tag

	if err := r.db.SelectContext(ctx, &tags, q, args...); err != nil {
		return nil, err
	}

	tagsPtr := make([]*dto.Tag, len(tags))
	for i := range tags {
		tagsPtr[i] = &tags[i]
	}

	return tagsPtr, nil
}
