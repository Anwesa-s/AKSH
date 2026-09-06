package health

import (
	"net/http"
	"time"
)

type ServiceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

func CheckService(name string, url string) ServiceStatus {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(url + "/health")

	if err != nil {
		return ServiceStatus{
			Name:   name,
			Status: "unhealthy",
			Reason: "service unreachable",
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ServiceStatus{
			Name:   name,
			Status: "unhealthy",
			Reason: "health check failed",
		}
	}

	return ServiceStatus{
		Name:   name,
		Status: "healthy",
	}
}
