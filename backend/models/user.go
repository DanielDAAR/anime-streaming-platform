package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRole defines the possible user roles
type UserRole string

const (
	RoleGuest UserRole = "guest"
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

// User represents a platform user
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username     string             `bson:"username" json:"username" validate:"required,min=3,max=30,alphanum"`
	Email        string             `bson:"email" json:"email" validate:"required,email,max=100"`
	PasswordHash string             `bson:"passwordHash" json:"-" validate:"required,min=60"` // bcrypt hash
	Role         UserRole           `bson:"role" json:"role" validate:"required,oneof=guest user admin"`
	Avatar       string             `bson:"avatar" json:"avatar"`
	IsActive     bool               `bson:"isActive" json:"isActive"`
	LastLogin    time.Time          `bson:"lastLogin" json:"lastLogin"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// UserResponse is the sanitized version for public API
type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      UserRole  `json:"role"`
	Avatar    string    `json:"avatar"`
	IsActive  bool      `json:"isActive"`
	LastLogin time.Time `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID.Hex(),
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		Avatar:    u.Avatar,
		IsActive:  u.IsActive,
		LastLogin: u.LastLogin,
		CreatedAt: u.CreatedAt,
	}
}

// PublicUser is a minimal user representation for comments
type PublicUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}
