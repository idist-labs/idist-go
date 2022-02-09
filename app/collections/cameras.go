package collections

import (
	"ai-camera-api-cms/app/providers/configProvider"
	"ai-camera-api-cms/helpers"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Camera struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	UserId       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"password"`
	Channel      string             `bson:"channel" json:"channel"`
	ConfigCamera ConfigCamera       `bson:"config" json:"config"`
	ProvinceId   int64              `bson:"province_id" json:"province_id"`
	DistrictId   int64              `bson:"district_id" json:"district_id"`
	WardId       int64              `bson:"ward_id" json:"ward_id"`
	VillageId    int64              `bson:"village_id" json:"village_id"`
	AreaId       primitive.ObjectID `bson:"area_id" json:"area_id"`
	IPAddress    string             `bson:"ip_address" json:"ip_address"`
	Status       bool               `bson:"status" json:"status"`
	IsActive     bool               `bson:"is_active" json:"is_active"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at" json:"-"`

	User     User      `bson:"-" json:"user"`
	Area     Area      `bson:"-" json:"area"`
	Province *Province `bson:"-" json:"province"`
	District *District `bson:"-" json:"district"`
	Ward     *Ward     `bson:"-" json:"ward"`
	Village  *Village  `bson:"-" json:"village"`
}
type ConfigCamera struct {
	ID           primitive.ObjectID `bson:"_id" json:"-"`
	MaxWeight    float64            `bson:"max_weight" json:"max_weight"`
	IsActiveTime bool               `bson:"is_active_time" json:"is_active_time"`
	TimeFrom     string             `bson:"time_from" json:"time_from"`
	TimeTo       string             `bson:"time_to" json:"time_to"`
	Dates        []string           `bson:"dates" json:"dates"`
}

type Cameras []Camera

func (u *Camera) CollectionName() string {
	return "cameras"
}

func (u *Camera) String() string {
	uJSON, _ := json.Marshal(u)
	return string(uJSON)
}

func (u *Camera) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Camera) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.ConfigCamera.ID = u.ID
	u.CreatedAt = helpers.Now()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Camera) Update() error {
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

func (u *Camera) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{
		"deleted_at": helpers.Now(),
	}}, nil); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Camera) Find(filter bson.M, opts ...*options.FindOptions) (Cameras, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	data := make(Cameras, 0)

	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Camera
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

func (u *Camera) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Camera) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "user" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var entry User
				_ = entry.First(bson.M{"_id": u.UserId, "deleted_at": nil})
				u.User = entry
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

type ResponseHls struct {
	Data struct {
		GetHlsStream struct {
			Link string `json:"link"`
		} `json:"getHlsStream"`
	} `json:"data"`
	Div string `json:"div"`
}

func (u *Camera) RenderStreamUrl(startTime string) string {
	conf := configProvider.GetConfig()
	query := ""
	var err error
	var req *http.Request
	var res *http.Response
	//API kiểu GraphQL của H-Factor để lấy link LIVE/PLAYBACK
	if startTime == "" {
		query = fmt.Sprintf(
			`{"query":"query{getHlsStream(user:\"%s\",pass:\"%s\",ip:\"%s\",port:554,channel:\"%s\"){link,div}}"}`,
			u.Username, u.Password, u.IPAddress, u.Channel,
		)
	} else {
		query = fmt.Sprintf(
			`{"query":"query{getHlsStream(user:\"%s\",pass:\"%s\",ip:\"%s\",port:554,channel:\"%s\",startTime:\"%s\"){link,div}}"}`,
			u.Username, u.Password, u.IPAddress, u.Channel, startTime,
		)
	}
	payload := strings.NewReader(query)
	client := &http.Client{}

	if req, err = http.NewRequest("GET", conf.GetString("camera.host"), payload); err != nil {
		return ""
	}
	req.Header.Add("token", conf.GetString("camera.token"))
	req.Header.Add("Content-Type", "application/json")

	if res, err = client.Do(req); err != nil || res.StatusCode != 200 {
		return ""
	}
	defer res.Body.Close()
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return ""
	}
	result := ResponseHls{}
	_ = json.Unmarshal(body, &result)
	return result.Data.GetHlsStream.Link
}
