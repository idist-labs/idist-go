package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type ConnectionStatus struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Alias     string             `bson:"alias" json:"alias"`
	DeletedAt *time.Time         `bson:"deleted_at" json:"-"`
}

type ConnectionStatuses []ConnectionStatus

func (u *ConnectionStatus) CollectionName() string {
	return "connectionStatuses"
}

func (u *ConnectionStatus) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *ConnectionStatus) FindByID(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, bson.M{"_id": ID}); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *ConnectionStatus) FindByAction(action string) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, bson.M{"action": action}); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *ConnectionStatus) FindAndCount(filter interface{}, opts ...*options.FindOptions) (int64, ConnectionStatuses, error) {
	var wg sync.WaitGroup
	var total int64
	var errCount, errFind error
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(ConnectionStatuses, 0)

	/* Khai báo chờ 2 runtime */
	wg.Add(2)
	/* Đếm số bản ghi */
	go func() {
		defer wg.Done()
		total, errCount = DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count())
	}()
	/* Lấy danh sách bản ghi */
	go func() {
		defer wg.Done()
		var cursor *mongo.Cursor
		if cursor, errFind = DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); errFind == nil {
			for cursor.Next(ctx) {
				var elem ConnectionStatus
				if errFind = cursor.Decode(&elem); errFind != nil {
					return
				}
				data = append(data, elem)
			}
			if errFind = cursor.Err(); errFind != nil {
				return
			}
			errFind = cursor.Close(ctx)
		}
	}()
	/* Chờ 2 truy vấn */
	wg.Wait()
	/* */
	if errCount != nil {
		return 0, data, errCount
	} else if errFind != nil {
		return 0, data, errFind
	}
	return total, data, nil
}

func (u *ConnectionStatus) FindByIDs(DB *mongo.Database, ids []primitive.ObjectID) (ConnectionStatuses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(ConnectionStatuses, 0)

	/* Lấy danh sách bản ghi */
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	if cursor, err := DB.Collection(u.CollectionName()).Find(ctx, filter, options.Find()); err == nil {
		for cursor.Next(ctx) {
			var elem ConnectionStatus
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

func (u *ConnectionStatus) Find(filter interface{}, opts ...*options.FindOptions) (ConnectionStatuses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(ConnectionStatuses, 0)
	/* Lấy danh sách bản ghi */
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem ConnectionStatus
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

func (u *ConnectionStatus) Update(data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if data != nil {
		if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
			"$set": data,
		}, options.Update()); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
			"$set": u,
		}, options.Update()); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (u *ConnectionStatus) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{
		"deleted_at": time.Now(),
	}}, nil); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *ConnectionStatus) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}
