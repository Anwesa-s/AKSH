package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Anwesa-s/AKSH/internal/redis"
)

func main() {

	client := redis.NewClient("localhost:6379")

	ctx := context.Background()

	// Store a value in Redis
	err := client.Set(ctx, "aksh:test", "hello-redis", 0).Err()
	if err != nil {
		log.Fatal("Failed to set value:", err)
	}

	fmt.Println("Value stored in Redis")

	// Retrieve the value
	value, err := client.Get(ctx, "aksh:test").Result()
	if err != nil {
		log.Fatal("Failed to get value:", err)
	}

	fmt.Println("Value retrieved:", value)
}