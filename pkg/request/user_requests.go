package request

type UserStore struct {
	Name     string `json:"name" validate:"required,min=5,max=32" error:"name is required, minimum length is 5, maximum length is 32"`
	Email    string `json:"email" validate:"required,email,min=6,max=32" error:"email is required & must be valid"`
	IsActive *bool  `json:"is_active" validate:"required" error:"is_active is required & must be true or false"`
}
