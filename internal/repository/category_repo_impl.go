package repository

import "gorm.io/gorm"

type CategoryRepoImplementation struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepoImplementation{
		db: db,
	}
}
