package redis

import (
	"crypto/tls"
	"time"

	"github.com/elvenworks/redis-conector/internal/delivery/actions"
	"github.com/elvenworks/redis-conector/internal/domain"
	"github.com/go-redis/redis/v8"
)

type InputClusterConfig struct {
	Hosts      []string
	User       string
	Password   string
	MaxRetries int
	Timeout    time.Duration
}

type RedisCluster struct {
	ConfigCluster  *redis.ClusterOptions
	clusterActions actions.IRedisClusterActions
}

func InitRedisCluster(config InputClusterConfig) *RedisCluster {

	RedisClusterConfig := RedisCluster{
		ConfigCluster: &redis.ClusterOptions{
			Addrs:       config.Hosts,
			Username:    config.User,
			Password:    config.Password,
			MaxRetries:  config.MaxRetries,
			DialTimeout: config.Timeout,
		},
	}

	if config.Password != "" {
		RedisClusterConfig.ConfigCluster.TLSConfig = &tls.Config{}
	}

	return &RedisCluster{
		ConfigCluster: RedisClusterConfig.ConfigCluster,
	}
}

func (r *RedisCluster) ClusterSetMessage(key string, message interface{}) error {

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

func (r *RedisCluster) ClusterGetMessage(key string) (string, error) {
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
