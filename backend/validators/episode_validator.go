package validators

import (
	"anime-streaming-platform/models"
	"errors"
	"strings"
)

// ValidateEpisode validates episode data
func ValidateEpisode(episode *models.Episode) error {
	if episode.AnimeID.IsZero() {
		return errors.New("animeId is required")
	}
	if episode.Number < 1 {
		return errors.New("episode number must be greater than 0")
	}
	if strings.TrimSpace(episode.Title) == "" {
		return errors.New("title is required")
	}
	if len(episode.Servers) == 0 {
		return errors.New("at least one server is required")
	}

	validQualities := map[string]bool{"360p": true, "480p": true, "720p": true, "1080p": true, "unknown": true}
	for i, server := range episode.Servers {
		if strings.TrimSpace(server.Name) == "" {
			return errors.New("server name is required")
		}
		if strings.TrimSpace(server.URL) == "" {
			return errors.New("server URL is required")
		}
		if !validQualities[server.Quality] {
			return errors.New("invalid quality for server " + string(rune(i)))
		}
	}

	return nil
}
