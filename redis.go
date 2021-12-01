package redis

import (
	"crypto/tls"
	"time"

	"github.com/elvenworks/redis-conector/internal/delivery/actions"
	"github.com/elvenworks/redis-conector/internal/domain"
	"github.com/go-redis/redis/v8"
)

type InputConfig struct {
	Host       string
	User       string
	Password   string
	DB         int
	MaxRetries int
	Timeout    time.Duration
}

type Redis struct {
	Config  *redis.Options
	actions actions.IRedisActions
}

func InitRedis(config InputConfig) *Redis {
	RedisConfig := Redis{
		Config: &redis.Options{
			Addr:        config.Host,
			Username:    config.User,
			Password:    config.Password,
			DB:          config.DB,
			MaxRetries:  config.MaxRetries,
			DialTimeout: config.Timeout,
		},
	}

	if config.Password != "" {
		RedisConfig.Config.TLSConfig = &tls.Config{}
	}

	return &Redis{
		Config: RedisConfig.Config,
	}
}

func (r *Redis) SetMessage(key string, message interface{}) error {

	inputMessage := domain.InputMessage{
		Key:     key,
		Message: message,
	}

	r.actions = actions.NewRedisActions(r.Config)

	defer r.actions.Close()

	err := r.actions.SetMessage(inputMessage)

	if err != nil {
		return err
	}

	return nil

}

func (r *Redis) GetMessage(key string) (string, error) {
	inputMessage := domain.InputMessage{
		Key: key,
	}

	r.actions = actions.NewRedisActions(r.Config)

	defer r.actions.Close()

	message, err := r.actions.GetMessage(inputMessage)

	if err != nil {
		return "", err
	}

	return message, nil
}
