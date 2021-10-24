package rediscache

import (
	"context"
	"github.com/go-redis/redis/v8"
)


type RedisCache struct {
	Client *redis.Client
}

func GetClient() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func (r *RedisCache) Get(key string) (map[string]string, error) {

	result := r.Client.HGetAll(context.Background(),key)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

func (r *RedisCache) Add(key string, value map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(),key, value)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *RedisCache) Delete(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(),key, fields ...)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *RedisCache) DeleteAll(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(),key)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}
