package repositories

import (
	"context"
	"time"

	"anime-streaming-platform/config"
	"anime-streaming-platform/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// HistoryRepository handles history data operations
type HistoryRepository struct {
	collection *mongo.Collection
}

// NewHistoryRepository creates a new history repository
func NewHistoryRepository() *HistoryRepository {
	return &HistoryRepository{
		collection: config.GetCollection("history"),
	}
}

// Upsert creates or updates a history entry
func (r *HistoryRepository) Upsert(ctx context.Context, history *models.History) error {
	history.UpdatedAt = time.Now()

	filter := bson.M{
		"userId":  history.UserID,
		"animeId": history.AnimeID,
	}

	update := bson.M{
		"$set": bson.M{
			"episodeId": history.EpisodeID,
			"progress":  history.Progress,
			"completed": history.Completed,
			"updatedAt": history.UpdatedAt,
		},
		"$setOnInsert": bson.M{
			"_id":        primitive.NewObjectID(),
			"createdAt":  time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

// GetByUserID retrieves user history with pagination
func (r *HistoryRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID, page, limit int) ([]models.History, int64, error) {
	skip := int64((page - 1) * limit)
	limit64 := int64(limit)

	filter := bson.M{"userId": userID}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit64).
		SetSort(bson.D{{Key: "updatedAt", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var history []models.History
	if err := cursor.All(ctx, &history); err != nil {
		return nil, 0, err
	}
	return history, total, nil
}
