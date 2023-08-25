package repository

import (
	"micro-company/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, v *domain.User) (*domain.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (*userRepository) Create(db *gorm.DB, v *domain.User) (*domain.User, error) {

	if er := db.Create(&v).Error; er != nil {
		return nil, er
	}

	return v, nil
}
