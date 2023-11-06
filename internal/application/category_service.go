package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type CategoryService struct {
	Repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		Repo: repo,
	}
}

func (s *CategoryService) CreateCategory(category *domain.Category) (*domain.Category, error) {
	result, err := s.Repo.CreateCategory(category)
	if err != nil {
		return nil, errors.New("failed to save category into database")
	}

	return result, nil
}

func (s *CategoryService) GetCategorys(TaskID uint) (*[]domain.Category, error) {
	categorys, err := s.Repo.FindCategorys(TaskID)
	if err != nil {
		return nil, errors.New("failed to get task category from database")
	}

	return categorys, nil
}

func (s *CategoryService) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	_, err := s.Repo.FindCategory(category.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("category with id %d id not found", category.ID)
	}
	if err != nil {
		return nil, errors.New("failed to get category from database")
	}

	category, err = s.Repo.UpdateCategory(category)
	if err != nil {
		return nil, errors.New("failed to update category to database")
	}

	return category, nil
}

func (s *CategoryService) DeleteCategory(category *domain.Category) error {
	_, err := s.Repo.FindCategory(category.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("category with id %d id not found", category.ID)
	}
	if err != nil {
		return errors.New("failed to get category from database")
	}

	if err := s.Repo.DeleteCategory(category); err != nil {
		return errors.New("we encountered an issue while trying to delete the category")
	}

	return nil
}
