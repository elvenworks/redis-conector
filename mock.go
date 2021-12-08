package redis

import (
	"github.com/stretchr/testify/mock"
)

type MockRedis struct {
	mock.Mock
}

func (mock MockRedis) SetMessage(key string, message interface{}) error {
	args := mock.Called(key, message)

	return args.Error(0)
}

func (mock MockRedis) GetMessage(key string) (string, error) {
	args := mock.Called(key)

	return args.Get(0).(string), args.Error(1)
}

type MockRedisCluster struct {
	mock.Mock
}

func (mock MockRedisCluster) ClusterSetMessage(key string, message interface{}) error {
	args := mock.Called(key, message)

	return args.Error(0)
}

func (mock MockRedisCluster) ClusterGetMessage(key string) (string, error) {
	args := mock.Called(key)

	return args.Get(0).(string), args.Error(1)
}
