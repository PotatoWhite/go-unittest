package user

import (
	"context"
	"github.com/google/uuid"
)

func NewQuery(r Repository) Query {
	return &query{
		repo: r,
	}
}

type query struct {
	repo Repository
}

func (q query) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	return q.repo.FindByID(ctx, id)
}
