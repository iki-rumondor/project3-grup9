package domain

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Full_name  string
	Email      string
	password   string
	Role       string
	Created_At time.Time
	Updated_At time.Time

	Tasks []Task
}
