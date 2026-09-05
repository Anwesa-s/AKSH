package ratelimit

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisLimiter struct {
	client *redis.Client
	limit  int
	window time.Duration
}

func NewRedisLimiter(client *redis.Client, limit int, window time.Duration) *RedisLimiter {
	return &RedisLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

func (r *RedisLimiter) Allow(key string) (bool, error) {

	ctx := context.Background()

	redisKey := fmt.Sprintf("aksh:rate:%s", key)

	count, err := r.client.Incr(ctx, redisKey).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		err := r.client.Expire(ctx, redisKey, r.window).Err()
		if err != nil {
			return false, err
		}
	}

	return count <= int64(r.limit), nil
}