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

// CommentRepository handles comment data operations
type CommentRepository struct {
	collection *mongo.Collection
}

// NewCommentRepository creates a new comment repository
func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		collection: config.GetCollection("comments"),
	}
}

// Create inserts a new comment
func (r *CommentRepository) Create(ctx context.Context, comment *models.Comment) error {
	comment.ID = primitive.NewObjectID()
	comment.Likes = 0
	comment.LikedBy = []primitive.ObjectID{}
	comment.IsDeleted = false
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

// GetByAnimeID retrieves comments for an anime with pagination
func (r *CommentRepository) GetByAnimeID(ctx context.Context, animeID primitive.ObjectID, page, limit int) ([]models.Comment, int64, error) {
	skip := int64((page - 1) * limit)
	limit64 := int64(limit)

	filter := bson.M{
		"animeId":  animeID,
		"parentId": bson.M{"$exists": false},
		"isDeleted": false,
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit64).
		SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

// GetReplies retrieves replies for a comment
func (r *CommentRepository) GetReplies(ctx context.Context, parentID primitive.ObjectID) ([]models.Comment, error) {
	filter := bson.M{
		"parentId":  parentID,
		"isDeleted": false,
	}

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var replies []models.Comment
	if err := cursor.All(ctx, &replies); err != nil {
		return nil, err
	}
	return replies, nil
}

// GetByID finds a comment by ID
func (r *CommentRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Comment, error) {
	var comment models.Comment
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("comment not found")
	}
	return &comment, err
}

// AddLike adds a like to a comment
func (r *CommentRepository) AddLike(ctx context.Context, commentID, userID primitive.ObjectID) error {
	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{
			"_id":     commentID,
			"likedBy": bson.M{"$ne": userID},
		},
		bson.M{
			"$addToSet": bson.M{"likedBy": userID},
			"$inc":      bson.M{"likes": 1},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("comment not found or already liked")
	}
	return nil
}

// SoftDelete marks a comment as deleted
func (r *CommentRepository) SoftDelete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"isDeleted": true, "updatedAt": time.Now()}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("comment not found")
	}
	return nil
}

// GetRecentComments retrieves recent comments for moderation
func (r *CommentRepository) GetRecentComments(ctx context.Context, page, limit int) ([]models.Comment, int64, error) {
	skip := int64((page - 1) * limit)
	limit64 := int64(limit)

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit64).
		SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, 0, err
	}
	return comments, total, nil
}

// CountByAnimeID counts comments for an anime
func (r *CommentRepository) CountByAnimeID(ctx context.Context, animeID primitive.ObjectID) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{"animeId": animeID, "isDeleted": false})
}
