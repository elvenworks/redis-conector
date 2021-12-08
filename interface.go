package redis

type IRedis interface {
	SetMessage(key string, message interface{}) error
	GetMessage(key string) (string, error)
}

type IRedisCluster interface {
	ClusterSetMessage(key string, message interface{}) error
	ClusterGetMessage(key string) (string, error)
}
