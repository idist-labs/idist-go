package collections

import (
	"ai-camera-api-cms/app/providers/configProvider"
	"ai-camera-api-cms/helpers"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Notification struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Title      string             `bson:"title" json:"title"`
	Message    string             `bson:"message" json:"message"`
	CameraId   primitive.ObjectID `bson:"camera_id" json:"camera_id"`
	AreaId     primitive.ObjectID `bson:"area_id" json:"area_id"`
	ProvinceId int64              `bson:"province_id" json:"province_id"`
	DistrictId int64              `bson:"district_id" json:"district_id"`
	WardId     int64              `bson:"ward_id" json:"ward_id"`
	VillageId  int64              `bson:"village_id" json:"village_id"`
	Type       string             `bson:"type" json:"type"`
	Status     int64              `bson:"status" json:"status"`
	Images     []string           ` bson:"images" json:"images"`
	Videos     []string           `bson:"videos" json:"videos"`
	Time       time.Time          `bson:"time" json:"time"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	DeletedAt  *time.Time         `bson:"deleted_at" json:"deleted_at"`
	Camera     Camera             `bson:"-" json:"camera"`
	Area       Area               `bson:"-" json:"area"`
}

type Notifications []Notification

func (u Notification) CollectionName() string {
	return "notifications"
}

func (u *Notification) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.CreatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Notification) ConvertMedia() {
	con := configProvider.GetConfig()
	for i := 0; i < len(u.Images); i++ {
		if u.Images[i] != "" && u.Images[i][0] == '/' {
			u.Images[i] = con.GetString("app.server.ai_host") + u.Images[i]
		}
	}
	for i := 0; i < len(u.Videos); i++ {
		if u.Videos[i] != "" && u.Videos[i][0] == '/' {
			u.Videos[i] = con.GetString("app.server.ai_host") + u.Videos[i]
		}
	}
}
func (u *Notification) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		err := result.Decode(&u)
		u.ConvertMedia()
		return err
	}
}

func (u *Notification) FindByAction(action string) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, bson.M{"action": action}); result.Err() != nil {
		return result.Err()
	} else {
		err := result.Decode(&u)
		u.ConvertMedia()
		return err
	}
}

func (u *Notification) Find(filter interface{}, opts ...*options.FindOptions) (Notifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Notifications, 0)

	/* Lấy danh sách bản ghi */
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Notification
			if err = cursor.Decode(&elem); err != nil {
				return data, err
			}
			elem.ConvertMedia()
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

func (u *Notification) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Notification) Delete() error {
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

func (u *Notification) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "camera" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var camera Camera
				_ = camera.First(bson.M{"_id": u.CameraId, "deleted_at": nil})
				u.Camera = camera
			}()
		}
		if property == "area" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var entry Area
				_ = entry.First(bson.M{"_id": u.AreaId, "deleted_at": nil})
				u.Area = entry
			}()
		}
	}
	wg.Wait()
}
