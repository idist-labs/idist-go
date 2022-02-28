package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"idist-core/app/providers/redisProvider"
	"time"
)

type Cache struct {
	Instance *redis.Client
}

var cacheInstance *Cache

func GetInstance() *Cache {
	if cacheInstance == nil {
		fmt.Println("Create instance redis cache")
		cacheInstance = new(Cache)
		cacheInstance.Instance = redisProvider.GetClient()
	}

	return cacheInstance
}

func (u *Cache) Get(key string, result interface{}) error {
	if value, err := u.Instance.Get(key).Result(); err == nil {
		return json.Unmarshal([]byte(value), &result)
	} else {
		return err
	}
}

func (u *Cache) Set(key string, value string, duration int) error {
	return u.Instance.Set(key, value, time.Minute*time.Duration(duration)).Err()
}
func (u *Cache) SetInterface(key string, data interface{}, duration int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return u.Instance.Set(key, value, time.Second*time.Duration(duration)).Err()
}

func (u *Cache) DelCache(key string) error {
	_, err := u.Instance.Del(key).Result()
	if err == redis.Nil {
		return nil
	} else {
		return err
	}
}
