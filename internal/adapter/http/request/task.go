package request

type CreateTask struct {
	Title       string `json:"title" valid:"required~title is required"`
	Description string `json:"description" valid:"required~description is required"`
	CategoryID  uint   `json:"category_id" valid:"required~category_id is required"`
}

type UpdateTask struct {
	Title       string `json:"title" valid:"required~title is required"`
	Description string `json:"description" valid:"required~description is required"`
}

type UpdateTaskStatus struct {
	Status bool `json:"status"`
}

type UpdateTaskCategory struct {
	CategoryID uint `json:"category_id"`
}
