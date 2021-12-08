package redis

import (
	"crypto/tls"
	"time"

	"github.com/elvenworks/redis-conector/internal/delivery/actions"
	"github.com/elvenworks/redis-conector/internal/domain"
	"github.com/go-redis/redis/v8"
)

type InputConfig struct {
	Hosts      []string
	Host       string
	User       string
	Password   string
	DB         int
	MaxRetries int
	Timeout    time.Duration
}

type Redis struct {
	Config         *redis.Options
	ConfigCluster  *redis.ClusterOptions
	actions        actions.IRedisActions
	clusterActions actions.IRedisClusterActions
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

	RedisClusterConfig := Redis{
		ConfigCluster: &redis.ClusterOptions{
			Addrs:       config.Hosts,
			Username:    config.User,
			Password:    config.Password,
			MaxRetries:  config.MaxRetries,
			DialTimeout: config.Timeout,
		},
	}

	if config.Password != "" {
		RedisConfig.Config.TLSConfig = &tls.Config{}
		RedisClusterConfig.ConfigCluster.TLSConfig = &tls.Config{}
	}

	return &Redis{
		Config:        RedisConfig.Config,
		ConfigCluster: RedisClusterConfig.ConfigCluster,
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

func (r *Redis) ClusterSetMessage(key string, message interface{}) error {

	inputMessage := domain.InputMessage{
		Key:     key,
		Message: message,
	}

	r.clusterActions = actions.NewRedisClusterActions(r.ConfigCluster)

	defer r.clusterActions.ClusterClose()

	err := r.clusterActions.ClusterSetMessage(inputMessage)

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

func (r *Redis) ClusterGetMessage(key string) (string, error) {
	inputMessage := domain.InputMessage{
		Key: key,
	}

	r.clusterActions = actions.NewRedisClusterActions(r.ConfigCluster)

	defer r.clusterActions.ClusterClose()

	message, err := r.clusterActions.ClusterGetMessage(inputMessage)

	if err != nil {
		return "", err
	}

	return message, nil
}
