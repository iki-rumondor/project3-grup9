package repository

import (
	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"gorm.io/gorm"
)

type TaskRepoImplementation struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &TaskRepoImplementation{
		db: db,
	}
}

func (r TaskRepoImplementation) CreateTask(task *domain.Task) (*domain.Task, error) {
	if err := r.db.Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepoImplementation) FindTasks() (*[]domain.Task, error) {
	var task []domain.Task
	if err := r.db.Preload("User").Find(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepoImplementation) UpdateTask(task *domain.Task) (*domain.Task, error) {
	var result domain.Task
	if err := r.db.Model(&task).Updates(&task).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TaskRepoImplementation) UpdateTaskStatus(task *domain.Task) (*domain.Task, error) {
	var result domain.Task
	if err := r.db.Model(&task).Update("status", task.Status).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TaskRepoImplementation) UpdateTaskCategory(task *domain.Task) (*domain.Task, error) {
	var result domain.Task
	if err := r.db.Model(&task).Update("category_id", task.CategoryID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TaskRepoImplementation) DeleteTask(task *domain.Task) error {
	if err := r.db.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

