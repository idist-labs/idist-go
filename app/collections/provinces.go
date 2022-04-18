package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"sync"
	"time"
)

type Province struct {
	ID             int64  `bson:"_id" json:"id"`
	Name           string `bson:"name" json:"name"`
	Domain         string `bson:"domain" json:"domain"`
	RegionId       int64  `bson:"region_id" json:"region_id"`
	Latitude       string `bson:"latitude" json:"latitude"`
	Longitude      string `bson:"longitude" json:"longitude"`
	Enable         bool   `bson:"enable" json:"enable"`
	TotalDistricts int64  `bson:"total_districts" json:"total_districts"`
	TotalUsers     int64  `bson:"total_users" json:"total_users"`
	TotalSchools   int64  `bson:"total_schools" json:"total_schools"`
	Color          string `bson:"color" json:"color"`
	Path           string `bson:"path" json:"path"`

	CreatedAt time.Time  `bson:"created_at" json:"-"`
	UpdatedAt time.Time  `bson:"updated_at" json:"-"`
	DeletedAt *time.Time `bson:"deleted_at" json:"-"`

	// Preload
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
				u.TotalDistricts = int64(len(entries))
				u.Districts = entries
			}()
		}
	}
	wg.Wait()

}
