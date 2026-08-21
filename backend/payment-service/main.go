package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payment struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Amount int    `json:"amount"`
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	payment := Payment{
		ID:     "P1001",
		Status: "successful",
		Amount: 49999,
	}

	json.NewEncoder(w).Encode(payment)
}

func main() {

	http.HandleFunc("/payments", paymentHandler)

	fmt.Println("💳 Payment Service running on http://localhost:9003")

	err := http.ListenAndServe(":9003", nil)
	if err != nil {
		fmt.Println(err)
	}
}