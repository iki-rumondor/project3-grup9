package request

type Register struct {
	FullName string `json:"full_name" valid:"required~field full_name is required"`
	Email    string `json:"email" valid:"required~field email is required, email"`
	Password string `json:"password" valid:"required~field password is required, length(6|99)~password at least 6 character"`
	// Role   string   `json:"role" valid:"required~field role_id is required, in(admin|member)"`
}

type Login struct {
	Email    string `json:"email" valid:"required~field email is required, email"`
	Password string `json:"password" valid:"required~field password is required"`
}

type UpdateUser struct {
	FullName string `json:"full_name" valid:"required~field full_name is required"`
	Email    string `json:"email" valid:"required~field email is required, email"`
}
