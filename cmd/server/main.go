package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Anwesa-s/AKSH/internal/config"
	"github.com/Anwesa-s/AKSH/internal/handlers"
	"github.com/Anwesa-s/AKSH/internal/middleware"
	"github.com/Anwesa-s/AKSH/internal/proxy"
	"github.com/Anwesa-s/AKSH/internal/ratelimit"
	"github.com/Anwesa-s/AKSH/internal/redis"
	"github.com/Anwesa-s/AKSH/internal/router"
)

func main() {

	// Load configuration
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Create router
	r := router.NewRouter()

	// Admin route
	r.Register(
		"/admin",
		middleware.Auth(
			middleware.AdminOnly(
				http.HandlerFunc(handlers.AdminHandler),
			),
		).ServeHTTP,
	)

	// Login route
	r.Register("/login", handlers.LoginHandler)

	// Health route
	r.Register("/health", handlers.HealthHandler)

	// Register routes from configuration
	for _, route := range cfg.Routes {

		service, exists := cfg.Services[route.Service]

		if !exists {
			log.Printf(
				"Service %s not found for route %s",
				route.Service,
				route.Path,
			)
			continue
		}

		serviceProxy, err := proxy.NewProxy(service.URL)
		if err != nil {
			log.Printf(
				"Failed to create proxy for %s: %v",
				route.Service,
				err,
			)
			continue
		}

		r.Register(route.Path, serviceProxy.ServeHTTP)

		fmt.Printf(
			"Registered route: %s → %s\n",
			route.Path,
			service.URL,
		)
	}

	// Create Redis client
	redisClient := redis.NewClient("localhost:6379")

	// Parse rate-limit window from configuration
	window, err := time.ParseDuration(cfg.RateLimit.Window)
	if err != nil {
		log.Fatal("Invalid rate limit window:", err)
	}

	// Create Redis-backed rate limiter
	limiter := ratelimit.NewRedisLimiter(
		redisClient,
		cfg.RateLimit.Limit,
		window,
	)
	// Add middleware
	handler := middleware.Logger(
		middleware.RateLimit(limiter)(
			middleware.Auth(r),
		),
	)

	fmt.Println("🚀 AKSH running on http://localhost:8080")

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
