package repository

import (
	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"gorm.io/gorm"
)

type AuthRepoImplementation struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepoImplementation{
		db: db,
	}
}

func (r *AuthRepoImplementation) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepoImplementation) UpdateUser(user *domain.User) (*domain.User, error) {
	var result domain.User
	if err := r.db.Model(&user).Updates(&user).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AuthRepoImplementation) DeleteUser(user *domain.User) error {
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepoImplementation) SaveUser(user *domain.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}
