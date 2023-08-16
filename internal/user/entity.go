package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);primary_key;"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null;"`
	CreateAt time.Time `json:"create_at" gorm:"not null;colum:create_at"`
}
