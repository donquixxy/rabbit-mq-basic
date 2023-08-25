package repository

import (
	"micro-company/internal/domain"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(db *gorm.DB, v *domain.Company) (*domain.Company, error)
}

type companyRepository struct {
}

func NewCommpanyRepository() CompanyRepository {
	return &companyRepository{}
}

func (*companyRepository) Create(db *gorm.DB, v *domain.Company) (*domain.Company, error) {
	if err := db.Create(&v).Error; err != nil {
		return nil, err
	}

	return v, nil
}
