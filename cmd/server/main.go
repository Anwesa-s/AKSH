package main

import (
	"fmt"
	"net/http"

	"github.com/Anwesa-s/AKSH/internal/handlers"
	"github.com/Anwesa-s/AKSH/internal/router"
)

func main() {

	r := router.NewRouter()

	r.Register("/", handlers.HomeHandler)
	r.Register("/health", handlers.HealthHandler)
	r.Register("/version", handlers.VersionHandler)

	r.Register("/users", handlers.UsersHandler)
  r.Register("/users/:id", handlers.UsersHandler)

  r.Register("/orders", handlers.OrdersHandler)
  r.Register("/orders/:id", handlers.OrdersHandler)

  r.Register("/payments", handlers.PaymentsHandler)
  r.Register("/payments/:id", handlers.PaymentsHandler)

	fmt.Println("🚀 AKSH running on http://localhost:8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}