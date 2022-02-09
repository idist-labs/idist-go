package collections

import (
	"ai-camera-api-cms/helpers"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Province struct {
	ID        int64     `bson:"_id" json:"id"`
	Alias     string    `bson:"alias" json:"alias"`
	Name      string    `bson:"name" json:"name"`
	Enable    bool      `bson:"enable" json:"enable"`
	UpdatedAt time.Time `bson:"updated_at" json:"-"`

	Districts []District `bson:"-" json:"districts"`
}

type Provinces []Province

func (u *Province) CollectionName() string {
	return "provinces"
}

func (u *Province) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Province) Update() error {
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
func (u *Province) Find(filter bson.M, opts ...*options.FindOptions) (Provinces, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Provinces, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Province
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

func (u *Province) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "districts" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				entry := District{}
				entries := Districts{}
				entries, _ = entry.Find(bson.M{"province_id": u.ID})
				u.Districts = entries
			}()
		}
	}
	wg.Wait()

}
