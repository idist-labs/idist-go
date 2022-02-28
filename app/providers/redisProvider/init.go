package redisProvider

import (
	"fmt"
	"github.com/go-redis/redis"
	"idist-core/app/providers/configProvider"
)

var client *redis.Client

func Init() {
	fmt.Println("------------------------------------------------------------")
	cf := configProvider.GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     cf.GetString("redis.addr"),
		Password: cf.GetString("redis.password"),
		DB:       0,
	})

	if cf.GetString("env") == "development" {
		if message, err := client.FlushAll().Result(); err != nil {
			fmt.Println(message)
		} else {
			fmt.Println("Redis flushed all cache.")
		}
	}

	if _, err := client.Ping().Result(); err != nil {
		fmt.Println("Redis server can't connect.")
	} else {
		fmt.Println("Redis server connected.")
	}
}

func GetClient() *redis.Client {
	return client
}
