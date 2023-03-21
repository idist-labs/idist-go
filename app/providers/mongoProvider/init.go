package mongoProvider

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"idist-core/app/providers/configProvider"
	"time"
)

var mongoDB *mongo.Database

var CTimeOut = 10 * time.Second

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), CTimeOut)
	defer cancel()
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Mongo: Đang khởi tạo kết nối...")
	CTimeOut = time.Duration(configProvider.GetConfig().GetInt64("mongo.timeout")) * time.Second
	c := configProvider.GetConfig()

	var uri = c.GetString("mongo.uri")
	if len(uri) == 0 {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
			c.GetString("mongo.username"),
			c.GetString("mongo.password"),
			c.GetString("mongo.host"),
			c.GetString("mongo.port"),
			c.GetString("mongo.database"),
			c.GetString("mongo.authSource"))
	}
	if c.GetString("mongo.username") == "" {
		uri = fmt.Sprintf("mongodb://%s:%s/%s?authSource=%s",
			c.GetString("mongo.host"),
			c.GetString("mongo.port"),
			c.GetString("mongo.database"),
			c.GetString("mongo.authSource"),
		)
	}

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			if configProvider.GetConfig().GetBool("log.query.detail") == true {
				fmt.Println(evt.Command.String())
			} else {
				switch evt.CommandName {
				case "find":
					fmt.Println("db.getCollection(" + evt.Command.Lookup(evt.CommandName).String() + ").find(" + evt.Command.Lookup("filter").String() + ")")
					return
				case "insert":
					fmt.Println("db.getCollection(" + evt.Command.Lookup(evt.CommandName).String() + ").insert(" + evt.Command.Lookup("documents").String() + ")")
					return
				case "update":
					fmt.Println("db.getCollection(" + evt.Command.Lookup(evt.CommandName).String() + ").update(" + evt.Command.Lookup("updates").String() + ")")
					return
				default:
					fmt.Println(evt.Command.String())
				}
			}
		},
	}
	clientOptions := options.Client().ApplyURI(uri).SetMonitor(cmdMonitor)
	if client, err := mongo.Connect(ctx, clientOptions); err != nil {
		fmt.Println("Mongo: Lỗi kết nối")
		panic(err)
	} else {
		fmt.Println("Mongo: Kết nối thành công")
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			fmt.Println("Mongo: Ping không thành công", zap.Error(err), zap.String("Mongo uri:", uri))
			panic(err)
		} else {
			fmt.Println("Mongo: {ping: Pong}")
		}
		mongoDB = client.Database(c.GetString("mongo.database"))
	}
}

func GetMongoDB() *mongo.Database {
	return mongoDB
}

func CloseMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), CTimeOut)
	defer cancel()
	_ = mongoDB.Client().Disconnect(ctx)
}
