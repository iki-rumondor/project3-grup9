package repository

import "github.com/iki-rumondor/project3-grup9/internal/domain"

type AuthRepository interface {
	FindByEmail(string) (*domain.User, error)
	SaveUser(*domain.User) error
	UpdateUser(*domain.User) (*domain.User, error)
	DeleteUser(*domain.User) error
}
