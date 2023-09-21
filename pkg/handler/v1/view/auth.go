package view

// LoginRequest represent the login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required" format:"email" binding:"required"`
	Password string `json:"password" validate:"required" binding:"required"`
} // @name LoginRequest

// LoginResponse represent the login response
type LoginResponse = Response[Auth] // @name LoginResponse

// Auth represent the auth response
type Auth struct {
	ID          int    `json:"id" validate:"required"`
	Email       string `json:"email" validate:"required" format:"email"`
	AccessToken string `json:"accessToken" validate:"required"`
} // @name Auth

// SignupRequest represent the signup request
type SignupRequest struct {
	Email    string `json:"email" validate:"required" format:"email" binding:"required,email"`
	Password string `json:"password" validate:"required" binding:"required"`
	FullName string `json:"fullName" validate:"required" binding:"required"`
	Avatar   string `json:"avatar"`
} // @name SignupRequest
