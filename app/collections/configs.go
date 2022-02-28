package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"strconv"
	"time"
)

type Config struct {
	ID        string     `bson:"_id" json:"id"`
	Value     string     `bson:"value" json:"value"`
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at" json:"deleted_at"`
}

type Configs []Config

func (u *Config) CollectionName() string {
	return "configs"
}

func (u *Config) GetInt64(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	filter := bson.M{
		"_id":        key,
		"deleted_at": nil,
	}
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return 0, result.Err()
	} else if result.Err() == mongo.ErrNoDocuments {
		u.CreatedAt = helpers.Now()
		u.UpdatedAt = helpers.Now()
		u.ID = key
		if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
			return 0, err
		} else {
			return 1, nil
		}
	} else {
		if err := result.Decode(&u); err != nil {
			return 0, err
		}
	}
	return strconv.ParseInt(u.Value, 10, 64)
}

func (u *Config) SetInt64(key string, value int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	filter := bson.M{
		"_id":        key,
		"deleted_at": nil,
	}
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return result.Err()
	} else if result.Err() == mongo.ErrNoDocuments {
		u.CreatedAt = helpers.Now()
		u.UpdatedAt = helpers.Now()
		u.ID = key
		u.Value = string(value)
		if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		if err := result.Decode(&u); err != nil {
			return err
		}
	}
	u.Value = strconv.FormatInt(value, 10)
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}
