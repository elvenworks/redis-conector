package actions

import (
	"github.com/elvenworks/redis-conector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockRedisClient struct {
	mock.Mock
}

func (mock MockRedisClient) SetMessage(m domain.InputMessage) error {
	args := mock.Called(m)

	return args.Error(0)
}

func (mock MockRedisClient) GetMessage(m domain.InputMessage) (string, error) {
	args := mock.Called(m)

	return args.Get(0).(string), args.Error(1)
}

func (mock MockRedisClient) Close() error {
	args := mock.Called()

	return args.Error(0)
}

type MockClusterClient struct {
	mock.Mock
}

func (mock MockClusterClient) ClusterSetMessage(m domain.InputMessage) error {
	args := mock.Called(m)

	return args.Error(0)
}

func (mock MockClusterClient) ClusterGetMessage(m domain.InputMessage) (string, error) {
	args := mock.Called(m)

	return args.Get(0).(string), args.Error(1)
}

func (mock MockClusterClient) ClusterClose() error {
	args := mock.Called()

	return args.Error(0)
}
