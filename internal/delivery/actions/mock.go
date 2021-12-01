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
