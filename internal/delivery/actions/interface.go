package actions

import "github.com/elvenworks/redis-conector/internal/domain"

type IRedisActions interface {
	SetMessage(m domain.InputMessage) error
	GetMessage(m domain.InputMessage) (string, error)
	Close() error
}

type IRedisClusterActions interface {
	ClusterSetMessage(m domain.InputMessage) error
	ClusterGetMessage(m domain.InputMessage) (string, error)
	ClusterClose() error
}
