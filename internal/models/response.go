package models

type Response struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
	Service string `json:"service,omitempty"`
	UserID string `json:"userId,omitempty"`
}