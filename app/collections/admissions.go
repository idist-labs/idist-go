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

type Admission struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Gender          string             `bson:"gender" json:"gender"`
	Birthday        string             `bson:"birthday" json:"birthday"`
	Ethnic          string             `bson:"ethnic" json:"ethnic"`
	Religion        string             `bson:"religion" json:"religion"`
	AverageSubjects bool               `bson:"average_subjects" json:"average_3_subjects"`
	AverageYears    bool               `bson:"average_years" json:"average_3_years"`
	GraduationYear  int64              `bson:"graduation_year" json:"graduation_year"`
	ProvinceId      int64              `bson:"province_id" json:"province_id"`
	Cmnd            string             `bson:"cmnd" json:"cmnd"`
	ClassName       string             `bson:"class_name" json:"class_name"`
	Academic        string             `bson:"academic" json:"academic"`
	Conduct         string             `bson:"conduct" json:"conduct"`
	IssuedAt        string             `bson:"issued_at" json:"issued_at"`
	PriorityObject  string             `bson:"priority_object" json:"priority_object"`
	KhuVuc          string             `bson:"khu_vuc" json:"khu_vuc"`
	IssuePlaceId    int64              `bson:"issue_place_id" json:"issue_place_id"`
	Address         string             `bson:"address" json:"address"`
	Class10         Class              `bson:"class_10" json:"class_10"`
	Class11         Class              `bson:"class_11" json:"class_11"`
	Class12         Class              `bson:"class_12" json:"class_12"`
	Email           string             `bson:"email" json:"email"`
	Phone           string             `bson:"phone" json:"phone"`
	ParentPhone     string             `bson:"parent_phone" json:"parent_phone"`
	Nganh1          Nganh              `bson:"nganh_1" json:"nganh_1"`
	Nganh2          Nganh              `bson:"nganh_2" json:"nganh_2"`
	Nganh3          Nganh              `bson:"nganh_3" json:"nganh_3"`
	Nganh4          Nganh              `bson:"nganh_4" json:"nganh_4"`
	Nganh5          Nganh              `bson:"nganh_5" json:"nganh_5"`
	Nganh6          Nganh              `bson:"nganh_6" json:"nganh_6"`
	KhaoSat         []KhaoSatItem      `bson:"khao_sat" json:"khao_sat"`
	IssuePlace      Province           `bson:"issue_place" json:"issue_place"`
	Province        Province           `bson:"-" json:"province"`

	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at" json:"deleted_at"`
}

type Admissions []Admission

func (u *Admission) CollectionName() string {
	return "admissions"
}

func (u *Admission) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Admission) Create() error {
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

func (u *Admission) Update() error {
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
func (u *Admission) Find(filter bson.M, opts ...*options.FindOptions) (Admissions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Admissions, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Admission
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

func (u *Admission) Delete() error {
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

func (u *Admission) Preload(properties ...string) {
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
