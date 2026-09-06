package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Anwesa-s/AKSH/internal/health"
	"github.com/Anwesa-s/AKSH/internal/models"
)

// Home endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{
		Name:    "AKSH",
		Message: "Welcome to AKSH API Gateway",
	})
}

type HealthResponse struct {
	Status   string                 `json:"status"`
	Services []health.ServiceStatus `json:"services"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	services := []health.ServiceStatus{
		health.CheckService("users", "http://localhost:9001"),
		health.CheckService("orders", "http://localhost:9002"),
		health.CheckService("payments", "http://localhost:9003"),
	}

	status := "healthy"

	for _, service := range services {
		if service.Status != "healthy" {
			status = "unhealthy"
			break
		}
	}

	response := HealthResponse{
		Status:   status,
		Services: services,
	}
	if status == "unhealthy" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	json.NewEncoder(w).Encode(response)
}

// Version endpoint
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{
		Version: "0.1.0",
	})
}

// Users endpoint
func UsersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path

	parts := strings.Split(path, "/")

	response := models.Response{
		Service: "Users",
	}

	if len(parts) > 2 {
		response.UserID = parts[2]
	}

	json.NewEncoder(w).Encode(response)
}

// Orders endpoint
func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{
		Service: "Orders",
	})
}

// Payments endpoint
func PaymentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{
		Service: "Payments",
	})
}
