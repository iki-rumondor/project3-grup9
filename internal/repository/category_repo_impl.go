package repository

import (
	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepoImplementation struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepoImplementation{
		db: db,
	}
}

func (r *CategoryRepoImplementation) CreateCategory(category *domain.Category) (*domain.Category, error) {
	if err := r.db.Save(category).First(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepoImplementation) FindCategories() (*[]domain.Category, error) {
	var categories []domain.Category
	if err := r.db.Preload("Tasks").Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}
