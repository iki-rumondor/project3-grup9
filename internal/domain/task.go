package domain

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not_null;varchar(120)"`
	Description string `gorm:"not_null;varchar(120)"`
	Status      bool
	UserID      uint
	CategoryID  uint
	User        User
	Category    Category

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {

	if err := tx.First(&User{}, "id = ?", t.UserID).Error; err != nil {
		return err
	}

	if err := tx.First(&Category{}, "id = ?", t.CategoryID).Error; err != nil {
		return err
	}

	return nil
}

func (t *Task) BeforeUpdate(tx *gorm.DB) error {

	if err := tx.First(&Task{}, "id = ? AND user_id = ?", t.ID, t.UserID).Error; err != nil {
		return err
	}

	if err := tx.First(&Category{}, "id = ?", t.CategoryID).Error; err != nil {
		return err
	}

	return nil
}

func (t *Task) BeforeDelete(tx *gorm.DB) error {

	if err := tx.First(&Task{}, "id = ? AND user_id = ?", t.ID, t.UserID).Error; err != nil {
		return err
	}

	return nil
}
