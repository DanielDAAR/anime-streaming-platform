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

// UserRepository handles user data operations
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: config.GetCollection("users"),
	}
}

// Create inserts a new user
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.Role = models.RoleUser
	user.IsActive = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

// GetByID finds a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

// GetByEmail finds a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

// GetByUsername finds a user by username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}

// GetAll retrieves paginated users
func (r *UserRepository) GetAll(ctx context.Context, page, limit int) ([]models.User, int64, error) {
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

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateLastLogin updates the last login timestamp
func (r *UserRepository) UpdateLastLogin(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"lastLogin": time.Now(), "updatedAt": time.Now()}},
	)
	return err
}

// UpdateRole changes user role
func (r *UserRepository) UpdateRole(ctx context.Context, id primitive.ObjectID, role models.UserRole) error {
	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"role": role, "updatedAt": time.Now()}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

// ToggleActive toggles user active status
func (r *UserRepository) ToggleActive(ctx context.Context, id primitive.ObjectID) error {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	_, err = r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"isActive": !user.IsActive, "updatedAt": time.Now()}},
	)
	return err
}

// CountByRole counts users by role
func (r *UserRepository) CountByRole(ctx context.Context, role models.UserRole) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{"role": role})
}
