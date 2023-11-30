package users

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	Occupation string `json:"occupation" binding:"required,min=3,max=100"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6,max=100"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email,min=3,max=100"`
}
