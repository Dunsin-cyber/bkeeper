package requests

type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"min=2,max=100"`
	LastName  string `json:"last_name" validate:"min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
	// Gender   string  

}