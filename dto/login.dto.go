package dto

// Define form validation when user hit /login API
type LoginCredential struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
