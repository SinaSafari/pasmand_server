package redis

import "github.com/go-redis/redis/v8"

var RedisClient *redis.Client

func SetupRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisClient = rdb
}
