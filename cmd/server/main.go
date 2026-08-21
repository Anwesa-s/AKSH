package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anwesa-s/AKSH/internal/config"
	"github.com/Anwesa-s/AKSH/internal/proxy"
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

	fmt.Println("🚀 AKSH running on http://localhost:8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}