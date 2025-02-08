package repositories

import (
	"my-procurement-system/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(userID uint) (*models.User, error)
	GetUsers() ([]models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}
