package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"primaryKey"`
	Email string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func  NewUser (email  string, password string) (User, error) {
	generatedPasswordHash, err := bcrypt.GenerateFromPassword([]byte (password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	return User {
		Email: email,
		PasswordHash: string(generatedPasswordHash),
		CreatedAt: time.Now(),
	}, nil
}