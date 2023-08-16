package user

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

func NewRepository(db *gorm.DB) Repository {

	// Migrate the schema
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Printf("failed to migrate schema: %v", err)
	}

	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Save(ctx context.Context, user User) error {
	user.CreateAt = time.Now()

	tx := r.db.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r repository) FindByID(ctx context.Context, id uuid.UUID) (User, error) {
	var user User
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}
