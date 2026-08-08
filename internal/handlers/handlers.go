package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

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

// Health endpoint
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{
		Status: "Healthy",
	})
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