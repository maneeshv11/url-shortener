package cache

import (
	"github.com/go-redis/redis"
	. "url-shortener/utils"
)

var redisCache *redis.Client

func GetInstance() *redis.Client {
	if redisCache != nil {
		return redisCache
	} else {
		redisCache = redis.NewClient(&redis.Options{
			Addr:     GetConfig().RedisUri,
			Password: "",
			DB:       0,
		})
		redisCache.Ping()
		return redisCache
	}

}

func Get(key string) (string, error) {
	return GetInstance().Get(key).Result()
}

func Set(key, value string) (err error) {
	_, err = GetInstance().Set(key, value, 0).Result()
	return
}

func RPush(key string, value ...string) (itemsAdded int64, err error) {

	itemsAdded, err = GetInstance().RPush(key, value).Result()
	return
}

func LPop(key string) (item string, err error) {

	item, err = GetInstance().LPop(key).Result()
	return
}

func RPop(key string) (item string, err error) {

	item, err = GetInstance().RPop(key).Result()
	return
}
