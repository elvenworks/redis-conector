package actions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/elvenworks/redis-conector/internal/domain"
	driver "github.com/elvenworks/redis-conector/internal/driver/redis"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type clusterClient struct {
	RedisClusterClient *redis.ClusterClient
}

func NewRedisClusterActions(config *redis.ClusterOptions) *clusterClient {

	client := driver.RedisNewClusterClient(config)

	return &clusterClient{
		RedisClusterClient: client,
	}
}

func (c *clusterClient) ClusterSetMessage(m domain.InputMessage) error {

	bytes, err := json.Marshal(m.Message)
	if err != nil {
		return err
	}

	errRedis := c.RedisClusterClient.Set(context.Background(), m.Key, bytes, 0)

	if errRedis != nil {
		return errRedis.Err()
	}

	return nil

}

func (c *clusterClient) ClusterGetMessage(m domain.InputMessage) (string, error) {

	redisResponse, err := c.RedisClusterClient.Get(context.Background(), m.Key).Result()

	if err == redis.Nil {
		logrus.Info("Key does not exist")
		e := fmt.Sprintf("%s does not exist", m.Key)
		return "", errors.New(e)
	} else if err != nil {
		logrus.Info(err)
		return "", err
	}

	return redisResponse, nil

}

func (c *clusterClient) ClusterClose() error {

	return c.RedisClusterClient.Close()
}
