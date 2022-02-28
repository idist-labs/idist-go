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

type Category struct {
	ID          primitive.ObjectID  `bson:"_id" json:"id"`
	Name        string              `bson:"name" json:"name"`
	Description string              `bson:"description" json:"description"`
	Slug        string              `bson:"slug" json:"slug"`
	ParentID    *primitive.ObjectID `bson:"parent_id" json:"parent_id"`
	IsActive    bool                `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time          `bson:"deleted_at" json:"-"`

	Parent        *Category `bson:"-" json:"parent"`
	CountArticles int64     `bson:"-" json:"count_articles"`
}

type Categories []Category

func (u *Category) CollectionName() string {
	return "districts"
}

func (u *Category) Find(filter interface{}, opts ...*options.FindOptions) (Categories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Categories, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Category
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

func (u *Category) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Category) Update() error {
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

func (u *Category) Delete() error {
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

func (u *Category) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Category) Create() error {
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

func (u *Category) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "parent" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				entry := Category{}
				_ = entry.First(bson.M{"_id": u.ParentID, "deleted_at": nil})
				u.Parent = &entry
			}()
		}
	}
	wg.Wait()

}
