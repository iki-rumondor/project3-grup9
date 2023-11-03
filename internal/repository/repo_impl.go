package repository

import "gorm.io/gorm"

type RepoImplementation struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository{
	return RepoImplementation{
		db: db,
	}
}