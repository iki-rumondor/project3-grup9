package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type CategoryRepository interface {
	CreateCategory(*domain.Category)
}
