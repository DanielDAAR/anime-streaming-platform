package services

import (
	"context"
	"errors"
	"math"
	"strings"

	"anime-streaming-platform/models"
	"anime-streaming-platform/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentService handles comment business logic
type CommentService struct {
	repo     *repositories.CommentRepository
	userRepo *repositories.UserRepository
}

// NewCommentService creates a new comment service
func NewCommentService() *CommentService {
	return &CommentService{
		repo:     repositories.NewCommentRepository(),
		userRepo: repositories.NewUserRepository(),
	}
}

// CreateComment creates a new comment
func (s *CommentService) CreateComment(ctx context.Context, animeID, userID primitive.ObjectID, content string) (*models.CommentResponse, error) {
	if content = strings.TrimSpace(content); content == "" {
		return nil, errors.New("content is required")
	}
	if len(content) > 1000 {
		return nil, errors.New("content must be less than 1000 characters")
	}

	comment := &models.Comment{
		AnimeID: animeID,
		UserID:  userID,
		Content: content,
	}

	if err := s.repo.Create(ctx, comment); err != nil {
		return nil, err
	}

	// Populate user data
	user, err := s.userRepo.GetByID(ctx, userID)
	if err == nil {
		comment.User = &models.PublicUser{
			ID:       user.ID.Hex(),
			Username: user.Username,
			Avatar:   user.Avatar,
			Role:     string(user.Role),
		}
	}

	resp := comment.ToResponse()
	return &resp, nil
}

// CreateReply creates a reply to a comment
func (s *CommentService) CreateReply(ctx context.Context, parentID, userID primitive.ObjectID, content string) (*models.CommentResponse, error) {
	if content = strings.TrimSpace(content); content == "" {
		return nil, errors.New("content is required")
	}
	if len(content) > 1000 {
		return nil, errors.New("content must be less than 1000 characters")
	}

	// Verify parent comment exists
	parent, err := s.repo.GetByID(ctx, parentID)
	if err != nil {
		return nil, err
	}

	// Check if parent is already a reply (only 1 level allowed)
	if parent.ParentID != nil {
		return nil, errors.New("cannot reply to a reply")
	}

	reply := &models.Comment{
		AnimeID:  parent.AnimeID,
		UserID:   userID,
		ParentID: &parentID,
		Content:  content,
	}

	if err := s.repo.Create(ctx, reply); err != nil {
		return nil, err
	}

	// Populate user data
	user, err := s.userRepo.GetByID(ctx, userID)
	if err == nil {
		reply.User = &models.PublicUser{
			ID:       user.ID.Hex(),
			Username: user.Username,
			Avatar:   user.Avatar,
			Role:     string(user.Role),
		}
	}

	resp := reply.ToResponse()
	return &resp, nil
}

// GetCommentsByAnimeID retrieves comments for an anime
func (s *CommentService) GetCommentsByAnimeID(ctx context.Context, animeID string, page, limit int) ([]models.CommentResponse, int, int, error) {
	objID, err := primitive.ObjectIDFromHex(animeID)
	if err != nil {
		return nil, 0, 0, err
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	comments, total, err := s.repo.GetByAnimeID(ctx, objID, page, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	// Populate user data and replies for each comment
	responses := make([]models.CommentResponse, len(comments))
	for i, comment := range comments {
		// Populate user
		user, err := s.userRepo.GetByID(ctx, comment.UserID)
		if err == nil {
			comment.User = &models.PublicUser{
				ID:       user.ID.Hex(),
				Username: user.Username,
				Avatar:   user.Avatar,
				Role:     string(user.Role),
			}
		}

		// Populate replies
		replies, err := s.repo.GetReplies(ctx, comment.ID)
		if err == nil && len(replies) > 0 {
			comment.Replies = replies
			// Populate user for replies
			for j := range comment.Replies {
				replyUser, err := s.userRepo.GetByID(ctx, comment.Replies[j].UserID)
				if err == nil {
					comment.Replies[j].User = &models.PublicUser{
						ID:       replyUser.ID.Hex(),
						Username: replyUser.Username,
						Avatar:   replyUser.Avatar,
						Role:     string(replyUser.Role),
					}
				}
			}
		}

		responses[i] = comment.ToResponse()
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return responses, int(total), totalPages, nil
}

// LikeComment adds a like to a comment
func (s *CommentService) LikeComment(ctx context.Context, commentID, userID primitive.ObjectID) error {
	return s.repo.AddLike(ctx, commentID, userID)
}

// DeleteComment soft deletes a comment
func (s *CommentService) DeleteComment(ctx context.Context, commentID primitive.ObjectID) error {
	return s.repo.SoftDelete(ctx, commentID)
}

// GetRecentComments retrieves recent comments for moderation
func (s *CommentService) GetRecentComments(ctx context.Context, page, limit int) ([]models.CommentResponse, int, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	comments, total, err := s.repo.GetRecentComments(ctx, page, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	responses := make([]models.CommentResponse, len(comments))
	for i, comment := range comments {
		user, err := s.userRepo.GetByID(ctx, comment.UserID)
		if err == nil {
			comment.User = &models.PublicUser{
				ID:       user.ID.Hex(),
				Username: user.Username,
				Avatar:   user.Avatar,
				Role:     string(user.Role),
			}
		}
		responses[i] = comment.ToResponse()
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return responses, int(total), totalPages, nil
}
