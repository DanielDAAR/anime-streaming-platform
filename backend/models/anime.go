package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Anime represents an anime series in the platform
type Anime struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Slug        string             `bson:"slug" json:"slug" validate:"required,min=3,max=100"`
	Title       string             `bson:"title" json:"title" validate:"required,min=1,max=200"`
	Description string             `bson:"description" json:"description" validate:"required,min=10,max=2000"`
	Genres      []string           `bson:"genres" json:"genres" validate:"required,min=1,dive,min=2,max=50"`
	Rating      float64            `bson:"rating" json:"rating" validate:"min=0,max=10"`
	Images      AnimeImages        `bson:"images" json:"images"`
	Status      string             `bson:"status" json:"status" validate:"required,oneof=ongoing completed upcoming cancelled"`
	Episodes    int                `bson:"episodesCount" json:"episodesCount"`
	Year        int                `bson:"year" json:"year" validate:"min=1900,max=2100"`
	Studio      string             `bson:"studio" json:"studio"`
	SEO         SEOMetadata        `bson:"seo" json:"seo"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// AnimeImages contains all image URLs for an anime
type AnimeImages struct {
	Poster    string `bson:"poster" json:"poster" validate:"required,url"`
	Banner    string `bson:"banner" json:"banner" validate:"url"`
	Thumbnail string `bson:"thumbnail" json:"thumbnail" validate:"url"`
}

// SEOMetadata contains search engine optimization data
type SEOMetadata struct {
	MetaTitle       string `bson:"metaTitle" json:"metaTitle"`
	MetaDescription string `bson:"metaDescription" json:"metaDescription"`
	Keywords        string `bson:"keywords" json:"keywords"`
	CanonicalURL    string `bson:"canonicalUrl" json:"canonicalUrl"`
}

// AnimeResponse is the sanitized version for public API
type AnimeResponse struct {
	ID          string      `json:"id"`
	Slug        string      `json:"slug"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Genres      []string    `json:"genres"`
	Rating      float64     `json:"rating"`
	Images      AnimeImages `json:"images"`
	Status      string      `json:"status"`
	Episodes    int         `json:"episodesCount"`
	Year        int         `json:"year"`
	Studio      string      `json:"studio"`
	CreatedAt   time.Time   `json:"createdAt"`
}

// ToResponse converts Anime to AnimeResponse
func (a *Anime) ToResponse() AnimeResponse {
	return AnimeResponse{
		ID:          a.ID.Hex(),
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Genres:      a.Genres,
		Rating:      a.Rating,
		Images:      a.Images,
		Status:      a.Status,
		Episodes:    a.Episodes,
		Year:        a.Year,
		Studio:      a.Studio,
		CreatedAt:   a.CreatedAt,
	}
}
