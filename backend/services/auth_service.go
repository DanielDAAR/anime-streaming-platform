package services

import (
	"context"
	"errors"
	"strings"

	"anime-streaming-platform/models"
	"anime-streaming-platform/repositories"
	"anime-streaming-platform/utils"
	"anime-streaming-platform/validators"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthService handles authentication business logic
type AuthService struct {
	userRepo *repositories.UserRepository
}

// NewAuthService creates a new auth service
func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repositories.NewUserRepository(),
	}
}

// Register creates a new user account
func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponse, error) {
	if err := validators.ValidateRegister(req); err != nil {
		return nil, err
	}

	// Check if email exists
	if _, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return nil, errors.New("email already registered")
	}

	// Check if username exists
	if _, err := s.userRepo.GetByUsername(ctx, req.Username); err == nil {
		return nil, errors.New("username already taken")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	// Create user
	user := &models.User{
		Username:     strings.TrimSpace(req.Username),
		Email:        strings.ToLower(strings.TrimSpace(req.Email)),
		PasswordHash: hashedPassword,
		Role:         models.RoleUser,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Generate token
	token, err := utils.GenerateToken(*user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.AuthResponse{
		Token: token,
		User:  user.ToResponse(),
	}, nil
}

// Login authenticates a user
func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponse, error) {
	if err := validators.ValidateLogin(req); err != nil {
		return nil, err
	}

	// Find user by email
	user, err := s.userRepo.GetByEmail(ctx, strings.ToLower(strings.TrimSpace(req.Email)))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Verify password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	// Update last login
	if err := s.userRepo.UpdateLastLogin(ctx, user.ID); err != nil {
		// Non-critical error, log but continue
	}

	// Generate token
	token, err := utils.GenerateToken(*user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.AuthResponse{
		Token: token,
		User:  user.ToResponse(),
	}, nil
}

// GetCurrentUser retrieves the current user from token claims
func (s *AuthService) GetCurrentUser(ctx context.Context, userID string) (*models.UserResponse, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.userRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	resp := user.ToResponse()
	return &resp, nil
}
