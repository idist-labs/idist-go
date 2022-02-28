package collections

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/providers/configProvider"
	"idist-core/helpers"
	"sync"
	"time"
)

type User struct {
	ID                    primitive.ObjectID `bson:"_id" json:"id"`
	Name                  string             `bson:"name" json:"name"`
	Avatar                string             `bson:"avatar" json:"avatar"`
	Username              string             `bson:"username" json:"username"`
	Password              string             `bson:"password_hash" json:"-"`
	NewPassword           string             `bson:"-" json:"password"`
	Dob                   string             `bson:"dob" json:"dob"`
	Phone                 string             `bson:"phone" json:"phone"`
	Address               string             `bson:"address" json:"address"`
	Authentication        string             `bson:"authentication" json:"authentication"`
	Image                 []string           `bson:"image" json:"image"`
	Lock                  bool               `bson:"lock" json:"lock"`
	RoleId                primitive.ObjectID `bson:"role_id" json:"role_id"`
	ProvinceId            int64              `bson:"province_id" json:"province_id"`
	DistrictId            int64              `bson:"district_id" json:"district_id"`
	WardId                int64              `bson:"ward_id" json:"ward_id"`
	VillageId             int64              `bson:"village_id" json:"village_id"`
	LockedAt              *time.Time         `bson:"locked_at" json:"locked_at"`
	Status                string             `bson:"status" json:"status" form:"status"`
	LoginFailedCount      int64              `bson:"login_failed_count" json:"login_failed_count"`
	LastActiveAt          *time.Time         `bson:"last_active_at" json:"last_active_at"`
	LastLoggedAt          *time.Time         `bson:"last_logged_at" json:"last_logged_at"`
	LastChangedPasswordAt *time.Time         `bson:"last_changed_password_at" json:"last_changed_password_at"`
	LastLoggedFailAt      *time.Time         `bson:"last_logged_fail_at" json:"last_logged_fail_at"`
	Province              *Province          `bson:"-" json:"province"`
	District              *District          `bson:"-" json:"district"`
	Ward                  *Ward              `bson:"-" json:"ward"`

	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at" json:"-"`
	CreatedBy primitive.ObjectID `bson:"created_by" json:"created_by"`
	UpdatedBy primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	DeletedBy primitive.ObjectID `bson:"deleted_by" json:"deleted_by"`
	Role      *Role              `bson:"-" json:"role"`
}

type Users []User

func (u *User) CollectionName() string {
	return "users"
}

func (u *User) String() string {
	uJSON, _ := json.Marshal(u)
	return string(uJSON)
}

func (u *User) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *User) Create() error {
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

func (u *User) Update() error {
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

func (u *User) Find(filter bson.M, opts ...*options.FindOptions) (Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Users, 0)

	/* Lấy danh sách bản ghi */
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem User
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

func (u *User) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).DeleteOne(ctx, bson.M{"_id": u.ID}, options.Delete()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *User) LoginFail() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	now := helpers.Now()
	u.UpdatedAt = now
	u.LastLoggedFailAt = &now
	u.LoginFailedCount++
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *User) LastActive() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": bson.M{
			"last_active_at": time.Now(),
		},
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *User) Preload(properties ...string) {
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
	}
	wg.Wait()

}

func (u *User) HasPermission() bool {
	config := configProvider.GetConfig()
	adminUsername := config.GetStringSlice("auth.admin")

	for _, entry := range adminUsername {
		if entry == u.Username {
			return true
		}
	}
	return false
}

func (u *User) IsAdmin() bool {
	roleAdmin, _ := primitive.ObjectIDFromHex(configProvider.GetConfig().GetString("role.admin_id"))
	if u.RoleId == roleAdmin {
		return true
	}
	adminUsername := configProvider.GetConfig().GetStringSlice("adminUsername")

	for _, entry := range adminUsername {
		if entry == u.Username {
			return true
		}
	}
	return false
}
