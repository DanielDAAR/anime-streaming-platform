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

// EpisodeRepository handles episode data operations
type EpisodeRepository struct {
	collection *mongo.Collection
}

// NewEpisodeRepository creates a new episode repository
func NewEpisodeRepository() *EpisodeRepository {
	return &EpisodeRepository{
		collection: config.GetCollection("episodes"),
	}
}

// Create inserts a new episode
func (r *EpisodeRepository) Create(ctx context.Context, episode *models.Episode) error {
	episode.ID = primitive.NewObjectID()
	episode.CreatedAt = time.Now()
	episode.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, episode)
	return err
}

// GetByID finds an episode by ID
func (r *EpisodeRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Episode, error) {
	var episode models.Episode
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&episode)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("episode not found")
	}
	return &episode, err
}

// GetByAnimeID retrieves episodes for a specific anime
func (r *EpisodeRepository) GetByAnimeID(ctx context.Context, animeID primitive.ObjectID, page, limit int) ([]models.Episode, int64, error) {
	skip := int64((page - 1) * limit)
	limit64 := int64(limit)

	filter := bson.M{"animeId": animeID}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit64).
		SetSort(bson.D{{Key: "number", Value: 1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var episodes []models.Episode
	if err := cursor.All(ctx, &episodes); err != nil {
		return nil, 0, err
	}

	return episodes, total, nil
}

// GetByAnimeAndNumber finds a specific episode by anime and number
func (r *EpisodeRepository) GetByAnimeAndNumber(ctx context.Context, animeID primitive.ObjectID, number int) (*models.Episode, error) {
	var episode models.Episode
	err := r.collection.FindOne(ctx, bson.M{"animeId": animeID, "number": number}).Decode(&episode)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("episode not found")
	}
	return &episode, err
}

// Update modifies an existing episode
func (r *EpisodeRepository) Update(ctx context.Context, id primitive.ObjectID, episode *models.Episode) error {
	episode.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"number":      episode.Number,
			"title":       episode.Title,
			"description": episode.Description,
			"servers":     episode.Servers,
			"duration":    episode.Duration,
			"thumbnail":   episode.Thumbnail,
			"updatedAt":   episode.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("episode not found")
	}
	return nil
}

// Delete removes an episode
func (r *EpisodeRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("episode not found")
	}
	return nil
}

// DeleteByAnimeID removes all episodes that belong to an anime.
func (r *EpisodeRepository) DeleteByAnimeID(ctx context.Context, animeID primitive.ObjectID) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"animeId": animeID})
	return err
}

// CountByAnimeID counts episodes for an anime
func (r *EpisodeRepository) CountByAnimeID(ctx context.Context, animeID primitive.ObjectID) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{"animeId": animeID})
}

// GetLatestEpisodes retrieves the latest episodes across all animes
func (r *EpisodeRepository) GetLatestEpisodes(ctx context.Context, limit int) ([]models.Episode, error) {
	opts := options.Find().SetLimit(int64(limit)).SetSort(bson.D{{Key: "createdAt", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var episodes []models.Episode
	if err := cursor.All(ctx, &episodes); err != nil {
		return nil, err
	}
	return episodes, nil
}
