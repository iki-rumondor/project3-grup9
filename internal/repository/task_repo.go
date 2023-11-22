package repository

import "github.com/iki-rumondor/project3-grup9/internal/domain"

type TaskRepository interface {
	CreateTask(*domain.Task) (*domain.Task, error)
	FindTasks() (*[]domain.Task, error)
	UpdateTask(*domain.Task) (*domain.Task, error)
	DeleteTask(*domain.Task) error
	UpdateTaskStatus(*domain.Task) (*domain.Task, error)
	UpdateTaskCategory(*domain.Task) (*domain.Task, error)
}
