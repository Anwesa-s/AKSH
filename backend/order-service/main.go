package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	ID      string `json:"id"`
	Product string `json:"product"`
	Status  string `json:"status"`
}

func orderHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	order := Order{
		ID:      "5001",
		Product: "Laptop",
		Status:  "confirmed",
	}

	json.NewEncoder(w).Encode(order)
}

func main() {

	http.HandleFunc("/orders", orderHandler)

	fmt.Println("📦 Order Service running on http://localhost:9002")

	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		fmt.Println(err)
	}
}