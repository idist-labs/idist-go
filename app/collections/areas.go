package collections

import (
	"ai-camera-api-cms/helpers"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Area struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ProvinceId  int64              `bson:"province_id" json:"province_id"`
	DistrictId  int64              `bson:"district_id" json:"district_id"`
	WardId      int64              `bson:"ward_id" json:"ward_id"`
	VillageId   int64              `bson:"village_id" json:"village_id"`
	Name        string             `bson:"name" json:"name"`
	CountCamera int64              `bson:"-" json:"count_camera"`
	Province    *Province          `bson:"province" json:"province"`
	District    *District          `bson:"district" json:"district"`
	Ward        *Ward              `bson:"ward" json:"ward"`
	Village     *Village           `bson:"village" json:"village"`

	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at" json:"deleted_at"`
}

type Areas []Area

func (u Area) CollectionName() string {
	return "areas"
}

func (u *Area) Create() error {
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

func (u *Area) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Area) FindByAction(action string) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, bson.M{"action": action}); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Area) Find(filter interface{}, opts ...*options.FindOptions) (Areas, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Areas, 0)

	/* Lấy danh sách bản ghi */
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Area
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

func (u *Area) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Area) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.UpdatedAt = time.Now()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Area) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	now := helpers.Now()
	u.DeletedAt = &now
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}
func (u *Area) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "province" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var province Province
				_ = province.First(bson.M{"_id": u.ProvinceId, "deleted_at": nil})
				u.Province = &province
			}()
		}
		if property == "district" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var district District
				_ = district.First(bson.M{"_id": u.DistrictId, "deleted_at": nil})
				u.District = &district
			}()
		}
		if property == "ward" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var ward Ward
				_ = ward.First(bson.M{"_id": u.WardId, "deleted_at": nil})
				u.Ward = &ward
			}()
		}
		if property == "village" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var village Village
				_ = village.First(bson.M{"_id": u.VillageId, "deleted_at": nil})
				u.Village = &village
			}()
		}
	}
	wg.Wait()

}
