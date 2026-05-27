package validators

import (
	"anime-streaming-platform/models"
	"anime-streaming-platform/utils"
	"errors"
	"strings"
)

// ValidateAnime validates anime data
func ValidateAnime(anime *models.Anime) error {
	if strings.TrimSpace(anime.Title) == "" {
		return errors.New("title is required")
	}
	if len(anime.Title) > 200 {
		return errors.New("title must be less than 200 characters")
	}
	if strings.TrimSpace(anime.Description) == "" {
		return errors.New("description is required")
	}
	if len(anime.Description) < 10 {
		return errors.New("description must be at least 10 characters")
	}
	if len(anime.Genres) == 0 {
		return errors.New("at least one genre is required")
	}
	if anime.Images.Poster == "" {
		return errors.New("poster image is required")
	}
	if anime.Status == "" {
		return errors.New("status is required")
	}
	validStatuses := map[string]bool{"ongoing": true, "completed": true, "upcoming": true, "cancelled": true}
	if !validStatuses[anime.Status] {
		return errors.New("invalid status")
	}
	if anime.Rating < 0 || anime.Rating > 10 {
		return errors.New("rating must be between 0 and 10")
	}

	// Auto-generate slug if not provided
	if anime.Slug == "" {
		anime.Slug = utils.GenerateSlug(anime.Title)
	}

	return nil
}
