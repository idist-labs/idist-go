package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"time"
)

type Ward struct {
	ID         int64  `bson:"_id" json:"id"`
	Name       string `bson:"name" json:"name"`
	DistrictId int64  `bson:"district_id" json:"district_id"`
	IsShow     bool   `bson:"is_show" json:"is_show"`

	CreatedBy primitive.ObjectID `bson:"created_by" json:"created_by"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedBy primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedBy primitive.ObjectID `bson:"deleted_by" json:"deleted_by"`
	DeletedAt *time.Time         `bson:"deleted_at" json:"deleted_at"`
}

type Wards []Ward

func (u *Ward) CollectionName() string {
	return "wards"
}

func (u *Ward) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Ward) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Ward) Find(filter bson.M, opts ...*options.FindOptions) (Wards, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Wards, 0)

	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Ward
			if err = cursor.Decode(&elem); err != nil {
				return data, err
			}
			data = append(data, elem)
		}
		if err = cursor.Err(); err != nil {
			return data, err
		}
		return data, cursor.Close(ctx)
	} else {
		return data, err
	}
}

func (u *Ward) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}
