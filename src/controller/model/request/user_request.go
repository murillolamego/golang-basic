package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=2,max=60"`
	Age      int8   `json:"age" binding:"required,min=18,max=140"`
}
