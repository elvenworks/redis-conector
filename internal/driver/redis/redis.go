package driver

import "github.com/go-redis/redis/v8"

func RedisNewClient(config *redis.Options) *redis.Client {

	redisClient := redis.NewClient(config)

	return redisClient

}
