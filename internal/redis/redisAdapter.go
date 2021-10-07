package redisAdapter

import (
	"time"

	"github.com/go-redis/redis"
)


var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func SetDict(key string, value []byte) *redis.StatusCmd{

	return rdb.HMSet(key, map[string]interface{}{ time.Now().String(): value})
}

func GetDict(key string) *redis.StringStringMapCmd{

	return rdb.HGetAll(key)
}

func DeleteDictField(key string, fields ...string)  *redis.IntCmd {

	return rdb.HDel(key, fields...)
}