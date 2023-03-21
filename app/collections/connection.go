package collections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"idist-core/app/providers/mongoProvider"
	"time"
)

const QueryTimeOut = 10 * time.Second

func DB() *mongo.Database {
	return mongoProvider.GetMongoDB()
}
