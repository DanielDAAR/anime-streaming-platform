package services

import (
	"context"
	"math"

	"anime-streaming-platform/models"
	"anime-streaming-platform/repositories"
	"anime-streaming-platform/utils"
	"anime-streaming-platform/validators"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AnimeService handles anime business logic
type AnimeService struct {
	repo        *repositories.AnimeRepository
	episodeRepo *repositories.EpisodeRepository
	commentRepo *repositories.CommentRepository
}

// NewAnimeService creates a new anime service
func NewAnimeService() *AnimeService {
	return &AnimeService{
		repo:        repositories.NewAnimeRepository(),
		episodeRepo: repositories.NewEpisodeRepository(),
		commentRepo: repositories.NewCommentRepository(),
	}
}

// CreateAnime creates a new anime
func (s *AnimeService) CreateAnime(ctx context.Context, anime *models.Anime) (*models.AnimeResponse, error) {
	if err := validators.ValidateAnime(anime); err != nil {
		return nil, err
	}

	// Generate slug if empty
	if anime.Slug == "" {
		anime.Slug = utils.GenerateSlug(anime.Title)
	}

	// Set default SEO if not provided
	if anime.SEO.MetaTitle == "" {
		anime.SEO.MetaTitle = anime.Title + " | Anime Streaming Platform"
	}
	if anime.SEO.MetaDescription == "" {
		anime.SEO.MetaDescription = anime.Description
	}
	if anime.SEO.CanonicalURL == "" {
		anime.SEO.CanonicalURL = "/anime/" + anime.Slug
	}

	if err := s.repo.Create(ctx, anime); err != nil {
		return nil, err
	}

	resp := anime.ToResponse()
	return &resp, nil
}

// GetAnimeBySlug retrieves an anime by slug
func (s *AnimeService) GetAnimeBySlug(ctx context.Context, slug string) (*models.AnimeResponse, error) {
	anime, err := s.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	resp := anime.ToResponse()
	return &resp, nil
}

// GetAnimeByID retrieves an anime by ID
func (s *AnimeService) GetAnimeByID(ctx context.Context, id string) (*models.AnimeResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	anime, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	resp := anime.ToResponse()
	return &resp, nil
}

// GetAllAnimes retrieves paginated animes with filters
func (s *AnimeService) GetAllAnimes(ctx context.Context, page, limit int, filters map[string]interface{}) ([]models.AnimeResponse, int, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	animes, total, err := s.repo.GetAll(ctx, page, limit, filters)
	if err != nil {
		return nil, 0, 0, err
	}

	responses := make([]models.AnimeResponse, len(animes))
	for i, anime := range animes {
		responses[i] = anime.ToResponse()
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return responses, int(total), totalPages, nil
}

// UpdateAnime updates an existing anime
func (s *AnimeService) UpdateAnime(ctx context.Context, id string, anime *models.Anime) (*models.AnimeResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := validators.ValidateAnime(anime); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, objID, anime); err != nil {
		return nil, err
	}

	updated, err := s.repo.GetByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	resp := updated.ToResponse()
	return &resp, nil
}

// DeleteAnime removes an anime and its episodes
func (s *AnimeService) DeleteAnime(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := s.episodeRepo.DeleteByAnimeID(ctx, objID); err != nil {
		return err
	}

	return s.repo.Delete(ctx, objID)
}

// GetLatestAnimes retrieves the latest animes
func (s *AnimeService) GetLatestAnimes(ctx context.Context, limit int) ([]models.AnimeResponse, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}

	animes, err := s.repo.GetLatest(ctx, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]models.AnimeResponse, len(animes))
	for i, anime := range animes {
		responses[i] = anime.ToResponse()
	}
	return responses, nil
}

// GetTopRatedAnimes retrieves top rated animes
func (s *AnimeService) GetTopRatedAnimes(ctx context.Context, limit int) ([]models.AnimeResponse, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}

	animes, err := s.repo.GetTopRated(ctx, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]models.AnimeResponse, len(animes))
	for i, anime := range animes {
		responses[i] = anime.ToResponse()
	}
	return responses, nil
}

// GetDashboardStats returns statistics for admin dashboard
func (s *AnimeService) GetDashboardStats(ctx context.Context) (map[string]interface{}, error) {
	latest, err := s.repo.GetLatest(ctx, 5)
	if err != nil {
		return nil, err
	}

	latestResp := make([]models.AnimeResponse, len(latest))
	for i, a := range latest {
		latestResp[i] = a.ToResponse()
	}

	return map[string]interface{}{
		"latestAnimes": latestResp,
	}, nil
}
