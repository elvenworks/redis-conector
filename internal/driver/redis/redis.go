package driver

import "github.com/go-redis/redis/v8"

func RedisNewClient(config *redis.Options) *redis.Client {

	redisClient := redis.NewClient(config)

	return redisClient

}

func RedisNewClusterClient(clusterConfig *redis.ClusterOptions) *redis.ClusterClient {

	redisClusterClient := redis.NewClusterClient(clusterConfig)

	return redisClusterClient
}
