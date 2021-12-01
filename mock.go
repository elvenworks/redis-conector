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
