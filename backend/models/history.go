package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// History tracks user viewing history
type History struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"userId" json:"userId" validate:"required"`
	AnimeID    primitive.ObjectID `bson:"animeId" json:"animeId" validate:"required"`
	EpisodeID  primitive.ObjectID `bson:"episodeId" json:"episodeId" validate:"required"`
	Progress   int                `bson:"progress" json:"progress" validate:"min=0"` // seconds watched
	Completed  bool               `bson:"completed" json:"completed"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// HistoryResponse with populated data
type HistoryResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	AnimeID   string    `json:"animeId"`
	EpisodeID string    `json:"episodeId"`
	Progress  int       `json:"progress"`
	Completed bool      `json:"completed"`
	UpdatedAt time.Time `json:"updatedAt"`
	Anime     *AnimeResponse `json:"anime,omitempty"`
	Episode   *EpisodeResponse `json:"episode,omitempty"`
}
