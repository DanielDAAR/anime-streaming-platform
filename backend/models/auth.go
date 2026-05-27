package models

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

// RegisterRequest represents registration data
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

// AuthResponse contains tokens and user data
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// JWTClaims represents custom JWT claims
type JWTClaims struct {
	UserID   string   `json:"userId"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Role     UserRole `json:"role"`
}
