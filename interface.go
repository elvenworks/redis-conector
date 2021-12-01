package redis

type IRedis interface {
	SetMessage(key string, message interface{}) error
	GetMessage(key string) (string, error)
}
