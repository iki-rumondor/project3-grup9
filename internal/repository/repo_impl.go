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

func (r *TaskRepoImplementation) FindTasks(UserID uint) (*[]domain.Task, error) {
	var task []domain.Task
	if err := r.db.Preload("UserProfile").Find(&task, "user_id = ?", UserID).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepoImplementation) UpdateTask(task *domain.Task) (*domain.Task, error) {
	var Tasks domain.Task
	if err := r.db.Model(&domain.Task{}).Where("id = ?", task.ID).Updates(&task).First(&Tasks).Error; err != nil {
		return nil, err
	}

	return &Tasks, nil
}

func (r *TaskRepoImplementation) UpdateStatusTask(task *domain.Task) (*domain.Task, error) {
	var Tasks domain.Task
	if err := r.db.Model(&domain.Task{}).Where("id = ?", task.ID).Updates(&task).First(&Tasks).Error; err != nil {
		return nil, err
	}

	return &Tasks, nil
}

func (r *TaskRepoImplementation) UpdateCategoryTask(task *domain.Task) (*domain.Task, error) {
	var Tasks domain.Task
	if err := r.db.Model(&domain.Task{}).Where("id = ?", task.ID).Updates(&task).First(&Tasks).Error; err != nil {
		return nil, err
	}

	return &Tasks, nil
}

func (r *TaskRepoImplementation) DeleteTask(task *domain.Task) error {
	if err := r.db.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepoImplementation) FindById(id uint) (*domain.Task, error) {
	var task domain.Task
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepoImplementation) FindUser(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *TaskRepoImplementation) FindTask(TaskID uint) (*domain.Task, error) {
	var task domain.Task
	if err := r.db.First(&task, "id = ? ", TaskID).Error; err != nil {
		return nil, err
	}

	return &task, nil
}
