package domain

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not_null;varchar(120)"`
	Description string `gorm:"not_null;varchar(120)"`
	Status      bool   `gorm:"not_null;varchar(120)"`
	User_Id     uint
	Category_Id uint
	Created_At  time.Time
	Updated_At  time.Time
	Users       User `gorm:"foreignKey:User_Id"`
	// Categorys   Category
}
