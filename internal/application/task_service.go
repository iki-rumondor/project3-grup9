package application

import (
	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"github.com/iki-rumondor/project3-grup9/internal/repository"
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
		return nil, err
	}

	return result, nil
}

func (s *TaskService) GetTasks() (*[]domain.Task, error) {
	tasks, err := s.Repo.FindTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) UpdateTask(task *domain.Task) (*domain.Task, error) {

	task, err := s.Repo.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) UpdateTaskStatus(task *domain.Task) (*domain.Task, error) {

	task, err := s.Repo.UpdateTaskStatus(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) UpdateTaskCategory(task *domain.Task) (*domain.Task, error) {

	task, err := s.Repo.UpdateTaskCategory(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}


func (s *TaskService) DeleteTask(task *domain.Task) error {

	if err := s.Repo.DeleteTask(task); err != nil {
		return err
	}

	return nil
}
