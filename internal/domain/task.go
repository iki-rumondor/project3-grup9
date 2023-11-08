package domain

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not_null;varchar(120)"`
	Description string `gorm:"not_null;varchar(120)"`
	Status      bool
	UserID      uint
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
	Category    Category
}
