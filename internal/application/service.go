package application

import "github.com/iki-rumondor/init-golang-service/internal/repository"

type Service struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) *Service{
	return &Service{
		Repo: repo,
	}
}