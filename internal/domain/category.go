package domain

import "time"

type Category struct {
	ID         uint   `gorm:"primaryKey"`
	Type       string `gorm:"not_null;varchar(120)"`
	Created_At time.Time
	Updated_At time.Time

	Tasks []Task
}
