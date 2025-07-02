package components

import "github.com/redis/go-redis/v9"

type Redis struct {
	Client *redis.Client
}

func NewRedis(addr string, password string) *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}),
	}
}
