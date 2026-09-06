package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {

		user := User{
			ID:   "101",
			Name: "AKSH User",
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	if r.Method == http.MethodPost {

		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		user.ID = "102"

		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

func main() {

	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("👤 User Service running on http://localhost:9001")

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println(err)
	}
}