package user

import (
	"context"
)

func NewCommand(r Repository) Command {
	return &command{
		repo: r,
	}
}

type command struct {
	repo Repository
}

func (c command) AddUser(ctx context.Context, user User) error {
	return c.repo.Save(ctx, user)
}
