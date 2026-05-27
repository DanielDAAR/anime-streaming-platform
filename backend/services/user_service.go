package services

import (
	"context"
	"math"

	"anime-streaming-platform/models"
	"anime-streaming-platform/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserService handles user business logic
type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		repo: repositories.NewUserRepository(),
	}
}

// GetAllUsers retrieves paginated users
func (s *UserService) GetAllUsers(ctx context.Context, page, limit int) ([]models.UserResponse, int, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	users, total, err := s.repo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return responses, int(total), totalPages, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	resp := user.ToResponse()
	return &resp, nil
}

// UpdateUserRole updates a user's role
func (s *UserService) UpdateUserRole(ctx context.Context, id string, role models.UserRole) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.repo.UpdateRole(ctx, objID, role)
}

// ToggleUserActive toggles user active status
func (s *UserService) ToggleUserActive(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.repo.ToggleActive(ctx, objID)
}

// GetStats returns user statistics
func (s *UserService) GetStats(ctx context.Context) (map[string]interface{}, error) {
	adminCount, err := s.repo.CountByRole(ctx, models.RoleAdmin)
	if err != nil {
		return nil, err
	}

	userCount, err := s.repo.CountByRole(ctx, models.RoleUser)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"totalAdmins": adminCount,
		"totalUsers":  userCount,
		"total":       adminCount + userCount,
	}, nil
}
