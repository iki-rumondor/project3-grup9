package domain

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not_null;varchar(120)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}
