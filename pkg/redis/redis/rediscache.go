package rediscache

import (
	"ClientWorkerService/pkg/helper"
	"context"
	"os"

	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type RedisCache struct {
	Client *redis.Client
}

func GetClient() *redis.Client {
	godotenv.Load()
	rdb := redis.NewClient(&redis.Options{
		Addr:     helper.ResolvePath("REDIS_HOST", "REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
	})
	return rdb
}

func (r *RedisCache) Get(key string) (map[string]string, error) {

	result := r.Client.HGetAll(context.Background(), key)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache RedisConnection ConnectRedis error: ": result.Err(),
		})
		return nil, result.Err()
	}
	return result.Val(), nil
}

func (r *RedisCache) Add(key string, value map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(), key, value)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache Redis error: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}

func (r *RedisCache) Delete(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(), key, fields...)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache Redis error: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}

func (r *RedisCache) DeleteAll(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(), key)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache Redis error: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}
