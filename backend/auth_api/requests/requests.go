package requests

type RegisterRequest struct{
	Email string `json:"email" binding:"required,isEmail,emailUnique"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoginRequest struct{
	Email string `json:"email" binding:"required,isEmail"`
	Password string `json:"password" binding:"required"`
}