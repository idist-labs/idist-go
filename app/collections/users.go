package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
	PasswordHash string             `bson:"password_hash" json:"password_hash"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at" json:"-"`

	// Tmp Value
	Password   string `bson:"-" json:"password"`
	RePassword string `bson:"-" json:"re_password"`
}
