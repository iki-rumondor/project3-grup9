package repository

import "github.com/iki-rumondor/project3-grup9/internal/domain"

type CategoryRepository interface {
	CreateCategory(*domain.Category) (*domain.Category, error)
	FindCategories() (*[]domain.Category, error)
	UpdateCategory(*domain.Category) (*domain.Category, error)
	DeleteCategory(*domain.Category) error
}
