package user

import (
	"context"
	"github.com/google/uuid"
)

type Query interface {
	GetUserById(ctx context.Context, id uuid.UUID) (User, error)
}

type Command interface {
	AddUser(ctx context.Context, user User) error
}

type Repository interface {
	Save(ctx context.Context, user User) error
	FindByID(ctx context.Context, id uuid.UUID) (User, error)
}
