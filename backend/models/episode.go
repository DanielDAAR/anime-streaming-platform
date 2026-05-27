package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Episode represents a single episode of an anime
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AnimeID     primitive.ObjectID `bson:"animeId" json:"animeId" validate:"required"`
	Number      int                `bson:"number" json:"number" validate:"required,min=1"`
	Title       string             `bson:"title" json:"title" validate:"required,min=1,max=200"`
	Description string             `bson:"description" json:"description" validate:"max=1000"`
	Servers     []EmbedServer      `bson:"servers" json:"servers" validate:"required,min=1,dive"`
	Duration    int                `bson:"duration" json:"duration" validate:"min=0"` // in minutes
	Thumbnail   string             `bson:"thumbnail" json:"thumbnail" validate:"url"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// EmbedServer represents an external streaming server
type EmbedServer struct {
	Name    string `bson:"name" json:"name" validate:"required,min=2,max=50"`
	URL     string `bson:"url" json:"url" validate:"required,url"`
	Quality string `bson:"quality" json:"quality" validate:"oneof=360p 480p 720p 1080p unknown"`
	Active  bool   `bson:"active" json:"active"`
}

// EpisodeResponse is the sanitized version for public API
type EpisodeResponse struct {
	ID          string        `json:"id"`
	AnimeID     string        `json:"animeId"`
	Number      int           `json:"number"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Servers     []EmbedServer `json:"servers"`
	Duration    int           `json:"duration"`
	Thumbnail   string        `json:"thumbnail"`
	CreatedAt   time.Time     `json:"createdAt"`
}

// ToResponse converts Episode to EpisodeResponse
func (e *Episode) ToResponse() EpisodeResponse {
	return EpisodeResponse{
		ID:          e.ID.Hex(),
		AnimeID:     e.AnimeID.Hex(),
		Number:      e.Number,
		Title:       e.Title,
		Description: e.Description,
		Servers:     e.Servers,
		Duration:    e.Duration,
		Thumbnail:   e.Thumbnail,
		CreatedAt:   e.CreatedAt,
	}
}
