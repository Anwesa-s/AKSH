package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Anwesa-s/AKSH/internal/ratelimit"
	"github.com/Anwesa-s/AKSH/internal/redis"
)

func main() {

	client := redis.NewClient("localhost:6379")

	limiter := ratelimit.NewRedisLimiter(
		client,
		5,
		10*time.Second,
	)

	key := "127.0.0.1"

	for i := 1; i <= 8; i++ {

		allowed, err := limiter.Allow(key)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"Request %d → allowed: %v\n",
			i,
			allowed,
		)
	}
}