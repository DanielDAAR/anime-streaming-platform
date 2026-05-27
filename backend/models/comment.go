package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment represents a user comment on an anime
type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AnimeID   primitive.ObjectID `bson:"animeId" json:"animeId" validate:"required"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId" validate:"required"`
	ParentID  *primitive.ObjectID `bson:"parentId,omitempty" json:"parentId,omitempty"`
	Content   string             `bson:"content" json:"content" validate:"required,min=1,max=1000"`
	Likes     int                `bson:"likes" json:"likes"`
	LikedBy   []primitive.ObjectID `bson:"likedBy" json:"-"`
	IsDeleted bool               `bson:"isDeleted" json:"isDeleted"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`

	// Populated fields (not stored in DB)
	User     *PublicUser  `bson:"-" json:"user,omitempty"`
	Replies  []Comment    `bson:"-" json:"replies,omitempty"`
}

// CommentResponse is the sanitized version for public API
type CommentResponse struct {
	ID        string          `json:"id"`
	AnimeID   string          `json:"animeId"`
	UserID    string          `json:"userId"`
	ParentID  string          `json:"parentId,omitempty"`
	Content   string          `json:"content"`
	Likes     int             `json:"likes"`
	IsDeleted bool            `json:"isDeleted"`
	CreatedAt time.Time       `json:"createdAt"`
	User      *PublicUser     `json:"user,omitempty"`
	Replies   []CommentResponse `json:"replies,omitempty"`
}

// ToResponse converts Comment to CommentResponse
func (c *Comment) ToResponse() CommentResponse {
	resp := CommentResponse{
		ID:        c.ID.Hex(),
		AnimeID:   c.AnimeID.Hex(),
		UserID:    c.UserID.Hex(),
		Content:   c.Content,
		Likes:     c.Likes,
		IsDeleted: c.IsDeleted,
		CreatedAt: c.CreatedAt,
	}

	if c.ParentID != nil {
		parentID := c.ParentID.Hex()
		resp.ParentID = parentID
	}

	if c.User != nil {
		resp.User = c.User
	}

	if len(c.Replies) > 0 {
		replies := make([]CommentResponse, len(c.Replies))
		for i, reply := range c.Replies {
			replies[i] = reply.ToResponse()
		}
		resp.Replies = replies
	}

	return resp
}
