package repositories

import (
	"context"
	"errors"
	"time"

	"anime-streaming-platform/config"
	"anime-streaming-platform/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AnimeRepository handles anime data operations
type AnimeRepository struct {
	collection *mongo.Collection
}

// NewAnimeRepository creates a new anime repository
func NewAnimeRepository() *AnimeRepository {
	return &AnimeRepository{
		collection: config.GetCollection("animes"),
	}
}

// Create inserts a new anime
func (r *AnimeRepository) Create(ctx context.Context, anime *models.Anime) error {
	anime.ID = primitive.NewObjectID()
	anime.CreatedAt = time.Now()
	anime.UpdatedAt = time.Now()
	anime.Episodes = 0

	_, err := r.collection.InsertOne(ctx, anime)
	return err
}

// GetByID finds an anime by ID
func (r *AnimeRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Anime, error) {
	var anime models.Anime
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&anime)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("anime not found")
	}
	return &anime, err
}

// GetBySlug finds an anime by slug
func (r *AnimeRepository) GetBySlug(ctx context.Context, slug string) (*models.Anime, error) {
	var anime models.Anime
	err := r.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&anime)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("anime not found")
	}
	return &anime, err
}

// GetAll retrieves paginated animes with optional filters
func (r *AnimeRepository) GetAll(ctx context.Context, page, limit int, filters map[string]interface{}) ([]models.Anime, int64, error) {
	skip := int64((page - 1) * limit)
	limit64 := int64(limit)

	// Build filter
	filter := bson.M{}
	if search, ok := filters["search"].(string); ok && search != "" {
		filter["$text"] = bson.M{"$search": search}
	}
	if genre, ok := filters["genre"].(string); ok && genre != "" {
		filter["genres"] = genre
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		filter["status"] = status
	}
	if year, ok := filters["year"].(int); ok && year > 0 {
		filter["year"] = year
	}

	// Count total
	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Find with pagination
	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit64).
		SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var animes []models.Anime
	if err := cursor.All(ctx, &animes); err != nil {
		return nil, 0, err
	}

	return animes, total, nil
}

// Update modifies an existing anime
func (r *AnimeRepository) Update(ctx context.Context, id primitive.ObjectID, anime *models.Anime) error {
	anime.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"slug":        anime.Slug,
			"title":       anime.Title,
			"description": anime.Description,
			"genres":      anime.Genres,
			"rating":      anime.Rating,
			"images":      anime.Images,
			"status":      anime.Status,
			"year":        anime.Year,
			"studio":      anime.Studio,
			"seo":         anime.SEO,
			"updatedAt":   anime.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("anime not found")
	}
	return nil
}

// Delete removes an anime
func (r *AnimeRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("anime not found")
	}
	return nil
}

// UpdateEpisodeCount updates the episode count for an anime
func (r *AnimeRepository) UpdateEpisodeCount(ctx context.Context, animeID primitive.ObjectID, count int) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": animeID},
		bson.M{"$set": bson.M{"episodesCount": count, "updatedAt": time.Now()}},
	)
	return err
}

// GetLatest retrieves the latest animes
func (r *AnimeRepository) GetLatest(ctx context.Context, limit int) ([]models.Anime, error) {
	opts := options.Find().SetLimit(int64(limit)).SetSort(bson.D{{Key: "createdAt", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var animes []models.Anime
	if err := cursor.All(ctx, &animes); err != nil {
		return nil, err
	}
	return animes, nil
}

// GetTopRated retrieves top rated animes
func (r *AnimeRepository) GetTopRated(ctx context.Context, limit int) ([]models.Anime, error) {
	opts := options.Find().SetLimit(int64(limit)).SetSort(bson.D{{Key: "rating", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var animes []models.Anime
	if err := cursor.All(ctx, &animes); err != nil {
		return nil, err
	}
	return animes, nil
}
