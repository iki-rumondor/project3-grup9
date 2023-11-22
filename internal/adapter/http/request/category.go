package request

type Category struct{
	Type string `json:"type" valid:"required~message is required"`
}