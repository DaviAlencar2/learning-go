package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"min=6,max=32"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
}
