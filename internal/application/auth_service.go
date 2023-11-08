package application

import (
	"errors"

	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"github.com/iki-rumondor/project3-grup9/internal/repository"
	"github.com/iki-rumondor/project3-grup9/internal/utils"
)

type AuthService struct {
	Repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) CreateUser(user *domain.User) error {

	if err := s.Repo.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) VerifyUser(user *domain.User) (string, error) {

	result, err := s.Repo.FindByEmail(user.Email)
	if err != nil {
		return "", errors.New("sorry, the provided email is not registered in our system")
	}

	if err := utils.ComparePassword(result.Password, user.Password); err != nil {
		return "", errors.New("whoops! password mismatch")
	}

	data := map[string]interface{}{
		"id":   result.ID,
		"role": result.Role,
	}

	jwt, err := utils.GenerateToken(data)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *AuthService) UpdateUser(user *domain.User) (*domain.User, error) {

	user, err := s.Repo.UpdateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) DeleteUser(user *domain.User) error {

	if err := s.Repo.DeleteUser(user); err != nil {
		return err
	}

	return nil
}
