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

type redisClient struct {
	RedisClient *redis.Client
}

func NewRedisActions(config *redis.Options) *redisClient {
	client := driver.RedisNewClient(config)

	return &redisClient{
		RedisClient: client,
	}
}

func (c *redisClient) SetMessage(m domain.InputMessage) error {

	bytes, err := json.Marshal(m.Message)
	if err != nil {
		return err
	}

	errRedis := c.RedisClient.Set(context.Background(), m.Key, bytes, 0)

	if errRedis != nil {
		return errRedis.Err()
	}

	return nil
}

func (c *redisClient) GetMessage(m domain.InputMessage) (string, error) {

	redisResponse, err := c.RedisClient.Get(context.Background(), m.Key).Result()

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
