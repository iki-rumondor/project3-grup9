package response

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Full_name  string
	Email      string
	password   string
	Role       string
	Created_At time.Time
	Updated_At time.Time
}

type Users struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Full_Name string `json:"full_name"`
}
