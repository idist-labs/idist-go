package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"sync"
	"time"
)

type School struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Latitude   string             `bson:"latitude" json:"latitude"`
	Longitude  string             `bson:"longitude" json:"longitude"`
	Enable     bool               `bson:"enable" json:"enable"`
	TotalUsers int64              `bson:"total_users" json:"total_users"`
	ProvinceId int64              `bson:"province_id" json:"province_id"`
	DistrictId int64              `bson:"district_id" json:"district_id"`
	WardId     int64              `bson:"ward_id" json:"ward_id"`

	CreatedAt time.Time  `bson:"created_at" json:"-"`
	UpdatedAt time.Time  `bson:"updated_at" json:"-"`
	DeletedAt *time.Time `bson:"deleted_at" json:"-"`

	// Preload
	Province Province  `bson:"-" json:"province"`
	District District  `bson:"-" json:"district"`
	Ward     Ward      `bson:"-" json:"ward"`
	Students []Student `bson:"-" json:"students"`
}

type Schools []School

func (u *School) CollectionName() string {
	return "schools"
}

func (u *School) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *School) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.CreatedAt = helpers.Now()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *School) Update() error {
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
func (u *School) Find(filter bson.M, opts ...*options.FindOptions) (Schools, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Schools, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem School
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

func (u *School) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.DeletedAt = helpers.PNow()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *School) Preload(properties ...string) {
	var wg sync.WaitGroup
	//for _, property := range properties {
	//	//if property == "districts" {
	//	//	wg.Add(1)
	//	//	go func() {
	//	//		defer wg.Done()
	//	//		entry := District{}
	//	//		entries := Districts{}
	//	//		entries, _ = entry.Find(bson.M{"province_id": u.ID})
	//	//		u.Districts = entries
	//	//	}()
	//	//}
	//}
	wg.Wait()

}
