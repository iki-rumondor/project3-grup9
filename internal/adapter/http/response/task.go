package response

import "time"

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTask struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
	User_Id     uint   `json:"user_id"`
	Category_Id uint   `json:"category_id"`
	Created_At  time.Time
}

type UpdateTask struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	User_Id     uint   `json:"user_id"`
	Category_Id uint   `json:"category_id"`
	Updated_At  time.Time
}

type UpdateStatusTask struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	User_Id     uint   `json:"user_id"`
	Category_Id uint   `json:"category_id"`
	Updated_At  time.Time
}

type UpdateCategoryTask struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	User_Id     uint   `json:"user_id"`
	Category_Id uint   `json:"category_id"`
	Updated_At  time.Time
}

type Tasks struct {
	Tasks []*Task
}
