package user

import (
	"github.com/google/uuid"
)

type UserDto struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CreateAt string    `json:"create_at"`
}
