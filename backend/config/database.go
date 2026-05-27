package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB     *mongo.Database
	Client *mongo.Client
)

// ConnectDB establishes connection to MongoDB
func ConnectDB(ctx context.Context) error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGODB_DB_NAME")
	if dbName == "" {
		dbName = "anime_streaming_db"
	}

	clientOptions := options.Client().ApplyURI(uri).
		SetServerSelectionTimeout(10 * time.Second).
		SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	Client = client
	DB = client.Database(dbName)

	log.Println("✅ Connected to MongoDB successfully")

	// Create indexes
	if err := createIndexes(ctx); err != nil {
		log.Printf("Warning: failed to create some indexes: %v", err)
	}

	return nil
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() error {
	if Client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return Client.Disconnect(ctx)
}

// GetCollection returns a MongoDB collection
func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database not initialized. Call ConnectDB first.")
	}
	return DB.Collection(name)
}

// createIndexes sets up database indexes for optimal performance
func createIndexes(ctx context.Context) error {
	// Animes indexes
	animeColl := GetCollection("animes")
	animeIndexes := []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"slug": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"title": "text", "description": "text", "genres": "text"},
		},
		{
			Keys: map[string]interface{}{"rating": -1},
		},
		{
			Keys: map[string]interface{}{"createdAt": -1},
		},
	}
	if _, err := animeColl.Indexes().CreateMany(ctx, animeIndexes); err != nil {
		return fmt.Errorf("anime indexes: %w", err)
	}

	// Episodes indexes
	episodeColl := GetCollection("episodes")
	episodeIndexes := []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"animeId": 1, "number": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"createdAt": -1},
		},
	}
	if _, err := episodeColl.Indexes().CreateMany(ctx, episodeIndexes); err != nil {
		return fmt.Errorf("episode indexes: %w", err)
	}

	// Users indexes
	userColl := GetCollection("users")
	userIndexes := []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    map[string]interface{}{"username": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"role": 1},
		},
	}
	if _, err := userColl.Indexes().CreateMany(ctx, userIndexes); err != nil {
		return fmt.Errorf("user indexes: %w", err)
	}

	// Comments indexes
	commentColl := GetCollection("comments")
	commentIndexes := []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"animeId": 1, "createdAt": -1},
		},
		{
			Keys: map[string]interface{}{"userId": 1},
		},
		{
			Keys: map[string]interface{}{"parentId": 1},
		},
	}
	if _, err := commentColl.Indexes().CreateMany(ctx, commentIndexes); err != nil {
		return fmt.Errorf("comment indexes: %w", err)
	}

	// History indexes
	historyColl := GetCollection("history")
	historyIndexes := []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"userId": 1, "animeId": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"updatedAt": -1},
		},
	}
	if _, err := historyColl.Indexes().CreateMany(ctx, historyIndexes); err != nil {
		return fmt.Errorf("history indexes: %w", err)
	}

	log.Println("✅ Database indexes created successfully")
	return nil
}
