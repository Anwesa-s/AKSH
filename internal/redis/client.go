package redis

import (
	"context"

	goredis "github.com/redis/go-redis/v9"
)

func NewClient(addr string) *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr: addr,
	})
}

func Ping(client *goredis.Client) error {
	return client.Ping(context.Background()).Err()
}