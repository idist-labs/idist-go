package collections

import (
	"ai-camera-api-cms/app/providers/mongoProvider"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const QueryTimeOut = 10 * time.Second

func DB() *mongo.Database {
	return mongoProvider.GetMongoDB()
}

type Repository interface {
}
