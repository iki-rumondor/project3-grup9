package response

import "time"

type Category struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Type       string `json:"type"`
	Created_At time.Time
	Updated_At time.Time
	Tasks      Task
}

type CreateCategory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Type       string `json:"type"`
	Created_At time.Time
}

type UpdateCategory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Type       string `json:"type"`
	Updated_At time.Time
}
