package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type District struct {
	ID         int64     `bson:"_id" json:"id"`
	Name       string    `bson:"name" json:"name"`
	Alias      string    `bson:"alias" json:"alias"`
	ProvinceID int64     `bson:"province_id" json:"province_id"`
	IsShow     bool      `bson:"is_show" json:"is_show"`
	UpdatedAt  time.Time `bson:"updated_at" json:"-"`
	Wards      []Ward    `bson:"-" json:"wards"`
}

type Districts []District

func (u *District) CollectionName() string {
	return "districts"
}

func (u *District) Find(filter interface{}, opts ...*options.FindOptions) (Districts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Districts, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem District
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

func (u *District) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *District) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *District) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *District) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "wards" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				entry := Ward{}
				entries := Wards{}
				entries, _ = entry.Find(bson.M{"district_id": u.ID})
				u.Wards = entries
			}()
		}
	}
	wg.Wait()

}
