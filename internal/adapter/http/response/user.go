package response

import "time"

type CreateUser struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUser struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}
