package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (s *TaskService) CreateTask(task *domain.Task) (*domain.Task, error) {

	result, err := s.Repo.CreateTask(task)
	if err != nil {
		return nil, errors.New("failed to save task into database")
	}

	return result, nil
}

func (s *TaskService) GetTask(UserID uint) (*[]domain.Task, error) {
	Taks, err := s.Repo.FindTasks(UserID)
	if err != nil {
		return nil, errors.New("failed to get user task from database")
	}

	return Taks, nil
}

func (s *TaskService) UpdateTask(task *domain.Task) (*domain.Task, error) {
	_, err := s.Repo.FindTask(task.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("photo with id %d id not found", task.ID)
	}
	if err != nil {
		return nil, errors.New("failed to get tasks from database")
	}

	task, err = s.Repo.UpdateTask(task)
	if err != nil {
		return nil, errors.New("failed to update tasks to database")
	}

	return task, nil
}

func (s *TaskService) UpdateStatusTask(task *domain.Task) (*domain.Task, error) {
	_, err := s.Repo.FindTask(task.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("photo with id %d id not found", task.ID)
	}
	if err != nil {
		return nil, errors.New("failed to get tasks from database")
	}

	task, err = s.Repo.UpdateStatusTask(task)
	if err != nil {
		return nil, errors.New("failed to update status tasks to database")
	}

	return task, nil
}

func (s *TaskService) UpdateCategoryTask(task *domain.Task) (*domain.Task, error) {
	_, err := s.Repo.FindTask(task.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("photo with id %d id not found", task.ID)
	}
	if err != nil {
		return nil, errors.New("failed to get tasks from database")
	}

	task, err = s.Repo.UpdateCategoryTask(task)
	if err != nil {
		return nil, errors.New("failed to update category tasks to database")
	}

	return task, nil
}

func (s *TaskService) DeleteTask(task *domain.Task) error {
	_, err := s.Repo.FindTask(task.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("tasks with id %d id not found", task.ID)
	}
	if err != nil {
		return errors.New("failed to get tasks from database")
	}

	if err := s.Repo.DeleteTask(task); err != nil {
		return errors.New("we encountered an issue while trying to delete the tasks")
	}

	return nil
}
