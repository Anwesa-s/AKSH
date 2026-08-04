package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Name:    "AKSH",
		Message: "Welcome to AKSH API Gateway",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Status: "Healthy",
	})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Version: "0.1.0",
	})
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	fmt.Println("🚀 AKSH running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}