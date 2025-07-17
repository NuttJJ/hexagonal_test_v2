package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password,omitempty" json:"-"` // optional สำหรับ OAuth
	GoogleID   string             `bson:"google_id,omitempty" json:"-"`
	Avatar     string             `bson:"avatar,omitempty" json:"avatar"`
	Provider   string             `bson:"provider" json:"provider"` // "local", "google"
	IsVerified bool               `bson:"is_verified" json:"is_verified"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
