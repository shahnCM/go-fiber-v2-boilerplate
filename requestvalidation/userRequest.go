package requestvalidation

type User struct {
	Name     string `json:"name" validate:"required,min=5,max=32"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	IsActive *bool  `json:"is_active" validate:"required"`
}
