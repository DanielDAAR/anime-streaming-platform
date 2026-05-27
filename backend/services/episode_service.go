package services

import (
	"context"
	"math"

	"anime-streaming-platform/models"
	"anime-streaming-platform/repositories"
	"anime-streaming-platform/validators"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EpisodeService handles episode business logic
type EpisodeService struct {
	repo      *repositories.EpisodeRepository
	animeRepo *repositories.AnimeRepository
}

// NewEpisodeService creates a new episode service
func NewEpisodeService() *EpisodeService {
	return &EpisodeService{
		repo:      repositories.NewEpisodeRepository(),
		animeRepo: repositories.NewAnimeRepository(),
	}
}

// CreateEpisode creates a new episode
func (s *EpisodeService) CreateEpisode(ctx context.Context, episode *models.Episode) (*models.EpisodeResponse, error) {
	if err := validators.ValidateEpisode(episode); err != nil {
		return nil, err
	}

	// Verify anime exists
	_, err := s.animeRepo.GetByID(ctx, episode.AnimeID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, episode); err != nil {
		return nil, err
	}

	// Update episode count
	count, err := s.repo.CountByAnimeID(ctx, episode.AnimeID)
	if err != nil {
		return nil, err
	}
	if err := s.animeRepo.UpdateEpisodeCount(ctx, episode.AnimeID, int(count)); err != nil {
		return nil, err
	}

	resp := episode.ToResponse()
	return &resp, nil
}

// GetEpisodeByID retrieves an episode by ID
func (s *EpisodeService) GetEpisodeByID(ctx context.Context, id string) (*models.EpisodeResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	episode, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	resp := episode.ToResponse()
	return &resp, nil
}

func (s *EpisodeService) resolveAnimeID(ctx context.Context, animeRef string) (primitive.ObjectID, error) {
	if objID, err := primitive.ObjectIDFromHex(animeRef); err == nil {
		return objID, nil
	}

	anime, err := s.animeRepo.GetBySlug(ctx, animeRef)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return anime.ID, nil
}

// GetEpisodesByAnime retrieves episodes for an anime by ObjectID or slug.
func (s *EpisodeService) GetEpisodesByAnime(ctx context.Context, animeRef string, page, limit int) ([]models.EpisodeResponse, int, int, error) {
	objID, err := s.resolveAnimeID(ctx, animeRef)
	if err != nil {
		return nil, 0, 0, err
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	episodes, total, err := s.repo.GetByAnimeID(ctx, objID, page, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	responses := make([]models.EpisodeResponse, len(episodes))
	for i, ep := range episodes {
		responses[i] = ep.ToResponse()
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return responses, int(total), totalPages, nil
}

// GetEpisodeByAnimeAndNumber retrieves a single episode by anime ObjectID/slug and episode number.
func (s *EpisodeService) GetEpisodeByAnimeAndNumber(ctx context.Context, animeRef string, number int) (*models.EpisodeResponse, error) {
	objID, err := s.resolveAnimeID(ctx, animeRef)
	if err != nil {
		return nil, err
	}

	episode, err := s.repo.GetByAnimeAndNumber(ctx, objID, number)
	if err != nil {
		return nil, err
	}

	resp := episode.ToResponse()
	return &resp, nil
}

// UpdateEpisode updates an existing episode
func (s *EpisodeService) UpdateEpisode(ctx context.Context, id string, episode *models.Episode) (*models.EpisodeResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := validators.ValidateEpisode(episode); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, objID, episode); err != nil {
		return nil, err
	}

	updated, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	resp := updated.ToResponse()
	return &resp, nil
}

// DeleteEpisode removes an episode
func (s *EpisodeService) DeleteEpisode(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Get episode to find animeId
	episode, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, objID); err != nil {
		return err
	}

	// Update episode count
	count, err := s.repo.CountByAnimeID(ctx, episode.AnimeID)
	if err != nil {
		return err
	}
	return s.animeRepo.UpdateEpisodeCount(ctx, episode.AnimeID, int(count))
}

// GetLatestEpisodes retrieves the latest episodes
func (s *EpisodeService) GetLatestEpisodes(ctx context.Context, limit int) ([]models.EpisodeResponse, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}

	episodes, err := s.repo.GetLatestEpisodes(ctx, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]models.EpisodeResponse, len(episodes))
	for i, ep := range episodes {
		responses[i] = ep.ToResponse()
	}
	return responses, nil
}
