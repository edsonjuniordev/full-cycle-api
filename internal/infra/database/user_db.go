package database

import (
	"github.com/edsonjuniordev/full-cycle-api/internal/entities"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
