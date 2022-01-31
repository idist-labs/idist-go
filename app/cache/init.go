package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"idist-go/app/providers/redisProvider"
	"time"
)

type Cache struct {
	Instance *redis.Client
}

func (u *Cache) Init() {
	fmt.Println("Redis instance created")
	u.Instance = redisProvider.GetClient()
}

func (u *Cache) Get(key string) (interface{}, error) {
	var result interface{}
	if value, err := u.Instance.Get(key).Result(); err == nil {
		err = json.Unmarshal([]byte(value), &result)
		return result, err
	} else {
		return nil, err
	}
}

func (u *Cache) SetCache(key string, value string, duration int) error {
	return u.Instance.Set(key, value, time.Minute*time.Duration(duration)).Err()
}

func (u *Cache) DelCache(key string) error {
	_, err := u.Instance.Del(key).Result()
	if err == redis.Nil {
		return nil
	} else {
		return err
	}
}
